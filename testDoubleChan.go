package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ce := &ChanEvent{}
	go ce.Send()
	go ce.Consumed()
	time.Sleep(4*time.Second)
}


type ChanEvent struct {
	once 		sync.Once
	chs 		chan chan int
}


func (c *ChanEvent) Send() {
	c.once.Do(func() {
		c.chs = make(chan chan int, 3)
	})
	for i := 1; i <= 3; i++ {
		intCh := make(chan int)
		c.chs <- intCh
		a := <- intCh
		fmt.Println("Send func recv", a)
	}
	close(c.chs)
}


func (c *ChanEvent)Consumed() {
	rd := rand.New(rand.NewSource(1))
	for ch := range c.chs {
		a := rd.Int()
		fmt.Println("Consumed func send", a)
		ch <- a
		close(ch)
	}
}