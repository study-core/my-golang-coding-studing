package main

import "fmt"

func main() {
	items := []int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}
	bubbleSort2(items)
	fmt.Println(items)
}

func bubblesort(items []int) {
	var (
		n       = len(items)
		swapped = true
	)
	for swapped {
		swapped = false
		for i := 0; i < n - 1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		n = n - 1
	}
}
func bubblesort2(items []int) {
	n := len(items)
	for j := 0; j < n - 1; j ++ {
		for i := 0; i < n - j - 1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
			}
		}
	}
}
func bubblesort3(items []int) {
	n := len(items)
	for i := 0; i < n - 1; i++ {
		for j := 0; j < n - 1; j++ {
			if items[j+1] < items[j] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
}