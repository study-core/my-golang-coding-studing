package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
)

type GavinMem struct {
	Name  string
	Term  uint32
}

type GavinHeap  []*GavinMem

func (h GavinHeap) Len() int           { return len(h) }
func (h GavinHeap) Less(i, j int) bool { return h[i].Term > h[j].Term } // term:  a.3 > c.2 > b.1,  So order is: a c b
func (h GavinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *GavinHeap) Push(x interface{}) {
	*h = append(*h, x.(*GavinMem))
}

func (h *GavinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *GavinHeap) IncreaseTerm () {
	for i := range *h {
		(*h)[i].Term++
	}
}
func (h *GavinHeap) DecreaseTerm () {
	for i := range *h {
		if (*h)[i].Term != 0 {
			(*h)[i].Term--
		}
	}
}

func main() {

	a := &GavinMem{
		Name: "a",
		Term: 4,
	}
	b := &GavinMem{
		Name: "b",
		Term: 1,
	}
	c := &GavinMem{
		Name: "c",
		Term: 2,
	}
	d :=  &GavinMem{
		Name: "d",
		Term: 34,
	}
	queue := new(GavinHeap)
	//queue := make(KallyHeap, 0)
	heap.Push(queue, c)
	heap.Push(queue, a)
	heap.Push(queue, d)
	heap.Push(queue, b)



	queueB, _ := json.Marshal(queue)
	fmt.Println("queue:", string(queueB))

	fmt.Println("\n---------------------------------------------------------\n")

	// 遍历 queue ，将 term > 5 的移除， 并加入 starveQueue 中
	i := 0
	for {
		if i == queue.Len() {
			break
		}
		bullet := (*(queue))[i]

		// When the task in the queue meets hunger, it will be transferred to starveQueue
		if bullet.Term == 4 {
			bulletByte, _ := json.Marshal(bullet)
			fmt.Println("i is ", i, "bullet:", string(bulletByte))
			heap.Remove(queue, i)
			i = 0
			continue
		}
		(*(queue))[i] = bullet
		i++
	}

	queueB2, _ := json.Marshal(queue)
	fmt.Println("queue:", string(queueB2))


}