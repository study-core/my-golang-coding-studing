package main

import (
	"unsafe"
	"sync/atomic"
)


// Cond实现了一个条件变量，一个等待或宣布事件发生的goroutines的集合点。
//
// 每个Cond都有一个相关的Locker L（通常是* Mutex或* RWMutex）。
type Cond struct {
	// 不允许复制,一个结构体,有一个Lock()方法,嵌入别的结构体中,表示不允许复制
	// noCopy对象，拥有一个Lock方法，使得Cond对象在进行go vet扫描的时候，能够被检测到是否被复制
	noCopy noCopy

	// 锁的具体实现，通常为 mutex 或者rwmutex
	L Locker

	// 通知列表,调用Wait()方法的goroutine会被放入list中,每次唤醒,从这里取出
	// notifyList对象，维护等待唤醒的goroutine队列,使用链表实现
	// 在 sync 包中被实现， src/sync/runtime.go
	notify  notifyList

	// 复制检查,检查cond实例是否被复制
	// copyChecker对象，实际上是uintptr对象，保存自身对象地址
	checker copyChecker

}

// NewCond方法传入一个实现了Locker接口的对象，返回一个新的Cond对象指针，
// 保证在多goroutine使用cond的时候，持有的是同一个实例
func NewCond(l Locker) *Cond {
	return &Cond{L: l}
}


// 等待原子解锁c.L并暂停执行调用goroutine。
// 稍后恢复执行后，Wait会在返回之前锁定c.L.
// 与其他系统不同，除非被广播或信号唤醒，否则等待无法返回。

// 因为等待第一次恢复时c.L没有被锁定，
// 所以当Wait返回时，调用者通常不能认为条件为真。
// 相反，调用者应该循环等待：

//    c.L.Lock()
//    for !condition() {
//        c.Wait()
//    }
//    ... make use of condition ...
//    c.L.Unlock()
//

// 调用此方法会将此routine加入通知列表,并等待获取通知,调用此方法必须先Lock,不然方法里会调用Unlock(),报错
//
func (c *Cond) Wait() {

	// 检查是否被复制; 如果是就panic
	// check检查，保证cond在第一次使用后没有被复制
	c.checker.check()
	// 将当前goroutine加入等待队列, 该方法在 runtime 包的 notifyListAdd 函数中实现
	// src/runtime/sema.go
	t := runtime_notifyListAdd(&c.notify)
	// 释放锁,
	// 因此在调用Wait方法前，必须保证获取到了cond的锁，否则会报错
	c.L.Unlock()

	// 等待队列中的所有的goroutine执行等待唤醒操作
	// 将当前goroutine挂起，等待唤醒信号
	// 该方法在 runtime 包的 notifyListWait 函数中实现
	// src/runtime/sema.go
	runtime_notifyListWait(&c.notify, t)
	// 被通知了,获取锁,继续运行
	c.L.Lock()
}




//
// 唤醒单个 等待的 goroutine
func (c *Cond) Signal() {
	// 检查c是否是被复制的，如果是就panic
	// 保证cond在第一次使用后没有被复制
	c.checker.check()
	// 通知等待列表中的一个
	// 顺序唤醒一个等待的gorountine
	// 在runtime 包的 notifyListNotifyOne 函数中被实现
	// src/runtime/sema.go
	runtime_notifyListNotifyOne(&c.notify)
}

// 唤醒等待队列中的所有goroutine。
func (c *Cond) Broadcast() {
	// 检查c是否是被复制的，如果是就panic
	// 保证cond在第一次使用后没有被复制
	c.checker.check()
	// 唤醒等待队列中所有的goroutine
	// 有runtime 包的 notifyListNotifyAll 函数实现
	// src\runtime\sema.go
	runtime_notifyListNotifyAll(&c.notify)
}

// copyChecker保持指向自身的指针以检测对象复制。
type copyChecker uintptr

// 检查c是否被复制，如果是则panic
/**
check方法在第一次调用的时候，会将checker对象地址赋值给checker，也就是将自身内存地址赋值给自身
 */
func (c *copyChecker) check() {
	/**
	因为 copyChecker的底层类型为 uintptr
	那么 这里的 *c其实就是 copyChecker类型本身，然后强转成uintptr
	和拿着 c 也就是copyChecker的指针去求 uintptr，理论上要想等
	即：内存地址为一样，则表示没有被复制
	 */
	 // 下述做法是：
	 // 其实 copyChecker中存储的对象地址就是 copyChecker 对象自身的地址
	 // 先把 copyChecker 处存储的对象地址和自己通过 unsafe.Pointer求出来的对象地址作比较，
	 // 如果发现不相等，那么就尝试的替换，由于使用的 old是0，
	 // 则表示c还没有开辟内存空间，也就是说，只有是首次开辟地址才会替换成功
	 // 如果替换不成功，则表示 copyChecker出所存储的地址和 unsafe计算出来的不一致
	 // 则表示对象是被复制了
	if uintptr(*c) != uintptr(unsafe.Pointer(c)) &&
		!atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c))) &&
		uintptr(*c) != uintptr(unsafe.Pointer(c)) {
		panic("sync.Cond is copied")
	}
}

// noCopy可以嵌入到结构中，在第一次使用后不得复制。
//
// 详细介绍请查看： https://github.com/golang/go/issues/8005#issuecomment-190753527
type noCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.
// Lock 是有 go vet 命令来判断是否有 copy 的检查的
func (*noCopy) Lock() {}



// sync/runtime.go
// 使用链表实现
type notifyList struct {

	wait   uint32		// 等待数
	notify uint32		// 唤醒数
	lock   uintptr		// 信号锁
	// 使用链表实现
	head   unsafe.Pointer	// 队列的当前头
	tail   unsafe.Pointer	// 队列的当前尾
}