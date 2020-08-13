package main

import "fmt"

// https://zhuanlan.zhihu.com/p/80325757
// https://www.cnblogs.com/nullzx/p/7499397.html   todo 这个的图画的最详细，一看就明白
func main() {
	pats :=[]string{"abcdef", "abhab", "bcd", "cde","cdfkcdf" }

	text := "bcabcdebcedfabcdefababkabhabk"

	ac :=NewAhoCorasickAutomation(pats)
	result := ac.find(text)

	for k, v := range result {
		fmt.Println("k :=", k, "v :=", v)
	}
}

// todo  ac 自动机  (多模式 字符串匹配算法)

// 我们经常用的字符串方法indexOf，都是判定两个字符串的包含关系，底层使用类似KMP，BM， Sunday这样的算法。
// todo 如果我们要判断一个长字符串是否包含多个短字符串呢？
// 比如在一篇文章找几个敏感词，在DNA串中找几个指定的基因对pattern进行预处理，
// 如果我们的模式串存在多个，则不适合了，我们就需要用到一种多模式匹配算法。

// AC自动机的核心算法仍然是寻找模式串内部规律，达到在每次失配时的高效跳转。这一点与单模式匹配KMP算法是一致的。
// todo 不同的是，AC算法寻找的是模式串之间的相同前缀关系。

// AC自动机在实现上要依托于Trie树（也称字典树）并借鉴了KMP模式匹配算法的核心思想。
// 实际上你可以把KMP算法看成每个节点都仅有一个孩子节点的AC自动机
//
// AC自动机的基础是Trie树。和Trie树不同的是，树中的每个结点除了有指向孩子的指针（或者说引用），
// 还有一个fail指针，它表示输入的字符与当前结点的所有孩子结点都不匹配时(注意，不是和该结点本身不匹配)，
// 自动机的状态应转移到的状态（或者说应该转移到的结点）。fail指针的功能可以类比于KMP算法中next数组的功能。

// 每个结点的fail指针表示由根结点到该结点所组成的字符序列的所有后缀　和　整个目标字符串集合（也就是整个Trie树）中的所有前缀 两者中最长公共的部分
// todo 注意上面这句话, 查看： https://www.cnblogs.com/nullzx/p/7499397.html
/**
todo 重要
由根结点到目标字符串“ijabdf”中的 ‘d’组成的字符序列“ijabd”的所有后缀
在整个目标字符串集{abd,abdk, abchijn, chnit, ijabdf, ijaij}的所有前缀中最长公共的部分就是abd，

而图中d结点（字符串“ijabdf”中的这个d）的fail正是指向了字符序列abd的最后一个字符.
*/

// todo AC自动机  其实就是  前缀Trie + KMP

/**
todo *******************************************************************************************************************
todo *******************************************************************************************************************
todo *******************************************************************************************************************
todo *******************************************************************************************************************
todo *******************************************************************************************************************
todo *******************************************************************************************************************

todo 它大概分为三个步骤：
		a、构建前缀树（生成goto表）
		b、添加失配指针（生成fail表）
		c、模式匹配（构造output表） 最后输出的 匹配结果集 个数和 []pat 一样


构建前缀树:  首先我们将所有的目标字符串插入到Trie树中，然后通过广度优先遍历为每个结点的所有孩子节点的fail指针找到正确的指向
			确定fail指针指向的问题和KMP算法中构造next数组的方式如出一辙。具体方法如下:

					1）将根结点的所有孩子结点的fail指向根结点，然后将根结点的所有孩子结点依次入列。

					2）若队列不为空：

					   2.1）出列，我们将出列的结点记为curr, failTo表示curr的fail指向的结点，即failTo = curr.fail

					   2.2) a.判断curr.child[i] == failTo.child[i]是否成立，

					           成立：curr.child[i].fail = failTo.child[i]，

					           不成立：判断 failTo == null是否成立

					                  成立： curr.child[i].fail == root

					                  不成立：执行failTo = failTo.fail，继续执行2.2）

					       b.curr.child[i]入列，再次执行再次执行步骤2)  todo A 的儿子入队，紧接着 B出队，再者 B的儿子入队，以此类推。

					   若队列为空:结束


todo 每个结点fail指向的解决顺序是按照【广度优先遍历】的顺序完成的，或者说【层序遍历】的顺序进行的，
	 也就是说我们是在解决当前结点的孩子结点fail的指向时，当前结点的fail指针一定已指向了正确的位置。

todo *******************************************************************************************************************
todo *******************************************************************************************************************
todo *******************************************************************************************************************
todo *******************************************************************************************************************
todo *******************************************************************************************************************
todo *******************************************************************************************************************
*/

/**
todo AC自动机的运行过程：

1）表示当前结点的指针指向AC自动机的根结点，即curr = root

2）从文本串中读取（下）一个字符

3）从当前结点的所有孩子结点中寻找与该字符匹配的结点，

   若成功：判断当前结点以及当前结点fail指向的结点是否表示一个字符串的结束，若是，则将文本串中索引起点记录在对应字符串保存结果集合中（索引起点= 当前索引-字符串长度+1）。curr指向该孩子结点，继续执行第2步

   若失败：执行第4步。

4）若fail == null（说明目标字符串中没有任何字符串是输入字符串的前缀，相当于重启状态机）curr = root, 执行步骤2，

   否则，将当前结点的指针指向fail结点，执行步骤3)
*/

