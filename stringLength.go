package main

import "fmt"

func main() {
	str := "0000000000000000000000000000000000000000000000000000000000000F01"
	fmt.Println(len(str)) // 64 个char (一个char = 4bit； 2个char == 8bit == 1个字节)  64/2 == 32 byte
}
