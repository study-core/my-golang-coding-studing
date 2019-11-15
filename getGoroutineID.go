package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("main", GoID())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i, GoID())
		}()
	}
	wg.Wait()
}

// import (
// 	"runtime"
// 	"time"
// 	"github.com/petermattis/goid"
// 	"fmt"
// )
// func BenchmarkASM() {
// 	idArr := make([]int64, 0)
// 	startTime := time.Now().UnixNano()
// 	for i := 0; i < 10000; i++ {
// 		idArr = append(idArr, goid.Get())
// 	}
// 	endTime := time.Now().UnixNano()
// 	fmt.Println(fmt.Sprint(startTime - endTime) + "\n" + fmt.Sprint(idArr))
// }

// func BenchmarkSlow() {
// 	idArr := make([]int64, 0)
// 	var buf [64]byte
// 	startTime := time.Now().UnixNano()

// 	for i := 0; i < 10000; i++ {
// 		idArr = append(idArr, goid.ExtractGID(buf[:runtime.Stack(buf[:], false)]))
// 	}
// 	endTime := time.Now().UnixNano()
// 	fmt.Println(fmt.Sprint(startTime - endTime) + "\n" + fmt.Sprint(idArr))
// }

// func main() {
// 	BenchmarkASM()
// 	BenchmarkSlow()
// }
