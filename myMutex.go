package main

import (
	"fmt"
	"time"
)

func main() {

	 m := NewMyMutex()

	 go func() {
	 	m.SetNum(2)
	 }()

	 go func() {
	 	m.SetNum(3)
	 }()

	go func() {
		m.SetNum(4)
	}()
	time.Sleep(4 * time.Second)
}

var i int = 1

func (m *MyMutex) SetNum(num int){
	//m.Lock()
	//defer m.Unlock()
	//fmt.Println("修改前:", i)
	//i = num
	//fmt.Println("修改后:", i)
	//if m.TryLock() {
	if m.TryLockTimeOut(1 * time.Second) {
		fmt.Println("修改前:", i)
		i = num
		fmt.Println("修改后:", i)
		m.Unlock()
	}
}

type MyMutex struct {
	ch chan struct{}
}

func NewMyMutex() *MyMutex {
	return &MyMutex{make(chan struct{},1)}
}

func (m *MyMutex) Lock() {
	m.ch <- struct{}{}
}

func (m *MyMutex) Unlock() {
	<-m.ch
}

func (m *MyMutex) TryLock() bool {
	select {
	case m.ch <- struct{}{}:
		return true
	default:
	}
	return false
}

func (m *MyMutex) TryLockTimeOut(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case m.ch <- struct{}{}:
		timer.Stop()
		return true
	case <-time.After(timeout):
	}
	return false
}


func (m *MyMutex) IsLocked() bool {
	return len(m.ch) >0
}