package main

import (
	"encoding/json"
	"fmt"
)

func main() {


	storage := make(map[string][]byte, 0)

	add := func (m map[string][]byte, key string, value []byte) {
		m[key] = value
	}

	storage["a"] = []byte("A")
	storage["b"] = []byte("B")
	add(storage, "c", []byte("C"))
	add(storage, "e", []byte("E"))
	add(storage, "f", []byte("F"))
	add(storage, "g", []byte("G"))
	add(storage, "a", []byte("AA"))

	fmt.Println("len:", len(storage))
	b, _ := json.Marshal(storage)
	fmt.Println("storage:", string(b))

	for k, v := range storage {
		fmt.Println("Key:", k, "Value:", string(v))
	}
}
