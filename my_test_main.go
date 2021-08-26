package main

import (
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"io"
)

func main() {


	ii := &InternalRLPTest{
		name: "I am kally",
		age: 27,
		cost: 2,
	}
	iib, err := rlp.EncodeToBytes(ii)
	if nil != err {
		fmt.Println("Rlp err", err)
		return
	}

	var ii2 InternalRLPTest

	if err := rlp.DecodeBytes(iib, &ii2); nil != err {
		fmt.Println("Failed to decode rlp", err)
	}

	jsonb, err := json.Marshal(ii2)
	if nil != err {
		fmt.Println("json err", err)
		return
	}
	fmt.Println(string(jsonb))
}

type InternalRLPTest struct {
	name string
	age  uint64
	cost uint64
}
type rlptest struct {
	Name string
	Age  uint64
}

// EncodeRLP implements rlp.Encoder.
func (i *InternalRLPTest) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, rlptest{Name: i.name, Age: i.age})
}

// DecodeRLP implements rlp.Decoder.
func (i *InternalRLPTest) DecodeRLP(s *rlp.Stream) error {
	var dec rlptest
	err := s.Decode(&dec)
	if err == nil {
		i.name, i.age = dec.Name, dec.Age
	}
	return err
}
