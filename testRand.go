package main

import (
	"math/rand"
	"fmt"
)

func main() {

	aArr := []int{2, 1, 9, 5, 6, 7}
	bArr := rand.Perm(len(aArr))
	fmt.Println(bArr)
}
