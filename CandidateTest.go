package main

import (
	"math/big"
	"fmt"
)

func main() {
	a :=  &Candidate{
		CandidateId:  "1",
		Deposit: new(big.Int).SetUint64(90),
		BlockNumber: new(big.Int).SetInt64(32),
		TxIndex: 	6}
	b :=  &Candidate{
		CandidateId:  "2",
		Deposit: new(big.Int).SetUint64(99),
		BlockNumber: new(big.Int).SetInt64(32),
		TxIndex: 	7}

	c :=  &Candidate{
		CandidateId:  "3",
		Deposit: new(big.Int).SetUint64(100),
		BlockNumber: new(big.Int).SetInt64(31),
		TxIndex: 	6}

	d :=  &Candidate{
		CandidateId:  "4",
		Deposit: new(big.Int).SetUint64(100),
		BlockNumber: new(big.Int).SetInt64(32),
		TxIndex: 	5}

	e := &Candidate{
		CandidateId:  "5",
		Deposit: new(big.Int).SetUint64(100),
		BlockNumber: new(big.Int).SetInt64(32),
		TxIndex: 	4}

	f :=  &Candidate{
		CandidateId:  "6",
		Deposit: new(big.Int).SetUint64(80),
		BlockNumber: new(big.Int).SetInt64(30),
		TxIndex: 	3}

	g :=  &Candidate{
		CandidateId:  "7",
		Deposit: new(big.Int).SetUint64(100),
		BlockNumber: new(big.Int).SetInt64(31),
		TxIndex: 	3}


	arr := []*Candidate{a, b, c, d, e, f, g}
	ids := make([]string, 0)
	canMap := make(map[string]*Candidate, 0)
	for _, can := range arr {
		ids = append(ids, can.CandidateId)
		canMap[can.CandidateId] = can
	}


	candidateSort(ids, canMap)
	fmt.Println("%+v", arr)  // [0xc04206c180 0xc04206c090 0xc04206c0f0 0xc04206c120 0xc04206c150 0xc04206c0c0]

	for _, v := range ids {
		fmt.Println("%v", v, "deposit:", canMap[v].Deposit, "blockNum:", canMap[v].BlockNumber.String(), "txIndex:", canMap[v].TxIndex)
	}
}


type Candidate struct {

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



func compare(c, can *Candidate) int {
	// 质押金大的放前面
	if c.Deposit.Cmp(can.Deposit) > 0 {
		return 1
	}else if c.Deposit == can.Deposit {
		// 块高小的放前面
		if i := c.BlockNumber.Cmp(can.BlockNumber); i > 0 {
			return -1
		}else if i == 0 {
			// tx index 小的放前面
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

// 候选人排序
func candidateSort(arr []string, candidates map[string]*Candidate) {
	quickRealSort(arr, candidates, 0, len(arr) - 1)
}
func quickRealSort (arr []string, candidates map[string]*Candidate, left, right int)  {
	if left < right {
		pivot := partition(arr, candidates, left, right)
		quickRealSort(arr, candidates, left, pivot - 1)
		quickRealSort(arr, candidates, pivot + 1, right)
	}
}
func partition(arr []string, candidates map[string]*Candidate, left, right int) int {
	for left < right {
		for left < right && compare(candidates[arr[left]], candidates[arr[right]]) >= 0 {
			fmt.Println(candidates[arr[left]].CandidateId, "Cmp",  candidates[arr[right]])
			right --
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left ++
		}
		for left < right && compare(candidates[arr[left]], candidates[arr[right]]) >= 0 {
			fmt.Println(candidates[arr[left]].CandidateId, "Cmp",  candidates[arr[right]])
			left ++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			right --
		}
	}
	return left
}

type CandidateSort struct {
	data 		[]*Candidate
	myLessFunc 	func (x, y *Candidate) bool  // overwrite less func of sort
}



func (self CandidateSort) Len() int {
	return len(self.data)
}

func (self CandidateSort) Less(i, j int) bool {
	return self.myLessFunc(self.data[i], self.data[j])
}

func (self CandidateSort) Swap(i, j int) {
	self.data[i], self.data[j] = self.data[j], self.data[i]
}