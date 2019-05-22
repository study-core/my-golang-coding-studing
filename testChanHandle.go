package main

import (
	"fmt"
	"sync"
)

func main() {

	/*ch := make(chan string, 10)

	//for _, v := range []string{"a", "b", "c", "d"} {
	//	ch <- v
	//}
	close(ch)
	i := 0
	for v := range ch {
		i++
		fmt.Println("第", i, "次循环,读出来的为:", v)
	}*/

	ch := make(chan string, 10)
	arr := []string{"a", "b", "c", "d"}
	var wg sync.WaitGroup
	wg.Add(len(arr))
	for _, v := range  arr {
		go func() {
			defer wg.Done()
			ch <- v
		}()
	}
	var v string
	for i:= 0; i < 10; i++ {
		if v = <- ch; "" != v {
			fmt.Println("第一次打印出来的值为:", v)
			break
		}
	}
	wg.Wait()

	ch = make(chan string, 10)
	slice := []string{"g", "v", "e", "f"}
	wg.Add(len(slice))
	for _, v := range  slice {
		go func() {
			defer wg.Done()
			ch <- v
		}()
	}

	for i:= 0; i < 10; i++ {
		if v = <- ch; "" != v {
			fmt.Println("第er次打印出来的值为:", v)
			break
		}
	}
	wg.Wait()
}
