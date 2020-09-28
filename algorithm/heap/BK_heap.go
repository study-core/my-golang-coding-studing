package main

import (
	"fmt"
	"strings"
)

// https://www.cnblogs.com/skywang12345/p/3655900.html
// https://www.cnblogs.com/skywang12345/p/3656098.html

// todo  二项堆
func main() {
	testMerge()
}

/**
todo 二项堆通常被用来实现【优先队列】, 它堆是指满足以下性质的【二项树】的集合.

	(01) todo 每棵二项树都满足最小堆性质。即，父节点的关键字 <= 它的孩子的关键字。
	(02) todo 不能有两棵或以上的二项树具有相同的度数(包括度数为0)。换句话说，具有度数k的二项树有0个或1个




  																		   B3
																		   6
																	 	 / |
											B2						   /  /|
											11				        18	9  23
											/|						/|	|
									     24 48				      20 35	31
				     B0     		      |						  |
		  		     13 		         52				          42

todo 包含n个节点的二项堆，表示成若干个2的指数和(或者转换成二进制)，则每一个2个指数都对应一棵二项树。
	例如，13(二进制是1101)的2个指数和为 13 = 2^3 + 2^2 + 2^0 = 8 + 4 + 1, 因此具有13个节点的二项堆由度数为3, 2, 0的三棵二项树组成.

*/

// todo 项堆是可合并堆，它的 `合并操作` 的复杂度是  O(log n)

/**
定义二项堆中每个节点的类型：

1. parent：指向父节点

2. sibling：指向右边的兄弟节点

3. child：定义该节点的子节点

4. degree:定义该节点的度数

5. 其他应用场景中需要的数据
*/
type BinomialNode struct {
	key    int           // 关键字(键值)
	degree int           // 度数  (节点 拥有 子节点的个数)
	child  *BinomialNode // 左孩子  todo 但是有多个 child 啊, 怎么整？ 那都是 记为 child 的 兄弟
	parent *BinomialNode // 父节点
	next   *BinomialNode // 兄弟
}

// todo 【拼接操作】
// 将某个树的 root 作为 另外一棵树root 的child, 就是拼接两棵树
func link(child, root *BinomialNode) {
	child.parent = root
	child.next = root.child // 这一句很重要,  子树的 root 成为, 新树 原来 child的 兄弟
	root.child = child
	root.degree++ // 子树的 度数没变,  新树的度数因为 root 加了一个 child 节点 所以度数 +1
}

// todo 【合并根表】
// 将h1, h2中的根表合并成一个按度数递增的链表，返回合并后的根节点
func mergeHeap(head1, head2 *BinomialNode) *BinomialNode {

	if nil == head1 && nil != head2 {
		return head2
	}

	if nil == head2 && nil != head1 {
		return head1
	}

	// head3: 新根表 的的头指针
	// root: 新根表的 root链表头指针
	// tmp: 中转变量
	var head3, root, tmp *BinomialNode

	for nil != head1 && nil != head2 {

		// 比较 头结点的  度数, 度数小的 游标往后移

		// 使用 tmp 采集 head1  和 head2 中的小者, 然后指针逐个往后移,  做到 将 head1 和head2 的元素从小到大的采集
		if head1.degree < head2.degree {
			tmp = head1
			head1 = head1.next // 将 指针往后移一位
		} else {
			tmp = head2
			head2 = head2.next // 将 指针往后移一位
		}
		if nil == head3 {
			head3 = tmp
			root = tmp // 第一次赋值, 得到根部
		} else {
			head3.next = tmp
			head3 = tmp // 等价于   head3 = head3.next  将 指针往后移一位
		}

		// todo 注意: 这个是最后 善后用的, 当某个 根表为空时, 则将剩余 根表的指针处 只想tmp的 next位置, 这样 head3 就能将剩余的  根表收完
		//		在 没有 根表出现 nil时, 直接将 head1 赋值给tmp.next其实也没问题的, 后续 随着指针的移动 会自动变化掉
		if nil != head1 {
			tmp.next = head1
		} else {
			tmp.next = head2
		}
	}
	return root

}

