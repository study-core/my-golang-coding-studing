package main

import (
	"fmt"
	"github.com/yangzhou/PlatON-Go/common"
)

func main() {
	// *nil 是会引发 空指针的
	if *GetNilPtr() == (common.Address{}) {
		fmt.Println("empty")
	}
}


func GetNilPtr () *common.Address {
	return nil
}