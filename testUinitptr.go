package main

import (
	"unsafe"
	"fmt"
	"sync/atomic"
)

func main() {
	c := Human{
		Name: "A",
	}
	fmt.Println("第一次使用")
	c.C.checkFunc(1)
	fmt.Println("第二次使用,但没有被复制")
	c.C.checkFunc(2)

	c1 := c
	fmt.Println("第三次使用 ... ")
	c1.C.checkFunc(3)

}

type check uintptr

func (c *check) checkFunc(i int) {
	fmt.Println(uintptr(*c))
	fmt.Println(uintptr(unsafe.Pointer(c)))
	if uintptr(*c) != uintptr(unsafe.Pointer(c)) &&
		!atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c))) &&
		uintptr(*c) != uintptr(unsafe.Pointer(c)) {
		fmt.Println("第" + fmt.Sprint(i) + "次使用, obj is copied")
	}
}

type Human struct {
	C		check
	Name 	string
}

