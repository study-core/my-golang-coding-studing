package main

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
)

//
//import (
//	"fmt"
//
//	//"github.com/PlatONnetwork/PlatON-Go/common"
//	"github.com/PlatONnetwork/PlatON-Go/rlp"
//	//"math/big"
//	//"reflect"
//)
//
////
////import (
////	"fmt"
////	"github.com/go-ethereum-analysis/rlp"
////)
////
////func main() {
////
////	type Params struct {
////		FuncName string
////		//Args    interface{}
////		Args    [] interface{}
////	}
////
////	type Args struct {
////		Name   string
////		Gender uint64
////		Age    uint64
////	}
////
////	type Input struct {
////		File string
////		Des  string
////		Num  uint64
////		Page uint64
////	}
////
////	input := &Input{
////		File: "Input",
////		Des:  "This is input",
////		Num:  34,
////		Page: 43,
////	}
////
////	ar := &Args{
////		Name:   "Args",
////		Gender: 23,
////		Age:    32,
////	}
////
////	inputB, _ := rlp.EncodeToBytes(input)
////	arB, _ := rlp.EncodeToBytes(ar)
////
////	//var p1, p2 Params
////
////	//if err := rlp.DecodeBytes(inputB, &p1); nil != err {
////	//	fmt.Println("input err:", err)
////	//}
////	//if err := rlp.DecodeBytes(arB, &p2); nil != err {
////	//	fmt.Println("args err:", err)
////	//}
////
////	kind, content, _, err := rlp.Split(inputB)
////
////	switch {
////	case err != nil:
////		fmt.Println(err)
////	case kind != rlp.List:
////		fmt.Println("input type error")
////	}
////
////	//name, _, err := rlp.SplitString(content)
////	_, name, _, err := rlp.Split(content)
////	if nil != err {
////		fmt.Println("input err", err)
////	}else {
////		fmt.Println("input name", string(name), "Input" == string(name))
////	}
////
////
////
////	kind, content, _, err = rlp.Split(arB)
////
////	switch {
////	case err != nil:
////		fmt.Println(err)
////	case kind != rlp.List:
////		fmt.Println("args type error")
////	}
////
////	_, name, _, err = rlp.Split(content)
////	if nil != err {
////		fmt.Println("args err", err)
////	}else {
////		fmt.Println("args name", string(name), "Args"== string(name))
////	}
////}
//
//func main() {
//	//type Message struct {
//	//	Head string
//	//	Body string
//	//	End  string
//	//}
//	//
//	//m := Message{
//	//	Head: "Gavin",
//	//	Body: "I am gavin",
//	//	End:  "finished",
//	//}
//	////arr := []Message{m}
//	//arr  := make([]*Message, 0)
//	//arr = append(arr, &m)
//	//b, _ := rlp.EncodeToBytes(arr)
//	//
//	////b, _ := rlp.EncodeToBytes(m)
//	//fmt.Println(b)
//	//
//	///** [218 133 71 97 118 105 110
//	//合约返回的: 				  [215 214 193 71 138 73 32 97 109 32 103 97 118 105 110 136 102 105 110 105 115 104 101 100]
//	//自己解出来: [219 218 133 71 97 118 105 110 138 73 32 97 109 32 103 97 118 105 110 136 102 105 110 105 115 104 101 100]
//	//*/
//
//
//	/*m := make(map[string]string, 0)
//	m["a"] = "A"
//
//	bm, err := rlp.EncodeToBytes(m)
//	if nil != err {
//		fmt.Println("rlp encode err:", err)
//	}
//
//	var mm map[string]string
//	err = rlp.DecodeBytes(bm, &mm)
//	if nil != err {
//		fmt.Println("rlp decode err:", err)
//	}
//
//
//	bmm, err := json.Marshal(mm)
//	if nil != err {
//		fmt.Println("json err:", err)
//	}
//
//
//	fmt.Println(string(bmm))*/
//
//
//
//	//type message struct {
//	//	Age uint64
//	//	Name string
//	//	Num *big.Int
//	//	Balance uint32
//	//}
//	//
//	//
//	//m := message{
//	//	Balance: 45,
//	//	Age: 34,
//	//	Name: "I love gavin",
//	//	Num: common.Big2,
//	//}
//
//
//	type mm struct {
//		N string
//	}
//
//	type message struct {
//
//		T1 string
//		//MM *mm
//		T2 string
//		T3 string
//	}
//
//	m := message{
//		T1: "t1",
//		T2: "t2",
//		T3: "t3",
//		//MM: &mm{N: "ss"},
//	}
//
//
//	b, err := rlp.EncodeToBytes(m)
//	if nil != err {
//		fmt.Println("rlp err:", err)
//	}
//
//
//	content, _, err := rlp.SplitList(b)
//
//
//	num, err := rlp.CountValues(content)
//	if nil != err {
//		fmt.Println("rlp CountValues err:", err)
//	} else {
//		fmt.Println("字段个数为: ", num)
//	}
//
//
//	decodeTopics := func(b []byte) ([]byte, []byte, error) {
//		member, rest, err := rlp.SplitString(b)
//		switch {
//		case err != nil:
//			return nil, nil, err
//		}
//		return member, rest, nil
//	}
//
//	for len(content) > 0 {
//		mem, tail, err := decodeTopics(content)
//		if nil != err {
//			panic(err)
//		}
//
//		fmt.Println("type: String, member:", string(mem), ", mem: ", mem)
//		content = tail
//	}
//
//
//
//}
//
//func decodeFuncAndParams(input []byte, i int) (rlp.Kind, []byte, []byte, error) {
//	fmt.Println("input:", input)
//	if i == 0 {
//		kind, content, _, err := rlp.Split(input)
//		switch {
//		case err != nil:
//			return kind, nil, nil, err
//		case kind != rlp.List:
//			return kind, nil, nil, fmt.Errorf("input type error")
//		}
//		input = content
//	}
//
//
//	//count, err := rlp.CountValues(input)
//	//if nil != err {
//	//	fmt.Println("获取整个list的元素个数: err", err)
//	//} else {
//	//	fmt.Println("获取整个list的元素个数: ", count)
//	//}
//
//
//	kind, member, rest, err := rlp.Split(input)
//
//	//fmt.Println("member:", common.BytesToUint64(member), "str:", string(member), "rest:", rest)
//	switch {
//	case err != nil:
//		return kind, nil, nil, err
//	}
//	return kind, member, rest, nil
//
//}

