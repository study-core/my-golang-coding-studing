package main

import "fmt"

// https://www.cnblogs.com/yjiyjige/p/3263858.html#commentform
func main() {

	//pat := "abbca"
	pat := "abcadabcabd"
	txt := "adfasfsfsaacvdfgsdeeabbabcabcabdabcabdcarergsdvdfabbcabbca"

	index := KMP(txt, pat)
	fmt.Println(index) // 20
}

// todo 正常的  kmp 算法， 使用 next数组

/**
 例如 我们有 被查找字符串  txt为: A B C A B C D H I J K
      匹配 字符串 pat为: A B C A B B

	那么我们有,

		0		  i
txt:	A B C A B C | D H I J K
pat:	A B C A B B |
		0		  j

当 pat的第j个字符匹配到 txt的第i个字符时出现 不匹配，这时候我们应该将 j在pat中往前回溯，
而不是将i在txt中往前回溯(将i在txt中往前回溯，是暴力匹配的做法)

这时候我们需要将 txt的 i固定， 然后将pat的j往回移动， 那么移动多少呢？如：


		0		  i
txt:	A B C A B C | D H I J K
pat:		  A B C | A B B
	      	  0	  j

我们可以大概看出一点端倪，当匹配失败时，j 要移动的下一个位置 k。 todo  存在着这样的性质：最前面的k个字符和j之前的最后k个字符是一样的.

对于 pat 来说，我们有：

	0   k-1       j-k        j
	A    B    C    A    B    B
			  k        j-1


为什么可以直接将j移动到k位置了。

因为:

当T[i] != P[j]时

有T[i-j ~ i-1] == P[0 ~ j-1]

由P[0 ~ k-1] == P[j-k ~ j-1]

必然：T[i-k ~ i-1] == P[0 ~ k-1]


todo 接下来就是重点了，怎么求这个（这些）k呢？

todo 因为在P的每一个位置都可能发生不匹配，也就是说我们要计算 【每一个位置 j 对应的 k 】，所以用一个数组next来保存，next[j] = k，表示当T[i] != P[j]时，j指针的下一个位置。


todo  上面的 退到 k 其实是 使用 pat中的 第k个字符和  txt的 第i个字符对比
 */

func getNext(pat string) []int {

	p := []rune(pat)

	plen := len(p)

	next := make([]int, plen)

	next[0] = -1

	j := 0
	k := -1 // 表示 已经是 pat的最起始位置了, 我们需要将 txt的坐标往后移, 从动作上感觉是 pat的窗口顺着 txt的轨道往后移一格

	// todo  构造 pat 的 next 数组
	//
	// todo 下面的循环 我们比喻为,  k  和 j 作为边界的 移动窗口，一直向后滑，并且会根据实际情况 调整  k 和 j 的坐标  由此推出各个 pat中的 字符需要跳转的位置
	//
	// todo 需要画图 演算才能明白怎么回事 ##############
	for j < plen-1 {
		if k == -1 || p[j] == p[k] {
			j++
			k++
			if p[j] == p[k] { // 当两个字符相等时要跳过
				next[j] = next[k]
				fmt.Println("j :=", j, "k :=", k, "进入 a, next[k] :=", next[k])
			} else {
				next[j] = k
				fmt.Println("j :=", j, "k :=", k, "进入 b")
			}
		} else { // 当 P[k] != P[j]
			k = next[k] // todo 这个操作 是在 调整 k
			fmt.Println("j :=", j, "k :=", k, "进入 c")
		}
	}
	/**
	由于KMP算法中指针i是不减的，因此j的指向位置只与模式串本身的结构有关。j的滑动位置的信息存放在next数组中。当匹配失败，就可以通过查询next数组的值得到下一次j滑动的位置。

	next数组存放的是模式串的移位信息，具体就是模式串的部分匹配值，next数组大小与模式串T等长

	 */
	 fmt.Println("next := ", next)
	return next
}

func KMP(txt, pat string) int {

	t := []rune(txt)

	p := []rune(pat)

	i := 0 // 主串的位置

	j := 0 // 模式串的位置

	next := getNext(pat)

	for i < len(t) && j < len(p) {
		if j == -1 || t[i] == p[j] { // 当j为-1时，要移动的是i，当然j也要归0
			i++
			j++
		} else {
			// i不需要回溯了
			// i = i - j + 1;
			j = next[j] // j回到指定位置
		}
	}

	if j == len(p) {
		return i - j
	} else {
		return -1

	}
}
