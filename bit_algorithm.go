package main

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
)

func main() {

	// 字母大小写转换
	//
	// golang的字符称为rune，等价于C中的char，可直接与整数转换
	// rune实际是整型，必需先将其转换为string才能打印出来，否则打印出来的是一个整数
	/*a := 'a'
	A := 'A'
	a ^= 32
	fmt.Println(string(a)) // A
	A ^= 32
	fmt.Println(string(A)) // a

	for _, char := range []rune("世界你好,hello world ！啦我") {
		fmt.Println(string(char))
	}

	c:='a'
	fmt.Println(c)
	fmt.Println(string(c))
	fmt.Println(string(97))

	cleanLastBit1(12) // 8
	cleanLastBit1(0)

	fmt.Println(findOnlyOne([]int{1,2,2,1,3,4,3}))

	fmt.Println(findNum([]int{1,1,2,3,3,3,2,2,4,1}))


	fmt.Println(abs(-45))*/

	a := 45
	b := 28

	fmt.Println(sum(a, b))

}

// 去除 x 的最后一个1 (不知道1在哪个 bit 上，只知道需要去除最后一个)
func cleanLastBit1(x uint32) uint32 {
	xb := common.Uint32ToBytes(x)
	fmt.Println(xb)

	// 1^2^3^4^5^1^2^3^4 = （1^1)^(2^2)^(3^3)^(4^4)^5= 0^0^0^0^5 = 5。

	/**
	x & (x-1)
	x = 1100
	x-1 = 1011
	x & (x-1) = 1000

	x:  1010100
	x-1: 1010001
	x&(x-1) == 1010000
	 */

	y := x&(x-1)
	yb := common.Uint32ToBytes(y)
	fmt.Println(yb)
	return y
}

// 除了一个数其他的都出现2次，找出这个数
func findOnlyOne (arr []int) int {

	var ans int

	for i := 0; i < 7; i++ {
		ans ^= arr[i]
	}

	return ans
}

// 除了一个数其他数都出现3次，找出这个数
func findNum (arr []int) int {
	var one int
	var two int

	for i := 0; i < len(arr); i ++ {
		two |= one&arr[i]
		one ^= arr[i]
		three := two&one

		one ^= three
		two ^= three


	}
	return one|two
}

// 取绝对值
func abs(a int) int {
	i := a >> 31
	// go 中不支持 ~ 符号， 取反就是 ^a
	if i == 0 { return a } else { return ^a + 1 }
}

// 不用`+`求和
func sum (a, b int) int {

	// a: 0010 == 2
	// b: 1010 == 10
	// a^b: 1000
	// a&b: 0010
	// a&b<<1: 0100
	// 1000+0100 == 1100 == 12
	// a + b == a^b + (a&b<<1)
	for b != 0 {
		// 使用for逐个追加 进位1
		a0 := a^b
		b0 := a&b<<1
		a = a0
		b = b0
	}
	return a
}