package main

import (
	"fmt"
	"math/bits"
)

func main() {

	// 0100 == 4
	// 1000 == 8
	// 1010 == 10
	// 1100 == 12
	// 1011 == 11
	// 1101 == 13

	l1 := bits.Len32(4)
	l2 := bits.Len32(8)
	l3 := bits.Len32(10)
	l4 := bits.Len32(12)
	l5 := bits.Len32(11)
	l6 := bits.Len32(13)

	fmt.Println(l1, l2, l3, l4, l5, l6)

	o1 := bits.OnesCount32(4)
	o2 := bits.OnesCount32(8)
	o3 := bits.OnesCount32(10)
	o4 := bits.OnesCount32(12)
	o5 := bits.OnesCount32(11)
	o6 := bits.OnesCount32(13)

	fmt.Println(o1, o2, o3, o4, o5, o6)
}
