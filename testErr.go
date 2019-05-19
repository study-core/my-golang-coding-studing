package main

import (
	"errors"
	"fmt"
)

func main() {
	text1, _ := Text()
	fmt.Println(text1)


	text2, err := Text()
	fmt.Println(text2, err)
}

func Text() (string, error) {
	return "a", errors.New("这个是个错误")
}