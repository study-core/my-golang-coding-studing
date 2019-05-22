package main

import "fmt"

func main() {

	var b typeInter

	b = &typeStruct{}
	fmt.Println("b", b)
	switch r := b.(type) {
	case *typeStruct:
		r = &typeStruct{}
		fmt.Println("r", r)
	default:
		fmt.Println("default")
	}
}


type typeInter interface {
	String() string
}

type typeStruct struct {
}

func (t *typeStruct) String () string {
	return "Gavin"
}