package main

import "fmt"

func main() {
	items := []int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}
	fmt.Println(TopK(items, 7)) // [27 43 97 100 202 506 1000]

	arr := []int{1, 5, 9, 22, 56, 78, 88, 94}
	fmt.Println(Bserch(arr,22)) // 3
	fmt.Println(SubSearch(arr,22))
	fmt.Println(fibSearch(arr,22))
}


// 二分练习
func Bserch (arr []int, key int) int {
	left, right := 0, len(arr) - 1
	for left <= right {
		mid := left & right + (left^right)>>1
		if key < arr[mid] {
			right = mid - 1
		}else if key > arr[mid] {
			left = mid + 1
		}else {
			return mid
		}
	}
	return -1
}

// 差值练习
func SubSearch(arr []int, key int) int {
	left, right := 0, len(arr) - 1
	for left <= right {
		mid := left + int((right - left) * ((key - arr[left])/(arr[right] - arr[left])))
		if key < arr[mid] {
			right = mid - 1
		}else if key > arr[mid] {
			left = mid + 1
		}else {
			return mid
		}
	}
	return -1
}

// 斐波那契练习
func fibSearch(arr []int, key int) int {
	fibArr := make([]int, 0)
	for i := 0; i <= 36; i++ {
		fibArr = append(fibArr, newFibMumber(i))
	}

	var (
		n = len(arr)
		k = 0
	)
	for n > fibArr[k] - 1 { // len(arr) = F(k) - 1
		k += 1
	}

	for i := n; i < fibArr[k]; i++ {
		arr = append(arr, 0)
	}
	left, right := 0, n - 1
	for left <= right {

		mid := left + fibArr[k - 1] - 1 // 死记

		if key < arr[mid] {
			right = mid - 1
			k -= 1
		}else if key > arr[mid] {
			left = mid + 1
			k -= 2
		}else {
			if mid < n {
				return mid
			} else {
				return n - 1
			}
		}
	}
	return -1
}

func newFibMumber(n int) int {
	if n < 2 {
		return n
	}
	arr := [3]int{0, 1, 0}
	for i := 2; i <= n; i++ {
		arr[2] = arr[0] + arr[1]

		arr[0] = arr[1]
		arr[1] = arr[2]
	}
	return arr[2]
}


/**
图 结构
 */
type GNode struct {
	Value interface{}
}

func (n *GNode) String() string {
	return fmt.Sprint(n.Value)
}

type ItemGraph struct {
	Nodes  	[]*GNode
	Edges	map[*GNode][]*GNode
}

func NewGraph () *ItemGraph {
	return &ItemGraph{Nodes: make([]*GNode, 0), Edges: make(map[*GNode][]*GNode, 0)}
}

func (g *ItemGraph) AddNode(node *GNode) {
	g.Nodes = append(g.Nodes, node)
}

func (g *ItemGraph) AddEdge(an, bn *GNode) {
	g.Edges[bn] = append(g.Edges[bn], an)
	g.Edges[an] = append(g.Edges[an], bn)
}

// 深度练习
type Stack struct {
	Item 	[]*GNode
}
func newStack() *Stack {
	return &Stack{ Item: make([]*GNode, 0)}
}

func (s *Stack) Push(node *GNode) {
	s.Item = append(s.Item, node)
}

func (s *Stack) Pop() *GNode {
	item := s.Item[len(s.Item) - 1]
	s.Item = s.Item[0:len(s.Item) - 1]
	return item
}

func (s *Stack) IsEmpty() bool {
	return len(s.Item) == 0
}



func (g *ItemGraph)DTS() {
	s := newStack()

	root := g.Nodes[0]

	s.Push(root)

	visited := make(map[*GNode]bool, 0)
	visited[root] = true

	for {
		if s.IsEmpty() {
			break
		}

		node := s.Pop()
		//if !visited[node] {
		//	visited[node] = true
		//}

		nodeArr := g.Edges[node]
		for _, n := range nodeArr {
			if !visited[n] {
				// 入队
				s.Push(n)
				visited[n] = true
			}
		}
		fmt.Println(node.String())
	}
}


// 广度练习
type Queue struct {
	Item []*GNode
}

func newQueue() *Queue {
	return &Queue{Item: make([]*GNode, 0)}
}

func (q *Queue)EnQueue(node *GNode) {
	q.Item = append(q.Item, node)
}

func (q *Queue) DeQueue() *GNode {
	item := q.Item[0]
	q.Item = q.Item[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.Item) == 0
}

func (g *ItemGraph) BTS (){
	q := newQueue()

	root := g.Nodes[0]

	q.EnQueue(root)

	visited := make(map[*GNode]bool, 0)

	for {
		if q.IsEmpty() {
			break
		}
		node := q.DeQueue()
		nodeArr := g.Edges[node]
		for _, n := range nodeArr {
			if !visited[n] {
				// 入队
				q.EnQueue(n)
				visited[n] = true
			}
		}

		fmt.Println(node.String())
	}


}
// 孤岛练习
func TopK(arr []int, topK int) []int {

	topArr := arr[0:topK]
	for i := len(topArr) /2 - 1; i >= 0; i -- {
		TopHeap(topArr, i, len(topArr))
	}

	for j := topK; j < len(arr); j ++ {
		if arr[j]  > topArr[0] {
			topArr[0] = arr[j]
			TopHeap(topArr, 0, len(topArr))
		}
	}
	return topArr
}

func TopHeap(arr []int, parent, length int ){
	i := parent
	for {
		left := 2 * parent + 1
		right := 2 * parent + 2

		if right < length && arr[right] < arr[i] {
			i = right
		}

		if left < length && arr[left] < arr[i] {
			i = left
		}

		if i != parent {
			arr[i], arr[parent] = arr[parent], arr[i]
			parent = i
		}else {
			break
		}
	}
}


// 二叉树练习
type Tnode struct {
	Value, Index 	int
	Left, Right 	*Tnode
}

func newNode(index, val int) *Tnode {
	return &Tnode{Index: index, Value: val}
}

type BTree struct {
	Root *Tnode
}

func newTree() *BTree {
	return &BTree{}
}

func (t *BTree) AddValNode(index, val int) {
	if nil == t.Root {
		t.Root = newNode(index, val)
		return
	}else {
		InsertNode(t.Root, newNode(index, val))
	}
}

func InsertNode (root, node *Tnode) {
	if node.Value < root.Value {
		if nil == root.Left {
			root.Left = node
		}else {
			InsertNode(root.Left, node)
		}
	}else {
		if nil == root.Right {
			root.Right = node
		}else{
			InsertNode(root.Right, node)
		}
	}
}

func (t *BTree) serch(val int) int {
	return BTreeSerch(t.Root, val)
}

func BTreeSerch(node *Tnode, val int) int{
	if node.Value == val {
		return node.Index
	}else if node.Value > val {
		return BTreeSerch(node.Left, val)
	}else {
		return BTreeSerch(node.Right, val)
	}
}