package main

import (
	"time"
	"fmt"
)

func main() {


	timer := time.NewTimer(0)
	// 丢弃初始刻度
	<-timer.C

	fmt.Println("a")
}
