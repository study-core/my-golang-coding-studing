package main

import (
	"fmt"
	"my-study/linknameA"
	"my-study/linknameC"
)

func main() {
	fmt.Println(linknameA.Hello())
	fmt.Println(linknameC.IsSpace(' '))
	//fmt.Println(linknameC.Hello()) // 这个会报错
	linknameC.Say()

}
