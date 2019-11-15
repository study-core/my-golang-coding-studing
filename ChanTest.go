package main

import "fmt"

func main() {

	ch1 := make(chan int, 1)

	chr := make(chan int, 1)
	ch1 <- 12

	chr <- <- ch1

	a := <- chr
	fmt.Println(a)

}

func skipN(done <-chan struct{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			// 是否关闭信号
			case <-done:
				return
			// 空消耗掉 前 num 个 元素
			case <-valueStream:
			}
		}
		for {
			select {
			case <-done:
				return
			// 接受后续的 元素  注意: <- <- 的写法
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}