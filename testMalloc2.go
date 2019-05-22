package main

import (
	"fmt"
	"time"

	"github.com/glycerine/offheap"
)

func main() {
	sliceLength := 1 << 10

	// 1MB
	buffLength := 1 << 20

	datas := make([][]byte, sliceLength)
	freeList := make([]*offheap.MmapMalloc, sliceLength)

	// Malloc
	for i := 0; i < sliceLength; i++ {
		mmap := offheap.Malloc(int64(buffLength), "")
		buff := mmap.Mem
		for i := 0; i < buffLength; i++ {
			buff[i] = 'G'
		}
		datas[i] = buff
		freeList[i] = mmap
	}

	// free
	for _, v := range freeList {
		v.Free()
	}

	fmt.Println("Cancel...")
	<-time.After(time.Hour)
}