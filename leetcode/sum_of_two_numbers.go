package main

import "fmt"


/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。



示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
 */
func main() {
	slice := []int{1, 7, 8, 4, 5, 6, 9, 2, 3, 10, 11}
	res := sum_tow_num(slice, 11)
	fmt.Println(res)
}

func sum_tow_num (nums []int, target int) [][]int {
	m := make(map[int]int, 0)  		// 用来做中转过滤的 map[元素]索引
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		sub := target - nums[i]			// a + b = target, 反思: a = target - b
		if index, ok := m[sub]; ok {	// 如果 a 存在 map中, 返回 a, b 的坐标
			res := []int{index, i}
			result = append(result, res)
		} else {
			m[nums[i]] = i
		}
	}
	return result
}