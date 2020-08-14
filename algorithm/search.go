package main

import (
	"math"
	"fmt"
)

func main() {
	a1 := []int{1, 2, 5, 7, 97, 15, 25, 30, 36, 39, 51, 91, 67, 78, 92, 80, 82, 85}
	fmt.Println(BSTsearch(a1, 25))
	fmt.Println(binarySearch4(a1, 25))
	// [0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597]
	//arr := make([]int, 0)
	//for i := 0; i < len(a1); i++ {
	//	arr = append(arr, Fibo1(i))
	//}
	//fmt.Println(arr)

	//HeapSearchK(a1, 3)
}

////////////////////////////////////////////////////////////////////////////   二分查找    ////////////////////////////////////////////////////////////////////////////

/**
基本思想：折半查找
 */
func binarySearch(arr []int, k int) int {
	left, right, mid := 1, len(arr), 0
	for {
		// mid向下取整
		mid = int(math.Floor(float64((left + right) / 2)))
		if arr[mid] > k {
			// 如果当前元素大于k，那么把right指针移到mid - 1的位置
			right = mid - 1
		} else if arr[mid] < k {
			// 如果当前元素小于k，那么把left指针移到mid + 1的位置
			left = mid + 1
		} else {
			// 否则就是相等了，退出循环
			break
		}
		// 判断如果left大于right，那么这个元素是不存在的。返回-1并且退出循环
		if left > right {
			mid = -1
			break
		}
	}
	// 输入元素的下标
	return mid
}

