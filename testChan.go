package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go sendChanFunc(c)
	go func() {
		v := <-c
		fmt.Println("I am", v)
		wg.Done()
	}()
	wg.Wait()
}
func sendChanFunc(send chan<- int) {
	send <- 123
}

func receiveChanFunc()  <- chan int {
	c := make(chan int)
	c <- 54
	return c
}