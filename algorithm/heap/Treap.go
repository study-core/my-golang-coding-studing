package main

// http://blog.imallen.wang/2015/11/15/2016-07-16-treapshu-ji-javashi-xian/

// 树堆   todo 和 AVL、RB 等等都是为了解决 二叉树退化为链表的
//
// Treap = Tree + Heap  是树和堆的合体
func main() {

}

// todo 原理:
// 		在树中维护一个”优先级“，”优先级“采用随机数的方法生成，
//		但是”优先级“必须满足根堆的性质，当然是“大根堆”或者“小根堆”都无所谓.

/**
1)节点中的值满足二叉查找树特性;
2)节点中的优先级满足最大堆特性;
 */

// Treap因在BST中加入了堆的性质, 在以随机顺序将节点插入二叉排序树时,
// 根据随机附加的优先级以旋转的方式维持堆的性质, todo 其特点是能基本实现随机平衡的结构.

// 相对于其他的平衡二叉搜索树, Treap的特点是实现简单, 且能基本实现随机平衡的结构.
// Treap维护堆性质的方法只用到了旋转, todo 只需要两种旋转, 编程复杂度比Splay要小一些.

type TreapNode struct {
	key, priority int // Treap每个节点记录两个数据, 一个是键值, 一个是随机附加的优先级.
	left, right *TreapNode
}

func rotateLeft (root *TreapNode) *TreapNode {
	x := root.right
	root.right = x.left
	x.left = root
	return x
}

func rotateRight (root *TreapNode) *TreapNode {
	x := root.left
	root.left = x.right
	x.right = root
	return x
}

func insert (root *TreapNode, key, priority int) *TreapNode {

	var x *TreapNode

	if nil == root {
		x = &TreapNode{
			key:      key,
			priority: priority,
			left:     nil,
			right:    nil,
		}
	} else if key < root.key {
		// 往左插入
		root.left = insert(root.left, key, priority)
		// 调整 堆序
		if root.left.priority < root.priority {  // 如果是 最大堆实现的话, 取反这里即可
			x = rotateRight(root)
		}

	} else {
		root.right = insert(root.right, key, priority)
		if root.right.priority < root.priority {  // 如果是 最大堆实现的话, 取反这里即可
			x = rotateLeft(root)
		}
	}
	return x
}

/**
 （1）找到相应的结点；
 （2）若该结点为叶子结点，则直接删除；
 （3）若该结点为只包含一个叶子结点的结点，则将其叶子结点赋值给它；
 （4）若该结点为其他情况下的节点，则进行相应的旋转，具体的方法就是每次找到优先级最小的儿子，
	  向与其相反的方向旋转，直到该结点为上述情况之一，然后进行删除。
 */
func remove (root *TreapNode, key int) *TreapNode {

}

type Treap struct {
	root *TreapNode
}

// todo 先按照  二叉树的性质将 key 插入到 leaf 上, 然后在按照 堆性质 做旋转 ??


/**
todo 注意两点:
	一、二叉堆必须是完全二叉树, 而Treap并不一定是;
	二、Treap并不严格满足平衡二叉排序树（AVL树）的要求,
		即树堆中每个节点的左右子树高度之差的绝对值可能会超过 1, 只是近似满足平衡二叉排序树的性质.
 */

// todo 本例子使用最小堆 实现

// 给节点随机分配一个优先级, 先和二叉排序树（又叫二叉搜索树）的插入一样, 先把要插入的点插入到一个叶子上, 然后再维护堆的性质.
func (t Treap) insert (key, priority int) {
	t.root = insert(t.root, key, priority)
}


//



// AVL和红黑树的编程实现的难度要比Treap大得多

// todo 应用:
// 	Treap 是一种高效的动态的数据容器,据此我们可以用它处理一些数据的动态统计问题