package main

import (
	"math/big"
	"encoding/json"
	"fmt"
)

func main() {

	a := &Infor{
		Deposit: big.NewInt(12),
		Lis:  &little{
			Name: "lala",
		},
		Count: 12,
	}


	b := &Infor{
		Deposit: big.NewInt(13),
		Lis:  &little{
			Name: "hehe",
		},

		Count: 14,
	}


	list := make([]*Infor, 0)

	list = append(list, a, b)

	byte1, _ := json.Marshal(list)
	fmt.Println("copy前:", string(byte1))

	// copy

	slice := make([]*Infor, len(list))

	copy(slice, list)

	byte2, _ := json.Marshal(slice)

	fmt.Println("copy后:", string(byte2))

	// change

	slice[1].Lis.Name = "OO"
	slice[1].Deposit = big.NewInt( 16)
	slice[1].Count = 16

	btye3, _ := json.Marshal(list)

	byte4, _ := json.Marshal(slice)

	fmt.Println("修改后 list:", string(btye3))
	fmt.Println("修改后 slice:", string(byte4))



	/*inMap :=  make(map[string][]*Infor, 0)

	inMap["A"] = list

	lalaMap := make(map[string][]*Infor, 0)

	for k, queue := range inMap {
		fmt.Println("k", k)

		btye4, _ := json.Marshal(queue)
		fmt.Println("val", string(btye4))
		lalaMap[k] = queue
	}


	qu := lalaMap["A"]
	qu[1].Deposit = big.NewInt( 18)
	qu[1].Lis.Name = "Emma"

	for k, queue := range inMap {
		fmt.Println("k", k)

		btye5, _ := json.Marshal(queue)
		fmt.Println("val", string(btye5))
		lalaMap[k] = queue
	}

	for k, queue := range lalaMap {
		fmt.Println("k", k)

		btye6, _ := json.Marshal(queue)
		fmt.Println("val", string(btye6))
		lalaMap[k] = queue
	}*/
}



type Infor struct {
	Deposit *big.Int

	Lis 	*little

	Count  uint32

}


type little struct {
	Name string
}