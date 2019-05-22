package main

import (
	"fmt"
	//"sync"
	//"sync/atomic"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"time"
)

func main() {

	//am := make(map[string]int, 1000)
	//
	//bm := make(map[string]int, 0)
	//
	//fmt.Println("am", len(am), am)
	//fmt.Println("bm", len(bm), bm)
	//
	//am["a"] = 1
	//bm["b"] = 2
	//
	//fmt.Println("am", len(am), am)
	//fmt.Println("bm", len(bm), bm)

	/*var flag  int = 0

	var wg sync.WaitGroup
	wg.Add(6)
	go func() {
		fmt.Println("我是1")

		flag |= 0
		wg.Done()
	}()
	go func() {
		fmt.Println("我是2")
		flag |= 0
		wg.Done()
	}()
	go func() {
		fmt.Println("我是3")
		flag |= 0
		wg.Done()
	}()
	go func() {
		fmt.Println("我是4")
		flag |= 0
		wg.Done()
	}()
	go func() {
		fmt.Println("我是5")
		flag |= 0
		wg.Done()
	}()
	go func() {
		fmt.Println("我是6")
		flag |= 0
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(flag)

	fmt.Println(len(ReturnMap()))*/




	//is := []int{1}
	//
	//fmt.Println(is, len(is))
	//for i, m := range is {
	//	if m == 1 {
	//		is = append(is[:i], is[i+1:] ...)
	//	}
	//}
	//
	//fmt.Println(is, len(is))

	nextIdArr := []int{1,5,7,8,10,12,16}

	curr_queue := []int{2, 50, 1, 8, 45, 61, 12, 5, 10, 47, 7, 31, 16}

	nextQueue := make([]int, len(nextIdArr))

	retry:
	for i, canId := range nextIdArr {
		for k := 0; k < len(curr_queue); k++ {
			id := curr_queue[k]
			if canId == id {

				nextQueue[i] = id
				curr_queue = append(curr_queue[:k], curr_queue[k+1:]...)
				goto retry
			}
		}
	}

	fmt.Println(curr_queue)

	fmt.Println(nextIdArr)
	fmt.Println(nextQueue)

	start := common.NewTimer()
	start.Begin()
	time.Sleep(time.Second)
	fmt.Println("等了", fmt.Sprintf("%v ms", start.End()))
}


func ReturnMap() map[string]int {
	return nil
}