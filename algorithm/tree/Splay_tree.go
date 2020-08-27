package main

// https://www.cnblogs.com/skywang12345/p/3604286.html
// https://blog.csdn.net/u012124438/article/details/78067998  todo 本文看这个

// todo  伸展树 (又叫 分裂树)
func main() {

}

/**
todo 特点：
	除了本身是棵二叉查找树；
	当某个节点被访问时，伸展树会通过旋转使该节点成为树根。
	这样做的好处是，下次要访问该节点时，能够迅速的访问到该节点。

这样就演化成了,  被访问频率高的 节点, 越靠近 树根
*/

type SplayNode struct {
	key int
	// parent *SplayNode todo 使用 自底向上 才需要存 parent 引用
	left  *SplayNode
	right *SplayNode
}

// todo Splay: 展开     rotate: 旋转

// 旋转策略有 两种:
//		自顶向下 todo (我们这里讲这种 ...) 这种实现 node中不需要存 parent 的引用
//		自底向上

// 旋转的方式有 三种:
//      直接旋转: 左旋 || 右旋 (zig)
// 		有 `Z 字旋转`: 先左旋, 再右旋 || 先右旋, 再左旋	(zig-zag)
// 		有 `一 字旋转`: 先左旋, 再左旋 || 先右旋, 再右旋   (zig-zig)

// todo 自顶向下  和  自底向上 的区别：
//		自下而上：您搜索树并在同一迭代中旋转
//      自上而下：您首先搜索，然后在另一个迭代中旋转
//

// 这些想法也适用于创建，当您使用自上而下的方式插入密钥时，就好像它是二叉树一样，然后在另一个迭代中将其移到头部。
// 【自底向上】是两次通过 (第一个作为BST遍历树, 然后向后旋转一个节点直到成为根), 而【自顶向下】是一次通过。

/**
todo 自顶向下的 展开

当我们沿着树向下搜索某个节点x时, 将搜索路径上的节点及其子树移走.
构建两棵临时的树——左树和右树. 没有被移走的节点构成的树称为中树.

(1) 当前节点x是中树的根
(2) 左树L保存小于x的节点
(3) 右树R保存大于x的节点

开始时候, x是树T的根, 左树L和右树R都为空.
*/
func Top2ButtonSplay(root *SplayNode, key int) *SplayNode {

	/*
	 * 旋转key对应的节点为根节点，并返回根节点。
	 *
	 * 注意：
	 *   (a)：伸展树中存在"键值为key的节点"。
	 *          将"键值为key的节点"旋转为根节点。
	 *   (b)：伸展树中不存在"键值为key的节点"，并且key < tree.key。
	 *      b-1 "键值为key的节点"的前驱节点存在的话，将"键值为key的节点"的前驱节点旋转为根节点。
	 *      b-2 "键值为key的节点"的前驱节点存在的话，则意味着，key比树中任何键值都小，那么此时，将最小节点旋转为根节点。
	 *   (c)：伸展树中不存在"键值为key的节点"，并且key > tree.key。
	 *      c-1 "键值为key的节点"的后继节点存在的话，将"键值为key的节点"的后继节点旋转为根节点。
	 *      c-2 "键值为key的节点"的后继节点不存在的话，则意味着，key比树中任何键值都大，那么此时，将最大节点旋转为根节点。
	 */

	if nil == root {
		return nil
	}

	// N: 代表一颗 空树的 root节点
	// l, r: 记录左右子树的 root节点临时量
	// c: 记录 中树的 root节点临时量
	var N, l, r, c *SplayNode
	N = &SplayNode{}
	l, r = N, N // 刚开始的时候,  l, r 子树为空

	for {

		if key < root.key {

			if nil == root.left {
				break
			}

			if key < root.left.key {
				c = root.left /* rotate right */
				root.left = c.right
				c.right = root
				root = c
				if nil == root.left {
					break
				}
			}
			r.left = root /* link right */
			r = root
			root = root.left

		} else if key > root.key {

			if nil == root.right {
				break
			}

			if key > root.right.key {
				c = root.right /* rotate left */
				root.right = c.left
				c.left = root
				root = c
				if nil == root.right {
					break
				}
			}

			l.right = root /* link left */
			l = root
			root = root.right
		} else {
			break
		}
	}

	l.right = root.left /* assemble */
	r.left = root.right
	root.left = N.right
	root.right = N.left

	return root
}

func Button2TopSplay(root *SplayNode, key int) *SplayNode {
	return nil
}

func insert(root, newNode *SplayNode) *SplayNode {

	var x, y *SplayNode
	x = root

	// 查找z的插入位置
	for nil != x {
		y = x

		if newNode.key < x.key {
			x = x.left
		} else if newNode.key > x.key {
			x = x.right
		} else {
			return root
		}
	}

	if nil == y {
		root = newNode
	} else {

		if newNode.key < y.key {
			y.left = newNode
		} else {
			y.right = newNode
		}
	}
	return root
}

func remove(root *SplayNode, key int) *SplayNode {

	var x *SplayNode

	if nil == root {
		return nil
	}

	// 查找键值为key的节点，找不到的话直接返回。
	node := search(root, key)
	if nil == node {
		return nil
	}

	// 将key对应的节点旋转为根节点。
	root = Top2ButtonSplay(root, key)

	if nil != root.left {
		// 将"tree的前驱节点"旋转为根节点
		x = Top2ButtonSplay(root.left, key)
		// 移除tree节点
		x.right = root.right
	} else {
		x = root.right
	}
	return x
}

// 最小的值都在 left 子树
func minimum (root *SplayNode) *SplayNode {
	if nil == root {
		return nil
	}

	for nil != root.left {
		root = root.left
	}
	return root
}

// 最大值都在 right 子树
func maximum (root *SplayNode) *SplayNode {
	if nil == root {
		return nil
	}

	for nil != root.right {
		root = root.right
	}
	return root
}



// ---------------------------------------------------------------------------------------------


// 从 某个节点作为起始, 开始查找 key所在的  node
//
// 就是个 递归的 普通二叉树查找
func search(root *SplayNode, key int) *SplayNode {
	if nil == root {
		return nil
	}
	if key < root.key {
		return search(root.left, key)
	} else if key > root.key {
		return search(root.right, key)
	} else {
		return root
	}
}

type SplayTree struct {
	root *SplayNode
}

func (self *SplayTree) Top2ButtonSplay(key int) {
	self.root = Top2ButtonSplay(self.root, key)
}

func (self *SplayTree) Button2TopSplay(key int) {
	self.root = Button2TopSplay(self.root, key)
}

// todo 插入操作
//
// 先按照 二叉树插入, 再将 key 旋转到 root 位置
// 思想: 最新插入的可能 一会就会被查找

func (self *SplayTree) insert(key int) {
	self.root = insert(self.root, &SplayNode{key: key})
	self.Top2ButtonSplay(key) // or self.Button2TopSplay(key)
}

// todo 删除操作
//    ----------------------------- 这个 不对吧?
func (self *SplayTree) remove(key int) {
	self.root = remove(self.root, key)
}

// todo  找最小值
func (self *SplayTree) minimum () *SplayNode {
	if node := minimum(self.root); nil != node {
		return node
	}
	return nil
}

// todo  找最大值
func (self *SplayTree) maximum () *SplayNode {
	if node := maximum(self.root); nil != node {
		return node
	}
	return nil
}