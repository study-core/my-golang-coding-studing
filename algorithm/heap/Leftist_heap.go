package main

import (
	"fmt"
)

// 左式堆
// https://www.cnblogs.com/skywang12345/p/3638384.html  (图很详细)

// todo 左倾堆 能保证 堆头是中位数 ??
func main() {
		arr := []int {1, 5, 8, 7, 6, 4, 12, 3, 14, 2}
		//arr := []int {10, 40, 24, 30, 36, 20, 12, 16 }
		heap := &leftistHeap{}
		for _, key := range arr {
			heap.insert(key)
		}
		//heap.print()
		heap.remove(14)
		heap.print()

		/**
			1, 5, 8, 7, 6, 4, 12, 3, 14, 2

			1(1) is root
			4(1) is 1's left child
			6(0) is 4's left child
			7(0) is 6's left child
			8(0) is 7's left child
			12(0) is 4's right child
			2(0) is 1's right child
			3(1) is 2's left child
			5(0) is 3's left child
			14(0) is 3's right child
		 */
}

// todo  能快速 解决  堆合并 慢问题

/**
左倾堆(leftist tree 或 leftist heap)，又被成为左偏树、左偏堆，最左堆等。
它和二叉堆一样，都是优先队列实现方式。当优先队列中涉及到 【对两个优先队列进行合并】 的问题时，二叉堆的效率就无法令人满意了，而本文介绍的左倾堆，则可以很好地解决这类问题。



一颗左倾树，它的节点除了和二叉树的节点一样具有左右子树指针外，还有两个属性：todo    键值   和   零距离。
(01) 键值的作用是来比较节点的大小，从而对节点进行排序。
(02) todo 零距离(英文名NPL，即Null Path Length     NPL)    则是从一个节点到 一个不具备 两个儿子的节点的  最短路径.      (具备0 或者 1 个 儿子的节点) 的NPL为0，    NULL节点 的NPL为-1

左倾堆有以下几个基本性质：
	[性质1] 节点的键值小于或等于它的左右子节点的键值。 todo  堆序
	[性质2] 节点的左孩子的NPL >= 右孩子的NPL。  todo  这个性质 保证了  左倾 这个特征
	[性质3] 节点的NPL = 它的右孩子的NPL + 1。 todo （因为 完全数的 左边 比右边 多一层 ?）
 */

// 左倾堆 是趋向  【非常不平衡】

type leftstNode struct {
	key, npl            int
	parent, left, right *leftstNode
}

// 合并
func mergeLeftistNode(rootA, rootB *leftstNode) *leftstNode {

	if nil == rootA {
		return rootB
	}

	if nil == rootB {
		return rootA
	}

	// todo 互换位置,  保证 小的在左边
	//
	// 合并x和y时，将x作为合并后的树的根；
	// 这里的操作是保证:  rootA  记录 堆的 root
	if rootA.key > rootB.key {
		tmp := rootA
		rootA = rootB
		rootB = tmp

		// 互换 parent
		rootB.parent = rootA.parent
	}

	// 将x的右孩子和y合并，"合并后的树的根"是x的右孩子。
	rootA.right = mergeLeftistNode(rootA.right, rootB)
	rootA.right.parent = rootA

	// todo 等全部 递归 完之后, 我们检查 左右儿子的 NPL 大小，看看是否需要 儿子对调  【大的 在左,  左倾】

	// 如果"x的左孩子为空" 或者 "x的左孩子的npl < 右孩子的npl"
	// 则，交换x和y
	if nil == rootA.left || rootA.left.npl < rootA.right.npl {
		tmp := rootA.left
		rootA.left = rootA.right
		rootA.right = tmp
	}

	// 最后 记录 新 root 的 npl
	if nil == rootA.right || nil == rootA.left {
		rootA.npl = 0
	} else {
		// todo 永远 用最小的 +1
		if rootA.left.npl > rootA.right.npl {
			rootA.npl = rootA.right.npl + 1
		} else {
			rootA.npl = rootA.left.npl + 1
		}
	}

	return rootA
}

func swapLeftistNode(root *leftstNode) {

	// 置换前先判断 right 存在与否， todo 因为都是用  右 替换 左
	if nil == root.right {
		return
	}


	if nil == root.left || root.left.npl < root.right.npl {
		tmp := root.left
		root.left = root.right
		root.right = tmp
	}

	// 最后 记录 新 root 的 npl
	if nil == root.right || nil == root.left {
		root.npl = 0
	} else {

		// todo 永远 用最小的 +1
		if root.left.npl > root.right.npl {
			root.npl = root.right.npl + 1
		} else {
			root.npl = root.left.npl + 1
		}
	}
}

func findLeftistNode(root *leftstNode, key int) *leftstNode {
	if nil == root {
		return nil
	}

	if key == root.key {
		return root
	} else {
		if node := findLeftistNode(root.left, key); nil != node {
			return node
		}
		if node := findLeftistNode(root.right, key); nil != node {
			return node
		}
	}
	return nil
}

/*
 * 打印"左倾堆"
 *
 * key        -- 节点的键值
 * direction  --  0，表示该节点是根节点
 *               -1，表示该节点是它的父结点的左孩子
 *                1，表示该节点是它的父结点的右孩子
 */
func printLeftistNode(root *leftstNode, parentKey, direction int) {

	if nil != root {

		if 0 == direction { // heap是根节点
			fmt.Printf("%d(%d) is root\n", root.key, root.npl)
		} else {
			// heap是分支节点
			if direction == 1 {
				fmt.Printf("%d(%d) is %d's right child\n", root.key, root.npl, parentKey)
			} else {
				fmt.Printf("%d(%d) is %d's left child\n", root.key, root.npl, parentKey)
			}
		}

		printLeftistNode(root.left, root.key, -1)
		printLeftistNode(root.right, root.key, 1)
	}
}

type leftistHeap struct {
	root *leftstNode
}

/**
todo  左倾堆合并

	todo 		(01) 如果一个空左倾堆与一个非空左倾堆合并，返回非空左倾堆。
	todo 		(02) 如果两个左倾堆都非空，那么比较两个根节点，取较小堆的根节点为新的根节点。将 【较小堆的根节点的右孩子】 和 【较大堆】 进行合并。
	todo 		(03) 如果新堆的右孩子的NPL > 左孩子的NPL，则交换左右孩子。
	todo 		(04) 设置新堆的根节点的NPL = 右子堆NPL + 1


todo  左倾堆 的基本操作 就是 合并,   插入只是 合并的特殊情况
 */
func (h *leftistHeap) merge(other *leftistHeap) {
	h.root = mergeLeftistNode(h.root, other.root)
}

// todo 插入是 合并的特例
func (h *leftistHeap) insert(key int) {
	node := &leftstNode{
		key: key,
		npl: 0,
	}

	h.root = mergeLeftistNode(h.root, node)
}

func (h *leftistHeap) removeRoot() {
	if nil == h.root {
		return
	}

	left := h.root.left
	right := h.root.right
	h.root = mergeLeftistNode(left, right)
	return
}

// 删除某个节点 (我自己写的)
func (h *leftistHeap) remove(key int) {

	node := findLeftistNode(h.root, key)

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

		node = mergeLeftistNode(left, right)
		if leftFlag {
			parent.left = node
		} else {
			parent.right = node
		}

		// 从 root 开始重新调整
		for nil != parent {
			swapLeftistNode(parent)
			parent = parent.parent
		}
	}
}


func (h *leftistHeap) print () {
	if nil != h.root {
		printLeftistNode(h.root, 0, 0)
	}
}