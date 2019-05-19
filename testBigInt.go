package main

import (
	"math/big"
	"fmt"
)

func main() {

	start := big.NewInt(10)
	end := big.NewInt(20)

	c := new(big.Int).Sub(end, start)

	fmt.Println("start", start.String(), "end", end.String(), "c", c.String())

	end.Sub(end, start)
	fmt.Println("start", start.String(), "end", end.String(), "c", c.String())
}
