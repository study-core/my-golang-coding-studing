package main

import (
	"fmt"
	"time"
)

//func main() {
//
//	ctx, _ := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
//
//	go func(ctx context.Context) {
//		<-ctx.Done()
//		fmt.Println("一号协程接收到信号...")
//	}(ctx)
//
//	go func(ctx context.Context) {
//		<-ctx.Done()
//		fmt.Println("二号协程接收到信号...")
//	}(ctx)
//
//	go func(ctx context.Context) {
//		<-ctx.Done()
//		fmt.Println("二号协程接收到信号...")
//	}(ctx)
//
//	go func(ctx context.Context) {
//		<-ctx.Done()
//		fmt.Println("三号协程接收到信号...")
//	}(ctx)
//
//	go func(ctx context.Context) {
//		<-ctx.Done()
//		fmt.Println("四号协程接收到信号...")
//	}(ctx)
//
//	go func(ctx context.Context) {
//		<-ctx.Done()
//		fmt.Println("五号协程接收到信号...")
//	}(ctx)
//
//	go func(ctx context.Context) {
//		<-ctx.Done()
//		fmt.Println("六号协程接收到信号...")
//	}(ctx)
//
//	time.Sleep(6*time.Second)
//}


func main() {

	ch := make(chan struct{})

	go func(ch chan struct{}) {
		<-ch
		fmt.Println("一号协程接收到信号...")
	}(ch)

	go func(ch chan struct{}) {
		<-ch
		fmt.Println("二号协程接收到信号...")
	}(ch)

	go func(ch chan struct{}) {
		<-ch
		fmt.Println("三号协程接收到信号...")
	}(ch)

	go func(ch chan struct{}) {
		<-ch
		fmt.Println("四号协程接收到信号...")
	}(ch)

	go func(ch chan struct{}) {
		<-ch
		fmt.Println("五号协程接收到信号...")
	}(ch)

	go func(ch chan struct{}) {
		<-ch
		fmt.Println("六号协程接收到信号...")
	}(ch)


	timer := time.NewTimer(2*time.Second)
	select {
	case  <- timer.C:
		ch <- struct{}{}
	}

	time.Sleep(7*time.Second)
}
