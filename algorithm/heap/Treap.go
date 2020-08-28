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


type TreapNode struct {
	key, priority int
	left, right *TreapNode
}


type Treap struct {
	root *TreapNode
}


// 当我们按照二叉查找树规则插入一个节点后, 优先级有可能不满足最大堆定义.
// 显然, 此时跟维护堆一样, 如果当前节点的优先级比根大就旋转,
// 如果当前节点是根的左儿子就右旋, 如果当前节点是根的右儿子就左旋.
func (t Treap) insert (key int) {

}