package main

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/ethdb"
	"strings"
)

func main() {


	db, _ := ethdb.NewLDBDatabase("/home/gavin/testLevelInterator", 1000, 1000)
	//
	//prefix := "Gavin"

	prefix := "Power"

	//db.Put([]byte("Emma"), []byte("Emma"))
	//
	//db.Put(append([]byte(prefix), []byte("Kally")...), []byte("Kally"))
	//
	//db.Put(append([]byte(prefix), []byte("ZhuYing")...), []byte("ZhuYing"))


	//
	//str := "0123456789abcdefghijklmnopqrstuvwxyz"
	//bytes := []byte(str)
	//
	//
	//for i := 0; i < 1000000; i++ {
	//	result := []byte{}
	//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//	for i := 0; i < 10; i++ {
	//		result = append(result, bytes[r.Intn(len(bytes))])
	//	}
	//
	//	value := []byte(result)
	//	key := value
	//	source := rand.NewSource(10)
	//	rd := rand.New(source)
	//	switch rd.Intn(10) {
	//	case 8, 9, 5, 2:
	//		db.Put(append([]byte(prefix), key...), value)
	//	default:
	//		db.Put(key, value)
	//	}
	//
	//}




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





	//fmt.Println(byteutil.BytesToString(byteutil.IntToBytes(12)))
	//fmt.Println(byteutil.BytesToInt([]byte("哈哈")))
}
