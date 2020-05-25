package main

import (
	"fmt"
	"github.com/pangu/PlatON-Go/rlp"
)

func main() {


	a := &ARLP{
		Id: 12,
		Name: "sddd",
		Context: "#$$%%",
	}


 	abyte,_:= rlp.EncodeToBytes(a)


 	var aa ARLP
 	err := rlp.DecodeBytes(abyte, &aa)
	if nil != err {
		fmt.Println(err)
	}

	fmt.Println(aa)

 	var b BRLP
 	err = rlp.DecodeBytes(abyte, &b)
 	if nil != err {
 		fmt.Println(err)
	}

 	fmt.Println(b)




	StakingWeight := [4]string{"S", "E", "456", "ERR"}

	powerByte, _ := rlp.EncodeToBytes(StakingWeight)

	var power [4]string
	rlp.DecodeBytes(powerByte, &power)
	fmt.Println(power)


}


type ARLP struct {
	Id 		uint32
	Name 	string
	Context string
}

type BRLP struct {
	Id 		uint32
	Name 	string
	Context string
}