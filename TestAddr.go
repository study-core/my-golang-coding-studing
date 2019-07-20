package main

import (
	"fmt"
	"math/big"

	//"github.com/PlatONnetwork/PlatON-Go/common"

	//"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/pangu/PlatON-Go/rlp"
	//"math/big"
)

func main() {

	//addrA := common.BigToAddress(big.NewInt(12))
	//addrB := common.BigToAddress(big.NewInt(13))
	//c := []byte(fmt.Sprint(78))
	//
	//
	//
	//fmt.Println(addrA.Hex())
	//fmt.Println(addrB.Hex())
	//fmt.Println(string(c))
	//
	//a_byte := addrA.Bytes()
	//b_byte := addrB.Bytes()
	//
	//arr := append(a_byte, append(b_byte, c...)...)
	//
	//fmt.Println(string(arr))
	//
	//
	//s1 := common.BytesToAddress(arr[0:common.AddressLength])
	//
	//fmt.Println(s1.Hex())


	//c := big.NewInt(12)
	//
	//cbyte, _ := rlp.EncodeToBytes(c)
	//
	//
	//fmt.Println()
	//
	//var d *big.Int
	//
	//rlp.DecodeBytes(cbyte, &d)
	//
	//fmt.Println(d.String())

	c := []byte("121245")

	arr := [6]byte{}
	for i, v := range c {
		arr[i] = v
	}

	fmt.Println(arr)

	arrbyte, _ := rlp.EncodeToBytes(arr)

	var d [6]byte
	rlp.DecodeBytes(arrbyte, &d)
	fmt.Println(d)

	var ds []byte
	for _, v := range d {
		ds = append(ds, v)
	}

	fmt.Println(string(ds))



	fmt.Println("sss|||||||||||||||||||||||||||||||||||||||||||")

	i := int(-1)



	ibyte, err := rlp.EncodeToBytes(i)
	if nil != err {
		fmt.Println(err)
	}

	var as int
	rlp.DecodeBytes(ibyte, &as)
	fmt.Println(as)



	ui := uint8(14)

	uiByte, err := rlp.EncodeToBytes(ui)
	if nil != err {
		fmt.Println(err)
	}

	var uu uint8
	rlp.DecodeBytes(uiByte, &uu)

	fmt.Println("uint8", uu)


	ff := new(big.Int).Sub(big.NewInt(5), big.NewInt(2))

	ffByte, err := rlp.EncodeToBytes(ff)

	if nil != err {
		fmt.Println(err)
	}

	var fuu *big.Int

	rlp.DecodeBytes(ffByte, &fuu)
	fmt.Println(fuu)


	a1 := &ash{
		A: -1,
		B: 2,
	}

	alByte, err := rlp.EncodeToBytes(a1)
	if nil != err {
		fmt.Println("struct", err)
	}

	var a1s *ash
	rlp.DecodeBytes(alByte, &a1s)
	fmt.Println(a1s)



	b1 := &bsh{
		A: new(big.Int).Sub(big.NewInt(1), big.NewInt(2)),
	}

	blByte, err := rlp.EncodeToBytes(b1)
	if nil != err {
		fmt.Println("struct", err)
	}

	var b1s *bsh
	rlp.DecodeBytes(blByte, &b1s)
	fmt.Println(b1s)


	flag := false

	flagByte, err := rlp.EncodeToBytes(flag)
	if nil != err {
		fmt.Println("bool", err)
	}

	var fflag bool

	rlp.DecodeBytes(flagByte, &fflag)
	fmt.Println(fflag)



	flo := float64(12.0)

	fByte, err := rlp.EncodeToBytes(flo)
	if nil != err {
		fmt.Println("float32", err)
	}

	var ffo float32

	rlp.DecodeBytes(fByte, &ffo)
	fmt.Println(ffo)

}


type ash struct {

	A int
	B uint
}


type bsh struct {
	A *big.Int
}
