package main

// https://www.cnblogs.com/nullzx/p/7968110.html#3857034   (超详细图解)

// todo Tarjan 算法 (塔扬 算法: 图的 `强连通分量算法`)   对比下 Kosaraju算法
//
// 主要用于  无向图 (连通图)
//
//
func main() {

}


/**
我们先来说下 什么是 无向图的 todo 割点与桥（割边）的定义

todo 割点：无向连通图中，去掉一个顶点及和它相邻的所有边，图中的连通分量数增加，则该顶点称为割点。

桥（割边）：无向联通图中，去掉一条边，图中的连通分量数增加，则这条边，称为桥或者割边。

割点与桥（割边）的关系：

1）有割点不一定有桥,有桥一定存在割点

2）桥一定是割点依附的边。
 */


/**
todo  原理： 判断一个顶点是不是割点除了从定义

假设DFS中我们从顶点U访问到了顶点V (此时顶点V还未被访问过), todo  U -> V    <父> -> <子>
那么我们称顶点U为顶点V的父顶点, V为U的孩子顶点.
在顶点U之前被访问过的顶点, 我们就称之为U的祖先顶点.


todo 普通节点：

todo 显然如果顶点U的所有孩子顶点可以不通过父顶点U而访问到U的祖先顶点,
	那么说明此时去掉顶点U不影响图的连通性, U就不是割点.

todo 相反, 如果顶点U至少存在一个孩子顶点, 必须通过父顶点U才能访问到U的祖先顶点,
	那么去掉顶点U后, 顶点U的祖先顶点和孩子顶点就不连通了, 说明U是一个割点.



todo  根节点 (DFS 起始节点, 随你自己选)

todo 根顶点是不是割点也很好判断, 如果从根顶点出发,
	一次DFS就能访问到所有的顶点,那么根顶点就不是割点.
	反之, 如果回溯到根顶点后, 还有未访问过的顶点,
	需要在邻接顶点上再次进行DFS, 根顶点就是割点.
 */

// 在具体实现Tarjan算法上，我们需要在DFS（深度优先遍历）中，todo 额外定义三个数组   dfn[]、low[]、parent[]
//
// TODO  dfn[]:
//		dnf数组的下标表示顶点的编号,
//		数组中的值表示该顶点在DFS中的遍历顺序(或者说时间戳),
//		每访问到一个未访问过的顶点, 访问顺序的值(时间戳)就增加1.
//		子顶点的dfn值一定比父顶点的dfn值大
//		(但不一定恰好大1, 比如父顶点有两个及两个以上分支的情况).
//		在访问一个顶点后, 它的dfn的值就确定下来了, 不会再改变.

// TODO low[]:
//		low数组的下标表示顶点的编号, 数组中的值表示DFS中该顶点不通过父顶点能访问到的祖先顶点中最小的顺序值 (或者说时间戳).
//		每个顶点初始的low值和dfn值应该一样, 在DFS中, 我们根据情况不断更新low的值.
//		假设由顶点U访问到顶点V. 【当从顶点V  <子> 回溯到顶点U <父> 时】,
//		if low[v] 子 < low[u] 父 {
//			low[u] 父 = low[v] 子
//		}
//		如果顶点U还有它分支, 每个分支回溯时都进行上述操作,
//		那么顶点low[u]就表示了不通过顶点U的父节点所能访问到的最早祖先节点.

// TODO parent[]:
//		下标表示顶点的编号, 数组中的值表示该顶点的父顶点编号,
//		它主要用于更新low值的时候排除父顶点, 当然也可以其它的办法实现相同的功能.


// TODO 找 割点：
//		只要有某个 儿子  low[v] <子> >= dnf[u] <父>
//
// todo 说明顶点V访问顶点U的祖先顶点, 必须通过顶点U,
// 		而不存在顶点V到顶点U祖先顶点的其它路径, 所以顶点U就是一个割点.
//		对于没有孩子顶点的顶点, 显然不会是割点.

// TODO 找 割边 (桥,  V-U桥)
//		只要 某个儿子  low[v] > dnf[u] 则 V-U 就是割边




// TODO 注意： Tarjan算法从图的任意顶点进行DFS都可以得出 `割点集` 和 `割边集`



//// 割边
//type CutVerEdge struct {
//	/*用于标记已访问过的顶点*/
//	marked []bool
//
//	/*三个数组的作用不再解释*/
//	low   []int
//	dfn   []int
//	parent   []int
//
//	/*用于标记是否是割点*/
//	isCutVer []bool
//
//	/*存储割点集的容器*/
//	listV []int
//
//	/*存储割边的容器，容器中存储的是数组，每个数组只有两个元素，表示这个边依附的两个顶点*/
//	listE [][]int
//
//	ug *TarjanUndirectedGraph
//
//	/*时间戳变量*/
//	visitOrder int
//}
//
///* todo 定义无向图 */
//type TarjanUndirectedGraph struct {
//	vtxNum   int /*顶点数量*/
//	edgeNum  int /*边数量*/
//
//	/*临接表*/
//	adj  [][]*TarjanEdge
//}
//
///**
//vtxNum: 定点数
//edgeNum: 边数
// */
//func  NewTarjanUndirectedGraph (vtxNum, edgeNum int, position []map[string]int) *TarjanUndirectedGraph {
//	graph := &TarjanUndirectedGraph{}
//	// todo 构造图
//	graph.vtxNum = vtxNum
//	graph.edgeNum = edgeNum
//
//	graph.adj = make([][]*TarjanEdge, vtxNum)
//
//	for i := 0; i < vtxNum; i++ {
//		graph.adj[i] = make([]*TarjanEdge, 0)
//	}
//
//	/* todo 无向图, 同一条边, 添加两次 */
//	for i := 0; i < edgeNum; i++ {
//		from := position[i]["from"]
//		to := position[i]["to"]
//		e1 :=  NewTarjanEdge(from, to)
//		e2 := NewTarjanEdge(to, from)
//
//		E1 := graph.adj[from]
//		E1 = append(E1, e1)
//		graph.adj[from] = E1
//
//		E2 := graph.adj[to]
//		E2 = append(E2, e2)
//		graph.adj[to] = E2
//
//	}
//	return graph
//}
//
//
//
//
//
//type TarjanEdge struct {
//	/*边起始顶点*/
//	from int
//	/*边终结顶点*/
//	to int
//}
//
//func NewTarjanEdge (from, to int) *TarjanEdge {
//	return &TarjanEdge{
//		from: from,
//		to:   to,
//	}
//}
//
//func (self *TarjanEdge) From() int {
//	return self.from
//}
//func (self *TarjanEdge) To() int {
//	return self.to
//}
//
//func (self *TarjanEdge) String() string {
//	return "[" + fmt.Sprint(self.from) + ", " + fmt.Sprint(self.to) +"]"
//}