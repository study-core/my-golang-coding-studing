package main

import (
	//"math/big"
	"Platon-go/common/hexutil"
	"fmt"
	"Platon-go/rlp"
	"bytes"
	"encoding/binary"
)

func main() {
	TestRlpData()
}


func TestRlpData()  {


	//from := []byte("0x740ce31b3fac20dac379db243021a51e80ad00d7")

	to := []byte("0x5a5c4368e2692746b286cee36ab0710af3efa6cf")

	params := make([][]byte, 0)
	params = append(params, []byte("transfer"))
	params = append(params, to)
	params = append(params, uint64ToBytes(500))


	var buf bytes.Buffer

	if err := rlp.Encode(&buf, params); nil != err {
		fmt.Println(err)
	} else {
		fmt.Println("transfer data: ", hexutil.Encode(buf.Bytes()))
	}

	// CALL
	args := make([][]byte, 0)
	args = append(args, []byte("balanceOf"))
	args = append(args, to)

	var call bytes.Buffer
	if err := rlp.Encode(&call, args); nil != err {
		fmt.Println(err)
	}else {
		fmt.Println("balanceOf data: ", hexutil.Encode(call.Bytes()))
	}

}


func uint64ToBytes(val uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, val)
	return buf[:]
}