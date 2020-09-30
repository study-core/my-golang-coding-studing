package main


// todo 配对堆
func main() {

}

// 配对堆 Pairing Heap 是一种实现简单、均摊复杂度 优越 的堆数据结构.
// 配对堆 是一种多叉树，并且可以被认为是一种 【简化的斐波那契堆】
// 配对堆 是一个支持插入，查询/删除最小值，合并，修改元素等操作的数据结构，也就是俗称的 【可并堆】

// 配对堆在 OI 界十分的冷门，但其实跑得比较快，也很好写，但不能可持久化，因为配对堆复杂度是势能分析出来的均摊复杂度。


type PairingNode struct {
	key int
	child, brother *PairingNode  // 使用 【左儿子 右兄弟】 的表示方法
	// 【原因】  由于配对堆是一棵多叉树，假设使用孩子表示法，对于每次新加入孩子，就需要动态 解开孩子节点，过程有点繁琐，所以我们使用 【左儿子 右兄弟】 来存储配对堆。

	parent *PairingNode  // 这个是为了 `descreaseNode()` 使用的
}

func NewPairingNode (key int) *PairingNode {
	return &PairingNode{
		key: key,
	}
}

func mergePairingNode (rootA, rootB *PairingNode) *PairingNode {

	if nil == rootA {
		return rootB
	}

	if nil == rootB {
		return rootA
	}

	// 以最小的 root 作为合并后的 新root
	if rootA.key > rootB.key {
		tmp := rootA
		rootA = rootB
		rootB = tmp
	}


	// 将b设为a的儿子
	rootB.brother = rootA.child
	rootA.child = rootB
	return rootA
}

func mergeBrotherPairingNode (node *PairingNode) *PairingNode {

	// 如果该树为空 或  没有兄弟（即他的父亲的儿子数小于2），就直接返回自身
	if nil == node || nil == node.brother {
		return node
	}

	first := node.brother			// 第一个兄弟
	second := first.brother			// 第二个兄弟

	// todo 为什么只要两个兄弟?  下面有递归.


	// 拆散
	node.brother = nil
	first.brother = nil

	// todo 核心 算法部分   （做到 两两 兄弟的 合并）
	return mergePairingNode(mergePairingNode(node, first), mergeBrotherPairingNode(second))
}

type PairingHeap struct {
	root *PairingNode
}

// todo 查询最小值
func (h *PairingHeap) min () *PairingNode {
	return h.root
}

// todo 合并
//
// 配对堆的合并操作极为简单，直接把根节点权值较大的那个配对堆设成另一个的儿子就好了
func (h *PairingHeap) union (other *PairingHeap) {
	h.root = mergePairingNode(h.root, other.root)
}

// todo 插入  (合并的特殊情况)
func (h *PairingHeap) insert (key int) {
	h.root = mergePairingNode(h.root, NewPairingNode(key))
}

// todo 删除最小值 (root)
//
// 我们拿掉根节点之后会发生什么，根节点原来的所有儿子构成了一片森林，所以我们要把他们合并起来。
// 一个很自然的想法是使用 `mergeBrother` 函数把儿子们一个一个并在一起，这样做的话正确性是显然的，但是会导致复杂度退化到 O(N).
// 为了保证删除操作的均摊复杂度为 O(log N), 我们需要：
// 	把儿子们 【从左往右】 两两配成一对，用 `mergeBrother` 操作把被配成同一对的两个儿子合并到一起, 再将新产生的堆 【从右往左】 暴力合并在一起。
func (h *PairingHeap) removeMin () {
	h.root = mergeBrotherPairingNode(h.root.child)
}