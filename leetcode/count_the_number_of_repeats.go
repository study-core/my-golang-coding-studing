package main

import "fmt"

/**


由 n 个连接的字符串 s 组成字符串 S，记作 S = [s,n]。例如，["abc",3]=“abcabcabc”。

如果我们可以从 s2 中删除某些字符使其变为 s1，则称字符串 s1 可以从字符串 s2 获得。例如，根据定义，"abc" 可以从 “abdbec” 获得，但不能从 “acbbe” 获得。

现在给你两个非空字符串 s1 和 s2（每个最多 100 个字符长）和两个整数 0 ≤ n1 ≤ 10^6 和 1 ≤ n2 ≤ 10^6。现在考虑字符串 S1 和 S2，其中 S1=[s1,n1] 、S2=[s2,n2] 。

请你找出一个可以满足使[S2,M] 是从 S1 获得出来的, 求最大整数 M 。



示例：

输入：
s1 ="acb",n1 = 4
s2 ="ab",n2 = 2

返回：
2

 */
func main() {
	m := getMaxRepetitions("abaacdbac", "adcbd", 100, 4)
	//m := getMaxRepetitions("acb", "ab", 4, 2)
	fmt.Println(m)
}

// todo 思路：
//	     找出 n1 个 s1 组成的序列 S1 中的字母序列是否可以 有一定规律出现 s2, todo 即 找出 [循环节]
//
//
// 假设:
// s1: "abaacdbac",  n1: 100
// s2: "adcbd",      n2: 4
//
// S1 = [s1, n1]
// S2 = [s2, n2]
//
//
//
// 对于 S1:    abaacdbac|abaacdbac|abaacdbac|abaacdbac|abaacdbac|...
// 对于 S2:	   a  	d  c  b   d|a       d  c  b   d|a       d  ...
//
// s2的索引:           2        0          2        0 ...
//
//
// 图上看, 当 n = 100 时,
//			[循环节]的个数为,  (100 -1)/2 = 49, todo 上面 第1个 s1 到 第三个 s1时才算一次循环, 所以 100-1
//			剩余 (100 -1)%2 = 1 个 s1
//
// 由上可见,
//			在出现 [循环节]之前, 就已经有一个 s1 了
// 			每个[循环节] 包含 2 个 s1, 包含 1 个 s2
//
//


/**
复杂度分析

时间复杂度：O(|s1|*|s2|)O(∣s1∣∗∣s2∣)。我们最多找过 |s2| + 1 个 s1，就可以找到循环节，最坏情况下需要遍历的字符数量级为 O(|s1|*|s2|)O(∣s1∣∗∣s2∣)。

空间复杂度：O(|s2|)O(∣s2∣)。我们建立的哈希表大小等于 s2 的长度

 */
func getMaxRepetitions(s1, s2 string, n1, n2 int) int {

	if n1 == 0 { // S1 [s1, 0], 所以 没有可匹配的  S2[s2, M]
		return 0
	}
	s1cnt, index, s2cnt := 0, 0, 0

	// recall 是我们用来找循环节的变量，它是一个哈希映射
	//
	// 我们如何找循环节？
	// 【一】假设我们遍历了 s1cnt 个 s1，此时匹配到了第 s2cnt 个 s2 中的第 index 个字符
	// 【二】如果我们之前遍历了 s1cnt' 个 s1 时，匹配到的是第 s2cnt' 个 s2 中【同样】的第 index 个字符，那么就有循环节了
	// 【三】我们用 (s1cnt', s2cnt', index) 和 (s1cnt, s2cnt, index) 表示【两次】包含相同 index 的匹配结果
	// 那么 Map 中的 `key` 就是 index，`value` 就是 (s1cnt', s2cnt') 这个二元组
	// 循环节就是；
	//    - 前 s1cnt' 个 s1 包含了 s2cnt' 个 s2
	//    - 以后的每 (s1cnt - s1cnt') 个 s1 包含了 (s2cnt - s2cnt') 个 s2
	// 那么还会剩下 (n1 - s1cnt') % (s1cnt - s1cnt') 个 s1, 我们对【这些剩下的s1】与 s2 进行暴力匹配
	//
	// todo 注意: s2 要从第 index 个字符开始匹配

	recall := make(map[int][2]int, 0)
	pre_loop, in_loop := [2]int{}, [2]int{}

	for {

		// 我们多遍历一个 s1，看看能不能找到循环节
		s1cnt += 1
		for i := 0; i < len(s1); i++ {

			ch := s1[i]          		// 逐个取出 s1的字符, 和s2的 第index 个字符匹配

			if ch == s2[index] {
				index += 1
				if index == len(s2) {	// 如果 s2 的字符被匹配完
					s2cnt += 1 			// 则, 累加 s2 的个数计数
					index = 0			// 重置 index, 从s2的第0个索引开始匹配
				}
			}
		}

		// 所有的 s1 就用完了
		if s1cnt == n1 {
			return s2cnt / n2 //todo  注意: 这里的  s2cnt 可能为0  (如果 没有 s2的循环节出现, 则 s2cnt 是0， 否则, 将看看 匹配的 s2个数是 S2的 几分之几, 我们需要 除数)
		}


		// todo 每当 一个 s1 字符遍历完, 我们就来看看 有没有 出现过 [循环节]
		//
		// 出现了之前的 index，表示找到了 [循环节]
		if _, ok := recall[index]; ok {

			item := recall[index]
			// todo 前 s1cnt' 个 s1 包含了 s2cnt' 个 s2
			pre_loop = item

			// todo 以后的每 (s1cnt - s1cnt') 个 s1 包含了 (s2cnt - s2cnt') 个 s2
			in_loop = [2]int{s1cnt - item[0], s2cnt - item[1]}
			break
		} else {

			// 否则, 我们记录 遍历到 s2 的当前索引字符时, 一共消耗了 多少 s1 及 s2
			//
			// {因为可能,  len(s1) >= len(s2) || len(s2) > len(s1) }
			recall[index] = [2]int{s1cnt, s2cnt}
		}
	}

	// ans 存储的是 S1 包含的 s2 的数量，考虑的之前的 pre_loop 和 in_loop
	//
	// 即: 遍历到当前 s1cnt 和 s2cnt 为止的 s2个数 (因为 可能没匹配 s2的完整字符, 可能是中间部分字符, 故是 当前 某个 s2的某个 索引)
	//
	// ans = s2cnt' + (n1 - s1cnt')/ (s1cnt - s1cnt') * (s2cnt - s2cnt')
	//	   = s2cnt' + (n1 - s1cnt') * (s2cnt - s2cnt') /(s1cnt - s1cnt')
	//
	ans := pre_loop[1] + (n1-pre_loop[0])/in_loop[0]*in_loop[1]


	// S1 的末尾还剩下一些 s1，我们暴力进行匹配
	//
	// rest = (n1 - s1cnt') % (s1cnt - s1cnt')
	rest := (n1 - pre_loop[0]) % in_loop[0]


	// 将 剩下的 s1 去和 s2暴力匹配完, todo 可能后面的字符是在没法匹配了, 也要这么做
	for i := 0; i < rest; i++ {
		for i := 0; i < len(s1); i++ {
			ch := s1[i]
			if ch == s2[index] {
				index += 1
				if index == len(s2) {
					ans += 1
					index = 0
				}
			}
		}
	}
	// S1 包含 ans 个 s2，那么就包含 ans / n2 个 S2
	return ans / n2
}
