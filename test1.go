package main

import (
	"fmt"
	//"runtime"
	//"sync"
	"errors"
)

//func main() {
//	const (
//		A int = 1 << iota   // 	0001
//		B					//	0010
//		C					//	0100
//	)
//
//	fmt.Println(A, B, C)  	//	1 2 4
//
//	fmt.Println((B & B) == 0)
//
//	fmt.Println(B & B)
//
//	fmt.Println(A != B)
//}


//func main() {
//	runtime.GOMAXPROCS(1)
//	wg := sync.WaitGroup{}
//	wg.Add(20)
//	for i := 0; i < 10; i++ {
//		go func() {
//			fmt.Println("A: ", i)
//			wg.Done()
//		}()
//	}
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			fmt.Println("B: ", i)
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

//func main() {
//	set := &threadSafeSet{}
//	set.s = []int{1, 4, 8}
//	re := set.Iter()
//	for v := range re {
//		fmt.Println(v)
//	}
//}
//
//type threadSafeSet struct {
//	s []int
//	sync.RWMutex
//}
//func (set *threadSafeSet) Iter() <-chan interface{} {
//	ch := make(chan interface{})
//	go func() {
//		set.RLock()
//
//		for _, elem := range set.s {
//			ch <- elem
//		}
//		close(ch)
//		set.RUnlock()
//
//	}()
//	return ch
//}




//type threadSafeSet struct {
//	sync.RWMutex
//	s []interface{}
//}
//
//func (set *threadSafeSet) Iter() <-chan interface{} {
//	// ch := make(chan interface{}) // 解除注释看看！
//	ch := make(chan interface{},len(set.s))
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
//func main() {
//
//	th := threadSafeSet{
//		s: []interface{}{"1", "2"},
//	}
//	v := <-th.Iter()
//	//fmt.Sprintf("%s%v", "ch", v)
//	fmt.Println(v)
//}





var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	err = nil
	if reallyDoIt {
		result, err := tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string,error)  {
	return "",ErrDidNotWork
}

func main() {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
}