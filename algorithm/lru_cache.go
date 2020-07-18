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

	Nodes      NodeMap

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
			cache.removeNode()
		}

		nd = new(lru_node)
		nd.Key = BytesToHash(key)
		nd.Value = value

		cache.addNode(nd)

	} else {
		nd.Value = value
		cache.moveNodeToHead(nd)
	}
}


func (cache *lru_cache) get(key []byte) []byte {

	nd, ok := cache.Nodes[BytesToHash(key)]
	if ok {
		cache.moveNodeToHead(nd)
		return nd.Value
	}
	return nil
}


func (cache *lru_cache) removeNode () {

	nd := cache.Tail.Pre
	cache.removeFromList(nd)

	delete(cache.Nodes, nd.Key)
	cache.Count --
}

func (cache *lru_cache) removeFromList(node *lru_node) {
	pre := node.Pre
	next := node.Next

	next.Pre = pre
	pre.Next = next

	node.Pre = nil
	node.Next = nil
}

func (cache *lru_cache) addNode (node *lru_node) {
	cache.addToHead(node)
	cache.Nodes[node.Key] = node
	cache.Count ++
}

func (cache *lru_cache) addToHead(node *lru_node) {
	next := cache.Head.Next

	next.Pre = node
	node.Next = next

	node.Pre = cache.Head
	cache.Head.Next = node

}

func (cache *lru_cache) moveNodeToHead (node *lru_node) {
	cache.removeFromList(node)
	cache.addToHead(node)
}
