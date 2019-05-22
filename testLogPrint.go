package main

import (
	//"fmt"
	//"github.com/PlatONnetwork/PlatON-Go/core/types"
	//"math/big"
	//"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	//"github.com/PlatONnetwork/PlatON-Go/common"
	"fmt"
	//"encoding/json"
	//"errors"
	//"math/big"
	//"math/big"
	"encoding/json"
)

/*
func main() {
*/
/*	pending := make(trs)

	ts := make(types.Transactions, 0)

	t := &types.Transaction{
	}
	ts = append(ts, t)
	pending[common.Address{}] = ts


	fmt.Println(fmt.Sprintf("lala %+v", pending))
	fmt.Println(fmt.Sprintf("ss %+v", ts))*//*


	var g *Genesis = nil
	if g != nil {
		fmt.Println("Full")
	}
	g.configOrDefault()

	//fmt.Println(g.Timestamp)


}

type Genesis struct {
	Timestamp  uint64              `json:"timestamp"`

}

func (g Genesis) configOrDefault()  {
	fmt.Println("ss")
}

//type trs map[common.Address]types.Transactions
*/

//func main() {
//
//	//arr := make(types.CandidateQueue, 0)
//	var arr types.CandidateQueue
//	arr = nil
//
//
//
//	candidate := &types.Candidate{
//		Deposit: 		new(big.Int).SetUint64(100),
//		BlockNumber:    new(big.Int).SetUint64(7),
//		CandidateId:   discover.MustHexID("0x01234567890121345678901123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345"),
//		TxIndex:  		6,
//		Host:  			"10.0.0.1",
//		Port:  			"8548",
//		Owner: 			common.HexToAddress("0x12"),
//
//	}
//
//
//	candidate2 := &types.Candidate{
//		Deposit: 		new(big.Int).SetUint64(99),
//		BlockNumber:    new(big.Int).SetUint64(7),
//		CandidateId:   discover.MustHexID("0x01234567890121345678901123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012341"),
//		TxIndex:  		5,
//		Host:  			"10.0.0.1",
//		Port:  			"8548",
//		Owner: 			common.HexToAddress("0x15"),
//
//	}
//
//
//	candidate3 := &types.Candidate{
//		Deposit: 		new(big.Int).SetUint64(99),
//		BlockNumber:    new(big.Int).SetUint64(6),
//		CandidateId:   discover.MustHexID("0x01234567890121345678901123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012342"),
//		TxIndex:  		5,
//		Host:  			"10.0.0.1",
//		Port:  			"8548",
//		Owner: 			common.HexToAddress("0x15"),
//
//	}
//
//
//	candidate4 := &types.Candidate{
//		Deposit: 		new(big.Int).SetUint64(120), // 99
//		BlockNumber:    new(big.Int).SetUint64(6),
//		CandidateId:   discover.MustHexID("0x01234567890121345678901123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012343"),
//		TxIndex:  		4,
//		Host:  			"10.0.0.1",
//		Port:  			"8548",
//		Owner: 			common.HexToAddress("0x15"),
//
//	}
//
//	arr = append(arr, candidate)
//	arr = append(arr, candidate2)
//	arr = append(arr, candidate3)
//	arr = append(arr, candidate4)
//	fmt.Println("排序前")
//	for _, can := range arr {
//		b, _ := json.Marshal(can)
//		fmt.Println(can.CandidateId.String(), "==", string(b))
//	}
//
//
//	arr.CandidateSort()
//	fmt.Println("排序后")
//	for _, can := range arr {
//		b, _ := json.Marshal(can)
//		fmt.Println(can.CandidateId.String(), "==", string(b))
//	}
//}

/*
func main() {

	e1 := errors.New("a")

	if e1 == errors.New("a") {
		fmt.Println("A")
	}else {
		fmt.Println("B")
	}
}*/


const (

	/** about candidate pool */
	// immediate elected candidate
	ImmediatePrefix     = "id"
	ImmediateListPrefix = "iL"
	// reserve elected candidate
	ReservePrefix     = "rd"
	ReserveListPrefix = "rL"
	// previous witness
	PreWitnessPrefix     = "Pwn"
	PreWitnessListPrefix = "PwL"
	// witness
	WitnessPrefix     = "wn"
	WitnessListPrefix = "wL"
	// next witness
	NextWitnessPrefix     = "Nwn"
	NextWitnessListPrefix = "NwL"
	// need refund
	DefeatPrefix     = "df"
	DefeatListPrefix = "dL"

	/** about ticket pool */
	// Remaining number key
	SurplusQuantity		= "sq"
	// Expire ticket prefix
	ExpireTicket		= "et"
	// candidate attach
	CandidateAttach	= "ca"
	// Ticket pool hash
	TicketPoolHash	= "tph"

	PREVIOUS_C = iota -1
	CURRENT_C
	NEXT_C


)

func main() {
	//fmt.Println(PREVIOUS_C, CURRENT_C, NEXT_C)
	//
	// str := "130000000000000000000"
	// //ste := "130000000000000000000"
	//a, _ := new(big.Int).SetString(str, 10)
	////a := new(big.Int).SetBytes([]byte(str))
	//fmt.Println(a)
	//
	////ii := int(12)
	////stt := string(ii)
	////fmt.Println(stt)
	//fmt.Sprint(int(789))
	//1000000000000000000000000
	//999999999999999999999999

	//threshold, _ := new(big.Int).SetString("1000000000000000000000000", 10)
	//
	//fmt.Println(threshold.String())
	//
	//a, _ := new(big.Int).SetString("999999999999999999999999", 10)
	//
	//fmt.Println(a.String())
	//
	//b := new(big.Int).Add(big.NewInt(1), a)
	//fmt.Println(a.Cmp(threshold))
	//fmt.Println(threshold.Cmp(a))
	//fmt.Println(threshold.Cmp(b))


	b, _ := json.Marshal(nil)

	fmt.Println("b", string(b))
	/*
	var i interface{} = int(12)

	if a, ok := i.(int); ok {
		fmt.Println(a)
	}*/

}