package main

// #include <stdlib.h>
/*
#cgo LDFLAGS: -lm
#include <math.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("调用C的随机函数", Random())
	f, err := Sqrt(0.2)
	fmt.Println("float: ", f, "err: ", err)
	Print("哈哈哈")
}

func Random() int {
	return int(C.rand())
}

func Seed(i int) {
	C.srand(C.uint(i))
}

func Sqrt(p float32) (float32, error) {
	n, err := C.sqrt(C.double(p))
	return float32(n), err
}

func Print(s string) {
	cs := C.CString(s)
	// 将 C 中的 free 函数放到 Go的 defer 中，这样在函数结束的时候， 内存就被释放了
	defer C.free(unsafe.Pointer(cs))
	C.myprint(cs)
}
