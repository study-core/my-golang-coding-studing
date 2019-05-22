package main

import (
	"fmt"
	"encoding/json"
)

func main() {

	//myFunc := func(queue []*mst) {
	//
	//	if len(queue) > 2 {
	//		queue = (queue)[:2]
	//
	//	}
	//
	//	arr, _ := json.Marshal(queue)
	//	fmt.Println("截取后的arr", string(arr))
	//}


	ma1 := &ma{
		Nme: 	"a",
	}

	ma2 := &ma{
		Nme: 	"b",
	}

	ma3 := &ma{
		Nme: 	"c",
	}

	mst1 := &mst{
			1,
		"NA",
		ma1,
	}

	mst2 := &mst{
			2,
		"NB",
		ma2,
	}

	mst3 := &mst{
			3,
		"NC",
		ma3,
	}

	arr := make([]*mst, 3)
	arr[0] = mst1
	arr[1] = mst2
	arr[2] = mst3

	//queue, _ := json.Marshal(arr)
	//fmt.Println("借去之前，外面：", string(queue))
	//
	//myFunc(arr)
	//
	//
	//queue2, _ := json.Marshal(arr)
	//fmt.Println("截取之后，外面：", string(queue2))


	marr := make([]*mst, 3)

	sli := []int{2, 3, 1}

	// copy
	retry:
	for i, id := range  sli {
		fmt.Println("现在是i==", i)
		for k := 0; k < len(arr); k ++ {
			a := arr[k]

			if a.Id == id {
				marr[i] = a
				arr = append(arr[:k], arr[k+1:]...)
				continue retry
			}

		}
	}

	queue3, _ := json.Marshal(arr)
	fmt.Println("Copy之后，原数组：", string(queue3))

	queue4, _ := json.Marshal(marr)
	fmt.Println("Copy之后，新数组：", string(queue4))

}


type mst struct {
	Id 		int
	Name 	string

	Ma 		*ma

}

type ma struct {
	Nme string
}