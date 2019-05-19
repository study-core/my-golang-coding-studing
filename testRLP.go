package main

import (
	"Platon-go/rlp"
	"Platon-go/p2p/discover"
	"fmt"
)

func main() {
	ids := make([]discover.NodeID, 0)
	if val, err := rlp.EncodeToBytes(ids); nil != err {
	} else {
		//state.StateDB{}.SetState(common.CandidateAddr, []byte("PwL"), val)
		fmt.Println(val)
		var aa []discover.NodeID
		rlp.DecodeBytes(val, &aa)
		fmt.Println(aa)
	}


	//b := []byte{192}

}
