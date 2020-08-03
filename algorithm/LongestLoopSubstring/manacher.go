package main

import "fmt"

func main() {

	str := "aaddxfbdfadadasfhhhhhwewewhhhhaasfas"
	res := manacher(str)
	fmt.Println(res)
}

// todo 马拉车算法
//
// 思路
//
// todo 步骤一, 先给字符全加上  特殊符号 及 首尾, 使  奇数串 和  偶数串  都变为 奇数
//
// 如:   abadab   =>  ^#a#b#a#d#a#b#$    (偶数 => 奇数)
// 		 abada    =>  ^#a#b#a#d#a#$      (奇数 => 奇数)
//
//  todo  这样我们就永远只在处理 奇数串 了

// todo  步骤二, 使用一个 和加了特殊字符之后的 str长度一样的 数组 p = []int
//
//		p[i] == i 为中心 的 回文子串的 `半径`
//
//	如:
//	i       0 1 2 3 4 5 6 7 8 9 10 11 12 13 14
//	arr[i]  ^ # c # a # b # b # a  #  f  #   $
//	p[i]      1 2 1 2 1 2 5 2 1 2  1  2  1
//
//   p[7] = 5, 最长回文子串长度  abba == 5-1 == 4 == p[7] - 1
//
// todo  (计算最长回文子串长度)
// 		最长回文半径 和 最长回文子串长度之间的关系：     int maxLength = p[i]-1     半径 - 1 = 回文子长度

// todo (计算最长回文子串起始索引,  【指的是 在没经过处理之前的 数组中的 下标哦】)
//
// 最长回文子串的长度，我们还需要知道它的起始索引值，这样才能截取出完整的最长回文子串
//
//   j         0 1 2 3 4 5
//  arr[j]     c a b b a f
//
// todo  最长回文子串的起始索引：
// 		int index = (i - p[i])/2       j<原数组下标> = (i<处理之后的数组下标> - p[i])/2      j = (7- p[7])/2 = (7-5)/2 == 1

//
// todo  初始化  var p []int
//
// i       0 1 2 3 4 5 6 7 8 9 10 11 12 13 14
// arr[i]  ^ # c # a # b # b # a  #  f  #  $
// p[i]      1 2 1 2 1 2 5 2 1 2  1  2  1

// 设置两个变量 `id` 和 `mx`,
// `id`:  是所有回文子串中, 能延伸到最右端位置的那个回文子串的中心点位置, mx是该回文串能延伸到的最右端的位置.
//
//	todo
//		当i等于7时, id等于7, p[id] = 5, 在以位置7为中心的回文子串中, 该回文子串的右边界是位置12       f的位置.
//		当i等于12时, id等于12, p[id] = 2, 在以位置12为中心的回文子串中, 该回文子串的右边界是位置14    $的位置.
//
//  由此我们可以得出回文子串右边界和其半径之间的关系：mx = p[id]+id


func  manacher(s string)  string {


	str := []rune(s)

	if len(str) < 2 {
		return s
	}

	// todo 第一步：预处理，将原字符串转换为新字符串
	tmpStr := "^"
	for  i := 0; i < len(str); i++ {
		tmpStr += "#" + string(str[i])
	}
	// 尾部再加上字符$，变为奇数长度字符串
	tmpStr += "#$"

	fmt.Println("添加特殊字符处理之后的 str:", tmpStr)

	tmpRune := []rune(tmpStr)

	// todo 第二步：计算数组p、起始索引、最长回文半径
	n := len(tmpRune)

	// p数组
	p := make([]int, n)

	var id, mx int

	// todo 最长回文子串的长度
	maxLength := -1

	// todo 最长回文子串的中心位置索引
	index := 0

	maxFn := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	for j :=1; j < n-1; j++ {

		// todo 参看前文第五部分  {很重要}
		if mx > j {
			p[j] = maxFn(p[2*id-j], mx-j)
		} else {
			p[j] = 1
		}

		// 向左右两边延伸，扩展右边界
		for tmpRune[j+p[j]] == tmpRune[j-p[j]] {
			p[j]++
		}

		// 如果回文子串的右边界超过了mx，则需要更新mx和id的值
		if mx < p[j] + j {
			mx = p[j] + j
			id = j
		}

		// 如果回文子串的长度大于maxLength，则更新maxLength和index的值
		if maxLength < p[j] - 1 {
			// 参看前文第三部分
			maxLength = p[j] - 1
			index = j
		}
	}
	// 第三步：截取字符串，输出结果
	// 起始索引的计算参看前文第四部分
	start := (index-maxLength)/2
	return string(str[start: start + maxLength])
}
