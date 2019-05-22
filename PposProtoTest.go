package main

import (
	"github.com/PlatONnetwork/PlatON-Go/core/ppos_storage"
	"github.com/golang/protobuf/proto"
	"math/big"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"strconv"
	"time"
	"sync"
	"fmt"
	"crypto/md5"
	"encoding/json"
)


func main() {

	nodeId := discover.MustHexID("0x01234567890121345678901123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345")


	pposStorage := ppos_storage.NewPPOS_storage()

	pposStorage.SetTotalRemain(51200)

	queue := make(types.CandidateQueue, 0)

	can := &types.Candidate{
		Deposit: big.NewInt(120),
		BlockNumber: big.NewInt(120000),
		TxIndex: 	21,
		CandidateId: 	nodeId,

		Host: 	"120.0.0.1",
		Port: 	"6789",

		Owner: 	common.HexToAddress("0x12"),

	}

	queue = append(queue, can)
	queue = append(queue, can)
	queue = append(queue, can)

	pposStorage.SetCandidateQueue(queue, ppos_storage.PREVIOUS)
	pposStorage.SetCandidateQueue(queue, ppos_storage.CURRENT)
	pposStorage.SetCandidateQueue(queue, ppos_storage.NEXT)
	pposStorage.SetCandidateQueue(queue, ppos_storage.IMMEDIATE)
	pposStorage.SetCandidateQueue(queue, ppos_storage.RESERVE)



	for i := 0; i < 51200; i++ {

		//now := time.Now().UnixNano()


		txHash := common.Hash{}
		txHash.SetBytes(crypto.Keccak256([]byte(strconv.Itoa(time.Now().Nanosecond() + i))))


		count := uint32(i) + uint32(time.Now().UnixNano())
		price := big.NewInt(int64(count))

		//pposStorage.SetExpireTicket(blockNumber, txHash)
		pposStorage.AppendTicket(nodeId, txHash, count, price)
	}


	blockNumber := big.NewInt(12)

	blockHash := common.Hash{}
	blockHash.SetBytes(crypto.Keccak256([]byte(strconv.Itoa(time.Now().Nanosecond() + 12))))



	var data []byte

	if ppos_temp := buildPBStorage(blockNumber, blockHash, queue); nil == ppos_temp {
		fmt.Println("Call Commit2DB FINISH !!!! , PPOS storage is Empty, do not write disk AND direct short-circuit ...")
		return
	}else{

		d, _ := json.Marshal(ppos_temp)
		fmt.Println("组装数据:", string(d))

		// write ppos_storage into disk with protobuf
		if da, err := proto.Marshal(ppos_temp); nil != err {
			fmt.Println("Failed to Commit2DB", "proto err", err)
			return
		}else {
			data = da
		}
	}

	fmt.Println("data len", len(data), "md5", md5.Sum(data))






	pb_pposTemp := new(ppos_storage.PB_PPosTemp)
	if err := proto.Unmarshal(data, pb_pposTemp); err != nil {
		fmt.Println("Failed to Call NewPPosTemp to Unmarshal Global ppos temp", "err", err)
		return
	}else {
		d, _ := json.Marshal(pb_pposTemp)
		fmt.Println("解析数据:", string(d))
	}
}




func buildPBStorage(blockNumber *big.Int, blockHash common.Hash, arr types.CandidateQueue) *ppos_storage.PB_PPosTemp {
	ppos_temp := new(ppos_storage.PB_PPosTemp)
	ppos_temp.BlockNumber = blockNumber.String()
	ppos_temp.BlockHash = blockHash.Hex()

	var empty int = 0  // 0: empty 1: no
	var wg sync.WaitGroup

	/**
	candidate related
	*/
	if len(arr) != 0  {

		canTemp := new(ppos_storage.CandidateTemp)


		wg.Add(5)
		// previous witness
		go func() {
			if queue := buildPBcanqueue(arr); len(queue) != 0 {
				canTemp.Pres = queue
				empty |= 1
			}
			wg.Done()
		}()
		// current witness
		go func() {
			if queue := buildPBcanqueue(arr); len(queue) != 0 {
				canTemp.Currs = queue
				empty |= 1
			}
			wg.Done()
		}()
		// next witness
		go func() {
			if queue := buildPBcanqueue(arr); len(queue) != 0 {
				canTemp.Nexts = queue
				empty |= 1
			}
			wg.Done()
		}()
		// immediate
		go func() {
			if queue := buildPBcanqueue(arr); len(queue) != 0 {
				canTemp.Imms = queue
				empty |= 1
			}
			wg.Done()
		}()
		// reserve
		go func() {
			if queue := buildPBcanqueue(arr); len(queue) != 0 {
				canTemp.Res = queue
				empty |= 1
			}
			wg.Done()
		}()



		wg.Wait()
		ppos_temp.CanTmp = canTemp
	}



	if empty == 0 {
		return nil
	}
	return ppos_temp
}


func buildPBcanqueue (canQqueue types.CandidateQueue) []*ppos_storage.CandidateInfo {
	if len(canQqueue) == 0 {
		return nil
	}

	pbQueue := make([]*ppos_storage.CandidateInfo, len(canQqueue))
	for i, can := range canQqueue {
		canInfo := &ppos_storage.CandidateInfo{
			Deposit: 		can.Deposit.String(),
			BlockNumber:	can.BlockNumber.String(),
			TxIndex:		can.TxIndex,
			CandidateId:	can.CandidateId.Bytes(),
			Host:			can.Host,
			Port:			can.Port,
			Owner:			can.Owner.Bytes(),
			Extra:			can.Extra,
			TxHash: 		can.TxHash.Bytes(),
			TOwner: 		can.TOwner.Bytes(),
		}
		pbQueue[i] = canInfo
	}
	return pbQueue
}



