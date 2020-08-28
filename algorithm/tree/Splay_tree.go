package main

import "math"

// http://blog.imallen.wang/2015/11/16/2016-07-17-shen-zhan-shu-ji-javashi-xian/ todo 讲得超级明白

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


// todo 【自底向上】需要利用旋转, 而【自顶向上】则除了旋转之外, 还需要进行连接操作.

// todo 自顶向下  和  自底向上 的区别：
//		自下而上：您搜索树并在同一迭代中旋转
//      自上而下：您首先搜索，然后在另一个迭代中旋转
//

// 这些想法也适用于创建，当您使用自上而下的方式插入密钥时，就好像它是二叉树一样，然后在另一个迭代中将其移到头部。
// 【自底向上】是两次通过 (第一个作为BST遍历树, 然后向后旋转一个节点直到成为根), 而【自顶向下】是一次通过。

/**
todo 自顶向下的 展开

	################### 看了很久, 真是看不懂 有顶向下 的思想 ###################

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
	 *   (b)：key < tree.key    【key < n: n为key的前驱节点】
	 *      	b-1 "键值为key的节点"的前驱节点存在的话，将"键值为key的节点"的前驱节点旋转为 子树根节点。
	 *      	b-2 "键值为key的节点"的前驱节点不存在的话，则意味着，key比树中任何键值都小，那么此时，将最小节点旋转为根节点。
	 *   (c)：key > tree.key    【key > n: n为key的后驱节点】
	 *      	c-1 "键值为key的节点"的后继节点存在的话，将"键值为key的节点"的后继节点旋转为 子树根节点。
	 *      	c-2 "键值为key的节点"的后继节点不存在的话，则意味着，key比树中任何键值都大，那么此时，将最大节点旋转为根节点。
	 */

	if nil == root {
		return nil
	}

	// N: 代表一颗 空树的 root节点
	// l, r: 记录左右子树的root节点临时量
	// curr: 记录 移动指针的临时量
	var N, l, r, curr *SplayNode
	N = &SplayNode{}
	l, r = N, N // 刚开始的时候,  l, r 子树为空

	for {

		// 如果 key 比root 小, 那么我们知道 key 在root的左子树
		if key < root.key {

			if nil == root.left {
				break
			}

			// 如果 key 比 Tree 左子树的root 还小
			if key < root.left.key { // zig-zig
				curr = root.left /* right rotate  */
				root.left = curr.right
				curr.right = root
				root = curr
				if nil == root.left {
					break
				}
			}

			// 在将 之前 key 的parent 和 grandpa 做了 右旋 之后, 从 新的 root 处做分裂,
			// 将分裂完的 左边部分作为 继续for (因为 目标 key 留在左边部分),
			// 将分裂完的 右边部分作为 之前 r 树的 左子树连接过去
			r.left = root /* link to right content tree */
			r = root           // 将 r树 下一次 link 时作为 link子树的 root的指针位移到新的连接点 (此时的 r树没有 做旋转哦, 只是位移了下指针哦)
			root = root.left   // root 为分裂后剩下的左部分 (包含 目标key 部分), 也是下次 for 中 key 作比较的 新指针起点

		} else if key > root.key {

			if nil == root.right {
				break
			}

			if key > root.right.key {
				curr = root.right /* left rotate  */
				root.right = curr.left
				curr.left = root
				root = curr
				if nil == root.right {
					break
				}
			}

			l.right = root /* link to left content tree */
			l = root
			root = root.right
		} else {
			break
		}
	}

	// 到最后找到目标 key了, 这时候 内存中有,   `l树`   和  `当前key作为root 的 中间树` 和 `r树`
	// 需要将中间树的 的左分支 作为 l树当前需要添加子树的 指针(临时root)出的 右子树
	// 将将中间树的 的右分支 作为 r树当前需要添加子树的 指针(临时root)出的 左子树
	// 最后以目标 key作为 整颗新树的 root, 将 l树追加到 root 的左边, 将 r树 追加到root 的右边

	l.right = root.left /* assemble */
	r.left = root.right

	// 因为最开始有, l, r = N, N
	// 后来, 每次分裂时都是将分裂的 左边部分加到 l树的右子树部分,
	// 将分裂的 右部分加到 r树的左子树部分

	// 这里为什么不用 l.right 和 r.left呢?
	//todo 因为之前每次分裂合并时, r 和 l 分别代表着下一次link的指针点一直在做位移,
	//		已经不是 真正的 l树和r树的 root了, 真正的 root 是第一次指针赋值时 ` l, r = N, N` 中的 N
	//		所以, 这里我们取值 N.right 和 N.left, 即最后内存中完整的  l树为 `N.right` 和 r树为 `N.left`
	root.left = N.right
	root.right = N.left

	// 返回目标 key (新的整棵树的 root)
	return root
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

	var newRoot *SplayNode

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

	// todo 这一步就是将 splay完之后的目标节点移除掉, 使用它的 前序节点作为新的 root
	if nil != root.left {
		// 将"tree的前驱节点"旋转为根节点
		newRoot = Top2ButtonSplay(root.left, key)
		// todo 移除掉 目标节点
		newRoot.right = root.right
	} else {

		// 不存在前驱节点, 则使用它的 right 节点作为新的 root
		newRoot = root.right
	}

	// 返回新的 root
	return newRoot
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

