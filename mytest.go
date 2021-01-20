package main

import (
	//"fmt"
	//"math/big"
	"sync"
	"time"
	"fmt"
)

func main() {

	/*// 对比 copy 和 append 的效率

	subVersion := math.MaxInt32 - xutil.CalcVersion(1793)
	sortVersion := common.Uint32ToBytes(subVersion)

	shares, _ := new(big.Int).SetString("14444444777777", 10)
	priority := new(big.Int).Sub(math.MaxBig104, shares)

	b104Len := len(math.MaxBig104.Bytes())
	zeros := make([]byte, b104Len)
	prio := append(zeros, priority.Bytes()...)

	num := common.Uint64ToBytes(uint64(456))
	txIndex := common.Uint32ToBytes(uint32(52))

	indexPre := len(staking.CanPowerKeyPrefix)
	indexVersion := indexPre + len(sortVersion)
	indexPrio := indexVersion + len(prio)
	indexNum := indexPrio + len(num)
	size := indexNum + len(txIndex)

	start2 := time.Now()
	key2 := append(staking.CanPowerKeyPrefix, append(sortVersion, append(prio,
		append(num, txIndex...)...)...)...)
	duration2 := time.Since(start2).Nanoseconds()

	start1 := time.Now()
	key1 := make([]byte, size)
	copy(key1[:len(staking.CanPowerKeyPrefix)], staking.CanPowerKeyPrefix)
	copy(key1[indexPre:indexVersion], sortVersion)
	copy(key1[indexVersion:indexPrio], prio)
	copy(key1[indexPrio:indexNum], num)
	copy(key1[indexNum:], txIndex)
	duration1 := time.Since(start1).Nanoseconds()

	fmt.Println("key1:\n", fmt.Sprint(key1), "duration1:\n", duration1)
	fmt.Println("key2:\n", fmt.Sprint(key2), "duration2:\n", duration2)
	fmt.Println(bytes.Equal(key1, key2))*/

	// 对比 fmt.Sprint 和 strconv.Itoa 及 strconv.FormatInt 的效率

	//start := time.Now()
	//for i := uint32(0); i < 10000; i++ {
	//	fmt.Sprint(i)
	//}
	//fmt.Printf("fmt.Sprint, %d \n", time.Since(start).Nanoseconds())
	//
	//start2 := time.Now()
	//for i := uint32(0); i < 10000; i++ {
	//	strconv.Itoa(int(i))
	//}
	//fmt.Printf("strconv.Itoa, %d \n", time.Since(start2).Nanoseconds())
	//
	//start3 := time.Now()
	//for i := uint32(0); i < 10000; i++ {
	//	strconv.FormatInt(int64(i), 10)
	//}
	//fmt.Printf("strconv.FormatInt, %d \n", time.Since(start3).Nanoseconds())

	/*total, _ := new(big.Int).SetString("262215742000000000000000000", 10)
	blocks, _ := new(big.Int).SetString("15759500", 10)

	balance := new(big.Int).Div(total, blocks)

	fmt.Println("balance:", balance)

	MillionLAT, _ := new(big.Int).SetString("1000000000000000000000000", 10)

	num := new(big.Int).Div(MillionLAT, balance)

	fmt.Println("num:", num)*/


	var mu sync.Mutex
	   go func() {
		   fmt.Println("里面直接 lock")
		       mu.Lock()
		   fmt.Println("里面已经 lock")
		       time.Sleep(10 * time.Second)
		   fmt.Println("里面直接 unlock")
		       mu.Unlock()
		   }()
	   time.Sleep(time.Second)
	   fmt.Println("外面直接 unlock")
	    mu.Unlock()
	fmt.Println("外面已经 unlock")
	   select {}
}
