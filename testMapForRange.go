package main

import (
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"time"
	"sync"
)

func main() {

	start := common.NewTimer()
	start.Begin()

	refundMap := make(map[discover.NodeID]struct{}, 0)

	nodeId_1 := discover.MustHexID("0x01234567890121345678901123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345")
	refundMap[nodeId_1] = struct{}{}
	nodeId_2 := discover.MustHexID("0x01234567890121345678901123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012346")
	refundMap[nodeId_2] = struct{}{}
	nodeId_3 := discover.MustHexID("0x01234567890121345678901123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012347")
	refundMap[nodeId_3] = struct{}{}
	nodeId_4 := discover.MustHexID("0x01234567890121345678901123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012348")
	refundMap[nodeId_4] = struct{}{}

	time.Sleep(1 * time.Second)

	fmt.Println("build data", start.End())

	//
	//arr := make([]string, len(refundMap))
	//
	//repeatMap := make(map[string]discover.NodeID, len(refundMap))

	for nodeId := range refundMap {
		fmt.Println("nodeId", nodeId.String())
	}

	time.Sleep(1 * time.Second)
	fmt.Println("handle finish", start.End())


	myM := sync.Map{}
	myM.Store("ss", "ss")
	myM.Store("bb", "bb")

	myM.Range(func(k, v interface{}) bool {
			fmt.Println("k", k, "v", v)
		return true
		})

}
