package main

import (
	"fmt"
)

func main() {

	// arr := []int {2, 4, 9, 6, 1, 8, 12, 5, 3, 7}  // root 4
	//arr := []int {2, 4, 9}  // root 4
	arr := []int {9, 6, 12, 1, 2, 4, 8, 5, 3, 7}
	key := 7

	index := RedBlackTreeSearch(arr, key)
	fmt.Println("index:", index)

}

/**
红黑树是2-3树的一种简单高效的实现
基本思路：红黑树的基本操作是添加、删除。在对红黑树进行添加或删除之后，都会用到旋转方法。为什么呢？道理很简单，添加或删除红黑树中的节点之后，红黑树就发生了变化，
可能不满足红黑树的5条性质，也就不再是一颗红黑树了，而是一颗普通的树。而通过旋转，可以使这颗树重新成为红黑树。简单点说，旋转的目的是让树保持红黑树的特性。
*/

// todo 动作:   左旋、  右旋、 颜色反转 (父和子)、  变色(自己)

/**
todo 定义：
	1、 每个节点或者是红色的，或者是黑色的
	2、 【根节点是黑色的】
	3、 每个叶子节点（红黑树中叶子节点为最后的空节点）是黑色的  todo 书上没这点   (网上说的, 所有的 叶子 指的是 nil 节点, 注意 黑节点也是可以有 两个  黑的 nil 叶子的)
	4、 如果一个节点是红色的，那么他的孩子都是黑色的
	5、 从任意一个节点到叶子节点 (网上说了 叶子 其实就是 nil 节点)  经过的黑色节点是一样的 todo 书上说是 任意一节点到 nil

todo 上面的性质 保证了:        从根到叶子的最长的可能路径不多于最短的可能路径的两倍长         结果是这个树大致上是平衡的

要知道为什么这些性质确保了这个结果?
	注意到 【性质4】 导致了路径不能有两个毗连的红色节点就足够了。最短的可能路径都是黑色节点，最长的可能路径有交替的红色和黑色节点。
			因为根据 【性质5】 所有最长的路径都有相同数目的黑色节点，这就表明了没有路径能多于任何其他路径的两倍长
 */

// 我们需要清楚红黑树的一个性质：todo  根节点必须为黑色的。一个实现红黑树的规则：新插入的节点永远为红色。

/**



推算出:  todo  红色的节点只能出现在一个黑色节点的左孩子处    (只是建立在我们在 2-3树 向 红黑树 变换的过程中)

				黑
			   /
	         红
 */
const (
	RED   = "RED"
	BLACK = "BLACK"
)

// 节点
type RBNode struct {
	Color               string
	Parent, Left, Right *RBNode
	Value, Index        int
}

type RBTree struct {
	Root *RBNode
}

/*
* 对红黑树的节点(x)进行左旋转
*
* 左旋示意图(对节点 x 进行左旋)：
*      px                              px
*     /                               /
*    x                               y
*   /  \      --(左旋)-.           / \                #
*  lx   y                          x  ry
*     /   \                       /  \
*    ly   ry                     lx  ly
*
*
 */
func (t *RBTree) leftRotateRedBlack(node *RBNode) *RBNode {
	// 先提出自己的 右子
	root := node.Right     				// 原来 root 的 原来右儿子 是新的 root
	node.Right = root.Left				// 原来 root 的 新右儿子 是 原来 右儿子的左儿子 (原来 root 的孙子)
	if nil != root.Left {
		root.Left.Parent = node			// 新的 root 的 左儿子 <现在是 新 root 的孙子了> 的 爹, 也就是原来 root 的 原来右儿子的 左儿子 的爹, 是原来的  root
	}
	root.Left = node					// 新的 root 的左儿子是 原来的 root

	root.Parent = node.Parent			// 新的 root 的爹 是 原来 root 的爹
	node.Parent = root					// 原来的 root 的爹, 是新的 root

	// 最后 处理边界 , 【由底向上】 当旋转到最后时, 其实root就是 树的root 了
	if nil == root.Parent {
		t.Root = root
	}

	return root
}

