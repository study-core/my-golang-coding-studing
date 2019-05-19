package main

import (
	"fmt"
	"myCryto-study/testMychain"
	"myCryto-study/testMybox"
	"github.com/go-errors/errors"
)

func main() {
	defer func() {
		if err := recover(); nil != err {
			msg := fmt.Sprintf("发生panic异常: %v\n", errors.Wrap(err, 2).ErrorStack())
			fmt.Println("panic", fmt.Sprint(err), "\nstack:", msg)
		}
	}()
	my := testMychain.Mychain{}


	fmt.Printf("lala %p", &my)
	fmt.Println(my.GetName())
	my.SetName("I Love Emma !!!!")

	fmt.Println(my.GetName())
	my.SetName("I Love Losi !!!!")
	my.SetName("I Love Kally !!!!")
	fmt.Println(my.GetName())

	m := testMybox.Mybox{}
	fmt.Println(m.GetName())
}




