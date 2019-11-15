package main

import (
	//"math/rand"
	"sync/atomic"
	"unsafe"
	"fmt"
)

func main() {
	m := &myAs{}
	sadrr := (*uint64)(unsafe.Pointer(&m.state1))
	//fmt.Println(uint64(18446744073709551615)>>32)
	//fmt.Println(uint64(4294967295) << 32)
	statep := m.stateFn()
	fmt.Println(statep, sadrr)
	fmt.Println(*statep)

	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))
	//fmt.Println(rand.Intn(1))

	done := int32(45)

	index := 0
	for i := 0; i < 46; i++ {
		atomic.AddInt32(&done, -1)
		index ++
	}

	fmt.Println("遍历了", fmt.Sprint(index), "次", "done 的值为 ", fmt.Sprint(done))





	//fmt.Println(uint64(4294967295)<<32)
	//// 这个数有点大哦，转成 uint32 和int32得到的不一样
	//state := atomic.AddUint64(statep, uint64(int(-1))<<32)
	//by := (m.state1)[:]
	//str := string(by)
	//if i, err := strconv.Atoi(str); nil != err {
	//	fmt.Println("err", err)
	//}else {
	//	fmt.Println(i, uint64(i))
	//}
	//fmt.Println(by, str)
	//fmt.Println("[]byte转uint64", )
	//
	//fmt.Println(state)
	//fmt.Println(m.state1)
	//fmt.Println((*uint64)(unsafe.Pointer(&m.state1)))
	//fmt.Println((*uint64)(unsafe.Pointer(&m.state1[0])))
	////fmt.Println(m.state1[0])
	//// 注意，这个是转了 int32
	////v := int32(state >> 32)
	//v := uint32(state >> 32)
	//// 注意，这个是转了 uint32
	//w := uint32(state)
	//fmt.Println(v, w)
	//
	////a := uint64(8)
	////
	////n := atomic.AddUint64(&a, uint64(7))
	////fmt.Println(n) // 15

}


/**
自动识别系统是多少位
*/
func  (m *myAs)stateFn() *uint64 {
	// 是否是 64位机器：因为64位机器站高8位
	if uintptr(unsafe.Pointer(&m.state1))%8 == 0 {
		fmt.Println("这是64位电脑")
		return (*uint64)(unsafe.Pointer(&m.state1))
	} else {
		// 如果是 32位机器， 则首元素从第 4 位开始：因为32 位机器占中间4位
		fmt.Println("这是32位电脑")
		return (*uint64)(unsafe.Pointer(&m.state1[4]))
	}
}


type myAs struct {
	state1 [12]byte
}

/*package main


import "fmt"

func  AppendSlice(list []int, i int) bool {
	if len(list) > 10 {
		return true
	} else {
		list = append(list, i)
		return AppendSlice(list, i+1)
	}
}

func main() {
	//var list = make([]int, 5)
	//AppendSlice(list, 1)
	//fmt.Println(list)


	i := 1
	defer fmt.Println(i)
	i++
	fmt.Println(i)
	i++
	//panic("exit")
	defer fmt.Println(i)
}*/

