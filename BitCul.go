package main

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"github.com/PlatONnetwork/PlatON-Go/x/xutil"

	//"github.com/PlatONnetwork/PlatON-Go/common/math"
	//"math/big"
)

const (
	invalided   = 1 << iota
	slashed
	notEnough
	valided  	= 0
)

func main() {

	fmt.Println(valided, invalided, slashed, notEnough)

	status := invalided|notEnough|slashed


	fmt.Println(status)


	fmt.Println("开始")

	fmt.Println(status&notEnough == notEnough)
	fmt.Println(status&invalided == invalided)
	fmt.Println(status|invalided == invalided)
	fmt.Println(status|invalided == invalided && status|valided != valided)
	fmt.Println(status|invalided == status&invalided)
	//
	//fmt.Println(math.MaxBig256.Cmp(big.NewInt(math.MaxInt64)))
	//
	//// 115792089237316195423570985008687907853269984665640564039457584007913129639935
	//fmt.Println(math.MaxBig256.String())
	//fmt.Println(math.MaxInt64) // 9223372036854775807
	//// fmt.Println(fmt.Sprint(math.MaxUint64)) // 18446744073709551615
	//
	//
	//fmt.Println(math.MaxBig63.String())
	//// 9223372036854775807
	//// 18446744073709551615
	//// 1000000000000000000000000
	//
	//// 10 00000000 000000000000000000
	//
	//str, _ := new(big.Int).SetString("1000000000000000000000000", 10)
	//fmt.Println(str)
	//
	//fmt.Println(big.MaxPrec)
	//
	//
	//// 这个是最大数
	//tt64      := math.BigPow(2, 128)
	//MaxBig64 := new(big.Int).Set(tt64)
	//fmt.Println(MaxBig64)


	fmt.Println(status)

	status &^= slashed
	fmt.Println(status)

	fmt.Println("对比")


	fmt.Println(status&notEnough, status&notEnough == notEnough)
	fmt.Println(status&invalided, status&invalided == invalided)
	fmt.Println(status&0xFF == notEnough)
	fmt.Println(status&0xFF == invalided)



	nodeID := discover.MustHexID("0x1f3a8672348ff6b789e416762ad53e69063138b8eb4d8780101658f24b2369f1a8e09499226b467d8bc0c4e03e1dc903df857eeb3c67733d21b6aaee2840e429")
	addr := common.HexToAddress("0x095e7baea6a6c7c4c2dfeb977efac326af552d87")

	if ad, err := xutil.NodeId2Addr(nodeID); nil != err {
		fmt.Println("解释:", err)
	}else {
		fmt.Println("生成的地址:", ad.Hex())
	}
	fmt.Println("NodeId:", nodeID.String())
	fmt.Println("addr:", addr.Hex())


	promoteVersion := uint32(6<<24 | 5<<16 | 2<<8 | 4)

	//fmt.Println("promoteVersion", promoteVersion)
	//
	promoteVersionByte := common.Uint32ToBytes(promoteVersion)
	fmt.Println(promoteVersionByte)
	//
	//promoteVersion2 := promoteVersion>>8
	//fmt.Println("promoteVersion2", promoteVersion2)
	//
	//promoteVersion2Byte := common.Uint32ToBytes(promoteVersion2)
	//fmt.Println(promoteVersion2Byte)


	fmt.Println("挑出对应为上的数字")

	tmp := promoteVersion<<8
	fmt.Println(common.Uint32ToBytes(tmp))
	tmp = tmp>>24
	fmt.Println(common.Uint32ToBytes(tmp))
	fmt.Println(tmp)

	aa := promoteVersion<<16
	fmt.Println(common.Uint32ToBytes(aa))
	aa = aa>>24
	fmt.Println(common.Uint32ToBytes(aa))
	fmt.Println(aa)


	tt := promoteVersion<<24
	fmt.Println(common.Uint32ToBytes(tt))
	tt = tt>>24
	fmt.Println(common.Uint32ToBytes(tt))
	fmt.Println(tt)



}