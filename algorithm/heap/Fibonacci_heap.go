package main
//
//
//// 斐波那契堆
////
//// https://www.cnblogs.com/skywang12345/p/3659122.html  （很详细）
//func main() {
//
//}
//
///**
//todo 斐波那契堆(Fibonacci heap)是一种可合并堆，可用于实现合并优先队列。     它比二项堆具有更好的平摊分析性能，它的合并操作的时间复杂度是O(1)。
//
//		与二项堆一样，它也是由一组堆最小有序树组成，并且是一种可合并堆。
//todo	与二项堆不同的是，斐波那契堆中的树不一定是二项树；而且二项堆中的树是有序排列的，但是斐波那契堆中的树都是有根而无序的。
//
//
//todo  斐波那契堆是由一组最小堆组成，这些最小堆的根节点组成了双向链表(后文称为"根链表")；斐波那契堆中的最小节点就是"根链表中的最小节点"！
//
//
//
//todo 斐波那契 堆的 root 链表 是一个 【双向链表】
//
//	对于  【插入】 等操作,  查到 root 的前面 其实也是 插到了 队列的 末尾, 因为假设我们 都是从 root 作为起点的话
// */
//
//type FibonacciNode struct {
//	key          int            // 关键字(键值)
//	degree       int            // 度数  (节点 拥有 子节点的个数)
//	leftBrother  *FibonacciNode // 左兄弟
//	rightBrother *FibonacciNode // 右兄弟
//	child        *FibonacciNode // 第一个孩子节点
//	parent       *FibonacciNode // 父节点
//	marked       bool           // 是否被删除第一个孩子   (marked在删除节点时有用)
//}
//
//func NewFibonacciNode (key int) *FibonacciNode {
//	node := &FibonacciNode{
//		key: key,
//		degree: 0,
//	}
//	// todo 每一个 裸节点  左右兄弟 都是自己 (因为是 双向链表)
//	node.leftBrother = node
//	node.rightBrother = node
//	return node
//}
//
///*
// * 将node堆结点加入root结点之前(循环链表中)
// *   a …… root
// *   a …… node …… root
//*/
//func addFibonacciNode(root, node *FibonacciNode) {
//
//	// todo 因为 双向 链表的原因,
//	//
//	//		对于 裸节点来说  leftBrother 和 rightBrother 都是 自己
//	//
//	//		node < = > node < = >  node    和    root < = > root < = >  root
//	//
//	//     	- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
//	//
//	//		node.leftBrother = root.leftBrother:
//	//
//	//				   root	<- node
//	//		- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
//	//
//	//		root.leftBrother.rightBrother = node:
//	//
//	//				原先   root.left.right == root  现在为  node
//	//
//	//				root -> node
//	//		- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
//	//
//	//		node.rightBrother = root
//	//		root.leftBrother = node
//	//
//	//		  ... < = > root < = > node < = > root < = > node < = > ...
//
//	node.leftBrother = root.leftBrother
//	root.leftBrother.rightBrother = node
//	node.rightBrother = root
//	root.leftBrother = node
//}
//
///*
//* 将双向链表b 链接到 双向链表a 的后面
//*/
//func catListFibonacci(a, b *FibonacciNode) {
//
//	// todo 类似 上面的 `addFibonacciNode()` 函数
//
//	tmp := a.rightBrother
//	a.rightBrother = b.rightBrother
//	b.rightBrother.leftBrother = a
//	b.rightBrother = tmp
//	tmp.leftBrother = b
//}
//
//
///*
// * 将node链接到root根结点
// */
//func link (root, node *FibonacciNode) {
//// 将node从双链表中移除
//removeNode(node);
//// 将node设为root的孩子
//if (root.child == null)
//root.child = node;
//else
//addNode(node, root.child);
//
//node.parent = root;
//root.degree++;
//node.marked = false;
//}
//
//type FibonacciHeap struct {
//	keyCount int            // 堆中 节点的总数
//	root     *FibonacciNode // root 即时整个堆中的最小 节点
//}
//
//// todo  插入
//func (h *FibonacciHeap) insert(key int) {
//	node := NewFibonacciNode(key)
//	if h.keyCount == 0 {
//		h.root = node
//	} else {
//		addFibonacciNode(h.root, node)
//		if node.key < h.root.key {
//			h.root = node
//		}
//	}
//	h.keyCount ++
//}
//
//// todo  堆合并
//func (h *FibonacciHeap) union(other *FibonacciHeap) {
//
//	if nil == other || nil == other.root {
//		return
//	}
//
//	if nil == h.root { // this无"最小节点"
//		h.root = other.root
//		h.keyCount = other.keyCount
//	} else { // this有"最小节点" && other有"最小节点"
//
//		// 将"other中根链表"添加到"this"中
//		catListFibonacci(h.root, other.root)
//		if h.root.key > other.root.key {
//			h.root = other.root
//			h.keyCount += other.keyCount
//		}
//	}
//}
//
//
///*
// * 删除结点node
// */
//func (h *FibonacciHeap) remove(node *FibonacciNode) {
//	key := h.root.key
//	decrease(node, m-1);
//	removeMin()
//}
//
///**
//todo 抽取最小节点
//	todo 抽取最小节点 的操作是斐波那契堆中较复杂的操作
//	(1）将要抽取最小结点的子树都直接串联在根表中
//	(2）合并所有degree相等的树，直到没有相等的degree的树
// */
//func (h *FibonacciHeap) popMin () *FibonacciNode {
//
//}