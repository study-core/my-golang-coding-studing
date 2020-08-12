package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// 线段树
func main() {

	// 基础数组中有六个点
	base := []int{5, 9, 7, 4, 6, 1}
	tree := NewSegmentTree(base)
	b, _ := json.Marshal(tree)
	fmt.Println(string(b))

	// 找出 [start, end] 中最小的数值
	min := tree.queryMin(1, 4)
	fmt.Println("找出最小值为", min)

	/*a := 12
	b := 15
	c := (a&b) | (a^b) >>1

	fmt.Println(c, a, b)
	c = (a+b)/2
	fmt.Println(c, a, b)*/
}

/**
todo 线段树的应用
	①更新序列的某个值。
	②查询序列的某个区间的最小值(最大值、区间和)线段树常用于解决区间统计问题。
	 求最值，区间和等操作均可使用该数据结构，本篇以求最小值为例。
	③更新序列的某个区间内的所有值。
 */

// 线段树的 node
type SegmentTreeNode struct {
	Start int // 区间左端点
	End   int // 区间右端点
	Data  int // 节点的权值  (最小值 树)
	Mark  int // 延迟更新的标记
}

type SegmentTree struct {
	Base  []int
	Nodes []*SegmentTreeNode
}

func NewSegmentTree(base []int) *SegmentTree {
	//存储线段树的数组
	tree := &SegmentTree{
		Base:  base,
		Nodes: make([]*SegmentTreeNode, len(base)<<1+2),
	}
	// 构建树
	tree.buildSegment(0)
	return tree
}

// 构造一颗线段树，传入下标
func (self *SegmentTree) buildSegment(index int) {

	// 取出该下标下的节点
	node := self.Nodes[index]

	// todo 根节点需要手动创建
	if nil == node {
		self.Nodes[index] = &SegmentTreeNode{Start: 0, End: len(self.Base) - 1}
		node = self.Nodes[index]
	}

	// 如果这个线段的左端点等于右端点则这个点是叶子节点
	if node.Start == node.End {
		node.Data = self.Base[node.Start]
	} else { //否则递归构造左右子树

		// 现在这个线段的中点
		//mid := (node.Start &  node.End)|(node.Start ^  node.End)>>1  // 想不明白, 明明求 中位数 是这样的, 这里却不行
		mid := (node.Start +  node.End)>>1

		left := index<<1 + 1
		right := index<<1 + 2
		//fmt.Println("nodes len:", len(self.Nodes), "currIndex:", index, "left:", left, "right:", right, "mid:", mid, "start:", node.Start, "end:", node.End)
		// 左孩子线段
		self.Nodes[left] = &SegmentTreeNode{Start: node.Start, End: mid}
		// 右孩子线段
		self.Nodes[right] = &SegmentTreeNode{Start: mid + 1, End: node.End}

		//构造左孩子
		self.buildSegment(left)

		//构造右孩子
		self.buildSegment(right)

		// todo 这个节点的值等于左右孩子中较小的那个
		node.Data = minFn(self.Nodes[left].Data, self.Nodes[right].Data)
	}
}

/**

todo 区间查询的核心思想:
	就是找到  `交集运算`  可以构成待查询区间的所有的子区间，并且使找到的子区间的大小尽量大。
	简单的说就是，找到一些区间，使其连接起来之后正好可以涵盖整个待查询的区间。
	我们只需要找到代表这些区间的节点的最小值即可。

通过二分的思想，把查询的复杂度降到O(logn)，我们在寻找这些子区间的时候，对于当前搜索到的子区间来说，

todo ********************************************************************************************************************
todo ********************************************************************************************************************
todo ********************************************************************************************************************
todo ********************************************************************************************************************

todo 有四种情况： (最后两种 可以看成是一种)
	对于第一种情况：【当前区间和被查询区间无交集】
				   当前区间肯定不是待查询区间的子集，所以这时候应该返回一个极大值。表示取不到这个区间。
	对于第二种情况：【待查询区间包含当前区间】
					待查询区间包含当前区间，这时候当前区间肯定是待查询区间的子集，所以应当返回当前区间的权值。
	对于第三种情况：【当前区间包含待查询区间】和【当前区间和待查询区间有交集但互不包含】
					当前区间包含待查询区间和当前区间和待查询区间有交集但互不包含，这时候当前区间的一部分是带查询区间的子集，
					另一部分不是，所以应当递归的去查询当前节点的左右子树，返回左右子树中较小的那个。

todo ********************************************************************************************************************
todo ********************************************************************************************************************
todo ********************************************************************************************************************
todo ********************************************************************************************************************
 */

