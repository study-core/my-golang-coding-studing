package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 12
	var ei interface{} = a

	var s MyStruct
	var ii MyInterface =fmt.Println(ei, ii)

}

type MyInterface interface {
	Print()
}

type MyStruct struct{}
func (ms MyStruct) Print() {}






// 需要和这些东东同步 ../cmd/link/internal/ld/decodesym.go:/^func.commonsize,
// ../cmd/compile/internal/gc/reflect.go:/^func.dcommontype and
// ../reflect/type.go:/^type.rtype.
type _type struct {
	size       uintptr	// 	表示类型的宽度/长度
	ptrdata    uintptr 	// 	包含所有指针的内存前缀的大小
	hash       uint32	//	Hash类型; 避免在哈希表中计算
	tflag      tflag	// 	额外类型信息标志
	align      uint8	//	该类型的变量对齐方式
	fieldalign uint8	//	该类型的结构字段对齐方式
	kind       uint8	//	C的枚举 (干嘛用的?)
	alg        *typeAlg	//	算法表
	//	gcdata存储垃圾收集器的GC类型数据
	//	如果 KindGCProg 位(用不同的bit来标识动作含义)被设置在kind字段中，则gcdata是GC程序。
	//	否则它是一个ptrmask 位图。 有关详细信息，请参阅 mbitmap.go。
	gcdata    *byte		// gc 的数据
	str       nameOff	// 字符串形式
	ptrToThis typeOff	// 指向此类型的指针的类型可以为零
}



// typeAlg is 总是 在 reflect/type.go 中 copy或使用.
// 并保持他们同步.
type typeAlg struct {
	// 算出该类型的Hash
	// (ptr to object, seed) -> hash
	hash func(unsafe.Pointer, uintptr) uintptr
	// 比较该类型对象
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
}

type nameOff int32
type typeOff int32