package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//ch := make(chan struct{})

	node := new(Node)
	node.init()
	//go node.Set("a", "A")
	//go node.Set("b", "B")
	//go fmt.Println(node.Get("a"))
	//node.clean()
	//go node.Set("c", "C")
	//go fmt.Println(node.Get("c"))

	go func() {
		fmt.Println("开始设置原始Map")
		node.Set("a", "A")
		node.Set("b", "B")
		fmt.Println(node.Get("a"))
	}()


	go func() {
		node.clean()
		node.Set("c", "C")
		fmt.Println(node.Get("c"))
	}()
	go func() {
		fmt.Println("直接过来拿Map")
		fmt.Println(node.Get("b"))
	}()
	//<- ch
	time.Sleep(4*time.Second)
}

type Node struct {
	lock 	sync.RWMutex
	round 	round
}

type round map[string]string

func(n *Node) init(){
	n.round = make(round, 0)
}

func (n *Node) Get (key string) string {
	n.lock.RLock()
	defer n.lock.RUnlock()
	a := n.round[key]
	return a
}


func (n *Node) Set(key, val string) {
	n.lock.Lock()
	defer n.lock.Unlock()
	n.round[key] = val
}

func (n *Node) clean(){
	fmt.Println("开始删除...Map")
	n.lock.Lock()
	n.round = make(round, 0)
	n.lock.Unlock()
}