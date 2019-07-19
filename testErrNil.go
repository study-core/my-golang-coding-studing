package main

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"github.com/PlatONnetwork/PlatON-Go/x/xutil"
)

func main() {

	//if err, ok := nil.Interface().(error); ok {
	//	return nil, err
	//}
	//return result[0].Bytes(),


	nodeStr := "enode://a6ef31a2006f55f5039e23ccccef343e735d56699bde947cfe253d441f5f291561640a8e2bbaf8a85a8a367b939efcef6f80ae28d2bd3d0b21bdac01c3aa6f2f@test-sea.platon.network:16791"

	node, _ := discover.ParseNode(nodeStr)
	addr, _ := xutil.NodeId2Addr(node.ID)
	fmt.Println("addr", addr.String())


	url := "enode://0x7bae841405067598bf65e7260ca693a964316e752249c4970085c805dbee738fdb41fc434e96e2b65e8bf1db2f52f05d9300d04c1e6129c26cb5d0f214b49968@platon.network:16791"
	node2, _ := discover.ParseNode(url)
	addr2, _ := xutil.NodeId2Addr(node2.ID)
	fmt.Println("addr2", addr2.String())





}
