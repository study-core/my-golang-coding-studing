package main

import "fmt"

func main() {
	arr := []int{4, 78, 65, 12, 57, 34, 1, 100, 43}
	QickSort2(arr)
	fmt.Println(arr)
}


// 快排练习
func QickSort (arr []int) {
	QuickRealSort(arr, 0, len(arr) - 1)
}

func QuickRealSort (arr []int, left, right int)  {
	if left < right {
		pivot := partition(arr, left, right)
		QuickRealSort(arr, left, pivot - 1)
		QuickRealSort(arr, pivot + 1, right)
	}
}

func partition(arr []int, left, right int) int {
	for left < right {
		for left < right && arr[left] <= arr[right] {
			right --
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left ++
		}
		for left < right && arr[left] <= arr[right] {
			left ++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			right --
		}
	}
	return left
}


// 堆排练习
func HeapSort(arr []int) {

	for i := len(arr)/2 -1; i >= 0; i -- {
		HeapAdjust2(arr, i, len(arr))
	}

	for i := len(arr) - 1; i > 0; i -- {
		arr[i], arr[0] = arr[0], arr[i]
		HeapAdjust2(arr, 0, i)
	}
}

func HeapAdjust2(arr []int, parent, length int) {
	i := parent
	for {
		left := 2 * parent + 1
		right := 2 * parent + 2

		if left < length && arr[left] > arr[i] {
			i = left
		}
		if right < length && arr[right] > arr[i] {
			i = right
		}

		if i != parent {
			arr[i], arr[parent] = arr[parent], arr[i]
			parent = i
		}else{
			break
		}
	}
}





// 冒泡练习
func bubbleSort(arr []int) {

	var (
		length = len(arr)
		flag = true
	)

	for flag{
		flag = false
		for i := 0; i < length - 1; i ++ {
			if arr[i] > arr[i + 1] {
				arr[i], arr[i + 1] = arr[i + 1], arr[i]
				flag = true
			}
		}
		length -= 1
	}
}

// 插入练习
func InsertSort(arr []int){

	for i := 0; i < len(arr) - 1; i ++ {
		for j := i + 1; j > 0; j -- {
			if arr[j] < arr [j -1] {
				arr[j], arr[j - 1] = arr[j - 1], arr[j]
			}else {
				break
			}
		}
	}
}

// 选择练习
func SelectSort(arr []int){

	for i := 0; i < len(arr) - 1; i ++ {
		min := i
		for j := i + 1; j < len(arr); j ++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}






// 从大到小

func QickSort2 (arr []int) {
	QuickRealSort2(arr, 0, len(arr) - 1)
}

func QuickRealSort2 (arr []int, left, right int)  {
	if left < right {
		pivot := partition2(arr, left, right)
		QuickRealSort2(arr, left, pivot - 1)
		QuickRealSort2(arr, pivot + 1, right)
	}
}

func partition2(arr []int, left, right int) int {
	for left < right {
		for left < right && arr[left] >= arr[right] {
			right --
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left ++
		}
		for left < right && arr[left] >= arr[right] {
			left ++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			right --
		}
	}
	return left
}
