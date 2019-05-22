package main

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
)

func main() {

	privateKey, _ := crypto.GenerateKey()

	addr := crypto.PubkeyToAddress(privateKey.PublicKey)

	nodeId := discover.PubkeyID(&privateKey.PublicKey)

	pubKey, _ := nodeId.Pubkey()

	a := /*privateKey.PublicKey.Curve == pubKey.Curve &&*/ privateKey.PublicKey.X.Cmp(pubKey.X) == 0 && privateKey.PublicKey.Y.Cmp(pubKey.Y) == 0


	fmt.Println("addr:=", addr.Hex(), "\nNodeId:=", nodeId.String(), "\nis only publickey：", a)
}
