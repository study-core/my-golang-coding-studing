package pkg

import (
	_ "unsafe"
)

//go:linkname println runtime·printnl
func println()

//go:linkname printInt runtime·printint
func printInt(a int)

func GavinMain()

/**
请使用 汇编实现下列内容的方法体

 	var a = 10
    println(a)

    var b = (a+a)*a
    println(b)

我们先分解成下面的形式，在讲解 汇编怎么实现出来
下面的形式就有点像汇编的思路了

	var a, b int

	a = 10
	runtime.printint(a)
	runtime.printnl()

	b = a
	b += b
	b *= a
	runtime.printint(b)
	runtime.printnl()

*/
