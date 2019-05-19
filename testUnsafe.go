package main

import (
	"unsafe"
	"fmt"
	"reflect"
)

// 结构体的成员在内存中的分配是一段连续的内存，结构体中 第一个成员 的地址就是这个结构体的地址
type Person2 struct {
	Name   string
	Height    int
}

//func main() {
//	a := Person2{"Jasper", 27}
//	pa := unsafe.Pointer(&a)
//	aname := *(*string)(unsafe.Pointer(uintptr(pa) + unsafe.Offsetof(a.Name))) // pname := (*string)(unsafe.Pointer(uintptr(pa))) 这样也是可以的
//	fmt.Println(aname)
//	pname := *(*string)(unsafe.Pointer(uintptr(pa)))
//	fmt.Println(pname)
//	aage := (*int)(unsafe.Pointer(uintptr(pa) + unsafe.Offsetof(a.Height)))
//  	//page := *(*int)(unsafe.Pointer(uintptr(pa)))
//  	fmt.Println(page)
//	aname = "Jasper2"
//	*aage = 28
//	fmt.Println(a) // {Jasper2 28}
//}

func main() {
	var v *V = new(V)
	//var i *int32 = (*int32)(unsafe.Pointer(v))
	//
	//*i = int32(1)
	//fmt.Println(*i)

	//var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + unsafe.Sizeof(int32(0))))  // 不是J
	var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + unsafe.Offsetof(v.j)))		// 是 J
	fmt.Println(*j)
	*j = int64(2)
	fmt.Println(v.j)
	v.PrintI()
	v.PrintJ()
}



type V struct {
	i int32
	j int64
}

func (this V) PrintI() {
	fmt.Printf("i=%d\n", this.i)
}

func (this V) PrintJ() {
	fmt.Printf("j=%d\n", this.j)
}


// 注意，这样虽然可以实现，但强烈推荐不要使用这种方法来转换类型，因为这样会导致修改转化过后的值会影响之前的变量
func string2byte(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func byte2string(b []byte) string{
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}