func minFn(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
	todo 时间复杂度  O(logn)

 *  查询某个区间的最小值
 * @param index 当前区间的下标 (nodes数组的下标)  一般 从0开始
 * @param start 待查询的区间的左端点
 * @param end 待查询的区间的左端点
 * @return 返回当前区间在待查询区间中的部分的最小值

todo 注意： 因为有 延迟更新，那么具体什么时候更新呢，就是在查询的时候。我们的查询代码 需要用到 pushDown
 */
func (self *SegmentTree) queryMin(start, end int) int {
	return self.query(0, start, end)
}
func (self *SegmentTree) query(index, start, end int) int {

	node := self.Nodes[index]
	b, _ := json.Marshal(node)
	fmt.Println("转向 nodes 数组的下标为:", index, " 对应的node:", string(b))

	// todo  情况一, 无交集, 返回最大数
	if node.Start > end || node.End < start {
		fmt.Println("情况一, 返回大数", math.MaxInt64)
		return math.MaxInt64
	}

	// todo  情况二, 待查询区间 包含当前区间, 返回当前区间权重值
	if node.Start >= start && node.End <= end {
		fmt.Println("情况二, 返回当前区间权重值", node.Data)
		return node.Data
	}


	self.pushDown(index) // todo 注意加了这一句！！！  在返回左右子树的最小值之前，进行扩展操作！

	// todo 情况三, 返回左右子树的最小值

	// 递归查询左子树和右子树  todo (求 最小值)
	min := minFn(self.query(index<<1+1, start, end), self.query(index<<1+2, start, end))
	fmt.Println("情况三, 返回左右子树的最小权重值", min)
	return min
}

/**
 *
 * @param currentIndex 当前节点的下标
 * @param updateIndex 需要被更新的节点下标
 * @param increaseData 更新增量
 */
func (self *SegmentTree) updateOne(currentIndex, updateIndex, increaseData int) {
	// 获取这个下标所对应的的节点
	node := self.Nodes[currentIndex]

	// todo 如果当前已经到  叶子节点啦, 且需要更新的 下标就是当前节点, 则更新
	if node.Start == node.End {
		if node.Start == updateIndex {
			node.Data += increaseData
			return
		}
	}

	// todo  二分法
	mid := (node.Start + node.End)>>1

	left := currentIndex<<1 + 1
	right := currentIndex<<1 + 2

	if updateIndex <= mid {
		// 待更新节点在左子树
		self.updateOne(left, updateIndex, increaseData)
	} else {
		// 待更新节点在右子树
		self.updateOne(right, updateIndex, increaseData)
	}
	// todo 更新当前节点的值
	node.Data = minFn(self.Nodes[left].Data, self.Nodes[right].Data)
}

/**
todo 区间修改，假设修改的值有m个，直接想到的一个办法就是执行m次单点更新，这时候的复杂度为O(mlogn)这不是我们所想看到的，
	假设所有的元素都更新，那么还不如直接重新构建整颗线段树。我们在这里用了一个伟大的思想，就是【延迟更新】

todo 延迟更新:
	延迟更新就是，更新的时候，我不进行操作，只是标记一下这个点需要更新，
	在我真正使用的时候我才去更新，这在我们进行一些数据库的业务的时候，也是很重要的一个思想。
	我们在封装节点的时候，有一个成员变量我们前面一直没有使用，那就是mark，现在就是使用这个成员变量的时候了。
	我们在进行区间修改的时候，我们把这个组成这个待修改区间的所有子区间都标记上.
	查找组成当前待修改区间的所有子区间的方法和查询方法是一样的，也是分三种情况.

     *
     * @param currentIndex 当前节点的下标
     * @param start 待更新的区间的左端点
     * @param end 待更新的区间的右端点
     * @param increaseData 增量值
     */
func (self *SegmentTree) updateByRange(currentIndex, start, end, increaseData int) {

	// 获取当前的节点
	node := self.Nodes[currentIndex]

	// todo 情况一: 无交集，则返回不处理
	if node.Start > end || node.End < start {
		return
	}

	// todo 情况二: 待查询区间 包含当前区间, 则当前区间需要被标记上  (标记累加)
	if node.Start >= start && node.End <= end {
		node.Data += increaseData
		node.Mark += increaseData
		return
	}

	// todo 【注意】 这一步哦   在更新左右子树之前进行扩展操作
	self.pushDown(currentIndex)

	left := currentIndex<<1 + 1
	right := currentIndex<<1 + 2

	// 更新左子树
	self.updateByRange(left, start, end, increaseData)
	// 更新右子树
	self.updateByRange(right, start, end, increaseData)

	// todo 情况三:
	node.Data = minFn(self.Nodes[left].Data, self.Nodes[right].Data)
}

// 需要在添加一个操作，就是在对某个节点的子节点进行标记的时候，
// 把本节点的已经被标记过的部分扩展到子节点中，
// 并把本节点的权值更新为子节点的权值的最小值。然后去除本节点的标记
//
// todo 把当前节点的标志值传给子节点   (其实就是将 自己提升, 将值转移到子节点上,  呈现的效果是, 自己升, 子节点下压)
func (self *SegmentTree) pushDown(currentIndex int) {

	// 获取该下标的节点
	node := self.Nodes[currentIndex]

	if node.Mark != 0 {

		left := currentIndex<<1 + 1
		right := currentIndex<<1 + 2

		self.Nodes[left].Mark += node.Mark  	//更新左子树的标志
		self.Nodes[right].Mark += node.Mark 	//更新右子树的标志
		self.Nodes[left].Data += node.Mark  	//左子树的值加上标志值
		self.Nodes[right].Data += node.Mark 	//右子树的值加上标志值
		node.Mark = 0                       	//清除当前节点的标志值
	}
}


// todo 新增


// todo 删除