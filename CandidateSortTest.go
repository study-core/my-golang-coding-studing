package main

import (
	"fmt"
	"math/big"
)

func main() {

	//arr := []int{4, 5, 8, 7, 9, 3, 6}
	//candidateSort2(arr)
	//fmt.Println(arr)


	a :=  &Candidate2{
		CandidateId:  "1",
		Deposit: new(big.Int).SetUint64(90),
		BlockNumber: new(big.Int).SetInt64(32),
		TxIndex: 	6}
	b :=  &Candidate2{
		CandidateId:  "2",
		Deposit: new(big.Int).SetUint64(99),
		BlockNumber: new(big.Int).SetInt64(32),
		TxIndex: 	7}

	c :=  &Candidate2{
		CandidateId:  "3",
		Deposit: new(big.Int).SetUint64(100),
		BlockNumber: new(big.Int).SetInt64(31),
		TxIndex: 	6}

	d :=  &Candidate2{
		CandidateId:  "4",
		Deposit: new(big.Int).SetUint64(100),
		BlockNumber: new(big.Int).SetInt64(32),
		TxIndex: 	5}

	e := &Candidate2{
		CandidateId:  "5",
		Deposit: new(big.Int).SetUint64(100),
		BlockNumber: new(big.Int).SetInt64(32),
		TxIndex: 	4}

	f :=  &Candidate2{
		CandidateId:  "6",
		Deposit: new(big.Int).SetUint64(80),
		BlockNumber: new(big.Int).SetInt64(30),
		TxIndex: 	3}

	g :=  &Candidate2{
		CandidateId:  "7",
		Deposit: new(big.Int).SetUint64(100),
		BlockNumber: new(big.Int).SetInt64(31),
		TxIndex: 	3}

	h :=  &Candidate2{
		CandidateId:  "8",
		Deposit: new(big.Int).SetUint64(99),
		BlockNumber: new(big.Int).SetInt64(32),
		TxIndex: 	8}

	i :=  &Candidate2{
		CandidateId:  "9",
		Deposit: new(big.Int).SetUint64(99),
		BlockNumber: new(big.Int).SetInt64(32),
		TxIndex: 	2}

	j :=  &Candidate2{
		CandidateId:  "10",
		Deposit: new(big.Int).SetUint64(99),
		BlockNumber: new(big.Int).SetInt64(32),
		TxIndex: 	3}


	arr := []*Candidate2{a, b, c, d, e, f, g, h, i, j}
	ids := make([]string, 0)
	canMap := make(map[string]*Candidate2, 0)
	for _, can := range arr {
		ids = append(ids, can.CandidateId)
		canMap[can.CandidateId] = can
	}

	candidateSort3(ids, canMap)

	//candidateSort2(arr)
	fmt.Println("%+v", arr)  // [0xc04208c1e0 0xc04208c140 0xc04208c0a0 0xc04208c0f0 0xc04208c050 0xc04208c000 0xc04208c190]

	//for _, v := range arr {
	//	fmt.Println("%v", v,  "id:", v.CandidateId, "deposit:", v.Deposit, "blockNum:", v.BlockNumber.String(), "txIndex:", v.TxIndex)
	//}
	for _, v := range ids {
		fmt.Println("%v", canMap[v], "id:", canMap[v].CandidateId, "deposit:", canMap[v].Deposit, "blockNum:", canMap[v].BlockNumber.String(), "txIndex:", canMap[v].TxIndex)
	}






}



type Candidate2 struct {

	// 抵押金额(保证金)数目
	Deposit			*big.Int
	// 发生抵押时的当前块高
	BlockNumber 	*big.Int
	// 发生抵押时的tx index
	TxIndex 		uint32
	// 候选人Id
	CandidateId 	string
	//
	Host 			string
	Port 			string



}

//func compare2(c, can *Candidate2) int {
//	// 质押金大的放前面
//	if c.Deposit.Cmp(can.Deposit) > 0 {
//		return 1
//	}else if c.Deposit.Cmp(can.Deposit) == 0 {
//		//return 0
//		if c.BlockNumber.Cmp(can.BlockNumber) > 0 {
//			return -1
//		}else if c.BlockNumber.Cmp(can.BlockNumber) == 0 {
//			//return 0
//			if c.TxIndex > can.TxIndex {
//				return -1
//			}else if c.TxIndex == can.TxIndex {
//				return 0
//			}else {
//				return 1
//			}
//
//		}else {
//			return 1
//		}
//	}else {
//		return -1
//	}
//}
//
func candidateSort2(arr []*Candidate2) {
	quickRealSort2(arr, 0, len(arr) - 1)
}
func quickRealSort2 (arr []*Candidate2, left, right int)  {
	if left < right {
		pivot := partition2(arr, left, right)
		quickRealSort2(arr, left, pivot - 1)
		quickRealSort2(arr, pivot + 1, right)
	}
}
func partition2(arr []*Candidate2, left, right int) int {
	for left < right {
		for left < right &&  compare2(arr[left], arr[right])  >= 0 {
			fmt.Println(arr[left], "Cmp",  arr[right])
			right --
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left ++
		}
		for left < right && compare2(arr[left], arr[right])  >= 0 {
			fmt.Println(arr[left], "Cmp",  arr[right])
			left ++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			right --
		}
	}
	return left
}





func compare2(c, can *Candidate2) int {
	// 质押金大的放前面
	if c.Deposit.Cmp(can.Deposit) > 0 {
		return 1
	}else if c.Deposit.Cmp(can.Deposit) == 0 {
		//return 0
		if c.BlockNumber.Cmp(can.BlockNumber) > 0 {
			return -1
		}else if c.BlockNumber.Cmp(can.BlockNumber) == 0 {
			//return 0
			if c.TxIndex > can.TxIndex {
				return -1
			}else if c.TxIndex == can.TxIndex {
				return 0
			}else {
				return 1
			}

		}else {
			return 1
		}
	}else {
		return -1
	}
}

func candidateSort3(arr []string, cans map[string]*Candidate2) {
	quickRealSort3(arr, cans, 0, len(arr) - 1)
}
func quickRealSort3 (arr []string, cans map[string]*Candidate2, left, right int)  {
	if left < right {
		pivot := partition3(arr, cans, left, right)
		quickRealSort3(arr, cans, left, pivot - 1)
		quickRealSort3(arr, cans, pivot + 1, right)
	}
}
func partition3(arr []string, cans map[string]*Candidate2, left, right int) int {
	for left < right {
		for left < right &&  compare2(cans[arr[left]], cans[arr[right]])  >= 0 {
			fmt.Println(arr[left], "Cmp",  arr[right])
			right --
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left ++
		}
		for left < right && compare2(cans[arr[left]], cans[arr[right]])  >= 0 {
			fmt.Println(arr[left], "Cmp",  arr[right])
			left ++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			right --
		}
	}
	return left
}