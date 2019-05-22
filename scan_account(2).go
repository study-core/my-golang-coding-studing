package main

import (
	"encoding/hex"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/rawdb"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/ethdb"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/trie"
	"os"
	"strconv"
)

func main() {
	file := os.Args[1]
	num, _ := strconv.Atoi(os.Args[3])
	address := common.HexToAddress(os.Args[21])
	addrHash := crypto.Keccak256Hash(address[:])
	fmt.Println("addrHash:", addrHash.String())

	ldb, err := ethdb.NewLDBDatabase(file, 0, 1)
	if err != nil {
		panic(err)
	}

	hash := rawdb.ReadCanonicalHash(ldb, uint64(num))
	fmt.Println("block hash:", hash.String())
	number := rawdb.ReadHeaderNumber(ldb, hash)
	fmt.Println("number:", *number)
	block := rawdb.ReadBlock(ldb, hash, uint64(num))
	if block == nil {
		panic(fmt.Sprintf("find block num:%d is nil", num))
	}

	root := block.Root()

	tr, err := trie.NewSecure(root, trie.NewDatabase(ldb), 0)
	var accountRoot common.Hash
	if err != nil {
		panic(err)
	}
	find := false
	iter := tr.NodeIterator(nil)
	for iter.Next(true) {
		if iter.Leaf() {
			var obj state.Account
			err := rlp.DecodeBytes(iter.LeafBlob(), &obj)
			if err != nil {
				panic(fmt.Sprintf("parse account error:%s", err.Error()))
			}
			value := iter.LeafKey()
			fmt.Println("account:", hex.EncodeToString(value), "nonce:", obj.Nonce)

			//accountTrie, err := trie.NewSecure(obj.Root, trie.NewDatabase(ldb), 0)
			//if err != nil {
			//	panic(fmt.Sprintf("open account err :%s", err.Error()))
			//}
			//fmt.Println("account trie:", accountTrie.Hash().String())
			//iter := accountTrie.NodeIterator(nil)
			//for iter.Next(true) {
			//	if iter.Leaf() {
			//
			//		valueKey := iter.LeafBlob()
			//		value, err := ldb.Get(valueKey)
			//		if err != nil {
			//			//panic(fmt.Sprintf("find value key:%s error:%s", hex.EncodeToString(valueKey), err.Error()))
			//		}
			//		fmt.Println("key:", hex.EncodeToString(iter.LeafKey()), hex.EncodeToString(valueKey), hex.EncodeToString(value))
			//	}
			//}
			if hex.EncodeToString(value) == hex.EncodeToString(addrHash[:]) {
				fmt.Println("find account ", address.String(), "addrHash:", addrHash.String(), "value:", hex.EncodeToString(value), "root:", obj.Root.String())
				accountRoot = obj.Root
				find = true
				break
			}
		}
	}

	if find {
		fmt.Println("find success account:", address, " root:", accountRoot.String())
		accountTrie, err := trie.NewSecure(accountRoot, trie.NewDatabase(ldb), 0)
		if err != nil {
			panic(fmt.Sprintf("open account err :%s", err.Error()))
		}
		fmt.Println("account trie:", accountTrie.Hash().String())
		iter := accountTrie.NodeIterator(nil)
		for iter.Next(true) {
			if iter.Leaf() {

				var valueKey []byte
				if err := rlp.DecodeBytes(iter.LeafBlob(), &valueKey); err != nil {
					panic(err)
				}

				seckeybuf := [43]byte{}
				secureKeyPrefix := []byte("secure-key-")
				buf := append(seckeybuf[:0], secureKeyPrefix...)
				buf = append(buf, valueKey...)

				value, err := ldb.Get(buf)
				if err != nil {
					fmt.Println("find value error key:", hex.EncodeToString(iter.LeafKey()), "valueKey:", hex.EncodeToString(valueKey), "error:", err.Error())
				} else {
					fmt.Println("key:", hex.EncodeToString(iter.LeafKey()), "valueKey:", hex.EncodeToString(valueKey), "value:", hex.EncodeToString(value))
				}
			}
		}
	} else {
		fmt.Println("not found address :", address.String())
	}
}
