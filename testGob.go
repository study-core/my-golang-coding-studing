package main

import (
	"encoding/gob"
	"bytes"
	"fmt"
)

func main() {


	type P struct {
		X, Y, Z int
		Name    string
	}


	var buf bytes.Buffer
	en := gob.NewEncoder(&buf)
	p := &P{
		X:	1,
		Y:	2,
		Z:	3,
		Name: "my",
	}
	en.Encode(p)

	data := buf.Bytes()

	fmt.Println(data)

	var sb bytes.Buffer
	sb.Write(data)

	de := gob.NewDecoder(&sb)
	var pp P
	de.Decode(&pp)
	fmt.Println(pp)
}
