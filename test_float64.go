package main

import (
	"bytes"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
)

func main() {

	//
	//a := float64(69947760.000 )
	//b := float64(4)
	//
	//c := b/a
	//fmt.Println(c)
	//log.Info("ss", "p", c)


	a := common.HexToAddress("0x9370379Fb8AdbD741D666F584d0fcDC15b1B4F9f")

	b := common.HexToAddress("0x9370379fb8adbd741d666f584d0fcdc15b1b4f9f")

	fmt.Println(bytes.Equal(a.Bytes(), b.Bytes()))
	fmt.Println(a == b)

}
