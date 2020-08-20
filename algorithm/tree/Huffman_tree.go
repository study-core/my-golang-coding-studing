package main

import (
	"fmt"
	"sort"
)

// todo 哈夫曼树
func main() {
	//arr := []int{7, 5, 2, 4}  // WPL: 35
	arr := []int{9, 3, 6, 8, 2} //  WPL: 61
	nodes := make(huffmanNodeQueue, 0)
	for i := 0; i < len(arr); i++ {
		nodes = append(nodes, &HuffmanNode{
			Value: arr[i],
			Left:  nil,
			Right: nil,
			Deep:  0,
		})
	}
	tree := BuildHuffmanTree(nodes)
	fmt.Println("Huffman WPL :=", tree.getWPL())
}


/**
todo 路径：
在一棵树中，一个结点到另一个结点之间的通路，称为路径。图 1 中，从根结点到结点 a 之间的通路就是一条路径。

todo 路径长度：
在一条路径中，每经过一个结点，路径长度都要加 1 。例如在一棵树中，规定根结点所在层数为1层，那么从根结点到第 i 层结点的路径长度为 i - 1 。图 1 中从根结点到结点 c 的路径长度为 3。

todo 结点的权：
给每一个结点赋予一个新的数值，被称为这个结点的权。例如，图 1 中结点 a 的权为 7，结点 b 的权为 5。

todo 结点的带权路径长度：
指的是从根结点到该结点之间的路径长度与该结点的权的乘积。例如，图 1 中结点 b 的带权路径长度为 2 * 5 = 10 。

todo 树的带权路径长度：
树的带权路径长度为树中所有叶子结点的带权路径长度之和。通常记作 “WPL” 。例如图 1 中所示的这颗树的带权路径长度为：
WPL = 7 * 1 + 5 * 2 + 2 * 3 + 4 * 3 = 35
 */

// 当用 n 个结点 (都做叶子结点且都有各自的权值) 试图构建一棵树时,
// 如果构建的这棵树的带权路径长度最小, 称这棵树为 "最优二叉树",
// 有时也叫“赫夫曼树”或者 "哈夫曼树".


/**
如:
					 O
				  /	    \
			    a (7)	  O
						/ 	 \
					  b (5)	   O
							/	   \
						 c (2)    d (4)



 */

// 在构建哈弗曼树时, 要使树的带权路径长度最小, 只需要遵循一个原则, 那就是:
// todo 权重越大的结点离树根越近.
// 在图中, 因为结点 【a】 的权值最大, 所以理应直接作为根结点的孩子结点.


// todo 构建 哈夫曼树
// 对于给定的有各自权值的 n 个结点，构建哈夫曼树有一个行之有效的办法：
//		1、在 n 个权值中选出两个最小的权值，对应的两个结点组成一个新的二叉树，且新二叉树的根结点的权值为左右孩子权值的和；
//		2、在原有的 n 个权值中删除那两个最小的权值，同时将新的权值加入到 n–2 个权值的行列中，以此类推； todo 且新队列需要重新排序
//		3、重复 1 和 2 ，直到所以的结点构建成了一棵二叉树为止，这棵树就是哈夫曼树。

// todo 一句话:
// 		其步骤是先构建一个包含所有节点的线性表，每次选取最小权值的两个节点，生成一个父亲节点，
//		该父亲节点的权值等于两节点权值之和，然后将该父亲节点加入到该线性表中.     新的线性表 需要重新排序

/**
如： 构建 哈夫曼树

    序列:  a (7),    b (5),   c (2),   d (4)

------------------------------------------------------

	a (7),    b (5),         O
						   /    \
						c (2)   d (4)

------------------------------------------------------

			a (7),		 O
					   /    \
					 b (5)   O
						   /    \
						c (2)   d (4)

------------------------------------------------------

						O
				     /     \
		           a (7)   	 O
						   /    \
						 b (5)   O
							   /    \
							c (2)   d (4)


 */

// todo 首先需要确定树中结点的构成。
// 		由于哈夫曼树的构建是从叶子结点开始，不断地构建新的父结点，直至树根，所以结点中应包含指向父结点的指针。
//		但是在使用哈夫曼树时是从树根开始，根据需求遍历树中的结点，因此每个结点需要有指向其左孩子和右孩子的指针。


type  HuffmanNode struct {
	Value int
	Left *HuffmanNode
	Right *HuffmanNode
	Deep int
}
func NewHuffmanExpansionNode(value int, left, right *HuffmanNode) *HuffmanNode {
	return &HuffmanNode{
		Value: value,
		Left:  left,
		Right: right,
		Deep:  0,
	}
}

type huffmanNodeQueue []*HuffmanNode

// 降序
func (self huffmanNodeQueue)Less(i, j int) bool {
	return self[i].Value > self[j].Value
}
func (self huffmanNodeQueue)Len() int {
	return len(self)
}
func  (self huffmanNodeQueue)Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self huffmanNodeQueue)sortNodes() {
	sort.Sort(self)
}


type HuffmanTree struct {
	Root *HuffmanNode
	Nodes huffmanNodeQueue
}

func BuildHuffmanTree(nodes huffmanNodeQueue) *HuffmanTree {
	if len(nodes) == 0 {
		return nil
	}
	tree := &HuffmanTree{}

	// 先来波 降序排序
	nodes.sortNodes()
	tree.Nodes = make(huffmanNodeQueue, len(nodes))
	copy(tree.Nodes, nodes)

	// 权值 从小到大 构造 哈夫曼树
	for len(nodes) != 0 {
		nodes.sortNodes()
		tail :=  nodes[len(nodes)-2:]
		nodes = nodes[:len(nodes)-2]


		littleOne := tail[0]
		secondOne := tail[1]
		parent := NewHuffmanExpansionNode(littleOne.Value + secondOne.Value, littleOne, secondOne)
		if len(nodes) == 0 {
			tree.Root = parent
			break
		} else {
			nodes = append(nodes, parent)
		}
	}
	return tree
}


func (self *HuffmanTree) getWPL() int {

	// 从 root 开始遍历, 累加 node 的带权路径长度, 求 tree 的带权路径长度
	queue := make(huffmanNodeQueue, 0)
	queue = append(queue, self.Root)

	var weight int

	for len(queue) != 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if nil != node.Left { // 为什么需要看 左节点, 请看上面的  Huffman tree 的构造示意图 自己明白
			node.Left.Deep = node.Deep + 1
			node.Right.Deep = node.Deep + 1
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			weight += node.Deep * node.Value
		}
	}
	return weight
}
