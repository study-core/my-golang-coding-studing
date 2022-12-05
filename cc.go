package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	var a uint64 = 56

	fmt.Println("old", a)
	atomic.CompareAndSwapUint64(&a, 56, 72)
	fmt.Println("new", a)

	atomic.AddUint64(&a, 1)

	fmt.Println("final", a)

	fmt.Println("whether change", atomic.CompareAndSwapUint64(&a, 56, 99), "after", a)



	m := make(map[string]*GavinClient, 0)

	v, ok := m["gavin"]

	fmt.Println(ok && v.IsConnect())
}


type GavinClient struct{
	Flag   bool
}

func (g *GavinClient) IsConnect() bool {
	fmt.Println("进来了啊")
	return g.Flag
}