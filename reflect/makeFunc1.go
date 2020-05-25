package main

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
)

//func main(){
//
//	// 一个互换函数
//	var swap = func(in []reflect.Value) []reflect.Value {
//		fmt.Println("函数替换")
//		fmt.Println(in[0].Interface(), in[1].Interface())
//		//1 2
//		return []reflect.Value{in[1], in[0]}
//	}
//
//	// 入参一个 func ptr
//	var makeSwap = func(fptr interface{}) {
//
//		// 间接返回 reflect.ValueOf(fptr) 指向的 值
//		var valueOf reflect.Value = reflect.Indirect( reflect.ValueOf(fptr))
//
//		// 根据 func sig 和 匿名func 返回一个新的func
//		var v reflect.Value = reflect.MakeFunc(valueOf.Type(), swap)
//		valueOf.Set(v)
//
//		fmt.Println("打印了一次 valueOf")
//		fmt.Println(valueOf)
//		//<func(int, int) (int, int) Value>
//
//		fmt.Println("打印了一次 v")
//		fmt.Println(v)
//		//<func(int, int) (int, int) Value>
//	}
//	var intSwap func(int, int) (int, int)
//	makeSwap(&intSwap)
//
//	fmt.Println("这里是")
//	fmt.Println(intSwap(1, 2))
//	//2 1
//
//
//
//	/*str := "hello"
//
//	r, err := rlp.EncodeToBytes([]byte(str))
//	if nil != err {
//		fmt.Println("err", err)
//	}
//
//	var b []byte
//
//	rlp.DecodeBytes(r, &b)
//	fmt.Println("rlpStr", string(b))
//	 */
//
//
//	//type Plan struct {
//	//	Epoch uint
//	//	Value  string
//	//}
//	//
//	//p1 := &Plan{
//	//	Epoch: 1,
//	//	Value: "A",
//	//}
//
//
//	//
//	//
//	//
//	//p2 := &Plan{
//	//	Epoch: 2,
//	//	Value: "B",
//	//}
//	//
//	////arr1 := make([]*Plan, 0)
//	////
//	////arr1 = append(arr1, p1)
//	////arr1 = append(arr1, p2)
//	//
//	//arr1 := []*Plan{p1, p2}
//	//
//	//b1, err := rlp.EncodeToBytes(&arr1)
//	//if nil != err {
//	//	fmt.Println("b1, err", err)
//	//}else {
//	//	fmt.Println("b1", b1)
//	//}
//	//
//	//
//	//i1 := []interface{}{uint(1),"A"}
//	//i2 := []interface{}{uint(2),"B"}
//	//
//	//arr2 := [][]interface{}{i1, i2}
//	//
//	//b2, err := rlp.EncodeToBytes(arr2)
//	//if nil != err {
//	//	fmt.Println("b2, err", err)
//	//}else {
//	//	fmt.Println("b2", b2)
//	//}
//
//
//	//pb1, _ := rlp.EncodeToBytes(p1)
//	//pb2, _ := rlp.EncodeToBytes(1)
//	//pb3, _ := rlp.EncodeToBytes("A")
//	//array := [][]byte{pb2, pb3}
//	//bbbb, _ := rlp.EncodeToBytes(array)
//	//
//	//
//	//fmt.Println(pb1)
//	//fmt.Println(bbbb)
//}



type WhtType struct {
	name   string
	age    int
	weight int
}

type WhtGroup struct {
	info   string
	number int
	member WhtType
}

func main() {
	var t WhtType
	var g WhtGroup

	one := WhtType{name: "jatel", age: 30, weight: 160}
	if one_code, err := rlp.EncodeToBytes(one); nil == err {
		fmt.Println(one_code)

		e := rlp.DecodeBytes(one_code, &t)
		if nil != e {
			fmt.Println("ee1", e)
		}else {
			fmt.Println(t)
		}


	}

	group := WhtGroup{info: "grou22", number: 3, member: one}

	if bb, err := rlp.EncodeToBytes(group); nil == err {
		fmt.Println(bb)

		e := rlp.DecodeBytes(bb, &g)

		if nil != e {
			fmt.Println("ee2", e)
		}else {
			fmt.Println(g)
		}
	}





}
