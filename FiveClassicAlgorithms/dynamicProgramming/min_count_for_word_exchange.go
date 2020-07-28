package main

import "fmt"

func main() {
	source := "horse"
	target := "ros"  // expect: 3
	count := minCountForWordExchange(source, target)
	fmt.Println(count)
	count = minCountForWordExchangeOptimization(source, target)
	fmt.Println(count)

}




/**
给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

	todo  a. 插入一个字符      b. 删除一个字符      c. 替换一个字符

	示例：
	输入: word1 = "horse", word2 = "ros"
	输出: 3
	解释:
	horse -> rorse (将 'h' 替换为 'r')
	rorse -> rose (删除 'r')
	rose -> ros (删除 'e')
 */


// todo 90% 的字符串问题都可以用动态规划解决，并且90%是采用二维数组

// O(n*m) 空间复杂度
func minCountForWordExchange(source, target  string) int {


	// 老套路, 三步走


	// todo 步骤一：定义数组元素的含义

	/**
	由于我们的目的求将 word1 转换成 word2 所使用的最少操作数 。

	todo 那我们就定义 dp[i] [j]的含义为：当字符串 word1 的长度为 i，字符串 word2 的长度为 j 时，
		将 word1 转化为 word2 所使用的最少操作次数为 dp[i] [j]
	 */
	m := len(source)
	n := len(target)
	//int[][] dp = new int[n1 + 1][n2 + 1]
	// todo 为什么需要 +1,   因为存在 [len(source) == 0  -> len(target) == n]    和   [len(source) == m  -> len(target) == 0]
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		item := make([]int, n+1)
		dp[i] = item
	}


	// todo 步骤二: 找出关系数组元素间的关系式  (状态方程)

	/**
	接下来我们就要找 dp[i] [j] 元素之间的关系了，比起其他题，这道题相对比较难找一点，但是，不管多难找，

	todo 大部分情况下，dp[i] [j] 和 dp[i-1] [j]、dp[i] [j-1]、dp[i-1] [j-1] 肯定存在某种关系.


	因为我们的目标就是，**从规模小的，通过一些操作，推导出规模大的**


	对于这道题，我们可以对 word1 进行三种操作

	todo  a. 插入一个字符      b. 删除一个字符      c. 替换一个字符

	由于我们是要让操作的次数最小，所以我们要寻找最佳操作。那么有如下关系式：

	todo 一、如果我们 word1[i] 与 word2 [j] 相等，这个时候不需要进行任何操作，显然有 dp[i] [j] = dp[i-1] [j-1]。（别忘了 dp[i] [j] 的含义哈）。

	todo 二、如果我们 word1[i] 与 word2 [j] 不相等，这个时候我们就必须进行调整，而调整的操作有 3 种，我们要选择一种。
		三种操作对应的关系试如下（注意字符串与字符的区别）：

	todo
		（1）、[len(word1) == len(word2) 时]    如果把字符 word1[i] `替换`成与 word2[j] 相等，则有 dp[i] [j] = dp[i-1] [j-1] + 1;
		（2）、[len(word1) < len(word2) 时]    如果在字符串 word1末尾 `插入`一个与 word2[j] 相等的字符，则有 dp[i] [j] = dp[i] [j-1] + 1;
		（3）、[len(word1) > len(word2) 时]    如果把字符 word1[i] `删除`，则有 dp[i] [j] = dp[i-1] [j] + 1;

	那么我们应该选择一种操作，使得 dp[i] [j] 的值最小，显然有

	todo
		dp[i] [j] = min(dp[i-1] [j-1], dp[i] [j-1], dp[[i-1] [j]]) + 1;

	todo 因为上面的分析可以看到 三种都是 +1, 取 min(dp[i-1] [j-1] + 1,  dp[i] [j-1] + 1,  dp[i-1] [j] + 1)  最小值的话
		其实就是:   min(dp[i-1] [j-1], dp[i] [j-1], dp[[i-1] [j]]) + 1

	于是，我们的关系式就推出来了，
	 */

	// todo 步骤三：找出初始值  (最优子结构)

	/**
	显然，当 dp[i] [j] 中，如果 i 或者 j 有一个为 0，那么还能使用关系式吗？
	答是不能的，因为这个时候把 i - 1 或者 j - 1，就变成负数了，数组就会出问题了，
	todo 所以我们的初始值是计算出所有的 dp[0] [0….n] 和 所有的 dp[0….m] [0]
	这个还是非常容易计算的，todo 因为当有一个字符串的长度为 0 时，转化为另外一个字符串，那就只能一直进行插入或者删除操作了
	 */


	minFn := func(a, b, c int) int {
		tmp := 0
		if a < b {
			tmp = a
		}else {
			tmp = b
		}

		if tmp < c {
			return tmp
		}
		return c
	}


	// 初始化

	// dp[0][0...n2]的初始值  二维矩阵 第一行的所有值
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j - 1] + 1
	}
	// dp[0...n1][0] 的初始值  二维矩阵 第一列的所有值
	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i - 1][0] + 1
	}

	// 通过公式推出 dp[m][n]
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {

			if source[i-1] == target[j-1] { // 如果 word1[i] 与 word2[j] 相等, 第 i 个字符 在 string中 对应下标是 i-1
				dp[i][j] = dp[i - 1][j - 1]   // todo 因为 word1[i] 与 word2[j] 相等, 所以 本次不需要操作, 即 操作数 还是引用上一个的 (用 dp[i - 1][j - 1]的)
			}else {

				// 如果需要发生, 替换|插入|删除 时, 使用状态方程
				dp[i][j] = minFn(dp[i - 1][j - 1], dp[i][j - 1], dp[i - 1][j]) + 1
			}
		}
	}
	return dp[m][n]


}




