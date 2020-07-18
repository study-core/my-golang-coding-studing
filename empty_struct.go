package main

import "fmt"

func main() {
	a := new(struct{})
	b := new(struct{})
	println("println: ", a, b, a == b) // 这里会遇到 编译优化, 使用go的 SSA 工具可以看

	c := new(struct{})
	d := new(struct{})
	fmt.Printf("fmt.Println: %p, %p, %v \n", c, d, c == d)
}
