package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go http.ListenAndServe(":8080", nil)
	ctx, _ := context.WithTimeout(context.Background(), (10 * time.Second))
	go testA(ctx)
	select {}
}

func testA(ctx context.Context) {
	ctxA, _ := context.WithTimeout(ctx, (5 * time.Second))
	ch := make(chan int)
	go testB(ctxA, ch)

	select {
	case <-ctx.Done():
		fmt.Println("testA Done")
		return
	case i := <-ch:
		fmt.Println(i)
	}

}

func testB(ctx context.Context, ch chan int) {
	//模拟读取数据
	sumCh := make(chan int)
	go func(sumCh chan int) {
		sum := 10
		time.Sleep(10 * time.Second)
		sumCh <- sum
	}(sumCh)

	select {
	case <-ctx.Done():
		fmt.Println("testB Done")
		<-sumCh
		return
		//case ch  <- <-sumCh: 注意这样会导致资源泄露
	case i := <-sumCh:
		fmt.Println("send", i)
		ch <- i
	}

}
