// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package mytoken

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"math/big"
	"testing"
	"time"

	"github.com/Dai0522/go-hash/bloomfilter"

	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/secp256k1"
)

func TestBloom(t *testing.T) {
	positive := []string{
		"testtest",
		"test",
		"hallo",
		"other",
	}
	negative := []string{
		"tes",
		"lo",
	}

	var bloom types.Bloom
	for _, data := range positive {
		startTime := time.Now()
		bloom.Add(new(big.Int).SetBytes([]byte(data)))
		fmt.Println("time", time.Since(startTime))
	}

	for _, data := range positive {
		if !bloom.TestBytes([]byte(data)) {
			t.Error("expected", data, "to test true")
		}
	}
	for _, data := range negative {
		if bloom.TestBytes([]byte(data)) {
			t.Error("did not expect", data, "to test true")
		}
	}
	result := make(map[float32]int)
	for i := 0; i < 100; i++ {
		fpp := testBloom9()
		if v, ok := result[fpp]; !ok {
			result[fpp] = 1
		} else {
			result[fpp] = v + 1
		}
	}
	for k, v := range result {
		fmt.Println("fpp：", k, "count：", v)
	}
}

func testBloom9() float32 {
	var tb types.Bloom
	err := 0
	number := 1075
	curve := secp256k1.S256()
	addrMap := make(map[string]int)
	for i := 0; i < number; i++ {
		privKey, _ := ecdsa.GenerateKey(curve, rand.Reader)
		//nodeId := discover.PubkeyID(&privKey.PublicKey)
		addr := crypto.PubkeyToAddress(privKey.PublicKey)
		if v, ok := addrMap[addr.String()]; ok {
			fmt.Println("相同元素", addr.String())
			addrMap[addr.String()] = v + 1
		} else {
			addrMap[addr.String()] = 0
		}
		if tb.TestBytes(addr.Bytes()) {
			err++
		} else {
			tb.Add(new(big.Int).SetBytes(addr.Bytes()))
		}
	}
	return float32(err) / float32(number)
}

func TestBloom_murmur3(t *testing.T) {
	result := make(map[float32]int)
	for i := 0; i < 1000; i++ {
		fpp := tBloomMurmur3()
		if v, ok := result[fpp]; !ok {
			result[fpp] = 1
		} else {
			result[fpp] = v + 1
		}
	}
	for k, v := range result {
		fmt.Println("fpp：", k, "count：", v)
	}
}

func tBloomMurmur3() float32 {
	number := 25
	bf, err := bloomfilter.New(uint64(number), 0.0001)
	if nil != err {
		panic(err)
	}
	errNumber := 0
	curve := secp256k1.S256()
	addrMap := make(map[string]int)
	for i := 0; i < number; i++ {
		privKey, _ := ecdsa.GenerateKey(curve, rand.Reader)
		//nodeId := discover.PubkeyID(&privKey.PublicKey)
		addr := crypto.PubkeyToAddress(privKey.PublicKey)
		if v, ok := addrMap[addr.String()]; ok {
			fmt.Println("相同元素", addr.String())
			addrMap[addr.String()] = v + 1
		} else {
			addrMap[addr.String()] = 0
		}
		if bf.MightContain(addr.Bytes()) {
			errNumber++
		} else {
			bf.Put(addr.Bytes())
			bf.Serialized()
		}
	}
	return float32(errNumber) / float32(number)
}

/*
import (
	"testing"

	"github.com/PlatONnetwork/PlatON-Go/core/state"
)

func TestBloom9(t *testing.T) {
	testCase := []byte("testtest")
	bin := LogsBloom([]state.Log{
		{testCase, [][]byte{[]byte("hellohello")}, nil},
	}).Bytes()
	res := BloomLookup(bin, testCase)

	if !res {
		t.Errorf("Bloom lookup failed")
	}
}


func TestAddress(t *testing.T) {
	block := &Block{}
	block.Coinbase = common.Hex2Bytes("22341ae42d6dd7384bc8584e50419ea3ac75b83f")
	fmt.Printf("%x\n", crypto.Keccak256(block.Coinbase))

	bin := CreateBloom(block)
	fmt.Printf("bin = %x\n", common.LeftPadBytes(bin, 64))
}
*/
