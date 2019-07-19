package main

import "fmt"

const (
	a = iota // 0
	b		 // 1
	c		 // 2

	d = iota + 1 	// 3 + 1
	e				// 4 + 1
	f				// 5 + 1

)

const (
	gg = iota // 0
	hh //1

)

func main() {
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(gg, hh)
}