package main 

import (
	"sync"
	"runtime"
	"time"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(4)

	//计数 锁
	var wg sync.WaitGroup

	//计数量 存储chan(需要具备缓冲)
	//可以看做是java的 new semaphore(2)
	semaphore := make(chan struct{}, 2)

	for i := 0; i < 5; i++ {
		wg.Add(1)

		//异步协程,模拟竞态
		go func (id int) {
			defer wg.Done()

		// 可以用和java的信号桶 放取相反的顺序
		// 如: 往信号桶里放一个信号 为获取信号量(java是从信号桶中拿掉一个信号)
		semaphore <- struct{}{}

		//声明 释放 
		defer func () {
			<- semaphore
		}()

		// 操作具体的业务 
		time.Sleep(time.Second * 2)
		fmt.Println("当前是次数:", (id + int(1)), ",当前时间为:", time.Now().Format("2006-01-02 03:04:05"))

		}(i)
	} 

	wg.Wait()
}