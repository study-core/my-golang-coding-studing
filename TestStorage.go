package main

import (
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"encoding/json"
	"fmt"
)

func main() {

	store := make(state.Storage)
	valStore := make(state.ValueStorage)

	store["A"] =  common.BytesToHash([]byte("a"))
	valStore[common.BytesToHash([]byte("a"))] = []byte("Aa")

	store["B"] =  common.BytesToHash([]byte("b"))
	valStore[common.BytesToHash([]byte("b"))] = []byte("Bb")

	store["C"] =  common.BytesToHash([]byte("c"))
	valStore[common.BytesToHash([]byte("c"))] = []byte("Cc")


	originStore := make(state.Storage)
	originValStore := make(state.ValueStorage)

	sb, _ := json.Marshal(store)
	svb, _ := json.Marshal(valStore)
	fmt.Println("store", string(sb))
	fmt.Println("valStore", string(svb))


	for key, valueKey := range store {
		value := valStore[valueKey]

		originStore[key] = valueKey
		originValStore[valueKey] = value
	}

	ob, _ := json.Marshal(originStore)
	ovb, _ := json.Marshal(originValStore)
	fmt.Println("originStore", string(ob))
	fmt.Println("originValStore", string(ovb))


}
