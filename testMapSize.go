package main

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto/sha3"
	"math/big"
	"strconv"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"fmt"
	//"sync"
	"time"
	"runtime"
)

func main() {





	ticketMap := make(map[common.Hash]*types.Ticket, 0)

	txHash := common.HexToHash("0xe894b9d0fe7abab6ef0dd1f5422a474b394c2cda2cb2859f642d32763b53ba6a")

	owner := common.HexToAddress("0xc1f330b214668beac2e6418dd651b09c759a4bf5")

	deposit := big.NewInt(1000000000000000000)

	nodeId := discover.MustHexID("0xa188edb6776931b5f18e228028aaab0d57217772753ac8d5bdaae585a4440cc94520c3b6f617c5cf60725893bc04326c87b5211d4b1d6c100dfc09f2c70917d8")

	blockNumber := big.NewInt(10000000000000000)

	cicle := 10000000

	//flag := make(chan int, cicle)
	//
	//var wg sync.WaitGroup
	//wg.Add(cicle)
	//
	//lock := sync.Mutex{}

	for index := 1; index <= cicle; index ++ {


		//go func(i int) {
			// generate ticket id
			value := append(txHash.Bytes(), []byte(strconv.Itoa(int(index)))...)
			ticketId := sha3.Sum256(value[:])

			ticket := &types.Ticket{
				TicketId:    ticketId,
				Owner:       owner,
				Deposit:     deposit,
				CandidateId: nodeId,
				BlockNumber: blockNumber,
			}
			//lock.Lock()
			ticketMap[ticketId] = ticket
			//lock.Unlock()
			//flag <- 0
			//wg.Done()
		//}(index)
	}

	//wg.Wait()
	//close(flag)


	fmt.Println("Map len", len(ticketMap))

	printMemStats()

	time.Sleep(10000000* time.Second)
}


func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("HeapAlloc = %v HeapIdel= %v HeapSys = %v  HeapReleased = %v\n", m.HeapAlloc/1024, m.HeapIdle/1024, m.HeapSys/1024,  m.HeapReleased/1024)
}
