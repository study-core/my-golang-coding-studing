package main

import (
	"fmt"

	//"github.com/PlatONnetwork/PlatON-Go/x/staking"
	"math/big"

	//"math/big"
)

func main() {

	//v_592, _ := new(big.Int).SetString("506f7765727ffff8ff00000000000000000000000000ffffff2c387869d4e4fbefffff000000000000025000000000", 16)
	//
	//v_582, _ := new(big.Int).SetString("506f7765727ffff8ff00000000000000000000000000ffffff2c3de43133125effffff000000000000024600000000", 16)
	//
	//v_01, _ := new(big.Int).SetString("506f7765727ffff8ff00000000000000000000000000ffffff2c3de43133125effffff000000000000000000000001", 16)
	//
	//
	//arr := make([]int, 1, 1)
	//fmt.Println(len(arr), cap(arr))
	////arr[0] = 5
	////arr = append(arr, 6)
	//fmt.Println(arr)
	//a := arr[1:]
	//fmt.Println(a)
	//fmt.Println(len(a), cap(a), len(arr), cap(arr))
	//
	//fmt.Println(12/13)
	//fmt.Println(12%13)
	//
	//str := "15e87df4d48d1294c7ca6485cd42b5f5d66b9f5561fd0f78b091ad6a43613921cab55951a6aa9289cdb76f80d6c42dedc673e03ff7e4b3f1b996db12351f3109"
	//
	//fmt.Println(len(str))
	//
	//
	//fmt.Println(uint32(0<<16 | 7<<8 | 1))

	/**
	262215742.48691600000000000000  激励池账户    // 应该为 262215742.48691650
	259096240.41867400000000000000  锁仓合约账户   // 我自己计算出来的 259096240.41867359
	70000000.00000000000000000000   质押合约账户
	1608688017.09441000000000000000 PlatON基金会账户
	8050000000.00000000000000000000 剩余总账户
	*/

	/*reward_tmp, _ := new(big.Float).SetString("262215742.48691600000000000000")
	first_tmp, _ := new(big.Float).SetString("62215742.4869160") // 62215742.4869165
	reward := new(big.Float).Sub(reward_tmp, first_tmp)
	fmt.Println("激励池gene", reward.String())

	//lock_tmp, _ := new(big.Float).SetString("259096240.41867400000000000000")
	lock_tmp, _ := new(big.Float).SetString("259096240.41867359") // 用自己算的
	real_lock := new(big.Float).Add(lock_tmp, first_tmp)

	fmt.Println("全部锁仓的钱", real_lock.String())
	stake_tmp, _ := new(big.Float).SetString("70000000.00000000000000000000")

	add := new(big.Float).Add(real_lock, stake_tmp)

	fmt.Println("锁仓+stake 的钱:", add.String())

	platon_tmp, _ := new(big.Float).SetString("1608688017.09441000000000000000")

	platon := new(big.Float).Add(platon_tmp, add)

	fmt.Println("platon基金会", platon.String())
	remain, _ := new(big.Float).SetString("8050000000.00000000000000000000")

	a := new(big.Float).Add(platon, remain)
	total := new(big.Float).Add(a, reward)

	fmt.Println("总数:", total.String())
*/




	//zeroEpoch  := new(big.Int).Mul(big.NewInt(622157424869165), big.NewInt(1E11))
	oneEpoch   := new(big.Int).Mul(big.NewInt(559657424869165), big.NewInt(1E11))
	twoEpoch   := new(big.Int).Mul(big.NewInt(495594924869165), big.NewInt(1E11))
	threeEpoch := new(big.Int).Mul(big.NewInt(429930862369165), big.NewInt(1E11))
	fourEpoch  := new(big.Int).Mul(big.NewInt(362625198306666), big.NewInt(1E11))
	fiveEpoch  := new(big.Int).Mul(big.NewInt(293636892642603), big.NewInt(1E11))
	sixEpoch   := new(big.Int).Mul(big.NewInt(222923879336939), big.NewInt(1E11))
	sevenEpoch := new(big.Int).Mul(big.NewInt(150443040698633), big.NewInt(1E11))
	eightEpoch := new(big.Int).Mul(big.NewInt(761501810943690), big.NewInt(1E10))


	a := new(big.Int).Add(oneEpoch, twoEpoch)

	b := new(big.Int).Add(a, threeEpoch)

	c := new(big.Int).Add(b, fourEpoch)
	d := new(big.Int).Add(c, fiveEpoch)
	e := new(big.Int).Add(d, sixEpoch)
	f := new(big.Int).Add(e, sevenEpoch)
	g := new(big.Int).Add(f, eightEpoch)


	fmt.Println(g)

	//fmt.Println("zeroEpoch ", zeroEpoch)
	//fmt.Println("oneEpoch ", oneEpoch)
	//fmt.Println("twoEpoch ", twoEpoch)
	//fmt.Println("threeEpoch ", threeEpoch)
	//fmt.Println("fourEpoch ", fourEpoch)
	//fmt.Println("fiveEpoch ", fiveEpoch)
	//fmt.Println("sixEpoch ", sixEpoch)
	//fmt.Println("sevenEpoch ", sevenEpoch)
	//fmt.Println("eightEpoch ", eightEpoch)
	//
	//fmt.Println("staking", "10000000000000000000000000")
	//
	//
	//reward  := new(big.Int).Mul(big.NewInt(2622157424869165), big.NewInt(1E11))
	//lock   := new(big.Int).Mul(big.NewInt(2590962404186735), big.NewInt(1E11))
	//platon := new(big.Int).Mul(big.NewInt(160868801709441), big.NewInt(1E13))
	//remain := new(big.Int).Mul(big.NewInt(8050000000 ), big.NewInt(1E18))
	//
	//
	//
	//
	//
	//fmt.Println("激励池账户 ", reward)
	//fmt.Println("锁仓合约账户 ", lock)
	//fmt.Println("PlatON基金会账户 ", platon)
	//fmt.Println("剩余总账户 ", remain)

	//arr := []int{2, 5, 9}
	//
	//for i := 0; i < len(arr); i++ {
	//	if /*arr[i] == 2 ||  arr[i] == 5 ||*/  arr[i] == 9 {
	//		arr = append(arr[:i], arr[i+1:]...)
	//		i--
	//	}
	//}
	//fmt.Println(arr)
	//
	//status := 0
	//
	//status |= staking.Invalided | staking.Withdrew |staking.LowRatio | staking.LowRatioDel | staking.DuplicateSign
	//fmt.Println(status)
	//


	arr := []int{5, 2, 5, 9, 8, 9}

	arr = arr[0:]

	//vrfQueue := []int{4, 8, 7, 6, 2, 4, 7}
	//vrfQueue := []int{5, 9}
	vrfQueue := []int{}

	next := shuffleQueue(arr, vrfQueue)
	fmt.Println(next, len(next), cap(next))

}


func shuffleQueue(remainCurrQueue, vrfQueue []int) []int {

	remainLen := len(remainCurrQueue)
	totalQueue := append(remainCurrQueue, vrfQueue...)

	for remainLen > int(6-4) && len(totalQueue) > 6 {
		totalQueue = totalQueue[1:]
		remainLen--
	}

	if len(totalQueue) > 6 {
		totalQueue = totalQueue[:6]
	}

	next := make([]int, len(totalQueue))

	copy(next, totalQueue)
	return next
}