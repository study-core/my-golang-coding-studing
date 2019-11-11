package main

import (
	"fmt"
	"my-study/linknameA"
	"my-study/linknameC"
)

func main() {
	fmt.Println(linknameA.Hello())
	fmt.Println(linknameC.IsSpace(' '))
	//fmt.Println(linknameC.Hello()) // 这个会报错
	linknameC.Say()

	var timeNow func() (int64, int32)
	err := forceexport.GetFunc(&timeNow, "time.now")
	if err != nil {
		// Handle errors if you care about name possibly being invalid.
	}
	// Calls the actual time.now function.
	sec, nsec := timeNow()
}
