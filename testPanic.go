package main

import "fmt"

func main() {

	err := fetchPanic2Err()
	if nil != err {
		fmt.Println(err.Error())
	}
}


// 使用 命名返参 是可以捕获 panic 的
func fetchPanic2Err () (err error) {

	defer func() {
		fmt.Println("1")
		if pnc := recover(); nil != pnc {
			fmt.Println("1 err")
			err = fmt.Errorf("this is panic： %s", fmt.Sprint(pnc))
		}
	}()

	defer func() {
		fmt.Println("2")
		if pnc := recover(); nil != pnc {
			fmt.Println("2 err")
			err = fmt.Errorf("there is panic： %s", fmt.Sprint(pnc))
		}
	}()

	panic("I am Gavin~")
	return nil
}
