package main

import "fmt"

func main() {
	str := "aaddxfbdfadadasfhhhhhwewewhhhhaasfas"
	res := longestPalindrome(str)
	fmt.Println(res) // hhhhwewewhhhh

}



// todo 中心扩展算法
//
// 时间复杂度：O(n²). 两层循环, 每层循环都是遍历每个字符.
//
// 空间复杂度：O(1)
func longestPalindrome(s string) string {
	str := []rune(s)
	if "" == s || len(str) < 1 {return ""}
	start := 0
	end := 0

	maxFn := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	// 由于存在奇数的字符串和偶数的字符串,
	// 所以我们需要从一个字符开始扩展,
	// 或者从两个字符之间开始扩展, 所以总共有 n + n - 1 个中心.
	//
	//  遍历每个中心, 然后判断对称位置是否相等
	//

	for i := 0; i < len(str); i++ {
		len1 := expandAroundCenter(s, i, i) 	//从一个字符扩展
		len2 := expandAroundCenter(s, i, i + 1) //从两个字符之间扩展

		// 为什么 取最长值, 如：
		//
		// 对于类似 aaaa 时,  当下标为 1 时,
		//  a   a   a   a
		//  0   1   2   3
		// 我们有
		//  a a a  (0 1 2)		和
		//  a a a a (0 1 2 3) 	两种 回文子串, 这时最长的为 第二种
		//
		// 对于类似 abaa 时, 当下标为 1 时,
		//	a  b  a  a
		//  0  1  2  3
		// 我们有
		//  a  b  a   (0 1 2)   和
		//  无 					两种  回文子串,  这时 最长的为 第一种
		//

		len := maxFn(len1, len2)
		//根据 i 和 len 求得字符串的相应下标
		if len > end - start {
			start = i - (len - 1) / 2
			end = i + len / 2
		}
	}


	return string(str[start: end + 1])
}

func expandAroundCenter (s string, left, right int) int {
 	L := left
 	R := right

 	str := []rune(s)

	for L >= 0 && R < len(str) && str[L] == str[R] {
		L--
		R++
	}
	return R - L - 1
}