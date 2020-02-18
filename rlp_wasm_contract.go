package main

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	"github.com/go-ethereum-analysis/rlp"
	"io/ioutil"
	"os"
)

func main() {


	buf, err := ioutil.ReadFile("./contract.wasm")
	if nil != err {
		fmt.Println("err:", err)
	}
	//fmt.Println("buf:", buf)

	params := struct {
		FuncName string
	}{
		FuncName: "init",
	}

	bparams, err := rlp.EncodeToBytes(params)
	if nil != err {
		fmt.Println("err:", err)
	}
	//fmt.Println("bparams:", bparams)

	arr := [][]byte{buf, bparams}

	barr, err := rlp.EncodeToBytes(arr)
	if nil != err {
		fmt.Println("err:", err)
	}
	//fmt.Println("barr:", barr)

	interp := []byte{0x00, 0x61, 0x73, 0x6d}

	input := append(interp, barr...)

	rlpData := hexutil.Encode(input)

	//fmt.Printf("rlp data = %s\n", rlpData)


	f, err := os.Create("./wasm_input.txt")
	if err != nil {
		fmt.Println(err)

		return
	}
	l, err := f.WriteString(rlpData)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}


}