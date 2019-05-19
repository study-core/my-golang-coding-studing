package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	cache := make(map[string]*stateCache, 0)

	st1 := stateDB{
		Name: "stA",
	}

	stCache := &stateCache{
		st: 	st1,
	}

	cache["a"] = stCache


	// read
	ca := cache["a"]
	s1 := &ca.st
	s1.Name = "B"

	by, _ := json.Marshal(cache["a"].st)
	fmt.Println(string(by))


	//st := &stateDB{
	//	Name: "A",
	//}
	//
	//(*st).Name = "B"
	//fmt.Println(st.Name)
}

type stateCache struct {
	st 		stateDB
}

type stateDB struct {
	Name 		string
}