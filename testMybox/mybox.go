package testMybox

import "myCryto-study/testfeed"

type Mybox struct {
	myFeed	testfeed.Myfeed
}

func (m *Mybox) GetName () string {
	return m.myFeed.GetName()
}

func (m *Mybox) SetName(member string) {
	m.myFeed.SetName(member)
}
