package main


// 斜堆 todo 是 左倾堆 的自调节形式

// todo  斜堆 和 左倾堆 的关系类似  AVL (平衡树) 和 Splay (伸展树) 一样

// todo 斜堆 具备 【堆序】的 二叉树,  但是保证不了 【树序】
func main() {

}

/**


todo  斜堆(Skew heap)也叫自适应堆(self-adjusting heap)，它是左倾堆的一个变种。

和左倾堆一样，它通常也用于实现优先队列；作为一种自适应的左倾堆，它的合并操作的时间复杂度也是O(lg n)。

todo 它与左倾堆的差别是：

	(01) 斜堆的节点没有"零距离"这个属性，而左倾堆则有。
	(02) 斜堆的合并操作和左倾堆的合并操作算法不同。

	todo (妈的, 上面这两句话, 不就是不一样了么)

todo 斜堆的合并操作

	(01) 如果一个空斜堆与一个非空斜堆合并，返回非空斜堆。
	(02) 如果两个斜堆都非空，那么比较两个根节点，取较小堆的根节点为新的根节点。将"较小堆的根节点的右孩子"和"较大堆"进行合并。
	(03) 合并后，交换新堆根节点的左孩子和右孩子。


        	todo 上面的 第(03)步 是斜堆和左倾堆的合并操作差别的关键所在，如果是左倾堆，则合并后要比较左右孩子的零距离大小，若右孩子的零距离 > 左孩子的零距离，则交换左右孩子；最后，在设置根的零距离。


todo  斜堆 的右路径 在任何时候 都可以 任意长

	因此,  最坏的 运行时间为 O(N). 和 伸展树 一样  M 次连续操作时间复杂度 最坏的情况为 O(M logN), 所以 摊还(均摊)开销 O(log N)

todo 每次 合并都互换 左右儿子的原因：
		由于合并都是沿着最右路径进行的,经过合并之后, 新斜堆的最右路径长度必然增加 【一个趋势】,这会影响下一次合并的效率。
		所以合并后，通过交换左右子树,使整棵树的最右路径长度非常小。然而斜堆不记录节点的距离,在操作时,从下往上,沿着合并的路径,在每个节点处都交换左右子树。
		通过不断交换左右子树,斜堆把最右路径甩向左边了。 【所以 不一定是 右边就比 左边小,  但是 大概率是】

 */

type SkewNode struct {
	key         int
	parent, left, right *SkewNode
}

func mergeSkewNode(rootA, rootB *SkewNode) *SkewNode {

	if nil == rootA {
		return rootB
	}

	if nil == rootB {
		return rootA
	}

	// 合并x和y时，将x作为合并后的树的根；
	// 这里的操作是保证: x的key < y的key   todo 和 左倾堆 一样, 始终用最小的 root 作为 新root
	if rootA.key > rootB.key {
		tmp := rootA
		rootA = rootB
		rootB = tmp

		// 互换 parent
		rootB.parent = rootA.parent
	}

	// todo 和 左倾堆 的区别     【直接互换  左右儿子】
	//
	// 将x的右孩子和y合并，
	// 合并后直接交换x的左右孩子，而不需要像左倾堆一样考虑它们的npl。
	tmp := mergeSkewNode(rootA.right, rootB)

	rootA.right = rootA.left
	rootA.left = tmp

	rootA.left.parent = rootA

	return rootA
}



func findSkewNode (root *SkewNode, key int)  *SkewNode {
	if nil == root {
		return nil
	}

	if key == root.key {
		return root
	} else {
		if node := findSkewNode(root.left, key); nil != node {
			return node
		}
		if node := findSkewNode(root.right, key); nil != node {
			return node
		}
	}
	return nil
}

type SkewHeap struct {
	root *SkewNode
}

// 合并
func (h *SkewHeap) merge(other *SkewHeap) {
	h.root = mergeSkewNode(h.root, other.root)
}

// 插入  (是合并的 特例)
func (h *SkewHeap) insert(key int) {
	node := &SkewNode{key: key}
	h.root = mergeSkewNode(h.root, node)
}

// 删除 root
func (h *SkewHeap) removeRoot() {

	if nil == h.root {
		return
	}

	left := h.root.left
	right := h.root.right
	h.root = mergeSkewNode(left, right) // 合并左右子树
	return
}

func (h *SkewHeap) remove (key int) {

	node := findSkewNode(h.root, key)
	if nil != node {

		parent := node.parent

		// 如果 key 就是 root 节点
		if nil == parent {
			h.removeRoot()
			return
		}

		var leftFlag bool
		if parent.left.key == node.key {
			leftFlag = true
		}

		// 否则, 将 node 的 left 和 right 做合并, 然后再接回 parant ，并充 root 开始做调整
		left := node.left
		right := node.right

		node = mergeSkewNode(left, right)
		if leftFlag {
			parent.left = node
		} else {
			parent.right = node
		}

		// 从 root 开始重新调整
		for nil != parent {
			mergeSkewNode(parent)
			parent = parent.parent
		}
	}
}