// todo 【合并 堆】
// 合并优先级: 度数小的 > 根小的
func union(heap1, heap2 *BinomialNode) *BinomialNode {

	// todo 先 合并两个 堆的 根表
	// 将h1, h2中的根表合并成一个按度数递增的链表root
	root := mergeHeap(heap1, heap2)
	if nil == root {
		return root
	}

	// todo 开始重头戏, 逐个根据度数 和 key 的大小开始合并 根表中的各个子树

	// 三个 临时指针, 用来遍历 跟链表, 逐个合并 堆的
	var pre, curr, next *BinomialNode
	curr = root
	next = curr.next

	// 新链表中"根节点度数相同的二项树"【连接】起来，直到所有根节点度数都不相同为止。
	// 在将新 链表中"根节点度数相同的二项树"连接起来时，可以将被连接的情况概括为4种
	for nil != next {
		// 一、 即，"当前节点的度数"与"下一个节点的度数"相等时。此时，不需要执行任何操作，继续查看后面的节点
		// 二、 即，"当前节点的度数"、"下一个节点的度数"和"下下一个节点的度数"都相等时。
		//		此时，暂时不执行任何操作，还是继续查看后面的节点。实际上，
		//		这里是将"下一个节点"和"下下一个节点"等到后面再进行整合连接。
		if curr.degree != next.degree || (nil != next.next && next.degree == next.next.degree) {
			// Case 1: curr.degree != next.degree
			// Case 2: curr.degree == next.degree == next.next.degree
			pre = curr
			curr = next

			// 三、 即，"当前节点的度数"与"下一个节点的度数"相等，并且"当前节点的键值"<="下一个节点的键值"。
			//		此时，将"下一个节点(对应的二项树)"作为"当前节点(对应的二项树)的左孩子"。
		} else if curr.key <= next.key {
			// Case 3: curr.degree == next.degree != next.next.degree
			//      && curr.key    <= next.key
			curr.next = next.next
			link(next, curr) // next 变成  curr 的child, 然后 curr 继续和 next.next 比较？

			// 四、 即，"当前节点的度数"与"下一个节点的度数"相等，并且"当前节点的键值">"下一个节点的度数"。
			//		此时，将"当前节点(对应的二项树)"作为"下一个节点(对应的二项树)的左孩子"。
		} else {
			// Case 4: curr.degree == next.degree != next.next.degree
			//      && curr.key    >  next.key

			// 因为 curr 要成为 next 的child了, 所以需要 看看是否是第一个 node , 需要处理好 root
			if nil == pre {
				root = next
			} else {
				pre.next = next
			}
			link(curr, next) // curr 变成  next 的 child, 然后  next 和 next.next 继续比较
			curr = next
		}

		// 在上面的 if else 中已经处理了  curr 和 curr.next了, 现在将 next 变量的指针指向 curr.next
		next = curr.next
	}
	return root
}

func testMerge() {
	head1 := &BinomialNode{
		key:    12,
		degree: 1,
		next: &BinomialNode{
			key:    3,
			degree: 3,
			next: &BinomialNode{
				key:    14,
				degree: 4,
				next: &BinomialNode{
					key:    2,
					degree: 7,
					next: &BinomialNode{
						key:    1,
						degree: 8,
						next:   nil,
					},
				},
			},
		},
	}

	head2 := &BinomialNode{
		key:    11,
		degree: 2,
		next: &BinomialNode{
			key:    6,
			degree: 5,
			next: &BinomialNode{
				key:    10,
				degree: 6,
				next: &BinomialNode{
					key:    21,
					degree: 8,
					next: &BinomialNode{
						key:    17,
						degree: 10,
						next:   nil,
					},
				},
			},
		},
	}

	arr1 := make([]string, 0)
	arr2 := make([]string, 0)
	for h1 := head1; nil != h1; {
		str1 := "{'degree': " + fmt.Sprint(h1.degree) + ", 'key': " + fmt.Sprint(h1.key) + "}"
		arr1 = append(arr1, str1)
		h1 = h1.next
	}
	for h2 := head2; nil != h2; {
		str2 := "{'degree': " + fmt.Sprint(h2.degree) + ", 'key': " + fmt.Sprint(h2.key) + "}"
		arr2 = append(arr2, str2)
		h2 = h2.next
	}
	fmt.Println("根表1: ", strings.Join(arr1, ","))
	fmt.Println("根表2: ", strings.Join(arr2, ","))

	root := mergeHeap(head1, head2)
	arr3 := make([]string, 0)
	for h3 := root; nil != h3; {
		str3 := "{'degree': " + fmt.Sprint(h3.degree) + ", 'key': " + fmt.Sprint(h3.key) + "}"
		arr3 = append(arr3, str3)
		h3 = h3.next
	}
	fmt.Println("根表3: ", strings.Join(arr3, ","))

	heap := &BinomialHeap{root: root}
	heap.print()
}

