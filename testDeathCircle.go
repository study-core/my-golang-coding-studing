package main

import (
	"runtime"
	"time"
)

func main() {

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {

				}
			}
		}()
	}

	for {
		time.Sleep(time.Second)
	}
}
