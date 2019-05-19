package main

import (
	"github.com/google/keytransparency/core/crypto/vrf/p256"
	"fmt"
	"encoding/hex"
)

func main() {
	/** 生成 公私钥对 */
	k, pk := p256.GenerateKey()
	//
	//m1 := []byte("Gavin")
	//m2 := []byte("Emma")
	//m3 := []byte("Kally")


	m1 := []byte("data1")
	m2 := []byte("data2")
	m3 := []byte("data3")

	/** 分别计算出三组不同的 随机数和证明 */
	index1, proof1 := k.Evaluate(m1)
	index2, proof2 := k.Evaluate(m2)
	index3, proof3 := k.Evaluate(m3)

	fmt.Println("随机数1：", hex.EncodeToString(index1[:]))
	fmt.Println("随机数2：", hex.EncodeToString(index2[:]))
	fmt.Println("随机数3：", hex.EncodeToString(index3[:]))

	fmt.Println("证明1：", hex.EncodeToString(proof1))
	fmt.Println("证明2：", hex.EncodeToString(proof2))
	fmt.Println("证明3：", hex.EncodeToString(proof3))


	/***  **/
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")

	/** 分别计算出三组不同的 随机数和证明 */
	index1, proof1 = k.Evaluate(m1)
	index2, proof2 = k.Evaluate(m2)
	index3, proof3 = k.Evaluate(m3)

	fmt.Println("随机数1：", hex.EncodeToString(index1[:]))
	fmt.Println("随机数2：", hex.EncodeToString(index2[:]))
	fmt.Println("随机数3：", hex.EncodeToString(index3[:]))

	fmt.Println("证明1：", hex.EncodeToString(proof1))
	fmt.Println("证明2：", hex.EncodeToString(proof2))
	fmt.Println("证明3：", hex.EncodeToString(proof3))

	fmt.Println("#################################################################################################################################")

	/** 分别计算出三组不同的 随机数和证明 */
	index1, proof1 = k.Evaluate(m1)
	index2, proof2 = k.Evaluate(m2)
	index3, proof3 = k.Evaluate(m3)

	fmt.Println("随机数1：", hex.EncodeToString(index1[:]))
	fmt.Println("随机数2：", hex.EncodeToString(index2[:]))
	fmt.Println("随机数3：", hex.EncodeToString(index3[:]))

	fmt.Println("证明1：", hex.EncodeToString(proof1))
	fmt.Println("证明2：", hex.EncodeToString(proof2))
	fmt.Println("证明3：", hex.EncodeToString(proof3))

	/** 遍历一个 匿名结构体 切片 */
	for i, tc := range []struct {
		m     []byte
		index [32]byte
		proof []byte
		err   error
	}{
		/** 前面三组都是没问题的 */
		{m1, index1, proof1, nil},
		{m2, index2, proof2, nil},
		{m3, index3, proof3, nil},

		/** 故意做些错误的信息去校验 */
		{m3, index3, proof2, nil},
		{m3, index3, proof1, p256.ErrInvalidVRF},
	} {

		index, err := pk.ProofToHash(tc.m, tc.proof)
		if got, want := err, tc.err; got != want {
			fmt.Printf("现在是第: " + fmt.Sprint(i) + " 个结构体， " + "Err ProofToHash: %v, want %v", got, want)
		}
		if err != nil {
			continue
		}
		if got, want := index, tc.index; got != want {
			fmt.Printf("现在是第: " + fmt.Sprint(i) + " 个结构体， " + "Err ProofToInex: %x, want %x", got, want)
		}
	}

}


