package main

import (
	"fmt"
)

//func main() {
//
//	var x int
//	threads := runtime.GOMAXPROCS(0)
//	for i := 0; i < threads; i++ {
//		go func() {
//
//			for {
//				x++
//			}
//
//		}()
//
//	}
//	time.Sleep(time.Second)
//	fmt.Println("x =", x)
//
//}

func main() {

	//var x int
	//
	//threads := runtime.GOMAXPROCS(0) - 1
	//
	//for i := 0; i < threads; i++ {
	//	go func() {
	//		for {
	//			x++
	//		}
	//	}()
	//}
	//time.Sleep(time.Second)
	//fmt.Println("x =", x)

	fmt.Println(fA())
	fmt.Println(fB())
	fmt.Println(fC())

}

func fA() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func fB() int {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func fC() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
