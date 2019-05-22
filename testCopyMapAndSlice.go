package main

import "fmt"

func main() {

	aM := make(map[string]int, 0)

	mc := myCache{
		m: &mapStruct{
			myMap: aM,
		},
	}

	mc.GetMyMap().myMap["a"] = 1

	fmt.Println(aM)

	mc.GetMyMap().myMap["b"] = 1

	fmt.Println(mc.GetMyMap().myMap)
}

type myCache struct {
	m *mapStruct
}

type mapStruct struct {
	myMap map[string]int
}


func (mc *myCache) GetMyMap () *mapStruct {
	return mc.m
}