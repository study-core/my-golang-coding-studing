package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	vm := &evm{}

	fmt.Println("开始测试：", vm.abort)

	var cancelFn context.CancelFunc
	timeout := 2 * time.Millisecond
	ctx := context.Background()

	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	} else {
		ctx, cancelFn = context.WithCancel(ctx)
	}
	defer cancelFn()

	go func(ctx context.Context) {
		<-ctx.Done()
		vm.terminate()
		fmt.Println("最后再次确认下vm.abort: ", vm.abort)
	}(ctx)

	start := time.Now()
	vm.doSomething()
	fmt.Println("结果vm.abort:", vm.abort, "耗时:", time.Since(start))

	//flag := false
	//
	//printString(!flag)
}

type evm struct {
	abort  bool
}

func (vm *evm) doSomething () {
	count := 1
	for !vm.abort && count <= 666666 {
		fmt.Println("I do something: ", count)
		count++
	}
}

func (vm *evm) terminate () {
	vm.abort = true
}


func printString(flag bool) {
	fmt.Println("mark:", flag)
}