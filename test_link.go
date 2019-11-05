package main

/*import (
	"fmt"
	_ "my-study/link/link"
	_ "unsafe"
)

func main() {
	fmt.Println(isSpace(' '))
	fmt.Println(isSpace2(' '))
	//pri("s")
}

//go:linkname isSpace fmt.isSpace
func isSpace(r rune) bool

//go:linkname isSpace2 fmt.isSpace
func isSpace2(r rune) bool

////go:linkname pri builtin.println
//func pri(args string)

//func pri() {
//	link.Pri()
//}*/

import (
	"fmt"
	_ "unsafe"
)

//var aaa = 1
//
////go:linkname a2 main.aaa
//var a2 int
//
////go:linkname a3 main.aaa
//var a3 int

func main() {
	//fmt.Println(aaa, a2, a3)
	//linkTest1()
	//linkTest2()
	linkTest3()
}

//go:linkname linkTarget main.linkTest3
func linkTarget() {
	fmt.Println("I am target~")
}

////go:linkname linkTest1 main.linkTarget
//func linkTest1()
//
////go:linkname linkTest2 main.linkTarget
//func linkTest2()

func linkTest3()
