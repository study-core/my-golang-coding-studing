package main

import (
	"time"
	"fmt"
)

func main() {

	c := make(chan int)

	go func(){
		// TODO
		time.Sleep(7 * time.Second)
		c <- 1
	}()

	select {
	case i := <- c:
		fmt.Println("I am ", i)
	//case  <- c:
	//	fmt.Println("done")

	case <- time.After(6 * time.Second):
		fmt.Println("time out")
	}
}
