package main

import "fmt"

func main() {
	var a = Accumulator()

	fmt.Printf("%d\n", a(1))

	fmt.Printf("%d\n", a(10))

	fmt.Printf("%d\n", a(100))

	fmt.Println("------------------------")

	var b = Accumulator()

	fmt.Printf("%d\n", b(1))

	fmt.Printf("%d\n", b(10))

	fmt.Printf("%d\n", b(100))

	/**
	(0xc000090010, 0) - 1
	(0xc000090010, 1) - 11
	(0xc000090010, 11) - 111
	------------------------
	(0xc000090058, 0) - 1
	(0xc000090058, 1) - 11
	(0xc000090058, 11) - 111
	*/
}

func Accumulator() func(int) int {
	var x int
	return func(delta int) int {
		fmt.Printf("(%+v, %+v) - ", &x, x)
		x += delta
		return x
	}
}
