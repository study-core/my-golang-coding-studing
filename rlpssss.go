package main

import (
	"fmt"
	"github.com/go-ethereum-analysis/rlp"
)

func main() {

	type Params struct {
		FuncName string
		//Args    interface{}
		Args    [] interface{}
	}

	type Args struct {
		Name   string
		Gender uint64
		Age    uint64
	}

	type Input struct {
		File string
		Des  string
		Num  uint64
		Page uint64
	}

	input := &Input{
		File: "Input",
		Des:  "This is input",
		Num:  34,
		Page: 43,
	}

	ar := &Args{
		Name:   "Args",
		Gender: 23,
		Age:    32,
	}

	inputB, _ := rlp.EncodeToBytes(input)
	arB, _ := rlp.EncodeToBytes(ar)

	//var p1, p2 Params

	//if err := rlp.DecodeBytes(inputB, &p1); nil != err {
	//	fmt.Println("input err:", err)
	//}
	//if err := rlp.DecodeBytes(arB, &p2); nil != err {
	//	fmt.Println("args err:", err)
	//}

	kind, content, _, err := rlp.Split(inputB)

	switch {
	case err != nil:
		fmt.Println(err)
	case kind != rlp.List:
		fmt.Println("input type error")
	}

	//name, _, err := rlp.SplitString(content)
	_, name, _, err := rlp.Split(content)
	if nil != err {
		fmt.Println("input err", err)
	}else {
		fmt.Println("input name", string(name), "Input" == string(name))
	}



	kind, content, _, err = rlp.Split(arB)

	switch {
	case err != nil:
		fmt.Println(err)
	case kind != rlp.List:
		fmt.Println("args type error")
	}

	_, name, _, err = rlp.Split(content)
	if nil != err {
		fmt.Println("args err", err)
	}else {
		fmt.Println("args name", string(name), "Args"== string(name))
	}
}