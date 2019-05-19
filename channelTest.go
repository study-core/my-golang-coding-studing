package main

import (
	"sync/atomic"
	"fmt"
)

func main() {

	d := &dpos{}
	atomic.AddInt64(&d.ticketTotal, int64(10))
	fmt.Println(d.ticketTotal)
	//s := int64(d.count)
	//atomic.AddInt64(&s, int64(10))
	//fmt.Println(d.count)
}


type dpos struct {
	//count int
	ticketTotal		int64
}