package testfeed

import "sync"

type Myfeed struct {
	one 	sync.Once
	names 	[]string
}

func (f *Myfeed) init() {
	f.names = []string{"I Am  Gavin !!!!"}

}

func (f *Myfeed) GetName () string {
	f.one.Do(f.init)
	if len(f.names) == 0 {
		return ""
	}
	last := f.names[len(f.names)-1]
	f.names = f.names[:len(f.names)-1]
	return last
}

func (f *Myfeed) SetName (member string) {
	f.one.Do(f.init)
	f.names = append(f.names, member)
}
