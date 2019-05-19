package main

import (
	"fmt"
	"time"
	"io"
	"reflect"
)

//type Money struct{Num int}
//
//func Producer(c chan<- Money) {
//	for i := 0; i < 10; i++ {
//		c <- Money{i}
//	}
//}
//
//func Consumer1(c <-chan Money) {
//	for m := range c {
//		fmt.Printf("oh, I get money: %v\n", m)
//	}
//}
//
//func Consumer2(c <-chan Money) {
//	for m := range c {
//		fmt.Printf("oh, I get money too: %v\n", m)
//	}
//}
//
//func main(){
//	c := make(chan Money, 2)
//
//	go Consumer1(c)
//	go Consumer2(c)
//
//	Producer(c)
//	time.Sleep(time.Second)
//}



//func main() {
//	x := 0
//	c := make(chan struct{}, 1)
//	go func() {
//		for i := 0; i < 10000; i++ {
//			c <- struct{}{}
//			x++
//			<-c
//		}
//	}()
//
//	go func() {
//		for i := 0; i < 10000; i++ {
//			c <- struct{}{}
//			x--
//			<-c
//		}
//	}()
//
//	time.Sleep(5*time.Second)
//	fmt.Println("x should be 0, and x =", x)
//}







//type pipefunc func(in io.Reader, out io.Writer)
//
//func bind(app func(in io.Reader, out io.Writer, args []string), args []string) pipefunc {
//	return func(in io.Reader, out io.Writer) {
//		app(in, out, args)
//	}
//}
//
//func pipe(apps ...pipefunc) pipefunc {
//	if len(apps) == 0 {return nil}
//	if len(apps) == 1 {return apps[0]}
//
//	app := apps[0]
//	for i := 1; i < len(apps); i++ {
//		app1, app2 := app, apps[i]
//
//		app = func(in io.Reader, out io.Writer) {
//			pr, pw := io.Pipe()
//			defer pw.Close()
//			go func() {
//				defer pr.Close()
//				app2(pr, out)
//			}()
//			app1(in, pw)
//		}
//	}
//	return app
//}
//
//func main() {
//	// pipe(bind(app1,args1), bind(app2, args2))
//	// tar, gzip for example
//	// func tar(in io.Reader, out io.Writer, files []string)
//	// func gzip(in io.Reader, out io.Writer)
//
//	pipe(bind(tar, files), gzip)(nil, out)    // 如此优雅
//}






//func or(channels ...<-chan interface{}) <-chan interface{} {
//	switch len(channels) {
//	case 0:
//		return nil
//	case 1:
//		return channels[0]
//	}
//
//	orDone := make(chan interface{})
//	go func() {
//		defer close(orDone)
//		var cases []reflect.SelectCase  // 构造 chan的反射实例
//		for _, c := range channels {
//			cases = append(cases, reflect.SelectCase{
//				Dir:  reflect.SelectRecv,
//				Chan: reflect.ValueOf(c),
//			})
//		}
//
//		index, recv, ok := reflect.Select(cases) // 获取所有通道中，所有有值的第一个
//	}()
//
//	return orDone
//}




//func or(channels ...<-chan interface{}) <-chan interface{} {
//	switch len(channels) {
//	case 0:
//		return nil
//	case 1:
//		return channels[0]
//	}
//
//	orDone := make(chan interface{})
//	go func() {
//		defer close(orDone)
//
//		switch len(channels) {
//		case 2:
//			select {
//			case <-channels[0]:
//			case <-channels[1]:
//			}
//		default:
//			m := len(channels) /2
//			select {
//			case <-or(channels[:m]...):
//			case <-or(channels[m:]...):
//			}
//		}
//	}()
//
//	return orDone
//}