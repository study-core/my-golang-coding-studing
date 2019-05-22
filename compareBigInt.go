package main

import (
	"math/big"
	"fmt"
)

func main() {
	a := new(big.Int).SetUint64(1)
	b := new(big.Int).SetUint64(5)
	c := new(big.Int).SetInt64(1)
	fmt.Println(a.Cmp(b)) // 1 cmp 5  < -1
	fmt.Println(a.Cmp(c)) // 1 cmp 1 == 0
	fmt.Println(b.Cmp(c)) // 5 cmp 1  > 1

}
