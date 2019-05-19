package testMychain

import (
	"myCryto-study/testfeed"
	"fmt"
	"runtime/debug"
)

type Mychain struct {
	myFeed	testfeed.Myfeed
}

func (m *Mychain) GetName () string {
	return m.myFeed.GetName()
}

func (m *Mychain) SetName(member string) {
	panic("fucking Emma's mouth")
	fmt.Println("stack", string(debug.Stack()))
	m.myFeed.SetName(member)
}
