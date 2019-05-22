package main

import (
	"fmt"
	"sort"
)

func main() {
	//arr := []int{1, 4, 5, 7, 8, 9, 14}
	arr := []int{1, 4}
	fmt.Println(solution(arr))
	fmt.Println(solution2(arr))
}


func solution (arr []int) bool {
	l := len(arr)
	if len(arr) < 3 {
		return false
	}

	for i := 0; i < l - 2; i++ {
		if arr[i] > 0 && arr[i] > arr[i + 2] - arr[i + 1] {
			return true
		}
	}
	return false
}



/**
思路: 三角形为 任意两边的和大于第三边，换句话，两个最小边的和大于最大边

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