/*
 * 对红黑树的节点(y)进行右旋转
 *
 * 右旋示意图(对节点 y 进行左旋)：
 *            py                               py
 *           /                                /
 *          y                                x
 *         /  \      --(右旋)-.            /  \                     #
 *        x   ry                           lx   y
 *       / \                                   / \                   #
 *      lx  rx                                rx  ry
 *
 */
func (t *RBTree) rightRotateRedBlack(node *RBNode) *RBNode {
	// 先提出自己的 左子
	root := node.Left
	node.Left = root.Right
	if nil != root.Right {
		root.Right.Parent = node
	}
	root.Right = node

	root.Parent = node.Parent
	node.Parent = root

	// 最后 处理边界 , 【由底向上】 当旋转到最后时, 其实root就是 树的root 了
	if nil == root.Parent {
		t.Root = root
	}
	return root
}

func insertValue(tree *RBTree, val, index int) {

	//tree.insertByButton2Top(val, index)
	tree.insertByTop2Button(val, index)  // todo 这个还有问题
}

// todo 由底向上 插入      二叉树 插入, 再调整树
func (t *RBTree) insertByButton2Top(val, index int) {

	node := &RBNode{Value: val, Index: index}
	// todo 1. 设置节点的颜色为红色     (新进来的节点一律设置为红色)
	node.Color = RED

	if nil == t.Root {
		t.Root = node
		return
	}

	var path bool
	var parent, current *RBNode
	parent = t.Root.Parent
	current = t.Root

	// todo 2. 将红黑树当作一颗二叉查找树，将节点添加到二叉查找树中。
	for nil != current {

		parent = current

		if node.Value < current.Value {
			current = current.Left
			path = false
		} else {
			current = current.Right
			path = true
		}

		// 找到 要插入的点了
		if nil == current {
			current = node
			current.Parent = parent
			if path {
				parent.Right = current
			} else {
				parent.Left = current
			}
			break
		}
	}

	// todo 3. 将它重新修正为一颗二叉查找树
	t.button2TopFixedUp(current)
}

