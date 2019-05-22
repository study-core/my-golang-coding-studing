package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := big.NewInt(10)

	b := big.NewInt(6)

	c := new(big.Int).Quo(a, b)

	c2 := new(big.Int).Div(a, b)

	//c := new(big.Int).And(a, b)   //10 && 3 == 0110 && 0011 == 0010 == 2

	//c := new(big.Int).AndNot(a, b)
	//
	//c := new(big.Int).Div(a, b)
	//
	//c, d := new(big.Int).DivMod(a, b, big.NewInt(1))
	//
	//c := new(big.Int).Exp(a, b, big.NewInt(1))
	//
	//c := new(big.Int).Binomial(10, 3)
	//
	//c := new(big.Int).GCD(a, b, big.NewInt(1), big.NewInt(2))
	//
	//c := new(big.Int).MulRange(10, 3)
	//



	fmt.Print(c, c2)
}