func main() {
	str := "493301712671ada506ba6ca7891f436d29185821"
	myint := uint32(246)

	e1Name := "mystr"
	e2Name := "myint"

	addr := common.HexToAddress("493301712671ada506ba6ca7891f436d29185821")

	strRlp, _ := rlp.EncodeToBytes(str)
	str256 := crypto.Keccak256(strRlp)

	e1NameRlp, _ := rlp.EncodeToBytes(e1Name)
	e1Name256 := crypto.Keccak256(e1NameRlp)

	e2NameRlp, _ := rlp.EncodeToBytes(e2Name)
	e2Name256 := crypto.Keccak256(e2NameRlp)


	addrRlp, _ := rlp.EncodeToBytes(addr)
	addr256 := crypto.Keccak256(addrRlp)

	intB := common.Uint32ToBytes(myint)

	fmt.Println("strByte:", str256, "\n十六进制的Hash", hexutil.Encode(common.BytesToHash(str256).Bytes()))
	fmt.Println("e1NameByte:", e1Name256, "len", len(e1Name256), "\n十六进制的Hash", hexutil.Encode(common.BytesToHash(e1Name256).Bytes()))
	fmt.Println("e2NameByte:", e2Name256, "len", len(e2Name256), "\n十六进制的Hash", hexutil.Encode(common.BytesToHash(e2Name256).Bytes()))

	fmt.Println("addrByte:", addr256, "len", len(addr256), "\n十六进制的Hash", hexutil.Encode(common.BytesToHash(addr256).Bytes()))

	fmt.Println("intByte:", intB, "\n十六进制的Hash", hexutil.Encode(common.BytesToHash(intB).Bytes()))

}