/**
一个节点的前驱节点:  是其左子树中的最大值，若无左子树，其前驱节点在从根节点到key的路径上，比key小的最大值。
一个节点的后继节点:  是右子树的最小值，若无右子树，其后继节点在从根节点到key的路径上，比key大的最小值。

如, 前驱:
	如果key所在的节点不存在, 则key没有前驱, 返回NULL.
	如果key所在的节点左子树不为空, 则其左子树的最大值为key的前驱.
	否则, key的前驱在从根节点到key的路径上, 在这个路径上寻找到比key小的最大值, 即为key的前驱.
 */


// // todo 【有顶向下的思想】 如果目标节点在右孩子中, 则将右子树保留在M中, 其余部分与之前的L树融合;
// 		    如果目标节点在左孩子中, 则将左子树保留在M中, 其余部分与之前的R树融合.
//		    一直循环直到找到节点, 然后将目标节点的左子树与L树融合, 右子树与R树融合.
//		    最后将使 `M.left=L,M.right=R` 即可.
//
// 但是有一个不足之处是只适应于节点一定存在的场合
/*
 * todo 完整的应该为:
 *   (a)：当前伸展树 root.key == key
 *          将"键值为key的节点"旋转为根节点
 *   (b)：key < tree.key    【key > max(n): max(n)为key的前驱节点】
 *      	b-1 "键值为key的节点"的前驱节点存在的话，将"键值为key的节点"的前驱节点旋转为 子树根节点。
 *      	b-2 "键值为key的节点"的前驱节点不存在的话，则意味着，key比树中任何键值都小，那么此时，将最小节点旋转为根节点。
 *   (c)：key > tree.key    【key < min(n): min(n)为key的后驱节点】
 *      	c-1 "键值为key的节点"的后继节点存在的话，将"键值为key的节点"的后继节点旋转为 子树根节点。
 *      	c-2 "键值为key的节点"的后继节点不存在的话，则意味着，key比树中任何键值都大，那么此时，将最大节点旋转为根节点。
 */
