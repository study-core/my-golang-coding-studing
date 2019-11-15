package linknameB

import (
	"fmt"
	_ "unsafe"
)

// 这里需要写完整的函数路径
// 导出用法
//go:linkname hello my-study/linknameA.Hello
func hello() string {
	return "hello world"
}

func say() {
	fmt.Println("I am b say~")
}
