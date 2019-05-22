package main

import (
	"fmt"
	//"encoding/binary"
	//"unsafe"
	//"sync/atomic"
	//"bytes"
	//"strconv"
	//"unsafe"
	"sync/atomic"
)

func main() {

	/*for _, v := range []int {1, 4, 5, 7, 8, 9, 6, 3} {
		vb :=
		fmt.Println("v", v)
		fmt.Println()
	}

	state1 := [12]byte{}*/

	/*i := uint64(12345678910111213145)

	fmt.Println(i)

	b := make([]byte, 12)
	binary.BigEndian.PutUint64(b, i)
	fmt.Println(b)
	state1 := [12]byte{}
	for i, v := range b {
		state1[i] = v
	}
	fmt.Println(state1)


	statep := (*uint64)(unsafe.Pointer(&state1))

	fmt.Println(*statep)

	state := atomic.AddUint64(statep, uint64(2)<<32)
	// 获取计数器的值
	v := int32(state >> 32)
	//获得wait（）等待次数
	w := uint32(state)

	fmt.Println(v, w)*/

/*
	var i1 int64 = 511 // [00000000 00000000 ... 00000001 11111111] = [0 0 0 0 0 0 1 255]

	s1 := make([]byte, 0)
	buf := bytes.NewBuffer(s1)

	// 数字转 []byte, 网络字节序为大端字节序
	binary.Write(buf, binary.BigEndian, i1)
	fmt.Println(buf.Bytes())

	// 数字转 []byte, 小端字节序
	buf.Reset()
	binary.Write(buf, binary.LittleEndian, i1)
	fmt.Println(buf.Bytes())

	// []byte 转 数字
	s2 := []byte{6: 1, 7: 255} // [0 0 0 0 0 0 1 255]
	buf = bytes.NewBuffer(s2)
	var i2 int64
	binary.Read(buf, binary.BigEndian, &i2)
	fmt.Println(i2)     // 511

	s3 := []byte{255, 1, 7:0}   // [255 1 0 0 0 0 0 0]
	buf = bytes.NewBuffer(s3)
	var i3 int64
	binary.Read(buf, binary.LittleEndian, &i3)
	fmt.Println(i3)     // 511



	ss := uint32(16777215)
	sd := byte(ss)
	fmt.Println(ss, sd)




	ab := []byte(strconv.Itoa(int(45)))
	ai, _ := strconv.Atoi(string(ab))
	fmt.Println(ab, ai)*/


}

/**
大端小端编码
 */

// 这里是大端模式
func f2() {
	var v2 uint32
	var b2 [4]byte
	v2 = 257
	// 将 256转成二进制就是
	// | 00000000 | 00000000 | 00000001 | 00000001 |
	// | b2[0]    | b2[1]   | b2[2]    | [3]      | // 这里表示b2数组每个下标里面存放的值

	// 这里直接使用将uint32l强转成uint8
	// | 00000000 0000000 00000001 | 00000001  直接转成uint8后等于 1
	// |---这部分go在强转的时候扔掉---|
	b2[3] = uint8(v2)

	// | 00000000 | 00000000 | 00000001 | 00000001 | 右移8位 转成uint8后等于 1
	// 下面是右移后的数据
	// |          | 00000000 | 00000000 | 00000001 |
	b2[2] = uint8(v2 >> 8)

	// | 00000000 | 00000000 | 00000001 | 00000001 | 右移16位 转成uint8后等于 0
	// 下面是右移后的数据
	// |          |          | 00000000 | 00000000 |
	b2[1] = uint8(v2 >> 16)

	// | 00000000 | 00000000 | 00000001 | 00000001 | 右移24位 转成uint8后等于 0
	// 下面是右移后的数据
	// |          |          |          | 00000000 |
	b2[0] = uint8(v2 >> 24)

	fmt.Printf("%+v\n", b2)
	// 所以最终将uint32转成[]byte数组输出为
	// [0 0 1 1]
}

// 这里是小端模式
// 在上面我们讲过，小端刚好和大端相反的，所以在转成小端模式的时候，只要将[]byte数组的下标首尾对换一下位置就可以了
func f3() {
	var v3 uint32
	var b3 [4]byte
	v3 = 257
	// 将 256转成二进制就是
	// | 00000000 | 00000000 | 00000001 | 00000001 |
	// | b3[0]    | b3[1]   | b3[2]    | [3]      | // 这里表示b3数组每个下标里面存放的值

	// 这里直接使用将uint32l强转成uint8
	// | 00000000 0000000 00000001 | 00000001  直接转成uint8后等于 1
	// |---这部分go在强转的时候扔掉---|
	b3[0] = uint8(v3)

	// | 00000000 | 00000000 | 00000001 | 00000001 | 右移8位 转成uint8后等于 1
	// 下面是右移后的数据
	// |          | 00000000 | 00000000 | 00000001 |
	b3[1] = uint8(v3 >> 8)

	// | 00000000 | 00000000 | 00000001 | 00000001 | 右移16位 转成uint8后等于 0
	// 下面是右移后的数据
	// |          |          | 00000000 | 00000000 |
	b3[2] = uint8(v3 >> 16)

	// | 00000000 | 00000000 | 00000001 | 00000001 | 右移24位 转成uint8后等于 0
	// 下面是右移后的数据
	// |          |          |          | 00000000 |
	b3[3] = uint8(v3 >> 24)

	fmt.Printf("%+v\n", b3)
	// 所以最终将uint32转成[]byte数组输出为
	// [1 1 0 0 ]
}


