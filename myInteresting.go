package main 

import (
	"fmt"
	"regexp"
)

func main() {

	str := "htt://ccvsfas48.jpg" 
   	imgReg := regexp.MustCompile("(http|https):.*?.(jpg|gif|png|bmp|jpeg)")

	fmt.Print("啦啦" + fmt.Sprint(imgReg.MatchString(str)))
}