// O(n*m) 空间复杂度优化成 O(n)
func minCountForWordExchangeOptimization (source, target string) int {

	// todo  思路 类似 climb_grid.go 的 climbGridOptimization()
	// todo 我们先根据原来的 二维状态方程,  dp[i][j] = minFn(dp[i - 1][j - 1], dp[i][j - 1], dp[i - 1][j]) + 1,  变换为 dp[n] 一维中的状态方程

	// todo 怎么变? 一维 dp[i-1] 是 原来二维的  dp[i][j-1], 而 dp[i]<旧值>  是原来二维的 dp[i-1][j], 那么一维数组 装不下 原来二维的第三个值  dp[i-1][j-1] 了， 故我们需要一个临时变量  pre

	// todo 得到新的状态方程 如下:
	//  	dp[i] = min(dp[i-1], pre, dp[i]) + 1

	m := len(source)
	n := len(target)
	dp := make([]int, n + 1)

	// dp[0...n]的初始值, todo 即 原 二维的 第一行
	for j := 0; j <= n; j++ {
		dp[j] = j  // 看下原来 二维的就懂了
	}

	minFn := func(a, b, c int) int {
		tmp := 0
		if a < b {
			tmp = a
		}else {
			tmp = b
		}

		if tmp < c {
			return tmp
		}
		return c
	}

	// 状态方程: dp[j] = min(dp[j-1], pre, dp[j]) + 1
	for i := 1; i <= m; i++  {
		temp := dp[0]

		// 相当于初始化
		dp[0] = i
		for j := 1; j <= n; j++ {

			// todo pre 相当于之前的 dp[i-1][j-1]
			pre := temp
			temp = dp[j]

			// 如果 word1[i] 与 word2[j] 相等, 第 i 个字符对应下标是 i-1
			if source[i-1] == target[j-1] {
				dp[j] = pre
			}else {

				// todo  启用 一维的状态方程
				dp[j] = minFn(dp[j - 1], pre, dp[j]) + 1
			}
		// 保存要被抛弃的值
		}
	}
	return dp[n]
}