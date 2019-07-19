package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/x/restricting"
	"github.com/PlatONnetwork/PlatON-Go/x/xcom"
	"github.com/PlatONnetwork/PlatON-Go/x/xutil"
	"github.com/pangu/PlatON-Go/common"
	"math/big"
)

func main() {

	nodeIdArr := []discover.NodeID{
		discover.MustHexID("0x1f3a8672348ff6b789e416762ad53e69063138b8eb4d8780101658f24b2369f1a8e09499226b467d8bc0c4e03e1dc903df857eeb3c67733d21b6aaee2840e429"),
		discover.MustHexID("0xa6ef31a2006f55f5039e23ccccef343e735d56699bde947cfe253d441f5f291561640a8e2bbaf8a85a8a367b939efcef6f80ae28d2bd3d0b21bdac01c3aa6f2f"),
		discover.MustHexID("0xc7fc34d6d8b3d894a35895aaf2f788ed445e03b7673f7ce820aa6fdc02908eeab6982b7eb97e983cc708bcec093b3bc512b0b1fbf668e6ab94cd91f2d642e591"),
		discover.MustHexID("0x97e424be5e58bfd4533303f8f515211599fd4ffe208646f7bfdf27885e50b6dd85d957587180988e76ae77b4b6563820a27b16885419e5ba6f575f19f6cb36b0"),
	}

	for i, nodeId := range nodeIdArr {

		addr, err := xutil.NodeId2Addr(nodeId)
		if nil != err {
			fmt.Printf("第几个: %d, nodeId转化addr 出了问题: %s", i, err)
		}
		fmt.Println("Addr:", addr.Hex())
	}



	res := xcom.Result{true, "1000", ""}
	data, err := rlp.EncodeToBytes(res)
	if nil != err {
		fmt.Println("编码失败:", err)
	}

	var r xcom.Result
	err = rlp.DecodeBytes(data, &r)
	if nil != err {
		fmt.Println("错误:", err)
	}

	rbyte, _ := json.Marshal(r)

	fmt.Println("结果:", string(rbyte))


	dbyte, _ := rlp.EncodeToBytes(r.Data)

	var aa uint32
	err = rlp.DecodeBytes(dbyte, &aa)
	if nil != err {
		fmt.Println("解码失败:", err)
	}else {
		fmt.Println(aa)
	}



	i := uint32(1000)
	ibyte, _ := rlp.EncodeToBytes(i)
	fmt.Println("rlp uint32:", ibyte)

	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, i)

	buf2 := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf2, i)

	fmt.Println("大端:", buf)
	fmt.Println("小端:", buf2)



	ii := uint16(1000)
	i16, _ := rlp.EncodeToBytes(ii)
	fmt.Println("rlp uint16:", i16)

	buf16 := make([]byte, 2)
	binary.BigEndian.PutUint16(buf16, ii)

	buf16_2 := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf16_2, ii)

	fmt.Println("大端:", buf16)
	fmt.Println("小端:", buf16_2)


	addr := common.HexToAddress("0x89dcade1e353984f4085ef99d3e24b5667a93aeb")
	fmt.Println("addr 20 byte:", addr)

	addrByte, _ := rlp.EncodeToBytes(addr)
	fmt.Println("addr rlp:", addrByte)




	aPlan := restricting.RestrictingPlan{
		Epoch: uint64(1),
		Amount: big.NewInt(1),
	}
	bPlan := restricting.RestrictingPlan{
		Epoch: uint64(2),
		Amount: big.NewInt(2),
	}

	arr := make([]restricting.RestrictingPlan, 0)

	arr = append(arr, aPlan)
	arr = append(arr, bPlan)

	planArrB, _ := rlp.EncodeToBytes(arr)

	fmt.Println("RestrictingPlan：", planArrB)


}

func ByteSlice(b []byte) []byte { return b }