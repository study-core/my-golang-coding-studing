package main

import (
	"encoding/json"
	"fmt"
)

func main() {


	aE := &sliceElm{

		Field:   "aE",
	}
	bE := &sliceElm{

		Field:   "bE",
	}
	cE := &sliceElm{

		Field:   "cE",
	}
	dE := &sliceElm{

		Field:   "dE",
	}

	a := &sliceCopy{
		Name: "A",

		Age: 1,

		E:  aE,
	}

	b := &sliceCopy{
		Name: "B",

		Age: 2,

		E:  bE,
	}

	c := &sliceCopy{
		Name: "C",

		Age: 3,

		E:  cE,
	}

	d := &sliceCopy{
		Name: "D",

		Age: 4,

		E:  dE,
	}


	arr := []*sliceCopy{a, b, c, d}

	arrStr, _ := json.Marshal(arr)

	fmt.Println("copy 前:", string(arrStr))

	copyArr := make([]*sliceCopy, len(arr))
	copy(copyArr, arr)

	arrStr2, _ := json.Marshal(arr)

	fmt.Println("copy 后:", string(arrStr2))

	copyArrStr, _ := json.Marshal(copyArr)

	fmt.Println("copy 的arr:", string(copyArrStr))

	for i, _ := range copyArr {
		if i == 2 {
			copyArr = append(copyArr[:i], copyArr[i+1:]...)
			break
		}
	}

	arrStr3, _ := json.Marshal(arr)

	fmt.Println("处理后:", string(arrStr3))

	copyArrStr2, _ := json.Marshal(copyArr)

	fmt.Println("处理后 copy 的arr:", string(copyArrStr2))

}


type sliceCopy struct {

	Name  string

	Age  int


	E 	*sliceElm

}

type sliceElm struct {

	Field   string
}