func binarySearch2(sortedArray []int, lookingFor int) int {
	var low int = 0
	var high int = len(sortedArray) - 1
	for low <= high {
		var mid int = low + (high-low)/2
		var midValue int = sortedArray[mid]
		if midValue == lookingFor {
			return mid
		} else if midValue > lookingFor {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// todo 推荐这种写法一
func binarySearch3(arr []int, k int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		// 这种写法防止两数和导致的内存溢出
		mid := low + (high-low)>>1 // avg=(a+b)>>1   右移表示除2，左移表示乘2
		if k < arr[mid] {
			high = mid - 1
		} else if k > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// todo 推荐这种写法 二, 这个最屌
func binarySearch4(arr []int, k int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		/**
		利用位与（&）提取出两个数相同的部分，利用异或（^）拿出两个数不同的部分的和，相同的部分加上不同部分的和除2即得到两个数的平均值
		异或： 相同得零，不同得1 == 男男等零，女女得零，男女得子

		avg = (a&b)  + (a^b)>>1;
		或者
		avg = (a&b)  + (a^b)/2;
		 */
		mid := (low&high) | (low^high)>>1
		if k < arr[mid] {
			high = mid - 1
		} else if k > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

////////////////////////////////////////////////////////////////////////////   差值查找    ////////////////////////////////////////////////////////////////////////////

/**
基本思想：折半查找的进化版，自适应中间值
根据 (关键值 - 起始值) / (末位值 - 起始值) 的比例来决定中间值的下标，这样能够快速的缩小查找范围，todo 会比直接折半好很多

todo 因为这样 更接近 关键值
 */
func insertSearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1
	for low < high {
		// 计算mid值是差值算法的核心代码
		mid := low + int((high-low)*(key-arr[low])/(arr[high]-arr[low]))
		if key < arr[mid] {
			high = mid - 1
		} else if key > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

////////////////////////////////////////////////////////////////////////////   斐波那契查找    ////////////////////////////////////////////////////////////////////////////

/**
基本思想：利用黄金分割 0.168 ：1 来确定中间值；也是二分查找一种改进版
用文字来说，就是费波那契数列由0和1开始，之后的费波那契系数就是由之前的两数相加而得出。
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233…… todo 特别指出：0不是第一项，而是第零项

数列的值为: F(0)=0，F(1)=1, F(n)=F(n-1)+F(n-2)（n>=2，n∈N*） n为数组下标

	|--------------- F(K)-1 ---------------|
	low					  mid        high
	|______________________|_______________|
	|------- F(K-1)-1 -----|--- F(K-2)-1 --|

他要求开始表中记录的个数为某个斐波那契数小1，即n = F(k)-1；开始将key值（要查找的数据）与第F(k-1)位置的记录进行比较(即mid = low + F(k-1) - 1)，比较结果也分为三种：
  （1）相等，mid位置的元素即为所求；
  （2）大于，low=mid+1，k-=2。说明：low=mid+1 :说明待查找的元素在[mid+1,high]范围内，k-=2 :说明范围[mid+1,high]内的元素个数为n-(F(k-1))= Fk-1-F(k-1)=Fk-F(k-1)-1=F(k-2)-1个，所以可以递归的应用斐波那契查找。
  （3）小于，high=mid-1，k-=1。说明：low=mid+1说明待查找的元素在[low,mid-1]范围内，k-=1 说明范围[low,mid-1]内的元素个数为F(k-1)-1个，所以可以递归 的应用斐波那契查找
 */
func fibonacciSearch(arr []int, key int) int {
	// 生成斐波那契数列，因为我们要满足 len(arr) = F(k) - 1
	fibArr := make([]int, 0)
	// 因为 斐波那契数列的性质我们知道数据递增的特别快，所以我们这里随机选择 生成的数列长度 36 够用了
	// [0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765 10946 17711 28657 46368 75025 121393 196418 317811 514229 832040 1346269 2178309 3524578 5702887 9227465 14930352]
	for i := 0; i <= 36; i++ {
		fibArr = append(fibArr, fibonacci(i))
	}
	//fmt.Println(fibArr)

	// 确定待查找数组在裴波那契数列的位置
	k := 0
	n := len(arr)

	// 此处 n > fib[k]-1 也是别有深意的
	// 若n恰好是斐波那契数列上某一项，且要查找的元素正好在最后一位，此时必须将数组长度填充到数列下一项的数字
	for n > fibArr[k]-1 {
		k = k + 1
	}
	//fmt.Println(k, fibArr[k])
	// 将待查找数组填充到指定的长度
	for i := n; i < fibArr[k]; i++ {
		arr = append(arr, 0)
	}
	low, high := 0, n-1
	for low <= high {
		// 获取黄金分割位置元素下标
		mid := low + fibArr[k-1] - 1
		if key < arr[mid] {
			// 若key比这个元素小, 则key值应该在low至mid - 1之间，剩下的范围个数为F(k-1) - 1
			high = mid - 1
			k -= 1
		} else if key > arr[mid] {
			// 若key比这个元素大, 则key至应该在mid + 1至high之间，剩下的元素个数为F(k) - F(k-1) - 1 = F(k-2) - 1
			low = mid + 1
			k -= 2
		} else {
			if mid < n {
				return mid
			} else {
				return n - 1
			}
		}
	}
	return -1
}

/**
生成 斐波那契数列
 */

// todo 最屌写法
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	var fibarry = [3]int{0, 1, 0}
	for i := 2; i <= n; i++ {
		fibarry[2] = fibarry[0] + fibarry[1]
		fibarry[0] = fibarry[1]
		fibarry[1] = fibarry[2]
	}
	return fibarry[2]
}

//递归实现
func Fibo1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n > 1 {
		return Fibo1(n-1) + Fibo1(n-2)
	} else {
		return -1
	}
}

//迭代实现
func Fibo2(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	} else {
		a, b := 1, 1
		result := 0
		for i := 3; i <= n; i++ {
			result = a + b
			a, b = b, result
		}
		return result
	}
}

//利用闭包
func Fibo3(n int) int {
	if n < 0 {
		return -1
	} else {
		f := Fibonacci()
		result := 0
		for i := 0; i < n; i++ {
			result = f()
		}
		return result
	}
}
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

////////////////////////////////////////////////////////////////////////////   二叉树查找    ////////////////////////////////////////////////////////////////////////////

/**
基本思路：先把数组构造出一颗二叉树的样纸，然后查找的时候在从root往下对比
 */
func BSTsearch(arr []int, key int) int {
	// 先在内存中构造 二叉树
	tree := new(Tree)
	for i, v := range arr {
		Insert(tree, v, i)
	}
	// 开始二叉树查找目标key
	return searchKey(tree.Root, key)
}

// 节点结构
type Node struct {
	Value, Index int // 元素的值和在数组中的位置
	Left, Right  *Node
}

// 树结构
type Tree struct {
	Root *Node
}

// 把数组的的元素插入树中
func Insert(tree *Tree, value, index int) {
	if nil == tree.Root {
		tree.Root = newNode(value, index)
	} else {
		InsertNode(tree.Root, newNode(value, index))
	}
}

// 把新增的节点插入树的对应位置
func InsertNode(root, childNode *Node) {
	// 否则，先和根的值对比
	if childNode.Value <= root.Value {
		// 如果小于等于跟的值，则插入到左子树
		if nil == root.Left {
			root.Left = childNode
		} else {
			InsertNode(root.Left, childNode)
		}
	} else {
		// 否则，插入到右子树
		if nil == root.Right {
			root.Right = childNode
		} else {
			InsertNode(root.Right, childNode)
		}
	}
}

func newNode(value, index int) *Node {
	return &Node{
		Value: value,
		Index: index,
	}
}

// 在构建好的二叉树中，从root开始往下查找对应的key 返回其在数组中的位置
func searchKey(root *Node, key int) int {
	if nil == root {
		return -1
	}
	if key == root.Value {
		return root.Index
	} else if key < root.Value {
		// 往左子树查找
		return searchKey(root.Left, key)
	} else {
		// 往右子树查找
		return searchKey(root.Right, key)
	}
}

////////////////////////////////////////////////////////////////////////////   2-3 树查找    ////////////////////////////////////////////////////////////////////////////

/**
2-3树 也叫 平衡树
基本思路：
 */

////////////////////////////////////////////////////////////////////////////   红黑树查找    ////////////////////////////////////////////////////////////////////////////

/**
红黑树是2-3树的一种简单高效的实现
基本思路：红黑树的基本操作是添加、删除。在对红黑树进行添加或删除之后，都会用到旋转方法。为什么呢？道理很简单，添加或删除红黑树中的节点之后，红黑树就发生了变化，
可能不满足红黑树的5条性质，也就不再是一颗红黑树了，而是一颗普通的树。而通过旋转，可以使这颗树重新成为红黑树。简单点说，旋转的目的是让树保持红黑树的特性。
 */


////////////////////////////////////////////////////////////////////////////   B/B+树查找    ////////////////////////////////////////////////////////////////////////////

/**
B/B+树是2-3树的另一种拓展，在文件系统和数据库系统中有着广泛的应用
基本思路：
 */

////////////////////////////////////////////////////////////////////////////   分块查找    ////////////////////////////////////////////////////////////////////////////

/**
基本思路：
 */

//func blockSearch(arr []int, key int) int {
//	return nil
//}

func newBlockArr(arr, indexArr []int, blockPhrase [][]int) {
	// 先根据业务场景规定每个块的长度，比方说 10
	// 先以第一个数为基准,递归分块
	key := arr[0]

	// 左和右数组
	leftArr, rightArr := make([]int, 0), make([]int, 0)
	for _, v := range arr {
		if v <= key {
			leftArr = append(leftArr, v)
		} else {
			rightArr = append(rightArr, v)
		}
	}
	//如果达到 10个元素了就把里面最大的加入 索引数组
	if len(leftArr) == 10 {
		max := 0
		for _, v := range leftArr {
			if max < v {
				max = v
			}
		}
		for i := 1; i < len(indexArr); i++ {
			if max >= indexArr[i-1] && max < indexArr[i] {
				tmpArr := []int{max}
				tmpArr = append(indexArr[:i], tmpArr...)
				indexArr = append(tmpArr, indexArr[i:]...)
			}
		}
	}

	if len(rightArr) == 10 {
		max := 0
		for _, v := range rightArr {
			if max < v {
				max = v
			}
		}
		for i := 1; i < len(indexArr); i++ {
			if max >= indexArr[i-1] && max < indexArr[i] {
				tmpArr := []int{max}
				tmpArr = append(indexArr[:i], tmpArr...)
				indexArr = append(tmpArr, indexArr[i:]...)
			}
		}
	}
}



////////////////////////////////////////////////////////////////////////////   Hash 查找    ////////////////////////////////////////////////////////////////////////////
