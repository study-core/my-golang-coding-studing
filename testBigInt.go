package main

import (
	"fmt"
	"math/big"
)

func main() {

	a :=  new(big.Int).SetInt64(int64(250))
	b :=  new(big.Int).SetInt64(int64(1))
	c :=  new(big.Int).SetInt64(int64(230))
	d :=  new(big.Int).SetInt64(int64(251))
	e :=  new(big.Int).SetInt64(int64(500))


	f2, g2 := new(big.Int).DivMod(a, a, new(big.Int)) // 250/250  1 0

	fmt.Println(f2, g2)


	f, g := new(big.Int).DivMod(b, a, new(big.Int)) // 1/250   0 1

	fmt.Println(f, g)


	f4, g4 := new(big.Int).DivMod(c, a, new(big.Int)) // 230/250  0 230

	fmt.Println(f4, g4)


	f3, g3 := new(big.Int).DivMod(d, a, new(big.Int)) // 251/250  1 1

	fmt.Println(f3, g3)


	f5, g5 := new(big.Int).DivMod(e, a, new(big.Int)) // 500/250  2 0

	fmt.Println(f5, g5)


	target := big.NewInt(254)
	tmp := big.NewInt(45)
	fmt.Println(target, tmp)
	target = mergeAmount(uint8(0), target, tmp)
	fmt.Println(target, tmp)
	target = mergeAmount(uint8(1), target, tmp)
	fmt.Println(target, tmp)
}


func mergeAmount (mark uint8, target, tmp *big.Int) *big.Int {
	if mark == 0 {
		target = new(big.Int).Add(target, tmp)
		return target
	}else {
		target = new(big.Int).Sub(target, tmp)
		return target
	}
}