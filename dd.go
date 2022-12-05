package main

import (
	"bytes"
	//"carrier-go/common/bytesutil"
	"encoding/binary"
	"fmt"
)

func main() {
	bf := make([]byte, 8)
	binary.BigEndian.PutUint64(bf, 33)
	a := bf[:]
	//a := bytesutil.Uint64ToBytes(33)
	x := true

	var b []byte
	if x {
		b = []byte{byte(1)}
	} else {
		b = []byte{byte(0)}
	}

	var buf bytes.Buffer
	buf.Write(a)
	buf.Write(b)

	fmt.Println(buf.Bytes())
	c := make([]byte, 0)
	c = append(c, a...)
	c = append(c, b...)
	fmt.Println(c)

}