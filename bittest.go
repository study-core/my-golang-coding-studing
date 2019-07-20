package main

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/x/xutil"
)

func main() {

	version := uint32(4<<16|5<<8|2)

	fmt.Println(version)

	fmt.Printf("version is %s \n", xutil.ProcessVerion2Str(version))

	ii := int(1)

	var bb YouInt

	bb =  ii
	fmt.Println(bb)


}


type YouInt = int