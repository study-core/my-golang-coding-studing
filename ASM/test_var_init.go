package main

import (
	"fmt"
	"my-study/ASM/pkg"
	"unsafe"
)

func main() {

	// 这里调用在汇编中初始化，但是在go中定义的变量 `Id`
	fmt.Println(pkg.Id)
	fmt.Println(pkg.Name)
	fmt.Println(pkg.Name2)
	fmt.Println(pkg.Int32Value)
	fmt.Println(pkg.Uint32Value)
	fmt.Println(pkg.Helloworld)
	fmt.Println(pkg.HelloGavin)
	fmt.Println("slice len", len(pkg.HelloGavin))
	fmt.Println("slice size", unsafe.Sizeof(pkg.HelloGavin))
	fmt.Println(pkg.GavinCh)
	fmt.Println(pkg.GavinMap)
	fmt.Println("测试只读变量", pkg.ConstId)
	//defer func() {
	//	if err := recover(); nil != err {
	//		fmt.Println("捕获异常", err)
	//	}
	//}()
	//pkg.ConstId = 2 // 这里会抛出链 recover 都捕获不了的 异常
	//fmt.Println("测试只读变量,被修改后", pkg.ConstId)
}
