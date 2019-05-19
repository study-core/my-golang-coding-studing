package main

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"fmt"
	"encoding/hex"
	"bytes"
)

func main() {

	valHash := common.BytesToHash([]byte{})

	//if common.BytesToHash([]byte{}) == valHash {
	if (valHash == common.Hash{}) {
		fmt.Println(valHash)
	}
	// 0000000000000000000000000000000000000000000000000000000000000000
	// 0000000000000000000000000000000000000000000000000000000000000000
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println(common.Hash{}.Bytes())
	fmt.Println(hex.EncodeToString(common.Hash{}.Bytes()))

	fmt.Println(bytes.Equal(common.Hash{}.Bytes(), common.Hash{}.Bytes()))
}