// 修正树  todo ( 由底到上 插入)
func (t *RBTree) button2TopFixedUp(node *RBNode) {

	/**
	todo 几种情况：

	一、插入的是 root, 直接变色为 黑
	二、插入的node 有 【黑父】, 啥都不做
	三、插入的node 有 【红父】
			【红叔】  parent 和 uncle 都变成 黑色,  爷爷 变成 红色, curr 指针指向 爷爷， 继续向上 递归
			【黑叔】
				node 是左节点， parent 也是左节点    parent 变黑  爷爷变红 且 parent 和 爷爷  右【单旋】      							todo 一字
				node 是左节点， parent 是右节点   node 和 parent 左【单旋】， 然后 node 变黑  爷爷变红 且 node再和 爷爷  右【单旋】   	todo Z字

				node 是右节点 【道理类似】


	todo 自底向上插入的缺点：
		需要加入 parent 指针，导致 编程复杂 和 结构变大
	 */

	var current, parent, gparent *RBNode // 父亲 和 祖父

	// todo 每次 插入后 需要先看下 父节点的颜色,
	// todo 有两种， 黑父  和  红父
	//
	// todo  【黑父】  插入立即完成
	// todo  【红父】 需要看叔叔的颜色  【黑叔】   【红叔】

	current = node

	// 若 【父节点存在，并且父节点的颜色是红色】
	for nil != current.Parent && nil != current.Parent.Parent && RED == current.Parent.Color {

		gparent = current.Parent.Parent // 爷爷
		parent = current.Parent    // 父亲

		//  若 【父节点”是“祖父节点的左孩子】
		if parent == gparent.Left {
			// Case 1条件：叔叔节点是红色
			uncle := gparent.Right
			if nil != uncle && RED == uncle.Color { // todo 【红叔】    这时候就是   current  和  父亲 和 叔叔 都是红色， 爷爷是黑色
				uncle.Color = BLACK  // todo 将 父亲 和 叔叔 都变味 黑色
				parent.Color = BLACK // todo
				gparent.Color = RED  // todo 将 爷爷 变为 红色
				current = gparent       // todo  并将 current 指针指向  爷爷, 逐步往上层 调整
				continue
			}

			// todo 【黑叔】

			// todo 重点
			// 如果 node 是 parent 的 右节点 需要 先 node 和 parent 左旋, 然后 再和 gparent 右旋
			// 如果 node 是 parent 的 左节点 需要 直接 parent 和 gparent 右旋

			// Case 2条件：叔叔是黑色，且当前节点是右孩子 todo  单旋 (Z 字旋转)
			if current == parent.Right {
				parent = t.leftRotateRedBlack(parent)
			}

			// Case 3条件：叔叔是黑色，且当前节点是左孩子。 todo 双旋 (一 字旋转)
			parent.Color = BLACK // 这里的 parent 是 原来的 node
			gparent.Color = RED  // 这个 不变的话 就出现了   parent 黑    gparent 红    uncle 黑
			current = t.rightRotateRedBlack(gparent) // todo 将 curr 指针到 三层树的 root, 继续递归遍历

		} else { //  若 【父节点”是“祖父节点的右孩子】
			// Case 1条件：叔叔节点是红色
			uncle := gparent.Left
			if nil != uncle && RED == uncle.Color {
				uncle.Color = BLACK
				parent.Color = BLACK
				gparent.Color = RED
				current = gparent
				continue
			}

			// Case 2条件：叔叔是黑色，且当前节点是左孩子  todo  单旋 (一字旋转)
			if current == parent.Left {
				parent = t.rightRotateRedBlack(parent)
			}

			// Case 3条件：叔叔是黑色，且当前节点是右孩子。 todo 双旋 (Z字旋转)
			parent.Color = BLACK
			gparent.Color = RED
			current = t.leftRotateRedBlack(gparent)
		}
	}
	// todo 记住, 无论如何, 在最后 都必须将 root 置为 黑色
	t.Root.Color = BLACK
}

