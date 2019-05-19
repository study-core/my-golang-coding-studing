package main

import (
	"unsafe"
	"runtime/race"
	"sync/atomic"
)


// 首次使用后不得复制池。
type WaitGroup struct {
	// noCopy可以嵌入到结构中，在第一次使用后不可复制,使用go vet作为检测使用
	noCopy noCopy


	// 64 bit：高32 bit是计数器，低32位是 阻塞的goroutine计数。
	// 64位的原子操作需要64位的对齐，但是32位。
	// 编译器不能确保它,所以分配了12个byte对齐的8个byte作为状态。

	// 共12个字节，低4字节用于记录wait等待次数，
	// 高8字节是计数器（64位机器是高8字节，32机器是中间4个字节，
	// 因为64位机器的原子操作需要64位的对齐，但是32位的编译器不能确保。）
	// 所以分配了12个byte对齐的8个byte作为状态
	// (即：在高低位上 总共用了12byte 代表8byte，为了完全覆盖 64位和32位机器)
	// 其实就是为了表达 高8byte，也即 高32位，仅此而已
	// byte=uint8范围：0~255，只取前8个元素。转为2进制：0000 0000，0000 0000... ...0000 0000
	state1 [12]byte
	// 信号量，用于唤醒goroution
	sema   uint32
}


// uintptr和unsafe.Pointer的区别就是：unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
// 而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象，uintptr类型的目标会被回收。
// state()函数可以获取到wg.state1数组中元素组成的二进制对应的十进制的值

// 根据编译器位数，获得标志位和等待次数的数据域
func (wg *WaitGroup) state() *uint64 {
	// 是否是 64位机器：因为64位机器站高8位
	if uintptr(unsafe.Pointer(&wg.state1))%8 == 0 {
		return (*uint64)(unsafe.Pointer(&wg.state1))
	} else {
		// 如果是 32位机器， 则首元素从第 4 位开始：因为32 位机器占中间4位
		return (*uint64)(unsafe.Pointer(&wg.state1[4]))
	}
}


func (wg *WaitGroup) Add(delta int) {

	// 获取到wg.state1数组中元素组成的二进制对应的十进制的值的指针
	statep := wg.state()
	if race.Enabled {
		_ = *statep
		if delta < 0 {

			race.ReleaseMerge(unsafe.Pointer(wg))
		}
		race.Disable()
		defer race.Enable()
	}

	// 将标记为加delta
	// 因为高32位是计数器
	// 所以把 delta的值左移32位，并从数组的首元素处开始赋值
	// statep 对于 [12]byte 来说：
	// 如果是64位的机器，那么首元素地址就是 0 下标处开始
	// 如果是32位机器，那么首元素地址是 4 下标处开始
	state := atomic.AddUint64(statep, uint64(delta)<<32)
	// 获取计数器的值： 转了 int32
	v := int32(state >> 32)
	//获得调用 wait（）等待次数：转了 uint32
	w := uint32(state)
	if race.Enabled {
		if delta > 0 && v == int32(delta) {
			race.Read(unsafe.Pointer(&wg.sema))
		}
	}

	// 计数器为负数，报panic
	//标记位不能小于0（done过多或者Add（）负值太多）
	if v < 0 {
		panic("sync: negative WaitGroup counter")
	}
	// 不能Add 与Wait 同时调用
	if w != 0 && delta > 0 && v == int32(delta) {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
	// Add 完毕
	if v > 0 || w == 0 {
		return
	}

	// 当等待计数器> 0时，而goroutine将设置为0。
	// 此时不可能有同时发生的状态突变:
	// - Add()不能与 Wait() 同时发生，
	// - 如果计数器counter == 0，不再增加等待计数器

	// 不能Add 与Wait 同时调用
	if *statep != state {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}

	// 所有状态位清零
	*statep = 0
	//唤醒等待的 goroutine
	for ; w != 0; w-- {
		// 目的是作为一个简单的wakeup原语，以供同步使用。true为唤醒排在等待队列的第一个goroutine
		runtime_Semrelease(&wg.sema, false)
	}
}

// Done方法其实就是Add（-1）
func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

// Wait 会一直阻塞到 计数器值为0为止
func (wg *WaitGroup) Wait() {

	// 获取到wg.state1数组中元素组成的二进制对应的十进制的值的指针
	statep := wg.state()
	if race.Enabled {
		_ = *statep
		race.Disable()
	}

	// cas算法
	//循环检查计数器V啥时候等于0
	for {
		// 获取 计数值
		state := atomic.LoadUint64(statep)
		// 高32位是计数器
		v := int32(state >> 32)
		// 计数值
		w := uint32(state)

		// v == 0 说明 goroutine 执行结束
		if v == 0 {
			if race.Enabled {
				race.Enable()
				race.Acquire(unsafe.Pointer(wg))
			}
			return
		}

		// 尚有未执行完的go程，等待标志位+1（直接在低位处理，无需移位）
		// 增加等待goroution计数，对低32位加1，不需要移位
		if atomic.CompareAndSwapUint64(statep, state, state+1) {
			if race.Enabled && w == 0 {
				race.Write(unsafe.Pointer(&wg.sema))
			}
			// 目的是作为一个简单的sleep原语，以供同步使用
			runtime_Semacquire(&wg.sema)
			// 在上一次Wait返回之前重新使用WaitGroup，即在之前的Done 中没有清空 计数量就会有问题
			if *statep != 0 {
				panic("sync: WaitGroup is reused before previous Wait has returned")
			}
			if race.Enabled {
				race.Enable()
				race.Acquire(unsafe.Pointer(wg))
			}
			return
		}
	}
}

