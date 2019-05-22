package main

import "fmt"

func main() {

	arr := []int{4, 7, 11, 18, 5, 6, 3, 2, 9}
	//QickSort(arr)
	HeapSort(arr)
	fmt.Println(arr)
}


func QickSort(arr []int){
	sort(arr, 0, len(arr) - 1)
}

func sort(arr []int, left, right int){
	if left < right {
		pivot := partion(arr, left, right)
		sort(arr, left, pivot - 1)
		sort(arr, pivot + 1, right)
	}
}


func partion(arr []int, left, right int) int {
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




func HeapSort(arr []int) {
	// 构造堆
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjust(arr, i, len(arr))
	}

	// 调整堆
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		adjust(arr, 0, i)
	}
}

func adjust(arr []int, parent, len int) {
	var i int
	for 2 * parent + 1 < len {
		lchild := 2 * parent + 1
		rchild := 2 * parent + 2
		i = lchild

		if rchild < len && arr[lchild] < arr[rchild] {
			i = rchild
		}
		if arr[i] > arr[parent] {
			arr[i], arr[parent] = arr[parent], arr[i]
			i = parent
		}else { break}
	}
}