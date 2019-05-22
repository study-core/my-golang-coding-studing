package main

import "fmt"

//func main() {
//	amap := make(map[string]string, 0)
//	amap["a"] = "A"
//	amap["b"] = "B"
//
//	for v := range amap {
//		fmt.Println(v)
//	}
//}

func main() {

	aCh := make(chan string, 0)
	ch := make(chan struct{}, 0)
	go func() {
		for _, v := range []string{"a", "b"} {
			aCh <- v
		}
		close(aCh)
	}()

	go func() {
		i := 1
		//for  s := range aCh {
		//fmt.Println("循环第" + fmt.Sprint(i) + "次~" + s)
		for  range aCh {
			fmt.Println("循环第" + fmt.Sprint(i) + "次~")
			i ++
		}
		close(ch)
	}()
	<- ch
}

//func main() {
//
//
//	aCh := make(chan string, 2)
//	for _, v := range []string{"a", "b"} {
//		aCh <- v
//	}
//	close(aCh)
//	for s := range aCh {
//		fmt.Println(s)
//	}
//}