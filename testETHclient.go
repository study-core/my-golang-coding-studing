package main

import (
	"log"

	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"myCryto-study/tokens"
)

var tokenAddr = common.HexToAddress("0x14501b5dD2d570d52aD4D7eA2D86f7db631c305C")
var toAddr = common.HexToAddress("0xf6c636aced08de7907ca6700437932e530da02c9")

func main() {
	c, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("dial err %v", err)
		return
	}
	filter, err := mytoken.NewMytokenFilterer(tokenAddr, c)
	if err != nil {
		log.Fatalf("new filter err %s", err)
	}
	//1. listen any event transfer  coming...
	ch := make(chan *mytoken.MytokenTransfer, 10)
	sub, err := filter.WatchTransfer(nil, ch, nil, []common.Address{toAddr})
	if err != nil {
		log.Fatalf("watch transfer err %s", err)
	}
	go func() {
		for {
			select {
			case <-sub.Err():
				return
			case e := <-ch:
				log.Printf("new transfer event from %s to %s value=%s,at %d",
					e.From.String(), e.To.String(), e.Value, e.Raw.BlockNumber)
			}
		}
	}()
	//2. get history of event transfer
	history, err := filter.FilterTransfer(&bind.FilterOpts{Start: 480000}, nil, []common.Address{toAddr})
	if err != nil {
		log.Fatalf("query history logs err %s", err)
	}
	for history.Next() {
		e := history.Event
		log.Printf("%s transfer to %s value=%s, at %d", e.From.String(), e.To.String(), e.Value, e.Raw.BlockNumber)
	}
	//transfer to this addr
	time.Sleep(time.Minute * 3)
	log.Printf("finished..")
	sub.Unsubscribe()
}