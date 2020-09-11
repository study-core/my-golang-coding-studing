package main

// http://blog.imallen.wang/2015/11/15/2016-07-16-treapshu-ji-javashi-xian/
// https://aimuke.github.io/algorithm/2019/06/28/algorithm-treap/#treap%E7%9A%84%E5%AE%9A%E4%B9%89 (这个 的插入|删除 且调整堆序 的图画的很清楚，及 使用场景)

// 树堆   todo 和 AVL、RB 等等都是为了解决 二叉树退化为链表的
//
// Treap = Tree + Heap  是树和堆的合体     todo 基本实现 【随机】平衡的结构
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

// Treap 的插入和删除， 在 先按照 二叉树插入(删除)，然后再按照堆序 维护堆，
// 其实最终结果 不会影响 二叉树的 树顺序 (经过旋转之后, min 或者 max 元素会到达对顶, 然后所有元素 比你大的还是在你右边， 比你小的还是在你左边, 仍然满足 二叉树性质)


/**
图解 插入

元素结构： 关键字(优先级)

     7(4)												          7(4)
        /     \											        /     \
    2(7)       8(5)										    2(7)       8(5)
   /    \         \                => 插入 3(25)   =>      /    \         \                => 满足堆序 不做后续操作
1(10)  5(23)       11(65)								1(10)  5(23)       11(65)
                  /										         /         /
              9(73)										     3(25)     9(73)



          7(4)											          7(4)                                  7(4)                         7(4)
        /     \											        /     \                                /     \                      /     \
    2(7)       8(5)										    2(7)       8(5)                        2(7)       8(5)               2(7)      8(5)
   /    \         \              => 插入  4(9)    =>	   /    \         \                       /    \         \              /   \          \
1(10)  5(23)       11(65)								1(10)  5(23)       11(65)     ==> 左旋 1(10)   5(23)      11(65)    1(10)   4(9)        11(65)
         /         /									         /         /                           /         /                 /   \        /
     3(25)     9(73)									     3(25)     9(73)                         4(9)     9(73)             3(25)  5(23)  9(73)
														        \                                    /
														        4(9)                                3(25)


继续插入 6(2)    =>  最终变成了
							          6(2)
							        /     \
							     2(7)       7(4)
							    /    \          \
							1(10)   4(9)         8(5)
							       /    \           \
							     3(25)  5(23)       11(65)
							                       /
							                   9(73)

										todo  从结果中可以看到，在 插入完 6(2)并最终调整完成后，不是【完全二叉树】，也不是【平衡二叉树】。
 */

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


func findMax(root *TreapNode) int {
	if nil != root.right {
		return findMax(root.right)
	}
	return root.key
}

func findMin(root *TreapNode) int {
	if nil != root.left {
		return findMin(root.left)
	}
	return root.key
}

// TODO 插入
//	时间复杂度：
//     由于旋转是O(1)的，最多进行h次(h是树的高度)，插入的复杂度是O(h)的，在期望情况下h=O(log n)，所以它的期望复杂度是O(log n)
//
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
TODO 删除

		todo 删除一个节点有两种方式，可以像删除二叉树节点那样删除一个节点，也可以像删除堆中节点那样删除 （主要是针对下面 第三点的哦）

 （1）若该结点为叶子 节点，则直接删除
 （2）若该结点为只包含一个叶子结点的结点，则将其叶子结点赋值给它 (即用叶子结点覆盖掉要删除的节点, 并清空叶子结点)
 （3）若该结点为其他情况下的节点，则进行相应的旋转，具体的方法就是每次找到优先级最小的儿子，
	  向与其相反的方向旋转 (左节点 右旋, 右节点 左旋)，直到该结点为上述 (1) (2) 情况之一，然后进行删除。 （严格来说 这是 按照堆形式删除）


下面再细说 按照 二叉树形式删除  和 堆形式删除

二叉树形式删除:
		找到左子树的 key 最大节点或者右子树的 key最小节点，然后copy元素的key值去覆盖当前被删除节点，但不拷贝其优先级（以免破坏堆属性)即 优先级使用当前被删除节点的优先级.

