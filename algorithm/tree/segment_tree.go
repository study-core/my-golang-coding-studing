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
	b := 13
	c := (a&b) | (a^b) >>1

	fmt.Println(c)
	c = (a+b)/2
	fmt.Println(c)*/
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
	Start int   // 区间左端点
	End   int	// 区间右端点
	Data  int   // 节点的权值  (最小值 树)
	Mark  int   // 延迟更新的标记
}

type SegmentTree struct {
	Base []int
	Nodes []*SegmentTreeNode
}


func NewSegmentTree(base []int) *SegmentTree {
	//存储线段树的数组
	tree := &SegmentTree{
		Base: base,
		Nodes: make([]*SegmentTreeNode, len(base)<<1+2),
	}
	// 构建树
	tree.buildSegment(0)
	return tree
}

//构造一颗线段树，传入下标
func (self *SegmentTree) buildSegment(index int) {

	//取出该下标下的节点
	node := self.Nodes[index]

	//根节点需要手动创建
	if nil == node {
		self.Nodes[index] = &SegmentTreeNode{Start: 0, End: len(self.Base) - 1}
		node = self.Nodes[index]
	}

	//如果这个线段的左端点等于右端点则这个点是叶子节点
	if node.Start == node.End {
		node.Data = self.Base[node.Start]
	} else { //否则递归构造左右子树

		//现在这个线段的中点
		mid := (node.Start + node.End) >> 1

		//左孩子线段
		self.Nodes[(index<<1)+1] = &SegmentTreeNode{Start: node.Start, End: mid}
		//右孩子线段
		self.Nodes[(index<<1)+2] = &SegmentTreeNode{Start: mid + 1, End: node.End}

		//构造左孩子
		self.buildSegment((index<<1)+1)

		//构造右孩子
		self.buildSegment((index<<1)+2)

		minFn := func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}
		//这个节点的值等于左右孩子中较小的那个
		node.Data = minFn(self.Nodes[(index<<1)+1].Data, self.Nodes[(index<<1)+2].Data)
	}
}


/**

todo 区间查询的核心思想:
	就是找到  `交集运算`  可以构成待查询区间的所有的子区间，并且使找到的子区间的大小尽量大。
	简单的说就是，找到一些区间，使其连接起来之后正好可以涵盖整个待查询的区间。
	我们只需要找到代表这些区间的节点的最小值即可。

通过二分的思想，把查询的复杂度降到O(logn)，我们在寻找这些子区间的时候，对于当前搜索到的子区间来说，

todo 有四种情况： (最后两种 可以看成是一种)
	对于第一种情况：【当前区间和被查询区间无交集】
				   当前区间肯定不是待查询区间的子集，所以这时候应该返回一个极大值。表示取不到这个区间。
	对于第二种情况：【待查询区间包含当前区间】
					待查询区间包含当前区间，这时候当前区间肯定是待查询区间的子集，所以应当返回当前区间的权值。
	对于第三种情况：【当前区间包含待查询区间】和【当前区间和待查询区间有交集但互不包含】
					当前区间包含待查询区间和当前区间和待查询区间有交集但互不包含，这时候当前区间的一部分是带查询区间的子集，
					另一部分不是，所以应当递归的去查询当前节点的左右子树，返回左右子树中较小的那个。

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
 */
func (self *SegmentTree) queryMin(start, end int) int {
	return self.query(0, start, end)
}
func (self *SegmentTree) query (index, start, end int) int {

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

	// todo 情况三, 返回左右子树的最小值

	// 递归查询左子树和右子树  todo (求 最小值)
	min := minFn(self.query(index<<1 + 1, start, end), self.query(index<<1 +2 , start, end))
	fmt.Println("情况三, 返回左右子树的最小权重值", min)
	return min
}



/**
 *
 * @param index 当前节点的下标
 * @param update 需要被更新的节点下标
 * @param date 更新增量
 */
func (self *SegmentTree) updateOne(index, update, date int) {
	//获取这个下标所对应的的节点
	node := self.Nodes[index]
	if node.Start == node.End  {
		if node.Start == update {
			node.Data+=date
			return
		}
	}

	// todo  二分法
	mid := (node.Start&node.End) | (node.Start^node.End) >>1

	if update<=mid {
		//待更新节点在左子树
		self.updateOne((index<<1)+1,update,date)
	} else {
		//待更新节点在右子树
		self.updateOne((index<<1)+2,update,date)
	}
	//更新当前节点的值
	node.Data =minFn(self.Nodes[(index<<1)+1].Data, self.Nodes[(index<<1)+2].Data)
}