// todo 代码实现

const ASCII_Size = 128

type acNode struct {

	/*如果该结点是一个终点，即，从根结点到此结点表示了一个目标字符串，则str != null, 且str就表示该字符串*/
	Str []rune

	/*ASCII == 128, 所以这里相当于128叉树*/
	Table [ASCII_Size]*acNode

	/*当前结点的孩子结点不能匹配文本串中的某个字符时，下一个应该查找的结点*/
	Fail *acNode
}

func (self *acNode) isWord() bool {
	return "" != string(self.Str)
}

type AhoCorasickAutomation struct {
	root  *acNode	  		//  AC自动机的根结点，根结点不存储任何字符信息
	pats  []string    		//  需要查找的  匹配串 集合
	res   map[string][]int 	//	key: pat,  value: []{pat在 txt出现的位置}

}

func NewAhoCorasickAutomation(patArr []string) *AhoCorasickAutomation {
	ac := &AhoCorasickAutomation{}
	ac.root = &acNode{}
	ac.pats = patArr
	ac.buildTrieTree()
	ac.buildAcFromTrie()
	return ac
}

/*
由目标字符串构建Trie树
*/
func (self *AhoCorasickAutomation) buildTrieTree() {

	// todo  使用  【广度遍历】
	for i := 0; i < len(self.pats); i++ {
		// 从 root 开始 构建
		curr := self.root
		targetArr := []rune(self.pats[i])

		for j := 0; j < len(targetArr); j++ {
			ch := targetArr[j]
			if nil == curr.Table[ch] {
				curr.Table[ch] = &acNode{}
			}
			curr = curr.Table[ch]
		}
		/* todo 将每个目标字符串的最后一个字符对应的结点变成终点 */
		curr.Str = []rune(self.pats[i])
	}
}

/*
由Trie树构建AC自动机，本质是一个自动机，相当于构建KMP算法的next数组
*/
func (self *AhoCorasickAutomation) buildAcFromTrie() {
	/* todo  广度优先遍历所使用的队列 */
	queue := make([]*acNode, 0)

	/*单独处理根结点的所有孩子结点*/
	for i := 0; i < len(self.root.Table); i++ {
		x := self.root.Table[i]
		if nil != x {
			/*根结点的所有孩子结点的fail都指向根结点*/
			x.Fail = self.root
			queue = append(queue, x)  /*所有根结点的孩子结点入列*/
		}
	}

	 for len(queue) != 0 {
		/* 确定出列结点的所有孩子结点的fail的指向 */
		 p := queue[0]
		 queue = queue[1:]

		for j := 0; j < len(p.Table); j++ {
			if nil != p.Table[j] {
				/*孩子结点入列*/
				queue = append(queue, p.Table[j])
				/*从p.fail开始找起*/
				 failTo := p.Fail
				for {
					/*说明找到了根结点还没有找到*/
					if nil == failTo {
						p.Table[j].Fail = self.root
						break
					}
					/*说明有公共前缀*/
					if nil != failTo.Table[j] {
						p.Table[j].Fail = failTo.Table[j]
						break
					}else{/*继续向上寻找*/
						failTo = failTo.Fail
					}
				}
			}
		}
	}
}

func (self *AhoCorasickAutomation) find(txt string) map[string][]int {

	/*创建一个表示存储结果的对象*/
	result := make(map[string][]int, )
	for i := 0; i < len(self.pats); i++ {
		result[self.pats[i]] = make([]int, 0)
	}

	// todo 从 root 开始查找
	curr := self.root
	i := 0
	txtRune := []rune(txt)
	for i < len(txtRune) {
		/*文本串中的字符*/
		ch := txtRune[i]

		/*文本串中的字符和AC自动机中的字符进行比较*/
		if nil != curr.Table[ch] {
			/*若相等，自动机进入下一状态*/
			curr = curr.Table[ch]
			if curr.isWord() {
				arr := result[string(curr.Str)]
				arr = append(arr, i - len(curr.Str) + 1)
				result[string(curr.Str)] = arr
			}

			/*这里很容易被忽视，因为一个目标串的中间某部分字符串可能正好包含另一个目标字符串，
			 * 即使当前结点不表示一个目标字符串的终点，但到当前结点为止可能恰好包含了一个字符串*/
			if nil != curr.Fail && curr.Fail.isWord() {
				arr := result[string(curr.Fail.Str)]
				arr = append(arr, i - len(curr.Fail.Str) + 1)
				result[string(curr.Fail.Str)] = arr
			}

			/*索引自增，指向下一个文本串中的字符*/
			i++
		}else{
			/*若不等，找到下一个应该比较的状态*/
			curr = curr.Fail

			/*到根结点还未找到，说明文本串中以ch作为结束的字符片段不是任何目标字符串的前缀，
			 * 状态机重置，比较下一个字符*/
			if nil == curr {
				curr = self.root
				i++
			}
		}
	}
	return result
}