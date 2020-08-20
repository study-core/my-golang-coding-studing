package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

// todo 哈夫曼编码
//
// 一种 `可变长编码`, 依据字符在需要编码文件中出现的概率提供对字符的唯一编码.
// 保证了可变编码的平均编码最短, 被称为最优二叉树, 有时又称为最佳编码.

// todo 特点： 哈夫曼编码就是在哈夫曼树的基础上构建的，这种编码方式最大的优点：  就是用最少的字符包含最多的信息内容
func main() {
	str := "abbccaadddrttyuuikk我我我我爱!_$$$$$%"
	fmt.Println("origin str :=", str)

	charArr := []rune(str)

	frequency := make(frequencyTable, 0)
	frequencyStr := make(map[string]int, 0)
	for i := 0; i < len(charArr); i++ {
		ch := charArr[i]
		frequencyStr[string(ch)]++
		frequency[ch]++
	}

	b, _ := json.Marshal(frequencyStr)
	fmt.Println("frequencyTable :=", string(b))

	queue := make(huffmanNodeQueue, 0)
	for k, v := range frequency {
		queue = append(queue, &HuffmanNode{
			char:    k,
			parent:   nil,
			left:     nil,
			right:    nil,
			priority: v,
			deep:     0,
		})
	}

	// 构造 Huffman Tree
	tree := BuildHuffmanTree(queue)
	fmt.Println("Tree WPL :=", tree.getWPL())

	codingTable := tree.BuildHuffmanCodingTable()
	for k, v := range codingTable {
		fmt.Println("Huffman coding, the char :=", string(k), ", code :=", v, ", 字符出现频率 :=", frequencyStr[string(k)])
	}

	source := "我爱a$_!!我ddd"

	encodeStr := huffmanEncode(codingTable, source)
	target := huffmanDecode(tree, encodeStr)
	fmt.Println("调试, 源字符串:", source, ", 编码值:", encodeStr, ", 解码回来:", target)
}

/**
todo 通过统计 `文本中相同字符的个数` 作为每个字符的权值，建立哈夫曼树

对于树中的每一个子树, 统一规定其左孩子标记为 0,  右孩子标记为 1.
这样, 用到哪个字符时, 从哈夫曼树的根结点开始, 依次写出经过结点的标记, 最终得到的就是该结点的哈夫曼编码.

todo 文本中字符出现的次数越多，在哈夫曼树中的体现就是越接近树根。编码的长度越短。

如:

						 O
					 /		 \
				   a (0)     O (1)
						   /      \
						b (0)      O (1)
       						     /      \
							   c (0)    d (1)


字符 a 用到的次数最多, 其次是字符 b. 字符 a 在哈夫曼编码是 0,  字符 b 编码为 10,  字符 c 的编码为 110, 字符 d 的编码为 111.



todo 树 的右节点大 还是左节点大， 这个看你喜欢怎么放啦，  所以 树可以是 左斜 也可以是 右斜
		严格来说， 也不能称之为 斜树, 因为 节点组成的 parent 在重新加入队列后， 队列需要再次排序的

	如:

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


##########################################################

	而 序列:   a (9),  b (3),  c (6),  d (8),  e (2)

------------------------------------------------------

			a (9),  c (6),  d (8),             O
										   /       \
										e (2)     b (3)

------------------------------------------------------

				    a (9),   d (8),          O
										/        \
								       O         c (6)
								   /       \
								 e (2)     b (3)

------------------------------------------------------

											  O
										/	       \
									  O			     O
								   /     \         /     \
								 O      c (6)    d (8)  a (9)
							  /     \
							e (2)   b (3)

	像上面 这颗就不是  斜树拉

todo 优点:

Huffman树可以根据输入的字符串中某个字符 `出现的次数` 来给某个字符设定一个权值，
然后可以根据权值的大小给一个给定的字符串编码，或者对一串编码进行解码，可以用于数据压缩或者解压缩，和对字符的编解码。

  todo 可是Huffman树的优点在哪？

  		todo 1、就在于它对出现次数大的字符（即权值大的字符）的编码比出现少的字符编码短，也就是说出现次数越多，编码越短，保证了对数据的压缩。

  		todo 2、保证编的码不会出现互相涵括，也就是不会出现二义性，
				比如a的编码是00100, b的编码是001, 而c的编码是00, 这样的话, 对于00100就可能是a, 也可能是b, c.
				todo 【而Huffman树编码方式不会出现这种问题】
不会出现二义性的原因是: todo 任何一个编码都不是另一个编码的前缀(prefix)
*/

