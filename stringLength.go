package main

import (
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/go-ethereum-analysis/common/hexutil"

	//"github.com/PlatONnetwork/PlatON-Go/common"
	//"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	//"github.com/PlatONnetwork/PlatON-Go/rlp"
	//"hash/fnv"
)
// 计算字符串长度
func main() {
	//str := "0000000000000000000000000000000000000000000000000000000000000F01"
	//str := "00000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000008000000000000000000000000000000000000000002000000000000000000004000000000020000000000000000000000000000000040000010000000000000000000000000000000000000000100000000000000000000000000000000000000010000000000000000000000000000000000000000008000000000000000000001000000000000000000000000000000000000000008000000000000000000008000000000000000000000010000000000000000000000000000000000000"
	//str := "070a8d6a982153cae4be29d434e8faef8a47b274a053f5a4ee2a6c9c13c31e5c 031b8ce914eba3a9ffb989f9cdd5b0f01943074bf4f0f315690ec3cec6981afc"
	//str := "070a8d6a982153cae4be29d434e8faef8a47b274a053f5a4ee2a6c9c13c31e5c"
	//str := "031b8ce914eba3a9ffb989f9cdd5b0f01943074bf4f0f315690ec3cec6981afc"
	//str := "0x1f3a8672348ff6b789e416762ad53e69063138b8eb4d8780101658f24b2369f1a8e09499226b467d8bc0c4e03e1dc903df857eeb3c67733d21b6aaee2840e429"
	//str := "290decd9548b62a8d60345a988386fc84ba6bc95"


	//hash := fnv.New64()
	//hash.Write([]byte("init"))
	//sum64 := hash.Sum64()
	//fmt.Println(sum64)
	////fmt.Println(len(str)) // 64 个char (一个char = 4bit； 2个char == 8bit == 1个字节)  64/2 == 32 byte
	////16840622013048855939
	////16840622013048855939
	//
	//
	//input, err := hexutil.Decode("0xd488e9b5e87ed4669d838a68656c6c6f776f726c64")
	//if nil != err {
	//	fmt.Println("err:", err)
	//}
	//
	//decodeFuncAndParams := func (input []byte) (uint64, []byte, error) {
	//	content, _, err := rlp.SplitList(input)
	//	if nil != err {
	//		return 0, nil, fmt.Errorf("failed to decode input funcName and params: %v", err)
	//	}
	//
	//	funcName, params, err := rlp.SplitString(content)
	//	if nil != err {
	//		return 0, nil, fmt.Errorf("failed to decode input funcName and params: %v", err)
	//	}
	//	return common.BytesToUint64(funcName), params, nil
	//
	//}
	//
	//
	//
	//funcName, _, err := decodeFuncAndParams(input)
	//if nil != err {
	//	fmt.Println("22 err:", err)
	//}
	//if funcName != sum64 {
	//	fmt.Println("false:", funcName)
	//}
	//
	//fmt.Println("funcName:", funcName, "\nsum64:", sum64, "funcName == sum64 ", funcName == sum64)


	//fmt.Println(17&1 == 1, !(17&1 == 1))
	//
	//fmt.Println(0&1 == 1, !(0&1 == 1))
	//
	//
	//nodeId, _ := discover.HexID("df672cbf413dc740036ef5bf2545180fe9309561189b22f220b77beae7fce61cba7faee821433eab0fce29b3a564288629cff5dee314195c33f601c165d1949a")
	//canAddr, _ := xutil.NodeId2Addr(nodeId)
	//fmt.Println(canAddr.String())


	//str := "0xf2900908010004000000000000001600000090024200400100000000000000057ffff58f39ccd3334ccccccccccccccd599998"
	str := "0xf2900f4240000400000000000000160000009003d090000100000000000000057fffed8f61a80000199999999999999a266664"

	b, _ := hexutil.Decode(str)

	content, _, err := rlp.SplitList(b)
	if nil != err {
		panic(err)
	}


	decodeTopics := func(b []byte) ([]byte, []byte, error) {
		member, rest, err := rlp.SplitString(b)
		if nil != err {
			panic(err)
		}
		return member, rest, nil
	}
	i := 1
	for len(content) > 0 {
		mem, tail, err := decodeTopics(content)
		if nil != err {
			panic(err)
		}
		//m := common.BytesToUint64(mem)

		fmt.Println("############## 第", i, "个 data", "v", mem, "hex:", hexutil.Encode(mem), )

		content = tail
		i++
	}










	type inner struct {
		Name string
		Age  uint64
	}

	type outer struct {
		Inner  *inner
		Desc  string
		Number  uint32
	}


	o := &outer{
		Inner:  &inner{
			Name: "我是里层 struct 的名称",
			Age: 12,
		},
		Desc:   "我是外层struct",
		Number: 0,
	}

	jsonByte, _ := json.Marshal(o)
	fmt.Println(string(jsonByte))





}
