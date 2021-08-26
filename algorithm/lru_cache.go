package main

import (
	"errors"
	"fmt"
)

func main() {
	cache, _ := NewLruCache(3)

	cache.Put([]byte("A"), []byte("a"))
	cache.Put([]byte("B"), []byte("b"))
	cache.Put([]byte("C"), []byte("c"))
	cache.Put([]byte("D"), []byte("d"))


	a := cache.get([]byte("A"))
	d := cache.get([]byte("D"))

	fmt.Println("得到:", string(a), "和", string(d))

}

type lru_node struct {
	Key   hash
	Value []byte
	Pre   *lru_node
	Next  *lru_node
}

type NodeMap      map[hash]*lru_node

type lru_cache struct {
	Capacity    uint64
	Count 		uint64

	Nodes      NodeMap  // 存放 (hash(key) -> node)

	// 头和尾 不是 有数据的 node， 只是 空node
	Head 		*lru_node
	Tail 		*lru_node
}

const HashLength  = 32

type hash [HashLength]byte

func BytesToHash(b []byte) hash {
	var h hash
	h.SetBytes(b)
	return h
}

func (h *hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}

func NewLruCache (capacity uint64) (*lru_cache, error) {
	if capacity == 0 {
		return nil, errors.New("Failed to new lru cache, the capacity param is invalid")
	}

	cache := &lru_cache{
		Capacity: capacity,
		Count:    0,
		Head:     nil,
		Tail:     nil,
		Nodes:  make(NodeMap),
	}

	// 头 和 尾 都是 空节点
	head :=  new(lru_node)
	tail := new(lru_node)

	head.Next = tail
	tail.Pre = head

	cache.Head = head
	cache.Tail = tail

	return cache, nil
}

func (cache *lru_cache) Put(key, value []byte) {

	nd, ok := cache.Nodes[BytesToHash(key)]

	if !ok {
		if cache.Count == cache.Capacity {
			cache.removeNode()  // 移除 队尾 的 node  (Tail的上一个)
		}

		nd = new(lru_node)
		nd.Key = BytesToHash(key)
		nd.Value = value

		cache.addNode(nd)  // 添加到队头 (Head的下一个)

	} else {
		nd.Value = value
		cache.moveNodeToHead(nd)
	}
}


func (cache *lru_cache) get(key []byte) []byte {

	nd, ok := cache.Nodes[BytesToHash(key)]
	if ok {
		cache.moveNodeToHead(nd) // 每次找到时， 都 移动到 队头
		return nd.Value
	}
	return nil
}

// 移除 队尾的 node
func (cache *lru_cache) removeNode () {

	nd := cache.Tail.Pre
	cache.removeFromList(nd)

	delete(cache.Nodes, nd.Key)
	cache.Count --
}

// 移除 当前节点
func (cache *lru_cache) removeFromList(node *lru_node) {
	pre := node.Pre
	next := node.Next

	next.Pre = pre
	pre.Next = next

	node.Pre = nil
	node.Next = nil
}
// 直接 添加一个节点
func (cache *lru_cache) addNode (node *lru_node) {
	cache.addToHead(node)
	cache.Nodes[node.Key] = node
	cache.Count ++
}
// 将 当前node 转移到 Head 的 下一个
func (cache *lru_cache) addToHead(node *lru_node) {
	next := cache.Head.Next

	next.Pre = node
	node.Next = next

	node.Pre = cache.Head
	cache.Head.Next = node

}

func (cache *lru_cache) moveNodeToHead (node *lru_node) {   // 将 当前node 转移到 Head 的 下一个
	cache.removeFromList(node) // 移除当前节点
	cache.addToHead(node)  // 将 当前node 转移到 Head 的 下一个
}
