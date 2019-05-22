package main

import (
	"fmt"
	"math/big"
)

func main() {

	str := "1000000000000000000000000000"

	bi, _ := new(big.Int).SetString(str, 10)



	fmt.Print([]byte(bi.Bytes()))
}
