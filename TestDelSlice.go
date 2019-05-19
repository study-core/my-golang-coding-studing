package main

import "fmt"

func main() {

	//arr := []int{4, 5, 8, 9, 6, 7, 2}
	arr := []int{4, 5, 8, 9, 6, 7, 2}

	//for i, v := range arr {
	//	fmt.Println("i", i, "len", len(arr))
	//	if v == 8 || v == 6 || v == 2 {
	//		arr = append(arr[:i], arr[i+1:]...)
	//	}
	//}

	for i := 0; i < len(arr); i ++ {
		fmt.Println("i", i, "len", len(arr))
		v := arr[i]
		if v == 8 || v == 6 || v == 2 {
			arr = append(arr[:i], arr[i+1:]...)
			i --
			fmt.Println("删了之后: len", len(arr))
		}
	}
	fmt.Println(arr)
}