// todo 由顶向下 插入        边调整 边查找 插入位置 【这个还有问题  不想了】
func (t *RBTree) insertByTop2Button(val, index int) {


	/**
	todo 使用自顶向下，改进主要是 针对 自底向上是  【红父  红叔】 的情况，这样使得 在自底向上 的  一字 和 Z字 旋转时， 新的父亲节点都是黑色，
	todo 不会遇上违反【每个红色节点的两个子节点都是黑色】，即旋转之后红黑树就完成了插入。

	做到了，
		a）无需添加父亲指针
		b）编码简单


	todo 自顶向下的实现方法：

		现在既然已经知道，要避免递归往上判断 (避免 使用 parent 指针)， 就需要避免情况 【红父  红叔】，就只需要一个办法：

		todo 【让 兄弟节点 永远是 黑色】

	todo 插入步骤：

	步骤1：. 从根节点往下，记录父亲节点P，祖父节点GP，祖祖父节点GGP，当前节点X。

	步骤2.：如果X有两个红色的孩子，那么就使两个孩子变成黑色，X变成红色。todo 这个过程【可能】使得X与P都是红色， 如果不满足这个条件，就跳到步骤4。

	步骤3.：如果出现了X与P都是红色，那么此时X的兄弟节点U必定是黑色，todo 因为从上往下的过程，我们已经确定了U肯定是黑色的。那么这就回到类似 【由底向上】时的 一字 或者 Z字。直接使用单旋转或者双旋转就可以解决。解决之后，
			让X 指向旋转之后的 root节点，此时X为黑色，两个孩子为红色，原本的X是指向这两个红色孩子中的其中一个的，我们在这里回退，目的是让GP,GGP随着X的下降，回到正常的值（此时不判定两个孩子是否都为红色，刚刚做的事情就是让两个孩子变成红色，根节点变成黑色，）。

	步骤4： 完成过程2(3)之后，继续往下前进 【顺着 二叉查找 key 的路线 往下走 每经过一个节点 都】重复过程2,4，直到到达key的节点，或者达到NULL，此时X为NULL。

	步骤5：如果到达了key，那么key已经存在，不能再插入，直接返回即可。如果到达NULL，那么X指向插入新的节点，并且设X为红色，并且判断此时的P是否是红色，如果是红色，那么兄弟U必然是黑色，那么再进行一次步骤3，就完成了插入。

	实现过程中需要保存GGP节点的原因是，G也会参与到旋转中，那么旋转之后，GGP需要指向新的旋转之后的根。


	*/

	if 5 == val {
		fmt.Println()
	}

	node := &RBNode{Value: val, Index: index, Color: RED}

	if nil == t.Root {
		t.Root = node
		t.Root.Color = BLACK
		return
	}

	// todo 一边查询 一边做 tree 的调整
	var current, parent, gparent, ggparent *RBNode

	current = t.Root
	parent = nil
	gparent = nil
	ggparent = nil

	for nil != current && val != current.Value { // 只有 key 没被插入 时 才会进入的 for  todo  第一个节点不会进 for 哦, 因为 root 为 nil


		/* todo 只有  左右 儿子都是 红 时， 达到了【让 兄弟节点 永远是 黑色】 的条件 */
		if nil != current.Left && nil != current.Right && current.Left.Color == RED && current.Right.Color == RED {
			// todo 变色 和 调整
			t.Top2ButtonFixedUp(ggparent, gparent, parent, current, val)
		}

		/* todo 由定向下 遍历定位到 需要插入的位置 */

		// 做个将 游标向下移动
		ggparent = gparent
		gparent = parent
		parent = current

		if val < current.Value { // 往左 下钻
			current = current.Left
		} else { // 往右下钻
			current = current.Right
		}

		// todo 找到当前 插入点了
		if nil == current {
			break
		}


	}

	// 遍历到最后, 如果 最终的 游标指针 不是停留在 nil 则,
	// 说明 key 之前已经 插入过了, 因为上面的 for
	if nil != current {
		return
	}

	// todo 下面才是插入节点哦

	// 遍历到了 叶子结点上了,  这时候需要在 当前位置插入 key 了
	//
	// 构造新节点
	current = node   // todo 注意 第一次进来时, root 是 nil 哦

	// 插入节点
	if nil != parent {
		if current.Value < parent.Value {
			parent.Left = current
		} else {
			parent.Right = current
		}
	}

	// todo 再次调整 (就是为了 放置 nil 节点为  黑色 ??)，
	//		第一次进来时因为 是放置 root 节点, 所以也需要 做颜色调整
	t.Top2ButtonFixedUp(ggparent, gparent, parent, current, val)
}

