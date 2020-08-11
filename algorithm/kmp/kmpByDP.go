package main

import "fmt"

//func main() {
//
//	str := "sa2233我呆&*特热"
//
//	// 打印 byte，非英文单个 char 可能是 多个 byte
//	fmt.Println("打印 str[i] 开始")
//	for i := 0; i < len(str); i++ {
//		fmt.Println(str[i])
//		fmt.Printf("%c \n", str[i])
//	}
//
//
//	r := []rune(str)
//
//	// 打印 char
//	fmt.Println("打印 rune[i] 开始")
//	for i := 0; i < len(r); i++ {
//		fmt.Println(r[i])
//		fmt.Printf("%c \n", r[i])
//	}
//
//}



// todo  使用 dp (动态规划) 求解 kmp

/**
todo
	编码	大小	支持语言
	ASCII	1个字节	英文
	Unicode	2个字节（生僻字4个）	所有语言
	UTF-8	1-6个字节，英文字母1个字节，汉字3个字节，生僻字4-6个字节	所有语言
 */


// https://zhuanlan.zhihu.com/p/83334559
func main() {

	pat := "abbca"
	txt := "adfasfsfsaacvdfgsdeeabbcarergsdvdfabbcabbca"

	k := new(kmp)
	k.initPattern(pat)
	index := k.search(txt)
	fmt.Println(index) // 20
}


// todo 使用 状态指针 +  dp   (有限状态自动机)



type kmp struct {
	dp  [][]int
	pat string
}


/**
传统的 KMP 算法是使用一个一维数组 next 记录前缀信息，
而本文是使用一个二维数组 dp 以状态转移的角度解决字符匹配问题，但是空间复杂度仍然是 O(256M) = O(M)
 */
func (self *kmp) initPattern(pat string) {


	// 初始化
	self.pat = pat

	// 匹配规则 小str的长度
	m := len(pat)


	// TODO 这个非常重要
	// 		`dp[状态][字符] = 下个状态`
	//
	// todo 使用 256 是因为 asscii 码只有 256 位 (0-255 或 -128-+127)
	//
	// todo ASCll不包含 汉字编码，包含汉字编码的是Unicode
	//
	//
	self.dp = make([][]int, m)
	for i := 0; i < m; i++ {
		self.dp[i] = make([]int, 256)  // todo 存储好， 每一种状态对 下一次出现的字符 (256个字符) 时, 跳转的状态为 多少
	}


	// todo 初始化 最开始的值 (最小子结构)
	//
	// todo 第0个状态 和 第0个 char 记录下一个 pat的 char状态为 1
	self.dp[0][int(byte(pat[0]))] = 1

	// 初始化 影子状态的值 todo 什么是 影子状态, 请查看 连接中的文章
	//
	// 因为是最开始，所以影子状态也是0
	x := 0

	// 初始化 保存状态的  二维 db 的全部值
	//
	// for 循环, 类似上面 `self.db[0][byte(pat[0])] = 1` 将后面的 pat各个场景的状态跳转值全部计算完
	// 全部 填充到 db 中
	for i := 1; i < m; i++ {  // 从 pat 的下标第 1 个字符将 pat各个字符下一个状态跳转的情况全部初始化到 db 中

		for char := 0; char < 256; char ++ {  // 将 256 个 asccii 码遍历完

			// todo 当前 asccii 字符 就是 当前 pat中的第 j 个字符时
			//		我们 记录pat的值 需要往后推一个, 也就是 状态 +1
			if char == int(byte(pat[i])) {
				self.dp[i][char] = i + 1
			} else { // todo 否则，当时 其他 char 时， 我们必须知道剩下 255 个字符应该往 哪个 状态 跳转

				// todo 最开始， 优先往 影子状态 跳转
				self.dp[i][char] = self.dp[x][char]
			}
			// todo 更新影子状态
			// 当前是状态 X，遇到字符 pat[j]，
			// pat 应该转移到哪个状态？
			x = self.dp[x][int(byte(pat[i]))]
		}
	}
}

func (self *kmp) search (txt string) int {

	m := len(self.pat)
	n := len(txt)

	// pat 的初始态为 0
	j := 0
	for  i:= 0; i < n; i++ {
		// 计算 pat 的下一个状态
		j = self.dp[j][int(byte(txt[i]))]
		// 到达终止态，返回结果
		if j == m {
			return i - m + 1
		}
	}
	// 没到达终止态，匹配失败
	return -1
}


