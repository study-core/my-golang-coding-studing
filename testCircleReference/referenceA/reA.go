package referenceA

import (
	//"myCryto-study/testCircleReference/referenceB"
	"fmt"
)

func ATest1(){
	//fn := <- FuncCh
	//fn()
	if nil != Func {
		Func()
	}
}

var (
	FuncCh =  make(chan func ())
	Func func ()
)
func Atest2 () {
	fmt.Println("我是A2")
}