// 递归查找  todo 现在第一颗树查找, 然后在下一颗树查找
func search(root *BinomialNode, key int) *BinomialNode {
	var parent, child *BinomialNode

	// 先从当前 heap 的第一个节点 (root 节点开始查找)
	parent = root
	for nil != parent {

		// 找到直接返回
		if parent.key == key {
			return parent
		} else {

			// 否则递归去 child 节点查找, 找到直接返回
			child = search(parent.child, key)
			if nil != child {
				return child
			}

			// 否则, 将指针移动到  root链表的下一个 `root node` 重复 for动作
			parent = parent.next
		}
	}
	return nil
}

func remove(root *BinomialNode, key int) *BinomialNode {
	if nil == root {
		return nil
	}

	node := search(root, key)

	// 如果该堆中没有对应的 key, 直接结束, 并返回原来的 root
	if nil == node {
		return root
	}

	// 将被删除的节点的数据数据上移到它所在的二项树的根节点
	parent := node.parent
	for nil != parent {
		// 交换数据
		tmp := node.key
		node.key = parent.key
		parent.key = tmp

		// 下一个父节点
		node = parent
		parent = node.parent
	}

	// 找到node的前一个根节点(prev)
	var prev, tmp *BinomialNode
	tmp = root
	for node != tmp {
		prev = tmp
		tmp = tmp.next
	}
	// 移除node节点
	if nil != prev {
		prev.next = node.next
	} else {
		root = node.next
	}
	root = union(root, reverse(node.child))
	return root
}

// 因为在移除 某棵树的 root 时, 这可原来的树的剩余节点就变成了一个 root链表 的二项式堆
// todo 参数中的 root 是这个 生成的 二项式堆的 root
// 我们可以知道这时候这个 堆中的 N颗树 的degree 是由大到小, 从左到右排序的
// 我们需要做 左右反转动作, 将degree 小的树放到最左边, degree 大的放到最右边
func reverse(root *BinomialNode) *BinomialNode {

	var next, tail *BinomialNode

	if nil == root {
		return root
	}
	root.parent = nil
	for nil != root.next {
		next = root.next
		root.next = tail
		tail = root
		root = next
		root.parent = nil
	}
	root.next = tail
	return root
}

func updateKey(node *BinomialNode, newKey int) {
	if nil == node {
		return
	}
	if newKey < node.key {
		decreaseKey(node, newKey)
	} else if newKey > node.key {
		increaseKey(node, newKey)
	}
}

// key 由大变小
func decreaseKey(node *BinomialNode, newKey int) {

	node.key = newKey

	var parent, child *BinomialNode

	parent = node.parent
	child = node

	// 保证最小的 节点一直往上提到 root todo 沿着一条线 往上提 ??
	for nil != parent && child.key < parent.key {
		// 交换parent和child的数据
		tmp := parent.key
		parent.key = child.key
		child.key = tmp

		child = parent
		parent = child.parent
	}
}

// key 由小变大
func increaseKey(node *BinomialNode, newKey int) {

	node.key = newKey

	var curr, child *BinomialNode
	curr = node
	child = curr.child

	// 一直将大的节点下沉, 将最小节点上提到 root

	// 但是 node 有多个 child 啊, 需要一一比较啊, 因为需要知道往哪个 child 的路径下沉node
	for nil != child {

		// todo 优先和自己的子节点作比较
		if curr.key > child.key {
			// 如果"当前节点" < "它的左孩子"，
			// 则在"它的孩子中(左孩子 和 左孩子的兄弟)"中，找出最小的节点；
			// 然后将"最小节点的值" 和 "当前节点的值"进行互换
			least := child // todo least是child和它的兄弟中的最小节点

			// 使用子节点和它的兄弟逐个作比较
			for nil != child.next {

				// 如果 子节点大于它的兄弟节点
				if least.key > child.next.key {
					least = child.next // 记录最小值有子节点的兄弟节点
				}

				// 否则继续比较子节点的  下一个兄弟节点
				child = child.next
			}
			// todo 交换最小节点和当前节点的值
			least.key, curr.key = curr.key, least.key

			// 交换数据之后，再对"原最小节点"进行调整，使它满足最小堆的性质：父节点 <= 子节点
			curr = least
			child = curr.child
		} else {

			// 否则, 直接看 子节点的兄弟节点
			child = child.next
		}
	}
}

/**
二项堆数据类型：

1. 根表的头节点

2. 其他应用需要的数据
*/
type BinomialHeap struct {
	root *BinomialNode // 根
	min  *BinomialNode // 业务数据, 可以放 最小值 之类的
}

