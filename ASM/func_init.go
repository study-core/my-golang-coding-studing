package main

import (
	"fmt"
	"my-study/ASM/pkg"
)

func main() {
	a := 1
	b := 2
	//c, d := pkg.GavinSwap(a, b)
	//fmt.Println("GavinSwap", c, d)
	c, d := pkg.GavinSwap2(a, b)
	fmt.Println("GavinSwap2", c, d)
}
