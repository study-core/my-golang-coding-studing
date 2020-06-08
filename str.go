package main

import (
	"fmt"
	"math/big"

	"github.com/go-ethereum-analysis/common/hexutil"
)

func main() {


	//str := "ABC"
	//b := []byte(str)
	//hex := hexutil.Encode(b)
	//fmt.Println("BBB:", hex)
	//
	//
	//b32 := common.Uint32ToBytes(32)
	//b3 := common.Uint32ToBytes(3)
	//b5 := common.Uint32ToBytes(5)
	//fmt.Println("b32:", b32, "\nb3:", b3, "\nb5:", b5)
	//h32 := common.BytesToHash(b32)
	//h3 := common.BytesToHash(b3)
	//h5 := common.BytesToHash(b5)
	//fmt.Println("hex32:", h32, "\nhex3:",  h3, "\nhex5:",  h5)
	//fmt.Println("hex32:", hexutil.Encode(h32.Bytes()), "\nhex3:",  hexutil.Encode(h3.Bytes()), "\nhex5:",  hexutil.Encode(h5.Bytes()))
	//
	//
	//
	//// z = x ** y mod | m |
	//r := new(big.Int).Exp(big.NewInt(32), big.NewInt(3), big.NewInt(5))
	//fmt.Println("32 ** 3 mod | 5 | == ", r, "byte:", common.BytesToHash(r.Bytes()).Bytes())


	x1 := hexutil.MustDecode("0x1c76476f4def4bb94541d57ebba1193381ffa7aa76ada664dd31c16024c43f59")
	y1 := hexutil.MustDecode("0x3034dd2920f673e204fee2811c678745fc819b55d3e9d294e45c9b03a76aef41")
	x2 := hexutil.MustDecode("0x209dd15ebff5d46c4bd888e51a93cf99a7329636c63514396b4a452003a35bf7")
	y2 := hexutil.MustDecode("0x04bf11ca01483bfa8b34b43561848d28905960114c8ac04049af4b6315a41678")
	x3 := hexutil.MustDecode("0x2bb8324af6cfc93537a2ad1a445cfd0ca2a71acd7ac41fadbf933c2a51be344d")
	y3 := hexutil.MustDecode("0x120a2a4cf30c1bf9845f20c6fe39e07ea2cce61f0c9bb048165fe5e4de877550")
	x4 := hexutil.MustDecode("0x111e129f1cf1097710d41c4ac70fcdfa5ba2023c6ff1cbeac322de49d1b6df7c")
	y4 := hexutil.MustDecode("0x2032c61a830e3c17286de9462bf242fca2883585b93870a73853face6a6bf411")
	x5 := hexutil.MustDecode("0x198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2")
	y5 := hexutil.MustDecode("0x1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed")
	x6 := hexutil.MustDecode("0x090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b")
	y6 := hexutil.MustDecode("0x12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa")

	input := make([]byte, 32*2*6)
	copy(input[32*0:32*1], x1)
	copy(input[32*1:32*2], y1)
	copy(input[32*2:32*3], x2)
	copy(input[32*3:32*4], y2)
	copy(input[32*4:32*5], x3)
	copy(input[32*5:32*6], y3)
	copy(input[32*6:32*7], x4)
	copy(input[32*7:32*8], y4)
	copy(input[32*8:32*9], x5)
	copy(input[32*9:32*10], y5)
	copy(input[32*10:32*11], x6)
	copy(input[32*11:32*12], y6)


	hex := hexutil.Encode(input)

	fmt.Println("hex:", hex)



	x1b := new(big.Int).SetBytes(x1)
	y1b := new(big.Int).SetBytes(y1)
	x2b := new(big.Int).SetBytes(x2)
	y2b := new(big.Int).SetBytes(y2)
	x3b := new(big.Int).SetBytes(x3)
	y3b := new(big.Int).SetBytes(y3)
	x4b := new(big.Int).SetBytes(x4)
	y4b := new(big.Int).SetBytes(y4)
	x5b := new(big.Int).SetBytes(x5)
	y5b := new(big.Int).SetBytes(y5)
	x6b := new(big.Int).SetBytes(x6)
	y6b := new(big.Int).SetBytes(y6)

	 fmt.Println("x1:=", x1b)
	 fmt.Println("y1:=", y1b)
	 fmt.Println("x2:=", x2b)
	 fmt.Println("y2:=", y2b)
	 fmt.Println("x3:=", x3b)
	 fmt.Println("y3:=", y3b)
	 fmt.Println("x4:=", x4b)
	 fmt.Println("y4:=", y4b)
	 fmt.Println("x5:=", x5b)
	 fmt.Println("y5:=", y5b)
	 fmt.Println("x6:=", x6b)
	 fmt.Println("y6:=", y6b)

	fmt.Println("")
	fmt.Println("bn256Mul")

	var ax, ay, scalar [32]byte
	copy(ax[:], hexutil.MustDecode("0x2bd3e6d0f3b142924f5ca7b49ce5b9d54c4703d7ae5648e61d02268b1a0a9fb7"))
	copy(ay[:], hexutil.MustDecode("0x21611ce0a6af85915e2f1d70300909ce2e49dfad4a4619c8390cae66cefdb204"))
	copy(scalar[:], hexutil.MustDecode("0x00000000000000000000000000000000000000000000000011138ce750fa15c2"))

	axb := new(big.Int).SetBytes(ax[:])
	ayb := new(big.Int).SetBytes(ay[:])
	scalarb := new(big.Int).SetBytes(scalar[:])

	fmt.Println("ax:=", axb)
	fmt.Println("ay:=", ayb)
	fmt.Println("scalar:=", scalarb)

	var bx, by [32]byte
	copy(bx[:], hexutil.MustDecode("0x070a8d6a982153cae4be29d434e8faef8a47b274a053f5a4ee2a6c9c13c31e5c"))
	copy(by[:], hexutil.MustDecode("0x031b8ce914eba3a9ffb989f9cdd5b0f01943074bf4f0f315690ec3cec6981afc"))
	bxb := new(big.Int).SetBytes(bx[:])
	byb := new(big.Int).SetBytes(by[:])
	fmt.Println("Mul得到的坐标")
	fmt.Println("bx:=", bxb)
	fmt.Println("by:=", byb)
	//




	fmt.Println("")
	fmt.Println("bn256Add")

	var ax1, ay1, bx2, by2 [32]byte
	copy(ax1[:], hexutil.MustDecode("0x17c139df0efee0f766bc0204762b774362e4ded88953a39ce849a8a7fa163fa9"))
	copy(ay1[:], hexutil.MustDecode("0x01e0559bacb160664764a357af8a9fe70baa9258e0b959273ffc5718c6d4cc7c"))
	copy(bx2[:], hexutil.MustDecode("0x039730ea8dff1254c0fee9c0ea777d29a9c710b7e616683f194f18c43b43b869"))
	copy(by2[:], hexutil.MustDecode("0x073a5ffcc6fc7a28c30723d6e58ce577356982d65b833a5a5c15bf9024b43d98"))

	ax1b := new(big.Int).SetBytes(ax1[:])
	ay1b := new(big.Int).SetBytes(ay1[:])
	bx2b := new(big.Int).SetBytes(bx2[:])
	by2b := new(big.Int).SetBytes(by2[:])
	fmt.Println("x1:=", ax1b)
	fmt.Println("y1:=", ay1b)
	fmt.Println("x2:=", bx2b)
	fmt.Println("y2:=", by2b)
}
