package main

import (
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"fmt"
	"encoding/json"
	"math/big"
)

func main() {


	delFunc := func(queue types.CandidateQueue) types.CandidateQueue {

		if len(queue) > 2 {

			temArr := (queue)[2:]
			queue = (queue)[:2]

			fmt.Println("tempArr len", len(temArr))
		}


		queueStr, _ := json.Marshal(queue)

		fmt.Println("queue len", len(queue), "queue", string(queueStr))

		return queue

	}


	a := &types.Candidate{
		Deposit:  big.NewInt(0),
	}

	b := &types.Candidate{
		Deposit:  big.NewInt(1),
	}

	c := &types.Candidate{
		Deposit:  big.NewInt(2),
	}

	d := &types.Candidate{
		Deposit:  big.NewInt(3),
	}

	arr := make(types.CandidateQueue, 0)
	arr = append(arr, a)
	arr = append(arr, b)
	arr = append(arr, c)
	arr = append(arr, d)

	arr = delFunc(arr)

	arrStr, _ := json.Marshal(arr)

	fmt.Println("arr len", len(arr), "arr", string(arrStr))
}



