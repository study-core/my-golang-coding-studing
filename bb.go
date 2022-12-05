package main

import (
	"RosettaFlow/carrier-go/common/bytesutil"
	"fmt"
)

func main() {

	//n := &Naa{
	//	when: 10,
	//}
	//
	//n.SetFn(func(when uint64) {
	//	fmt.Println("old when in callback::::::::::::", n.GetWhen())
	//	n.SetWhen(when)
	//	fmt.Println("new when in callback::::::::::::", n.GetWhen())
	//})
	//
	//fmt.Println("old when before call fn and set", n.GetWhen())
	//
	//n.SetWhen(2)
	//
	//fmt.Println("old when before call fn and after set", n.GetWhen())
	//
	//f := n.fn
	//f(12)
	//fmt.Println("old when after call fn", n.GetWhen())


	v := []byte{}


	val := bytesutil.BytesToUint32(v)
	val |= OnConsensusExecuteTaskStatus.Uint32()

	fmt.Println(val)

	v = bytesutil.Uint32ToBytes(OnConsensusExecuteTaskStatus.Uint32())

	val = bytesutil.BytesToUint32(v)
	val |= OnRunningExecuteStatus.Uint32()

	fmt.Println(val)
}


type Naa struct {
	when  uint64

	fn    func (when uint64)
}

func (n *Naa) GetWhen() uint64 { return n.when }
func (n *Naa) SetWhen(when uint64) { n.when = when }
func (n *Naa) SetFn(f func (when uint64)) { n.fn = f }



const (
	/**
	######   ######   ######   ######   ######
	#   THE LOCAL NEEDEXECUTE TASK STATUS    #
	######   ######   ######   ######   ######
	*/
	OnConsensusExecuteTaskStatus   LocalTaskExecuteStatus = 1 << iota 	// 0001: the execute task is on consensus period now.
	OnRunningExecuteStatus                               				// 0010: the execute task is running now.
	OnTerminingExecuteStatus                               				// 0010: the execute task is termining now.
	UnKnownExecuteTaskStatus       = 0        							// 0000: the execute task status is unknown.
)

type LocalTaskExecuteStatus uint32

func (s LocalTaskExecuteStatus) Uint32() uint32 { return uint32(s) }
