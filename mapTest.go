package main

import "fmt"

func main() {

	map1 := map[string]int{
		"a": 4,
		"b": 5,
		"c": 6,
		"d": 7,
	}

	map2 := map1

	delete(map1, "a")

	fmt.Println(map1)
	fmt.Println(map2)
}
