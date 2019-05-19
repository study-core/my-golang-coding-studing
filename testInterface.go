package main

import "fmt"

func main() {
	b := BI{}
	var a1 A = &b
	fmt.Println(a1)
}

type A interface {
	Sum()
}

type B interface {
	Sum()
	Add()
}


type  AI struct {
}
func (a *AI) Sum()  {
}


type BI struct {
}
func(b *BI)Sum(){
}

func(b *BI) Add(){
}