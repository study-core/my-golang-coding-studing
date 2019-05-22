package main

import (
	//"fmt"
	//"strings"
	"fmt"
)

func main() {
	//prefix := []byte("im")
	//s := make([]byte, len(prefix))
	//copy(s, prefix)
	//fmt.Println(string(s))
	//fmt.Println(s, prefix)
	//prefix = append(prefix, []byte("sde")...)
	//fmt.Println(string(prefix))
	//fmt.Println(strings.Trim(string(prefix), string(s)))
	//fmt.Println(prefix[len(s):])
	//fmt.Println([]byte("sde"))
	//fmt.Println(prefix)


	arr := make([]int, 3)

	list := []int{12, 5, 74, 8, 9, 6}
	copy(arr, list)
	fmt.Println(arr, list)
}



