package main

import (
	//"bytes"
	"fmt"
	//"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	//"math"

	//"math"

	//"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	//"github.com/PlatONnetwork/PlatON-Go/x/restricting"
	//"github.com/pangu/PlatON-Go/common"
	//"github.com/pangu/PlatON-Go/rlp"
	//"math/big"

	//"encoding/json"
	//"github.com/PlatONnetwork/PlatON-Go/common"
	//"github.com/PlatONnetwork/PlatON-Go/ethdb"
	//"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	//"github.com/PlatONnetwork/PlatON-Go/x/staking"
	//"math/big"
)

func main() {

	/*db, _ := ethdb.NewLDBDatabase("/home/gavin/testLevelInterator", 1000, 1000)

	can := staking.Candidate{
		NodeId:             discover.NodeID{},
		StakingAddress:     common.ZeroAddr,
		BenifitAddress:     common.ZeroAddr,
		StakingBlockNum:    uint64(15),
		StakingTxIndex:     uint32(47),
		StakingEpoch:       uint32(74),
		Shares:             big.NewInt(1442),
		Released:           big.NewInt(10),
		ReleasedHes:        big.NewInt(1442),
		RestrictingPlan:    big.NewInt(1442),
		RestrictingPlanHes: big.NewInt(1442),
		Status:             uint32(0),

		Description: 		staking.Description{
			NodeName:	"Anode",
			ExternalId: "xxxx",
			Website: 	"www.gavin.cn",
			Details: 	"This a good man",
		},
	}

	val, err := rlp.EncodeToBytes(can)
	if nil != err {
		fmt.Println("encode err:", err)
	}

	key := []byte("can")
	if err := db.Put(key, val); nil != err {
		fmt.Println("put db err:", err)
	}

	if valByte, err := db.Get(key); nil != err {
		fmt.Println("Get db err:", err)
	}else {
		var can staking.Candidate

		if err := rlp.DecodeBytes(valByte, &can); nil != err {
			fmt.Println("decode err:", err)
		}else {

			b, _ := json.Marshal(can)
			fmt.Println(string(b))
		}
	}*/

	//
	//aa := uint16(1000)
	//
	//aaByte, _ := rlp.EncodeToBytes(aa)
	//fmt.Println("aa rpl:", aaByte)
	//
	//
	//addr := common.HexToAddress("0x493301712671Ada506ba6Ca7891F436D29185821")
	//
	//addrByte, _ := rlp.EncodeToBytes(addr)
	//fmt.Println("Addr rlp:", addrByte)
	//
	//
	//nodeId := discover.MustHexID("0xb96194c3c48d7b94ccd4c782ce19e034cd9da00e1537e85aa1ed2791836a9ca03061c1c35463d21c07a6db5a388d97706f9edaa4535fe46b2f816fd7f4c1d962")
	//
	//nodeIdByte, _ := rlp.EncodeToBytes(nodeId)
	//fmt.Println("NodeId rlp:", nodeIdByte)



	//var plans []restricting.RestrictingPlan
	//
	//p1 := restricting.RestrictingPlan{
	//	Epoch:  12,
	//	Amount:  big.NewInt(45),
	//}
	//
	//p2 := restricting.RestrictingPlan{
	//	Epoch:  24,
	//	Amount:  big.NewInt(90),
	//}
	//
	//plans = append(plans, p1)
	//plans = append(plans, p2)
	//
	//pjson, _ := json.Marshal(plans)
	//fmt.Println("plan 元数据:", string(pjson))
	//
	//pByte, _ := rlp.EncodeToBytes(plans)
	//fmt.Println("Plans rlp:", pByte)
	//
	//
	//p1B, _ := rlp.EncodeToBytes(p1)
	//fmt.Println("P1 rlp:", p1B)
	//
	//epoch := uint64(12)
	//eb, _ := rlp.EncodeToBytes(epoch)
	//fmt.Println("epoch rlp:", eb)
	//
	//amount := big.NewInt(45)
	//ab, _ := rlp.EncodeToBytes(amount)
	//fmt.Println("amount rlp:", ab)
	//
	//
	//
	//p2B, _ := rlp.EncodeToBytes(p2)
	//fmt.Println("P2 rlp:", p2B)
	//
	//epoch2 := uint64(24)
	//eb2, _ := rlp.EncodeToBytes(epoch2)
	//fmt.Println("epoch2 rlp:", eb2)
	//
	//amount2 := big.NewInt(90)
	//ab2, _ := rlp.EncodeToBytes(amount2)
	//fmt.Println("amount2 rlp:", ab2)





	/*var plans []restricting.RestrictingPlan

	balance, _ := new(big.Int).SetString("50000000000000000000", 10)

	p1 := restricting.RestrictingPlan{
		Epoch:  1,
		Amount:  balance,
	}



	plans = append(plans, p1)

	pjson, _ := json.Marshal(plans)
	fmt.Println("plan 元数据:", string(pjson))

	pByte, _ := rlp.EncodeToBytes(plans)
	fmt.Println("Plans rlp:", pByte)


	p1B, _ := rlp.EncodeToBytes(p1)
	fmt.Println("P1 rlp:", p1B)

	epoch := uint64(1)
	eb, _ := rlp.EncodeToBytes(epoch)
	fmt.Println("epoch rlp:", eb)

	amount := balance
	ab, _ := rlp.EncodeToBytes(amount)
	fmt.Println("amount rlp:", ab)*/

	//var pp []restricting.RestrictingPlan
	//
	//rlp.DecodeBytes(pByte, &pp)
	//
	//ppjson, _ := json.Marshal(pp)
	//fmt.Println("反编译回来:", string(ppjson))








	//paramMap := make(map[string]interface{}, 0)
	//
	//paramMap["a"] = "a"
	//
	//paramMap["f"] = "f"
	//
	//paramMap["qeq"] = 1122
	//paramMap["e"] = 33
	//paramMap["b"] = 444
	//paramMap["c"] = 55
	//
	//a1, _ := json.Marshal(paramMap)
	//fmt.Println(a1, "\n", string(a1))


	//paramMap2 := make(map[string]interface{}, len(paramMap))
	//
	//
	//paramMap2["f"] = "f"
	//
	//paramMap2["b"] = 444
	//paramMap2["e"] = 33
	//
	//paramMap2["c"] = 55
	//paramMap2["a"] = "a"
	//paramMap2["qeq"] = 1122
	//
	//
	//a2, _ := json.Marshal(paramMap2)
	//fmt.Println(a2, "\n", string(a2))
	//
	//
	//fmt.Println(bytes.Equal(a1, a2))
	//
	//fmt.Println(fmt.Sprint(math.MaxUint32))
	//fmt.Println(fmt.Sprint(math.MaxUint32/(21000)))
	//
	//
	//resCh  := make(chan int, 10)
	//
	//for i := 0; i <4; i++ {
	//	resCh <- i
	//}
	//close(resCh)
	//
	//fmt.Println(len(resCh))
	//
	//
	//
	//
	//ar := []int{2, 8, 9}
	//for i := 0; i < len(ar); i++ {
	//	if ar[i] == 2 || ar[i] == 8 || ar[i] == 9 {
	//		ar = append(ar[:i], ar[i+1:]...)
	//		i--
	//	}
	//}
	//
	//fmt.Println(ar)
	//
	//
	//aq := []int{2, 8, 9}
	//for i, mem := range aq {
	//	if mem == 2 || mem == 8 || mem == 9 {
	//		aq = append(aq[:i], aq[i+1:]...)
	//	}
	//}
	//
	//fmt.Println(aq)



	type test struct {
		Name string
		Age  uint64
	}

	arr := make([]*test, 0)

	t1 := &test{
		Name: "Eee",
		Age:1,
	}

	t2 := &test{
		Name: "eee",
		Age: 2,
	}

	arr = append(arr, t1)
	arr = append(arr, t2)

	b, _ := rlp.EncodeToBytes(arr)
	fmt.Println(b)

	// go [204 197 131 69 101 101 1 197 131 101 101 101 2]
	// py \xcc\xc5\x83Eee\x01\xc5\x83eee\x02
	// py 0xcc 0xc5 0x83 0x45 0x65 0x65 0x1 0xc5 0x83 0x65 0x65 0x65 0x2
	// py [204 197 131 69 101 101 1 197 131 101 101 101 2]
	// py [204, 197, 131, 69, 101, 101, 1, 197, 131, 101, 101, 101, 2]

}