// TODO 实现Huffman树的编解码需要三种数据类型，
// 		一个是优先级队列，用来保存树的结点， (可以 递增 或者 递减)
//		二是树，用来解码，
//		三是表，用来当作码表编码.

type frequencyTable map[rune]int // 存储 map{char -> 频率}

type HuffmanNode struct {
	char     rune         // 单个字符
	parent   *HuffmanNode // 为了方便从 字符往 root 找 编码  (由下往上 找)
	left     *HuffmanNode // root 找到字符 (上往下 找)
	right    *HuffmanNode
	priority int // 权重
	deep     int // 深度 (路径长)
}

type huffmanNodeQueue []*HuffmanNode

// 构造 扩充节点 (parent)
func NewHuffmanExpansionNode(priority int, left, right *HuffmanNode) *HuffmanNode {
	return &HuffmanNode{
		priority: priority,
		left:     left,
		right:    right,
	}
}

// 降序
func (self huffmanNodeQueue) Less(i, j int) bool {
	return self[i].priority > self[j].priority
}
func (self huffmanNodeQueue) Len() int {
	return len(self)
}
func (self huffmanNodeQueue) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self huffmanNodeQueue) sortNodes() {
	sort.Sort(self)
}

type HuffmanCodingTable map[rune]string // 存放 ASCII 256个字符对应的编码表   如: 010 -> x 之类

type HuffmanTree struct {
	root *HuffmanNode
}

func BuildHuffmanTree(nodes huffmanNodeQueue) *HuffmanTree {
	if len(nodes) == 0 {
		return nil
	}

	tree := &HuffmanTree{}
	nodes.sortNodes()

	for len(nodes) != 0 {
		nodes.sortNodes()
		tail := nodes[len(nodes)-2:]
		nodes = nodes[:len(nodes)-2]

		littleOne := tail[0]
		secondOne := tail[1]
		parent := NewHuffmanExpansionNode(littleOne.priority+secondOne.priority, littleOne, secondOne)
		if len(nodes) == 0 {
			tree.root = parent
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
	queue = append(queue, self.root)

	var weight int

	for len(queue) != 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if nil != node.left { // 为什么需要看 左节点, 请看上面的  Huffman tree 的构造示意图 自己明白
			node.left.deep = node.deep + 1
			node.right.deep = node.deep + 1
			queue = append(queue, node.left)
			queue = append(queue, node.right)
		} else {
			weight += node.deep * node.priority
		}
	}
	return weight
}

func  find (node *HuffmanNode, currIndex int, codeArr []rune) (int, string) {
	if currIndex == len(codeArr) -1 {
		return -1, string(node.char)
	}
	next := currIndex + 1
	if string(codeArr[currIndex]) == "0" && nil != node.left {  // 往左转
		return find (node.left, next, codeArr)
	}

	if string(codeArr[currIndex]) == "1" && nil != node.right {  // 往右转
		return find (node.right, next, codeArr)
	}
	return next, string(node.char)
}

// 构建 编码字典表, 这里使用 由上到下, 由root到leaf获取 编码值
func (self *HuffmanTree) BuildHuffmanCodingTable() HuffmanCodingTable {
	table := make(HuffmanCodingTable, 0)
	codingValue(self.root, "", table)
	return table
}

func codingValue(node *HuffmanNode, value string, table HuffmanCodingTable) {
	if 0 != node.char {
		table[node.char] = value
		return
	}
	left := node.left
	right := node.right

	// 左路径 用 0
	if nil != left {
		lvalue := value + "0"
		codingValue(left, lvalue, table)
	}

	// 右路径 用 1
	if nil != right {
		rvalue := value + "1"
		codingValue(right, rvalue, table)
	}
}

func huffmanEncode (table HuffmanCodingTable, str string) string {
	if "" == str {
		return ""
	}
	charArr := []rune(str)
	res := ""

	for i := 0; i < len(charArr); i++ {
		res += table[charArr[i]]
	}
	return res
}

func huffmanDecode (tree *HuffmanTree, str string) string {
	if "" == str {
		return ""
	}
	charArr := []rune(str)
	res := ""

	for len(charArr) != 0 {
		next, str := find(tree.root, 0, charArr)
		res += str
		if next != -1 { // 已经 在本次中遍历完了, 就结束 for
			charArr = charArr[next:]
		} else {
			break
		}
	}
	return res
}