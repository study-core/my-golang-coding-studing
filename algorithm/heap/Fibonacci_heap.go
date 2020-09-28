package main

// 斐波那契堆
//
// https://www.cnblogs.com/skywang12345/p/3659122.html  （很详细）
func main() {

}

/**
todo 斐波那契堆(Fibonacci heap)是一种可合并堆，可用于实现合并优先队列。     它比二项堆具有更好的平摊分析性能，它的合并操作的时间复杂度是O(1)。

		与二项堆一样，它也是由一组堆最小有序树组成，并且是一种可合并堆。
todo	与二项堆不同的是，斐波那契堆中的树不一定是二项树；而且二项堆中的树是有序排列的，但是斐波那契堆中的树都是有根而无序的。


todo  斐波那契堆是由一组最小堆组成，这些最小堆的根节点组成了双向链表(后文称为"根链表")；斐波那契堆中的最小节点就是"根链表中的最小节点"！
 */

type FibonacciNode struct {
	key          int            // 关键字(键值)
	degree       int            // 度数  (节点 拥有 子节点的个数)
	leftBrother  *FibonacciNode // 左兄弟
	rightBrother *FibonacciNode // 右兄弟
	child        *FibonacciNode // 第一个孩子节点
	parent       *FibonacciNode // 父节点
	marked       bool           // 是否被删除第一个孩子   (marked在删除节点时有用)
}

/*
 * 将node堆结点加入root结点之前(循环链表中)
 *   a …… root
 *   a …… node …… root
*/
func addFibonacciNode(root, node *FibonacciNode) {
	node.leftBrother = root.leftBrother
	root.leftBrother.rightBrother = node
	node.rightBrother = root
	root.leftBrother = node
}

type FibonacciHeap struct {
	keyCount int // 堆中 节点的总数
	root     *FibonacciNode  // root 即时整个堆中的最小 节点
}

// todo  插入
func (h *FibonacciHeap) insert (key int) {
	node := &FibonacciNode{key: key}
	if h.keyCount == 0 {
		h.root = node
	} else {
		addFibonacciNode(h.root, node)
		if node.key < h.root.key {
			h.root = node
		}
	}
	h.keyCount ++
}

// todo  堆合并
func (h *FibonacciHeap) union (other *FibonacciHeap) {

}