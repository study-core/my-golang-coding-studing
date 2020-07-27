package main

import "fmt"

// https://zhuanlan.zhihu.com/p/91582909
func main() {
	num := climbStairs(6)
	fmt.Println(num)
}


// TODO 爬楼梯


// 思路:
//
// todo 动态规划的三大步骤
//
//
func climbStairs(n int) int { // n为 楼梯的阶数

	if n <= 1 {
		return n
	}

	// todo 先创建一个数组来保存历史数据  (保存 历史解法数量)
	dp := make([]int, n+1) // n+1 是因为, 当 n的取值为 0 -> n 而不是 1 -> n

	// todo 给出初始值
	dp[0] = 0    // 当台阶数 n == 0时, 只有 0种解法
	dp[1] = 1	 // 当台阶数 n == 1时, 只有 1中解法


	// todo 通过关系式来计算出 dp[n]
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]  // dp[i] 的所有解法数 == dp[i-1] + dp[i-2]
	}
	// 把最终结果返回
	return dp[n]
}