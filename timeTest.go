package main

import (
	"fmt"
	"time"
)

func main() {
	duration := uint64(60 * (time.Second.Nanoseconds()))
	fmt.Println(duration)
	fmt.Println(time.Second.Nanoseconds())
	fmt.Println(time.Second.String())


}
