package main

import (
	"encoding/hex"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"github.com/PlatONnetwork/PlatON-Go/x/staking"
	"github.com/PlatONnetwork/PlatON-Go/x/xutil"
)

func main() {

	nodeId := discover.MustHexID("38953e713af7d4c20b6b26b1e6413b387efdb26e559f14ae65381e1b9de563d0b15f6f50e9baf5ba4a8b81216ea4fe4ae0830c916f0e58520d1194e7a7843337")

	canAddr, _ := xutil.NodeId2Addr(nodeId)


	key := staking.CandidateKeyByAddr(canAddr)
	fmt.Println("DelCandidateStore key", hex.EncodeToString(key))
	// 43616e6d7c4f102996913ad73470f7cf1c05adda1fb5ac
	// 43616e6d7c4f102996913ad73470f7cf1c05adda1fb5ac


	/*func Is_Valid(status uint32) bool {
		return !Is_Invalid(status)
	}

	func Is_Invalid(status uint32) bool {
		return status&Invalided == Invalided
	}*/

	var status uint32
	status |= 0
	fmt.Println(staking.Is_Invalid(status))
	fmt.Println(!staking.Is_Invalid(status))
	fmt.Println(staking.Is_Valid(status))
}