堆形式删除:
		只需要把要删除的节点旋转到叶节点上，然后直接删除就可以了。具体的方法就是每次找到优先级最小的儿子，向与其相反的方向旋转(左节点 右旋, 右节点 左旋)，直到那个节点被旋转到了叶节点，然后直接删除。

todo 时间复杂度：
 	最多进行O(h)次旋转，期望复杂度是 O(log n)
 */

// todo 按照 二叉树形式删除
func directRemove(root *TreapNode, key int) *TreapNode {
	if nil == root {
		return nil
	}

	// 二叉树查找 比较
	if key < root.key {
		root .left = directRemove(root.left, key)
	}else if key > root.key{
		root.right = directRemove(root.right, key)
	}else{

		// todo 找到了 要被删除的节点, 这里开始  二叉树形式删除的 精髓
		// 第三种情况, 节点同时具备 left 和 right
		if nil != root.left && nil != root.right {

			// 找到左子树的最大节点或者右子树的最小节点， 这里我们使用 左子树的做大值 (指的是 key 而不是 priority)

			leftMaxVal := findMax(root.left)
			root.key = leftMaxVal

			// 在将 左子树的最大key 替换完 root 的 key之后, 我们需要 删除掉这个左子树的 最大值.
			// 所以 在左子树中 使用 leftMaxVal 作为继续删除的关键字我们继续用 递归
			directRemove(root.left, leftMaxVal)
		}else{

			// 最终递归到 leaf 了, 我们将这个 leaf 删除掉
			if nil == root.left {
				root = root.right
			} else {
				root = root.left
			}
		}
	}
	return root
}


// todo 按照 堆形式删除
func rotateRemove(root *TreapNode, key int) *TreapNode {

	if nil == root {
		return nil
	}

	// 二叉树查找 比较
	if key < root.key {
		root.left = rotateRemove(root.left, key)
	}else if key > root.key {
		root.right = rotateRemove(root.right, key)
	}else{

		// todo 找到了 要被删除的节点, 这里开始  堆形式删除的 精髓
		// 第三种情况, 节点同时具备 left 和 right
		if nil != root.left && nil != root.right {
			// 若左孩子priority更大，则右旋
			if root.left.priority < root.right.priority{
				root = rotateRight(root)
			}else{
				//反之左旋
				root = rotateLeft(root)
			}

			// 因为需要将 被删除节点 旋转到 某个 leaf 上, 所以在上面的 旋转完之后, 我们继续递归
			root = rotateRemove(root, key)
		}else{

			// todo 最终到了 leaf 或者只剩一个 leaf 时
			if nil == root.left {
				root = root.right
			} else {
				root = root.left
			}
		}
	}
	return root
}

// todo 查找
//	使用 二叉树查找
//		根据Treap具有二叉搜索树的性质，可以快速查找所需节点。 时间复杂度： 期望复杂度是O(log n)
func find (root *TreapNode, key int) *TreapNode {
	if key < root.key {
		return find(root.left, key)
	} else if key > root.key {
		return find(root.right, key)
	} else {
		return root
	}
}



// ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------


type Treap struct {
	root *TreapNode
}

// todo 先按照  二叉树的性质将 key 插入到 leaf 上, 然后在按照 堆性质 做旋转


/**
todo 注意两点:
	一、二叉堆必须是完全二叉树, 而Treap并不一定是;
	二、Treap并不严格满足平衡二叉排序树（AVL树）的要求,
		即树堆中每个节点的左右子树高度之差的绝对值可能会超过 1, 只是近似满足平衡二叉排序树的性质.
 */

// todo 本例子使用最小堆 实现

// 给节点随机分配一个优先级, 先和二叉排序树（又叫二叉搜索树）的插入一样, 先把要插入的点插入到一个叶子上, 然后再维护堆的性质.
func (t *Treap) insert (key, priority int) {
	t.root = insert(t.root, key, priority)
}

func (t *Treap) remove(key int) {
	// t.root = directRemove(t.root, key)
	t.root = rotateRemove(t.root, key)
}

func (t *Treap) find (key int) *TreapNode {
	return find(t.root, key)
}

// AVL和红黑树的编程实现的难度要比Treap大得多

// todo 应用:
// 		Treap 是一种高效的动态的数据容器,据此我们可以用它处理一些数据的动态统计问题