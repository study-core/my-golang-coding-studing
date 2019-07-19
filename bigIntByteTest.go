package main

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/ethdb"
	"github.com/PlatONnetwork/PlatON-Go/x/staking"
	"math/big"
	"strings"
)

func main() {



	amount1, _ := new(big.Int).SetString("9000000000000000000000000", 10)
	amount2, _ := new(big.Int).SetString("60000000000000000000000000", 10)
	amount3, _ := new(big.Int).SetString("1300000000000000000", 10)

	sub1 := len(math.MaxBig256.Bytes()) - len(amount1.Bytes())
	sub2 := len(math.MaxBig256.Bytes()) - len(amount2.Bytes())
	sub3 := len(math.MaxBig256.Bytes()) - len(amount3.Bytes())

	ar1 := make([]byte, sub1)
	ar2 := make([]byte, sub2)
	ar3 := make([]byte, sub3)


	res1 := append(ar1, amount1.Bytes()...)
	res2 := append(ar2, amount2.Bytes()...)
	res3 := append(ar3, amount3.Bytes()...)


	db, _ := ethdb.NewLDBDatabase("/home/gavin/testSort", 1000, 1000)


	prefix := staking.CanPowerPrefixStr

	db.Put(append([]byte(prefix), res1...), []byte("9000000000000000000000000"))
	db.Put(append([]byte(prefix), res2...), []byte("60000000000000000000000000"))
	db.Put(append([]byte(prefix), res3...), []byte("1300000000000000000"))

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
			fmt.Print("key:=", key, "\t\t")
			fmt.Println("Value:=", string(iter.Value()))
			num ++
		}

		count ++
	}
	fmt.Println("耗时:", timer.End(), "ms")
	fmt.Println("最后遍历了:", count, "次！！", "\n有Gavin前缀的kv只有:", num, "个！！")


	prefix = "bytes"

	b1 := common.Uint32ToBytes(uint32(456))
	b2 := common.Uint32ToBytes(uint32(2))
	b3 := common.Uint32ToBytes(uint32(99999))

	db.Put(append([]byte(prefix), b1...), []byte("456"))
	db.Put(append([]byte(prefix), b2...), []byte("2"))
	db.Put(append([]byte(prefix), b3...), []byte("99999"))


	iter2 := db.NewIteratorWithPrefix([]byte(prefix))

	/*
		耗时: 467.933146 ms
		最后遍历了: 1089736 次！！
		有Gavin前缀的kv只有: 20002 个！！
	*/
	//iter := db.NewIterator()





	for iter2.Next() {
		key := string(iter2.Key())
		if strings.Contains(key, prefix){
			fmt.Print("key:=", key, "\t\t")
			fmt.Println("Value:=", string(iter2.Value()))
		}

	}
}
