package main

import (
	"myCryto-study/testCircleReference/referenceA"
	"myCryto-study/testCircleReference/referenceB"
	"time"
)

func main() {
	referenceB.BTest2()
	referenceA.ATest1()
	time.Sleep(2*time.Second)
}
