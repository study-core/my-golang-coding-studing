package main

import (
	"fmt"
	"time"
)

func main() {


	now := time.Now()

	if then, err := time.Parse("2006-01-02 15:04:05", "2019-05-30 12:00:00"); nil != err {
		fmt.Println(err.Error())
	}else {
		fmt.Println(now.After(then))
	}


}
