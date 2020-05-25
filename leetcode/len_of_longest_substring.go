package main

import "fmt"

// 找出字符串中的 最长无重复子串的长度

func main() {
	str := "abcabcbb"
	fmt.Println("最长无重复子串len:", lengthOfLongestSubstring(str))
}


/**
思路:

利用移动窗口 (可伸缩的)
 */
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}

	n := len(s)
	// rk: 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	// ans: 最长的长度记录值
	rk, ans := -1, 0

	// i 表示 移动窗口的左指针
	for i := 0; i < n; i++ {

		// 能够到这里, 说明下面 伸长右指针的 for已经出现了 重复字符了
		//
		// 移除当前左指针对应的元素, 继续走下面 右指针的for，如果还重复说明当前右移左指针覆盖到重复的元素，需要继续右移
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}

		// 一直遍历 伸长移动窗口的右指针         m[s[rk+1]] == 1 时，说明出现了重复元素, 这时候需要终止 伸长有指针的for， 变而转向 外层右移 左指针的for
		for rk + 1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk - i + 1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}


