package main

import (
	"bytes"
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	"github.com/PlatONnetwork/PlatON-Go/rlp"

	//"encoding/hex"
	//"github.com/go-ethereum-analysis/common"
	//"github.com/go-ethereum-analysis/p2p/discover"

	//"bytes"
	"fmt"
	//"github.com/PlatONnetwork/PlatON-Go/crypto"
	//"github.com/go-ethereum-analysis/common/hexutil"

	//"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	//"github.com/PlatONnetwork/PlatON-Go/rlp"
	//"time"
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
	//str := "493301712671ada506ba6ca7891f436d29185821"
	//myint := uint32(246)
	//
	//e1Name := "mystr"
	//e2Name := "myint"
	//
	//addr := common.HexToAddress("493301712671ada506ba6ca7891f436d29185821")
	//
	//strRlp, _ := rlp.EncodeToBytes(str)
	//str256 := crypto.Keccak256(strRlp)
	//
	//e1NameRlp, _ := rlp.EncodeToBytes(e1Name)
	//e1Name256 := crypto.Keccak256(e1NameRlp)
	//
	//e2NameRlp, _ := rlp.EncodeToBytes(e2Name)
	//e2Name256 := crypto.Keccak256(e2NameRlp)
	//
	//
	//addrRlp, _ := rlp.EncodeToBytes(addr)
	//addr256 := crypto.Keccak256(addrRlp)
	//
	//intB := common.Uint32ToBytes(myint)
	//
	//fmt.Println("strByte:", str256, "\n十六进制的Hash", hexutil.Encode(common.BytesToHash(str256).Bytes()))
	//fmt.Println("e1NameByte:", e1Name256, "len", len(e1Name256), "\n十六进制的Hash", hexutil.Encode(common.BytesToHash(e1Name256).Bytes()))
	//fmt.Println("e2NameByte:", e2Name256, "len", len(e2Name256), "\n十六进制的Hash", hexutil.Encode(common.BytesToHash(e2Name256).Bytes()))
	//
	//fmt.Println("addrByte:", addr256, "len", len(addr256), "\n十六进制的Hash", hexutil.Encode(common.BytesToHash(addr256).Bytes()))
	//
	//fmt.Println("intByte:", intB, "\n十六进制的Hash", hexutil.Encode(common.BytesToHash(intB).Bytes()))

	rlpData := ""

	var params [][]byte
	params = make([][]byte, 0)

	fnType, _ := rlp.EncodeToBytes(uint16(2009))
	params = append(params, fnType)
	version := uint32(0<<16 | 11<<8 | 0)
	//version := uint32(0<<16 | 10<<8 | 1)
	fmt.Println("version", version)
	addr, _ := rlp.EncodeToBytes(version)
	params = append(params, addr)

	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, params)
	if err != nil {
		fmt.Printf("encode rlp data fail: %v \n", err)
	} else {
		rlpData = hexutil.Encode(buf.Bytes())
		fmt.Printf("rlp data = %s \n", rlpData)
	}

	//
	//start := time.Now()
	//
	//time.Sleep(time.Second*10)
	//
	//end := time.Since(start)
	//
	//fmt.Println("Duration:", end)


	/*codeStr := "0x6080604052600436106100955763ffffffff60e060020a60003504166306fdde0381146100a7578063095ea7b31461013157806318160ddd1461016957806323b872dd14610190578063313ce567146101ba57806354fd4d50146101e557806370a08231146101fa57806395d89b411461021b578063a9059cbb14610230578063cae9ca5114610254578063dd62ed3e146102bd575b3480156100a157600080fd5b50600080fd5b3480156100b357600080fd5b506100bc6102e4565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100f65781810151838201526020016100de565b50505050905090810190601f1680156101235780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561013d57600080fd5b50610155600160a060020a0360043516602435610372565b604080519115158252519081900360200190f35b34801561017557600080fd5b5061017e6103d9565b60408051918252519081900360200190f35b34801561019c57600080fd5b50610155600160a060020a03600435811690602435166044356103df565b3480156101c657600080fd5b506101cf6104cc565b6040805160ff9092168252519081900360200190f35b3480156101f157600080fd5b506100bc6104d5565b34801561020657600080fd5b5061017e600160a060020a0360043516610530565b34801561022757600080fd5b506100bc61054b565b34801561023c57600080fd5b50610155600160a060020a03600435166024356105a6565b34801561026057600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610155948235600160a060020a031694602480359536959460649492019190819084018382808284375094975061063f9650505050505050565b3480156102c957600080fd5b5061017e600160a060020a03600435811690602435166107da565b6003805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561036a5780601f1061033f5761010080835404028352916020019161036a565b820191906000526020600020905b81548152906001019060200180831161034d57829003601f168201915b505050505081565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35060015b92915050565b60005481565b600160a060020a038316600090815260016020526040812054821180159061042a5750600160a060020a03841660009081526002602090815260408083203384529091529020548211155b80156104365750600082115b156104c157600160a060020a03808416600081815260016020908152604080832080548801905593881680835284832080548890039055600282528483203384528252918490208054879003905583518681529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35060016104c5565b5060005b9392505050565b60045460ff1681565b6006805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561036a5780601f1061033f5761010080835404028352916020019161036a565b600160a060020a031660009081526001602052604090205490565b6005805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561036a5780601f1061033f5761010080835404028352916020019161036a565b3360009081526001602052604081205482118015906105c55750600082115b156106375733600081815260016020908152604080832080548790039055600160a060020a03871680845292819020805487019055805186815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060016103d3565b5060006103d3565b336000818152600260209081526040808320600160a060020a038816808552908352818420879055815187815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a383600160a060020a031660405180807f72656365697665417070726f76616c28616464726573732c75696e743235362c81526020017f616464726573732c627974657329000000000000000000000000000000000000815250602e019050604051809103902060e060020a9004338530866040518563ffffffff1660e060020a0281526004018085600160a060020a0316600160a060020a0316815260200184815260200183600160a060020a0316600160a060020a03168152602001828051906020019080838360005b8381101561077f578181015183820152602001610767565b50505050905090810190601f1680156107ac5780820380516001836020036101000a031916815260200191505b509450505050506000604051808303816000875af19250505015156107d057600080fd5b5060019392505050565b600160a060020a039182166000908152600260209081526040808320939094168252919091522054905600a165627a7a723058206e603c634e6d2ac6f989b2e2055cfbca683a30ca87aac8e0a9a6bb65e858050e0029"

	code, err := hexutil.Decode(codeStr)
	if nil != err {
		fmt.Println(err)
	}
	codeHash := crypto.Keccak256Hash(code)
	hash := hex.EncodeToString(codeHash[:])
	fmt.Println(hash)*/



	/*addrHashs := map[string]struct{}{

		"158a5868701d23cbb986316b306ca4c579c7c50f74f8efc129da1539e594a65a": {},
		"3bc7204a1d5e9a79689bd995b99addf9f9d96b9945dac728a0c4d69f76e59864": {},
		"40638507490bb06791d2ed0c0aafb54748d71d1dc6bd1e7279ebf9a025efcac0": {},
		"c35d4e937cb07793341b95f9a5b53299e6ce4c0b394bbac77611ce83b2989c6a": {},
		"fc4309a490605cca5aa8ee4f6f8f34c932e8dc639dfe70da95708f9d65317778": {},
		"65c8eae1c398b26f2142697cbc6cb5c266ab8391cabb310d640659c46af7a5e6": {},
		"85c34b6555280d3ec1f09bc3c6ed6e95da3dee4881785c51a03945f30c66748e": {},
		"86ff74a19a7a3866664b9befe51f5b2b63a4bf3e97dab8efafc2541daf83fbca": {},
		"a685afb4b40e30f9c82591c3968e8bd68c0b3342626258c8202e9adbe836f70f": {},
		"b4dc9441ef63c5f599dd6bacc725c75d6e907f2d70bbfeea0431548cdc051d6f": {},
		"c685dae6d7d935ed6556b6a5490e029aaf319b20f346322253a8d3cdb04c2678": {},
		"cfd8f594f67b53c2f7deb970354576ff010b10a393947acdcea1304978517101": {},
	}
//  0x32bec384344c2dc1ea794a4e149c1b74dd8467ef

	adds := map[string]struct{}{
		"0x0f2a1795f7605a96910c6c782ea5d3d291fd77fc": {},
		"0x18253041ce1d238f42f685ff1714153ea9c97699": {},
		"0x187b3a7d5f790a30f59703338eba42ee11e584fa": {},
		"0x1e4419a4a0c96bab21004c258586c9e172c98ed6": {},
		"0x3800f5390e5921059a1ae817e12c224813cdd33a": {},
		"0x471b4f5e00bf612766b1ead6df6e658244e6c179": {},
		"0xab782161cf50b8282afd717d27b5a99bc80f909b": {},
		"0xb3dce06660b0d1b6c89e217a6ec61a074029e5d3": {},
		"0xc1f4f1fa20461564e031c6e189018adc11270156": {},
		"0xdd26d10d54e3860a62d268bb877952292a6fcc48": {},
		"0xfbe13f9f86a7bb272f0d6479beb2b0f4b4114ed8": {},
		"0xfd783cbf2603b1ddf0740512aac47dc954a13897": {},
	}


	for addStr, _ := range adds {
		addr := common.HexToAddress(addStr)
		addHash := crypto.Keccak256Hash(addr[:])
		hash := hex.EncodeToString(addHash[:])

		if _, ok := addrHashs[hash]; !ok {
			fmt.Println("db中没有加载到的Addr", "addr:", addr.String(), "hash:", hash )
		} else {
			fmt.Println("addStr:", addStr, "addrHash:", hash)
		}




	}

	for k, _ := range addrHashs {

		flag := 0

		for addStr, _ := range adds {
			addr := common.HexToAddress(addStr)
			addHash := crypto.Keccak256Hash(addr[:])
			hash := hex.EncodeToString(addHash[:])

			if k == hash {
				flag |= 1
			}


		}

		if flag != 1 {
			fmt.Println("发现 db 中找找不到的:", k)
		}
	}

	//fmt.Println(len(addrHashs))
*/














}