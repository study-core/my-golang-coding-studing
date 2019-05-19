//package main
//
//import (
//	"sync"
//	"fmt"
//)
//
////下面的迭代会有什么问题？
//
//type threadSafeSet struct {
//	sync.RWMutex
//	s []interface{}
//}
//
//func (set *threadSafeSet) Iter() <-chan interface{} {
//	 ch := make(chan interface{}) // 解除注释看看！
//	//ch := make(chan interface{},len(set.s))
//	go func() {
//		set.RLock()
//
//		for elem,value := range set.s {
//			ch <- elem
//			println("Iter:",elem,value)
//		}
//
//		close(ch)
//		set.RUnlock()
//
//	}()
//	return ch
//}
//
//func main()  {
//
//	th:=threadSafeSet{
//		s:[]interface{}{"1","2"},
//	}
//	v:=<-th.Iter()
//	fmt.Sprintf("%s%v","ch",v)
//}
//


package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	//var peo People = Stduent{}
	var peo People = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
