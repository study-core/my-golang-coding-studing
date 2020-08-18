package main


// todo 滑动窗口算法可以用以解决  数组/字符串的子元素问题, 它可以将嵌套的循环问题, 转换为单循环问题, 降低时间复杂度.

// todo 滑动窗口实际上是,  利用决策单调性来实现复杂度优化.
//		滑动窗口适用的题目一般具有单调性
//		滑动窗口、双指针、单调队列和单调栈经常配合使用

// todo 应用:
//		流量控制
// 		TCP协议中, 还有就是为服务的限流, 滑动窗口算法功能上相当于令牌桶算法.
func main() {

}



/**
todo 题 1

给定一个整数数组，计算长度为 'k' 的连续子数组的最大总和。

如:

输入：arr [] = {100,200,300,400},   k = 2

输出：700

解释：300 + 400 = 700
 */

func slidingWindow_maxCount(arr []int, k int) int {

	if len(arr) < k {
		return -1
	}

	// 初始化最大值, 最初的窗口也就是最大值
	var maxCount int
	for i := 0; i < k; i++ {
		maxCount += arr[i]
	}

	// 初始化最初窗口的 总和值 <窗口的宽度为 k>
	windowSum := maxCount

	maxFn := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	// 开始将窗口往右滑动
	for i := k; i < len(arr); i++ {
		// 新窗口的和 = 前一个窗口的和 + 新进入窗口的值 - 移出窗口的值
		windowSum += arr[i] - arr[i -k]
		maxCount = maxFn(maxCount, windowSum)
	}
	return maxCount
}


/**
todo 题 2

给定一个字符串 S 和一个字符串 T，请在 S 中找出包含 T 所有字母的最小子串。

输入: S = "ADOBECODEBANC", T = "ABC"

输出: "BANC"
 */




/**
todo 题 3

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度

输入: "abcabcbb"

输出: 3

解释: 因为无重复字符的最长子串是 "abc" 或者是 "cab" 或者是 "bca" 所以其长度为 3

 */