package main

/*
int PlusOne(int n)
{
    return n + 1;
}
*/
import "C"

import (
	"fmt"
)

func main() {
	var n int = 10
	var m int = int(C.PlusOne(C.int(n))) // 类型要转换
	fmt.Println(m)                       // 11

}