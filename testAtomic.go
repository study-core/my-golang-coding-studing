package main

import (
	"sync/atomic"
	"fmt"
	//"math"
	"sync"
)

func main() {

	var myInt32 int32 = 100
	fmt.Println(myInt32)
	newInt := atomic.AddInt32(&myInt32, -32)
	fmt.Println(newInt)

	// Uint 原子减
	var myUint32 uint32 = 100
	fmt.Println(myUint32)
	newUint := atomic.AddUint32(&myUint32, ^uint32(-(-32) - 1))  // 官方做法
	//newUint := atomic.AddUint32(&myUint32, uint32(int32(-32))) // 不行
	/*
	// 非官方做法
	nn := -32 // 或者 nn := int32(-32)
	newUint := atomic.AddUint32(&myUint32, uint32(nn)) // 只对变量有效
	const NN  = -32
	newUint := atomic.AddUint32(&myUint32, NN&math.MaxUint32) // 只对常量有效
	*/
	fmt.Println(newUint)

	var myUint64 uint64 = 100
	fmt.Println(myUint64)
	fmt.Println(atomic.CompareAndSwapUint64(&myUint64, 101, 102)) // old 与原先值相等才会被 new替换
	fmt.Println(myUint64)

	//var m sync.Mutex

}
