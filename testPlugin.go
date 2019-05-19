//package main
//
//import "fmt"
//
//func main() {
//	//var a myIn
//	var b myAlias
//
//	c := int(1)
//	//a = c
//	b = c
//	fmt.Println(b)
//	//fmt.Println(c)
//}
//
//
//type myIn  int
//
//type myAlias = int


package main

type MyInt = int

func (i *MyInt) Increase(a int) { // Error: cannot define new methods on non-local type int
	*i = *i + MyInt(a)
}

func main() {
	var mi MyInt = 6
	mi.Increase(5)
}