func (self *SplayTree) Top2ButtonSplay(key int) {
	self.root = Top2ButtonSplay(self.root, key)
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
//
//    它会先在伸展树中查找键值为key的节点。若没有找到的话，则直接返回。
//    若找到的话，则将该节点旋转为根节点，然后再删除该节点，之后将它的【前驱节点】作为根节点；
//   如果它的前驱节点不存在，则根为它的右孩子。
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


func (self *SplayTree) print() {

}



// -----------------------------------------------

// 由底向上
type SplayNode2 struct {
	key int
	parent *SplayNode2 // todo 使用 自底向上 才需要存 parent 引用
	left  *SplayNode2
	right *SplayNode2
}

// 递归实现
func Button2TopSplay1(root *SplayNode2, key int) *SplayNode2 {

	if nil == root {
		return nil
	}

	if key < root.key {
		root.left = Button2TopSplay1(root.left, key)

		// 右旋
		tmp := root
		root = root.left  // 旧 root 旧左子 是新 root
		tmp.left = root.right // 旧 root 新左子 是新root 旧右子
		root.right = tmp // 新 root 新右子 是旧 root

	}else if key > root.key {
		root.right=Button2TopSplay1(root.right, key)

		// 左旋
		tmp := root
		root = root.right
		tmp.right = root.left
		root.left = tmp

	}
	return root
}

// 非递归实现 (for 代替 递归)
func Button2TopSplay2(root *SplayNode2, key int) *SplayNode2 {

	if nil == root {
		return nil
	}

	//如果只有一个节点,则无须伸展
	if nil == root.left && nil == root.right {
		return root
	}

	// 使用 【栈】
	parentStack := NewSplayNodeStack(math.MaxInt64)
	curr := root

	// 沿着 二叉树查找路径 做 DFS, 逐个将 DFS 路径上的节点加入 parent 栈
	for {
		if key < curr.key {
			if nil != curr.left {
				parentStack.Push(curr)
				curr = curr.left
			}else{
				return root
			}
		}else if key > curr.key {
			if nil != curr.right {
				parentStack.Push(curr)
				curr = curr.right
			}else{
				return root
			}
		}else{
			break
		}
	}

	rotateRight := func(root *SplayNode2) *SplayNode2 {
		// 右旋
		tmp := root
		root = root.left  // 旧 root 旧左子 是新 root
		tmp.left = root.right // 旧 root 新左子 是新root 旧右子
		root.right = tmp // 新 root 新右子 是旧 root
		return root
	}

	rotateLeft := func(root *SplayNode2) *SplayNode2 {
		// 左旋
		tmp := root
		root = root.right
		tmp.right = root.left
		root.left = tmp
		return root
	}

	// 逐个将 stack 中的parent 弹出
	var parent, grandpa *SplayNode2
	for !parentStack.IsEmpty() {

		// 弹出 parent 和 grandpa
		parent = parentStack.Pop()
		if parentStack.IsEmpty() {
			grandpa = nil
		}else{
			grandpa = parentStack.Peek()
		}

		// 如果要查找的 node是当前parent 的左节点
		if parent.left == curr {
			// 需要 右旋
			if nil != grandpa {   // todo  zig-zag
				if parent == grandpa.left {

					// 右旋
					parent = rotateRight(parent)

					// 将旋转完成的 新root (就是要查找的 node )接到grandpa的左节点上
					// 继续 for循环, 就这样一层层的将 目标node往最顶层的 root 提
					grandpa.left = parent

					//grandParentNode.left=rotateRight(parentNode);
				}else{
					parent = rotateRight(parent)
					grandpa.right=parent
					//grandParentNode.right=rotateRight(parentNode);
				}
			}else{  // 单次 zig

				// 如果 没有 grandpa, 则只需要旋转一次即可
				parent = rotateRight(parent)
			}
		}else{

			// 左旋
			if nil !=grandpa {
				if parent == grandpa.left {
					parent = rotateLeft(parent)
					grandpa.left=parent
				}else{
					parent = rotateLeft(parent)
					grandpa.right=parent
				}
			}else{
				parent = rotateLeft(parent)
			}
		}

		// 移动 curr 指针, 因为经过旋转之后的 子树  curr 被提到 子树的root位置, 也就是往上提了一层
		curr = parent
	}

	// 提到最后, curr就是整颗 新树的 root,
	// 将 root变量的指针指向 新树的root (也就是 curr)
	root = curr
	return root
}


type SplayTree2 struct {
	root *SplayNode2
}

// todo 【由底向上的思想】如果被查找的节点在左边, 则进行右旋;
// 		如果节点在右边, 则进行左旋.
//		不过需要注意的是旋转过程必须是自底向上的, 反过来则不行.
func (self *SplayTree2) Button2TopSplay(key int) {
	//self.root = Button2TopSplay1(self.root, key)  // 递归实现
	self.root = Button2TopSplay2(self.root, key)	// 非递归
}


// --------------

type splayNodeStack struct {
	capacity, size, curr int
	arr []*SplayNode2
}

func NewSplayNodeStack(capacity int) *splayNodeStack {
	return &splayNodeStack{
		capacity: capacity,
		size:     0,
		curr:     -1,
		arr: make([]*SplayNode2, 0),
	}
}
func (self *splayNodeStack) IsEmpty() bool {
	return self.size == 0
}
func (self *splayNodeStack) Size() int {
	return self.size
}
func (self *splayNodeStack) Capacity() int {
	return self.capacity
}
func (self *splayNodeStack) Push(node *SplayNode2) bool {

	if self.size == self.capacity {
		return false
	}
	self.arr = append(self.arr, node)
	self.curr = len(self.arr) - 1
	self.size ++
	return true
}
func (self *splayNodeStack) Peek() *SplayNode2 {
	return self.arr[self.curr]
}
func (self *splayNodeStack) Pop() *SplayNode2 {
	if self.size == 0 {
		return nil
	}
	node := self.arr[len(self.arr) - 1]
	self.arr = self.arr[:len(self.arr) - 1]
	self.curr = len(self.arr) - 1
	self.size --
	return node
}