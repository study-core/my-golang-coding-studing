package main

import (
	"encoding/json"
	"fmt"
	"go1.14.6-analysis/src/container/heap"
)

type KallyMem struct {
	Name  string
	Term  uint32
}

type KallyHeap  []*KallyMem

func (h KallyHeap) Len() int           { return len(h) }
func (h KallyHeap) Less(i, j int) bool { return h[i].Term > h[j].Term } // term:  a.3 > c.2 > b.1,  So order is: a c b
func (h KallyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *KallyHeap) Push(x interface{}) {
	*h = append(*h, x.(*KallyMem))
}

func (h *KallyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *KallyHeap) IncreaseTerm () {
	for i := range *h {
		(*h)[i].Term++
	}
}
func (h *KallyHeap) DecreaseTerm () {
	for i := range *h {
		if (*h)[i].Term != 0 {
			(*h)[i].Term--
		}
	}
}

func main() {

	a := &KallyMem{
		Name: "a",
		Term: 4,
	}
	b := &KallyMem{
		Name: "b",
		Term: 1,
	}
	c := &KallyMem{
		Name: "c",
		Term: 2,
	}
	d :=  &KallyMem{
		Name: "d",
		Term: 34,
	}
	queue := new(KallyHeap)
	//queue := make(KallyHeap, 0)
	heap.Push(queue, c)
	heap.Push(queue, a)
	heap.Push(queue, d)
	heap.Push(queue, b)


	x := &KallyMem{
		Name: "x",
		Term: 104,
	}
	y := &KallyMem{
		Name: "y",
		Term: 201,
	}
	z := &KallyMem{
		Name: "z",
		Term: 428,
	}
	starveQueue := new(KallyHeap)
	heap.Push(starveQueue, x)
	heap.Push(starveQueue, y)
	heap.Push(starveQueue, z)

	queueB, _ := json.Marshal(queue)
	fmt.Println("queue:", string(queueB))

	starveQueueB, _ := json.Marshal(starveQueue)
	fmt.Println("starveQueue:", string(starveQueueB))

	fmt.Println("\n---------------------------------------------------------\n")

	i := 0
	for {
		if i == queue.Len() {
			break
		}
		bullet := (*(queue))[i]
		bullet.Term++

		// When the task in the queue meets hunger, it will be transferred to starveQueue
		if bullet.Term >= 5 {
			heap.Push(starveQueue, bullet)
			heap.Remove(queue, i)
			i = 0
			continue
		}
		(*(queue))[i] = bullet
		i++
	}

	queueB2, _ := json.Marshal(queue)
	fmt.Println("queue:", string(queueB2))

	starveQueueB2, _ := json.Marshal(starveQueue)
	fmt.Println("starveQueue:", string(starveQueueB2))

	//x := heap.Pop(queue)
	//fmt.Println(x)
	//
	//queue.IncreaseTerm()
	//y := heap.Pop(queue)
	//fmt.Println(y)
	//
	//queue.IncreaseTerm()
	//z := heap.Pop(queue)
	//fmt.Println(z)
}