// todo 【合并操作】
// 	合并操作是二项堆的重点，二项堆的【添加操作】也是基于【合并操作】来实现的

/**
todo 步骤
(01) 将两个二项堆的根链表合并成一个链表。合并后的新链表按照”节点的度数”单调递增排列。  【合并 根表】
(02) 将新链表中”根节点度数相同的二项树”连接起来，直到所有根节点度数都不相同。			【合并 堆】
*/
func (self *BinomialHeap) union(other *BinomialHeap) {
	if nil != other && nil != other.root {
		self.root = union(self.root, other.root)
	}
}

// todo 【插入操作】
//		插入操作就相当简单了。插入操作可以看作是将"要插入的节点"和当前已有的堆进行合并
func (self *BinomialHeap) insert(key int) {
	node := &BinomialNode{key: key}
	self.root = union(self.root, node)
}

// todo 【删除操作】

/**
 步骤:
	(01) 将"该节点"交换到"它所在二项树"的根节点位置。
		方法是，从"该节点"不断向上(即向树根方向)"遍历，
		不断交换父节点和子节点的数据，直到被删除的键值到达树根位置。
	(02) 将"该节点所在的二项树"从二项堆中移除；将该二项堆记为heap。 todo 树 的 root 从 堆的 root链表中 脱离出来
	(03) 将"该节点所在的二项树"进行反转。反转的意思，
		todo 就是将根的所有孩子独立出来，并将这些孩子整合成二项堆，将该二项堆记为child。
	(04) 将child和heap进行合并操作。

todo 总的思想，就是将被"删除节点"从它所在的二项树中孤立出来，然后再对二项树进行相应的处理
*/

func (self *BinomialHeap) remove(key int) {
	self.root = remove(self.root, key)
}

// todo 【更新操作】
// 		1、值变小
//		2、值变大
func (self *BinomialHeap) update(oldKey, newKey int) {
	node := search(self.root, oldKey)
	if nil != node {
		updateKey(node, newKey)
	}
}

// todo 【找堆中最小key】
func (self *BinomialHeap) minimum() *BinomialNode {
	if nil == self.root {
		return nil
	}

	// prev, curr 用来做双指针移动用,  min 用来记录实时的 最小值用
	var prev, curr, min *BinomialNode
	prev = self.root
	curr = prev.next
	min = self.root
	// 找到最小节点
	for nil != curr {
		if curr.key < min.key {
			min = curr
		}
		prev = curr
		curr = curr.next
	}
	return min
}

// todo 【移除堆中最小key】
func (self *BinomialHeap) extractMinimum() {

	if nil == self.root {
		return
	}

	var prev, curr, preMin, min *BinomialNode
	prev = self.root
	curr = prev.next
	min = self.root
	// 找到最小节点
	for nil != curr {
		if curr.key < min.key {
			min = curr
			preMin = prev
		}
		prev = curr
		curr = curr.next
	}

	// 特殊处理下
	if nil == preMin { // root的根节点就是最小根节点
		self.root = self.root.next
	} else { // root的根节点不是最小根节点
		preMin.next = min.next // 因为 移除了 min, 所以原来 min上面的 next 接到 preMin上
	}

	// 反转最小节点的左孩子，得到最小堆child.  todo 因为 在min所在的 树中, 当 min 被移除之后得到的 二项式堆的 root 就是原来树的 child
	// 这样，就使得最小节点所在二项树的孩子们都脱离出来成为一棵独立的二项树(不包括最小节点)
	child := reverse(min.child)
	// 将"删除最小节点的二项堆child"和"root"进行合并  todo 两个堆 合并
	self.root = union(self.root, child)
}

// todo 【打印二项式堆】
func (self *BinomialHeap) print() {

	if nil == self.root {
		fmt.Println("heap: {}")
		return
	}

	p := self.root
	fmt.Printf("heap degree: {")
	for nil != p {
		fmt.Printf("B%d ", p.degree)
		p = p.next
	}
	fmt.Printf("}\n")

	var i int
	p = self.root
	for nil != p {
		i++
		fmt.Printf("%d. Binomial Tree B%d: \n", i, p.degree)
		fmt.Printf("\tkey:%2d{degree:%d} is root\n", p.key, p.degree)
		p = p.next
	}
}

func (self *BinomialHeap) contains(key int) bool {
	if node := search(self.root, key); nil != node {
		return true
	}
	return false
}
