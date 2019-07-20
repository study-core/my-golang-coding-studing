package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	v2 := "tangs"
	fmt.Println("v2's value is : ", v2, ", type is : ", reflect.TypeOf(v2))
	ty := reflect.ValueOf(&v2).Elem()


	s := ty.Interface()
	s = "tangs"
	fmt.Println("s's value is : ", s, ", type is : " ,reflect.TypeOf(s))



	typ := reflect.TypeOf(main)
	name1 := typ.Name()
	fmt.Println("Name of function" + name1)

	name2 := runtime.FuncForPC(reflect.ValueOf(main).Pointer()).Name()
	fmt.Println("Name of function : " + name2)


	pc, _, _, _ := runtime.Caller(0)
	fmt.Println("Name of function: " + runtime.FuncForPC(pc).Name())
	fmt.Println()

	// or, define a function for it
	fmt.Println("Name of function: " + funcName())
	x()

	fmt.Println("lala:", reflect.TypeOf(stringStr).In(0).)
}

func funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func x() {
	fmt.Println("Name of function: " + funcName())
	y()
}

func y() {
	fmt.Println("Name of function: " + funcName())
	z()
}
func z() {
	fmt.Println("Name of function: " + funcName())
}

func stringStr(ss string) string {
	return ss
}