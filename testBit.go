package main

import (
	"unsafe"
	"sync/atomic"
	"fmt"
	"strconv"
)

func main() {
	m := &myAs{}
	sadrr := (*uint64)(unsafe.Pointer(&m.state1))
	//fmt.Println(uint64(18446744073709551615)>>32)
	//fmt.Println(uint64(4294967295) << 32)
	statep := m.stateFn()
	fmt.Println(statep, sadrr)
	fmt.Println(*statep)
	fmt.Println(uint64(4294967295)<<32)
	// 这个数有点大哦，转成 uint32 和int32得到的不一样
	state := atomic.AddUint64(statep, uint64(int(-1))<<32)
	by := (m.state1)[:]
	str := string(by)
	if i, err := strconv.Atoi(str); nil != err {
		fmt.Println("err", err)
	}else {
		fmt.Println(i, uint64(i))
	}
	fmt.Println(by, str)
	fmt.Println("[]byte转uint64", )

	fmt.Println(state)
	fmt.Println(m.state1)
	fmt.Println((*uint64)(unsafe.Pointer(&m.state1)))
	fmt.Println((*uint64)(unsafe.Pointer(&m.state1[0])))
	//fmt.Println(m.state1[0])
	// 注意，这个是转了 int32
	//v := int32(state >> 32)
	v := uint32(state >> 32)
	// 注意，这个是转了 uint32
	w := uint32(state)
	fmt.Println(v, w)

	//a := uint64(8)
	//
	//n := atomic.AddUint64(&a, uint64(7))
	//fmt.Println(n) // 15

}

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