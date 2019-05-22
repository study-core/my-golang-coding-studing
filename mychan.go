package main 

import (
	"sync"
)

//定义接收 chan的外层包装struct
type receiver struct{

	//用匿名字段(只写 类型)
	sync.WaitGroup

	data chan int
}


//自定义接收函数 (返回一个 接收chan的指针)
func newReceiver() *receiver {
	//初始化一个 chan
	r := &receiver{
		data: make(chan int),
	}

	//自增 计数量
	r.Add(1)

	//使用协程异步接收 chan的消息

	go func () {

		//释放 计数量
		defer r.Done()

		//遍历 chan的消息
		for x := range r.data{
			println("receiv:", x)
		}
	}()

	//返回 chan (不把r.Wait()写到函数中就是 可以让goroutine 执行前先把 chan返回出去,让调用方自己去 阻塞 代码块)
	return r
}


func main() {
	
	//先 获取 chan的外层(工厂实例)
	mychan := newReceiver()

	//写入消息
	mychan.data <- 100
	mychan.data <- 260

	//关闭 chan 
	close(mychan.data)

	//阻塞 代码块
	mychan.Wait()

}