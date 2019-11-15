package main

/*
#include <stdlib.h>

// 这个是 C 中的开辟内存的方法啊
void* makeslice(size_t memsize) {
    return malloc(memsize);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

/**
Go 中访问 C的内存
*/

func makeByteSlize(n int) []byte {

	// 调用 C 的函数开辟内存空间
	p := C.makeslice(C.size_t(n))

	a := (*[1<<32 + 1]byte)(p)

	// 2147483648
	// 4294967297
	fmt.Println("len p: ", unsafe.Sizeof(p))
	fmt.Println("Size a: ", unsafe.Sizeof(a))
	fmt.Println("len a: ", len(a))
	fmt.Println("n: ", n)
	fmt.Println("n/2: ", n/2)
	return (a)[0:n:n]
}

func freeByteSlice(p []byte) {

	// 调用C的函数，释放内存
	C.free(unsafe.Pointer(&p[0]))
}

func main() {

	sli := make([]byte, 1<<32+1)
	fmt.Println(len(sli))

	s := makeByteSlize(1<<32 + 1)
	s[len(s)-1] = 255
	fmt.Println("切片的大小:", len(s), "末尾元素: ", s[len(s)-1])
	fmt.Printf("大小: %d G", len(s)/(1024*1024*1024))
	freeByteSlice(s)
}
