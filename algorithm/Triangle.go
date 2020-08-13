package main

import (
	"fmt"
	"sort"
)

// todo 找出 数组中是否存在 有三个数 组成三角形
func main() {
	arr := []int{1, 4, 5, 7, 8, 9, 14}
	//arr := []int{1, 4}
	fmt.Println(solution(arr))
	fmt.Println(solution2(arr))
}

// todo 这种解法 最优
func solution (arr []int) bool {
	sort.Ints(arr)
	l := len(arr)
	if len(arr) < 3 {
		return false
	}

	// 思路: 三角形为 任意两边的和大于第三边，换句话，todo 两个最小边的和 大于  最大边

	// todo 对比下面的解法，总感觉这种解法 覆盖不全，有漏洞？

	// todo  细细想， 其实这种是正确的， 排序后的  a  < b  < c 其中 a 和 b 离 c 最近，
	// 		你想想， 最近的两个数的和都没比 c 大，那么 a 和 b 之前的 随意两数和 会比 c 大？
	for i := 0; i < l - 2; i++ {
		if arr[i] > 0 && arr[i] + arr[i + 1]  > arr[i + 2] {  // todo  精髓在这里, a + b > c, 其中  a < c, b < c
			return true
		}
	}
	return false
}



/**
思路: 三角形为 任意两边的和大于第三边，换句话，todo 两个最小边的和 大于  最大边

先对数组排序，逐个以 大的数作为基准，然后用左右指针先对应第二大的数和最小数
求和与最大数作比较，然后再用第二大数和第二小数求和和最大数作比较，依次这样比较完。
然后，再以第二大数作为新基准，用第三大数和最小数求和去和第二大数作比较，
然后再用第三大数和第二小数求和和第二大数作比较，依次比完，
就这样一直更换新的基准数逐个比完数组
*/
func solution2(arr []int) bool {
	sort.Ints(arr)
	for i := len(arr) -1; i >= 2; i-- {
		left, right := 0, i - 1
		for left < right {
			if arr[left] + arr[right] > arr[i] {
				right --
				return true
			}else {
				left ++
			}
		}
	}
	return false
}