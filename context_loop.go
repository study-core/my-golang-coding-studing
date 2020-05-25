package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {

	vm := NewVM(Context{})
	ctx := context.Background()
	var cancelFn context.CancelFunc
	if true {
		timeout := time.Duration(1) * time.Second
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	} else {
		ctx, cancelFn = context.WithCancel(ctx)
	}
	defer cancelFn()
	vm.Ctx = ctx

	//fmt.Println("ret:", recursiveCall(vm, 100))
	fmt.Println("ret:", vm.Call())

}

var abortErr = errors.New("just be aborted")


type Context struct {
	Ctx context.Context
}

type VM struct {
	Context
	interpreters []Interpreter
	interpreter  Interpreter
}

func NewVM(ctx Context) *VM {
	vm := &VM{
		Context:      ctx,
		interpreters: make([]Interpreter, 0, 1),
	}
	vm.interpreters = append(vm.interpreters, NewEVMInterpreter(vm))
	vm.interpreters = append(vm.interpreters, NewWASMInterpreter(vm))
	vm.interpreter = vm.interpreters[0]
	return vm
}

//func recursiveCall (vm *VM, num int64) error {
//	if num > 0 {
//		return recursiveCall(vm, num -1)
//	}
//	return vm.Call()
//}

func (vm *VM) Call () error {
	return run(vm)
}



func run(vm *VM) error {

	for _, interpreter := range vm.interpreters {
		if interpreter.CanRun() {
			if vm.interpreter != interpreter {
				// Ensure that the interpreter pointer is set back
				// to its current value upon return.
				defer func(i Interpreter) {
					vm.interpreter = i
				}(vm.interpreter)

				vm.interpreter = interpreter
			}
			fmt.Println("进入这里了...")
			return interpreter.Run()
		}
	}
	return nil
}

type Interpreter interface {
	Run() error
	CanRun() bool
}


type evmInterpreter struct {
	vm *VM
	abort bool
}
func NewEVMInterpreter(vm *VM) *evmInterpreter {
	return &evmInterpreter{
		vm:      vm,
	}
}

func (in *evmInterpreter) Run ()  error {

	fmt.Println("进入 evm 啦...")
	go func(ctx context.Context) {
		<-ctx.Done()
		in.abort = true
	}(in.vm.Ctx)

	time.Sleep(time.Second *3)
	if in.abort {
		fmt.Println("evm 超时 ...")
		return fmt.Errorf("evm: %v", abortErr)
	}
	return nil
}

func (in *evmInterpreter) CanRun() bool {
	return false
}



type wasmInterpreter struct {
	vm *VM
}

func NewWASMInterpreter(vm *VM) *wasmInterpreter {
	return &wasmInterpreter{
		vm: vm,
	}
}

func (in *wasmInterpreter)  Run() error {
	creator, err := NewWasmEngineCreator("Wagon")
	if err != nil {
		return err
	}
	engine := creator.Create(in.vm)

	err = engine.run()
	return  err
}

func (in *wasmInterpreter) CanRun() bool {
	return true
}


var creators = map[string]wasmEngineCreator{
	"Wagon": &wagonEngineCreator{},
}

func NewWasmEngineCreator(vm string) (wasmEngineCreator, error) {

	if creator, ok := creators[vm]; ok {
		return creator, nil
	}
	return nil, fmt.Errorf("unsupport wasm type: %d", vm)
}

type wasmEngineCreator interface {
	Create(evm *VM) wasmEngine
}




type wagonEngineCreator struct {
}

func (w *wagonEngineCreator) Create(vm *VM) wasmEngine {
	return &wagonEngine{
		vm:      vm,
	}
}

type wasmEngine interface {
	run() error
}

type wagonEngine struct {
	vm      *VM
	abort bool
}

func (engine *wagonEngine) run() error {

	fmt.Println("进入 wasm ...")

	go func(ctx context.Context) {
		<-ctx.Done()
		engine.abort = true
	}(engine.vm.Ctx)

	time.Sleep(time.Second*3)
	if engine.abort {
		fmt.Println("wasm 超时 ...")
		return fmt.Errorf("wasm: %v", abortErr)
	}
	return  nil
}

