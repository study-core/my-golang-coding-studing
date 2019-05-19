package main

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"encoding/json"
	"fmt"
)

func main() {

	mb := &MB{
		Name: 		"lala",
		TOwner: 	common.HexToAddress("0x493301712671ada506ba6ca7891f436d29185821"),
		THash: 		common.HexToHash("0x353b241a560c665f57597b4840e5e0f5f97f5d4050971e66b754ccaea654f5e7"),
	}

	mbStr1, _ := json.Marshal(mb)

	mbCopy := *mb

	mbCopy.THash = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")

	mbStr2, _ := json.Marshal(mb)
	mbCopyStr, _ := json.Marshal(&mbCopy)


	fmt.Println(string(mbStr1))
	fmt.Println(string(mbStr2))
	fmt.Println(string(mbCopyStr))
}


type MB struct {
	Name 	string
	TOwner 	common.Address
	THash 	common.Hash
}

