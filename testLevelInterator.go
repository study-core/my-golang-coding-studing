package main

import (

	"fmt"
	"github.com/GavinXu520/platon-go/common"
	"github.com/PlatONnetwork/PlatON-Go/x/staking"
	"math/big"

	//"github.com/PlatONnetwork/PlatON-Go/x/staking"
	//"math/big"
	"math/rand"
	"strings"
	"time"

	//"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/ethdb"
)

func main() {


	db, _ := ethdb.NewLDBDatabase("/home/gavin/testLevelInterator", 1000, 1000)








	/*prefix := "Gavin"
	db.Put([]byte("Emma"), []byte("Emma"))

	db.Put(append([]byte(prefix), []byte("Kally")...), []byte("Kally"))

	db.Put(append([]byte(prefix), []byte("ZhuYing")...), []byte("ZhuYing"))

	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)


	for i := 0; i < 2000000; i++ {
		result := []byte{}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < 10; i++ {
			result = append(result, bytes[r.Intn(len(bytes))])
		}

		value := []byte(result)
		key := value
		source := rand.NewSource(10)
		rd := rand.New(source)
		switch rd.Intn(10) {
		case 8, 9, 5, 2:
			db.Put(append([]byte(prefix), key...), value)
		default:
			db.Put(key, value)
		}

	}
*/


	//prefix := "Power"
	//
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//for i := 0; i < 200; i++ {
	//	target := r.Intn(999)
	//
	//	block := r.Intn(99)
	//	//key := []byte(fmt.Sprint(target))
	//	val := 999 - target
	//
	//	//bl := 99 - block
	//
	//	keyBytes := make([]byte, 8)
	//	binary.BigEndian.PutUint64(keyBytes[:], uint64(val))
	//
	//	blockBytes := make([]byte, 8)
	//	binary.BigEndian.PutUint64(blockBytes, uint64(block))
	//
	//	db.Put(append(append([]byte(prefix), keyBytes...), blockBytes...), []byte(fmt.Sprint(val) + "_" + fmt.Sprint(target) + "_" + fmt.Sprint(block)))
	//}





	prefix := staking.CanPowerPrefixStr

	balanceStr := []string {

		"9000000000000000000000000",
		"60000000000000000000000000",
		"1300000000000000000000000",
		"1100000000000000000000000",
		"1000000000000000000000000",
		"4879000000000000000000000",
		"1800000000000000000000000",
		"1000000000000000000000000",
		"1000000000000000000000000",
		"70000000000000000000000000",
		"5550000000000000000000000",
		"44488850000000000000000000000",
		"650073899000000000000000000",

	}

	initProcessVersion := uint32(1<<16 | 0<<8 | 0) // 65536
	for i := 0; i < 1000; i ++ {

		var index int
		if i >= len(balanceStr) {
			index = i%(len(balanceStr)-1)
		}

		//t.Log("Create Staking num:", index)

		balance, _ := new(big.Int).SetString(balanceStr[index], 10)

		rand.Seed(time.Now().UnixNano())

		weight := rand.Intn(1000000000)


		balance = new(big.Int).Add(balance, big.NewInt(int64(weight)))

		key := staking.TallyPowerKey(balance, uint64(i), uint32(index), initProcessVersion)
		val := fmt.Sprint(initProcessVersion) + "_" + balance.String() + "_" + fmt.Sprint(i) + "_" + fmt.Sprint(index)
		db.Put(key, []byte(val))
	}





	timer := common.NewTimer()
	timer.Begin()
	iter := db.NewIteratorWithPrefix([]byte(prefix))

	/*
		耗时: 467.933146 ms
		最后遍历了: 1089736 次！！
		有Gavin前缀的kv只有: 20002 个！！
	*/
	//iter := db.NewIterator()

	count := 0
	num := 0
	for iter.Next() {
		key := string(iter.Key())
		if strings.Contains(key, prefix){
			//fmt.Print("key:=", key, "\t\t")
			fmt.Println("Value:=", string(iter.Value()))
			num ++
		}

		count ++
	}
	//fmt.Println("耗时:", timer.End(), "ms")
	//fmt.Println("最后遍历了:", count, "次！！", "\n有Gavin前缀的kv只有:", num, "个！！")





	//fmt.Println(byteutil.BytesToString(byteutil.IntToBytes(12)))
	//fmt.Println(byteutil.BytesToInt([]byte("哈哈")))
}