func (t *RBTree) Top2ButtonFixedUp(ggparent, gparent, parent, current *RBNode, val int) {



	// todo  这个 函数 最终导致 需要插入的点被一步步的往上 提 一点 (当然 不是 一直提到顶哦)

	// 先将 当前 节点 变为 红, 及他的两个 儿子变为 黑
	current.Color = RED
	if nil != current.Left {
		current.Left.Color = BLACK
	}
	if nil != current.Right {
		current.Right.Color = BLACK
	}


	// 如果 【红父】
	if nil != parent && parent.Color == RED {

		// 将 爷爷 变成 红色
		if nil != gparent {
			gparent.Color = RED
		}

		//
		// 下面是决定做 单旋 还是 双旋 (一字 还是 Z字)
		//

		if (val < parent.Value) != (val < gparent.Value) { // 一种很骚的写法 直接 覆盖所有的  Z字

			// todo 旋转
			/* parent = t.Rotate(gparent, node.Value) // 使用下面的形式 写明白的 */
			if val < parent.Value {
				parent = RightRotateRedBlackNode(parent) // 右旋
				if nil != gparent {
					gparent.Right = parent   // 这里需要给 gparent 的 Right 指针 使用 新的parent重新赋值
				}
			} else {
				parent = LeftRotateRedBlackNode(parent) // 左旋
				if nil != gparent {
					gparent.Left = parent
				}
			}

		}

		/* current = t.Rotate(ggparent, node.Value) // 使用下面的形式 写明白的  */

		// todo 看 【由底向上】 部分, 和 <算法与数据结构> 书上 363 页 的图 自然清楚
		//
		// 这里 设计到 做了变色和 旋转后,  的 current  游标指针的 转移
		//
		// 如果是 一字型的话, 最新的 current 指向 原来 parent node, 原parent 现在是 三层树的 新root
		//
		// 如果是 Z字型的话, 最新的 current 指向 原来的 current node, 原current 现在是 三层树的 新root
		if val < gparent.Value {
			gparent = RightRotateRedBlackNode(gparent)    // 原先 gparent 的位置 有新元素 替代
		} else {
			gparent = LeftRotateRedBlackNode(gparent)
		}

		current = gparent   // todo 将 curr 指针指向 新的 gparent

		if nil != ggparent {
			if val < ggparent.Value {
				ggparent.Left = gparent
			} else {
				ggparent.Right = gparent
			}
		}


		// 调整 树结构  重新 遍历新节点  (在这颗 小的 三层树中 的 root 我们置为 黑)
		// todo 请看 【由底向上】部分,  这个和 红黑树的 root是黑的 性质无关, 因为这不是整棵树的 root 而是我们局部旋转的 三层小树的 新root
		current.Color = BLACK
	}
	t.Root.Color = BLACK // todo 最后不管怎么调整 整棵树的  root 一定是  黑色
}

func LeftRotateRedBlackNode(node *RBNode) *RBNode {
	root := node.Right
	node.Right = root.Left
	root.Left = node

	return root
}

func RightRotateRedBlackNode(node *RBNode) *RBNode {
	root := node.Left
	node.Left = root.Right
	root.Right = node

	return root
}

//func (t *RBTree) Rotate(node *RBNode, key int) *RBNode {
//
//	if key < node.Value {
//
//		if left := node.Left; key < left.Value {
//
//			// 右 单旋
//			tmp := left.Left
//			left.Left = tmp.Right
//			tmp.Right = left
//			node.Left = tmp
//		} else {
//
//			// 左 单旋
//			tmp := left.Right
//			left.Right = tmp.Left
//			tmp.Left = left
//			node.Left = tmp
//		}
//		return node.Left
//	} else {
//		if right := node.Right; key < right.Value {
//			// 右 单旋
//			tmp := right.Left
//			right.Left = tmp.Right
//			tmp.Right = right
//			node.Right = tmp
//		} else {
//			// 左 单旋
//			tmp := right.Right
//			right.Right = tmp.Left
//			tmp.Left = right
//			node.Right = tmp
//		}
//		return node.Right
//	}
//}

/**
红黑树查找
*/
func RedBlackTreeSearch(arr []int, key int) int {
	// 先构造树
	tree := new(RBTree)
	for i, v := range arr {
		insertValue(tree, v, i)
	}

	fmt.Println("tree root:", tree.Root.Value, "color:", tree.Root.Color)

	// 开始二叉树查找目标key
	return tree.serch(key)
}

func (t *RBTree) serch(key int) int {
	return serch(t.Root, key)
}

// 就是 普通二叉树查找
func serch(node *RBNode, key int) int {
	if nil == node {
		return -1
	}
	if key < node.Value {
		return serch(node.Left, key)
	} else if key > node.Value {
		return serch(node.Right, key)
	} else {
		return node.Index
	}
}
