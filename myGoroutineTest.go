package main 

import (
	"fmt"
	"sync"
)

func main() {
	testGoroutine()
}


func testGoroutine() {
	myChan := make(chan uint32)
	var wg sync.WaitGroup

	mySlice := []uint32{12, 13, 47, 58, 78, 98}

	go func() {
		for {
			member, ok := <- myChan
			if !ok {
				break;
			}
			fmt.Println("这是元素:", member)
		}
	}()

	for _, value := range mySlice{
		wg.Add(1)
		go func (v uint32) {
			defer  wg.Done()
			myChan <- v
		}(value)
	}

	wg.Wait()
	close(myChan)

	
}