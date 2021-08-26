package main

import (
	"fmt"
	//"math/rand"
)

func main() {
	items := []int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}

	//items := make([]int, 0)
	//myRand := rand.New(rand.NewSource(2^256 -1))
	//
	//for i := 0; i < 1000000; i++ {
	//	items = append(items, myRand.Int())
	//}


	//items := []int{1892000, 1935000, 1978000, 2021000, 2053250, 1838250, 1881250, 1924250, 1967250,
	//	2010250, 1580250, 1623250, 1666250, 1709250, 1795250, 1569500, 1655500, 1698500, 1741500, 1784500,
	//	1644750, 1752250, 1827500, 1988750, 2031750, 2074750}

	//items := []int{2, 3, 1}

	//heapSort3(items)
	//insertSort1(items)
	//quickSort2(items)
	quickSort4(items)
	//quickSort3(items)
	fmt.Println(items)
	//fmt.Println(items[len(items) - 7:])
	//HeapSearchK(items, 7)
}
//////////////////////////////////////////////////////////////////////////////// 冒泡排序 ////////////////////////////////////////////////////////////////////////
/**
冒泡排序
基本思想：两个数比较大小，较大的数下沉，较小的数冒起来
 */
func bubbleSort(arr []int) {
	var (
		n       = len(arr)
		swapped = true
	)

	// todo 最好的 冒泡, 这个效率最高

	// todo 由小到大 排序, 逐个将大数往后推
	for swapped {
		swapped = false
		for i := 0; i < n - 1; i++ {
			/**
			因为外层是个死循环，内层为值遍历到数组的倒数第二个，然后遍历到的每一个都是用当前的和后面一个最对比
			如果发现只要有需要把大的往后推的就把标识位改为true，但是如果在遍历完本次之后发现就没有 i > i + 1 的，那么就属于数组全部都排序好的了
			如果该次数组的遍历有大的就把他逐个往后推到末尾，然后标识位为 true 且 在进入下一轮 遍历之前把下一轮遍历的数组元素遍历个数 减1，【因为最后一个已经是排好的最大的了】
			todo 这个做法比 方法2的好，因为增加了标识位的判断避免了多余的 for 的遍历次数
			而方法2是预定好一定会做 n - 1次对数组的for，导致真正整体循环的次数是 n的阶乘，
			todo 而本方法根据实际情况(由标识位来控制)确定是否需要是n的完整阶乘    还是一半额阶乘


			todo  因为标识位的存在, 就变成了 根据 数组的实际顺序情况 决定是否继续

			todo  而 方法二, 是不管 数组是否已经提前有序, 都要继续走完 预设的 for
			*/
			if arr[i] > arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
		n -= 1
	}
}

func bubbleSort2(arr []int) {
	n := len(arr)
	for j := 0; j < n - 1; j ++ { //这个代表 需要做多少次对 数组的for
		for i := 0; i < n - j - 1; i++ { //每一次对数组的for
			if arr[i] > arr[i+1] {  // todo 对数组的各个元素逐个对比，把大的值一直推到末尾
				/**
				对于 for 中的 n - j - 1 解释
				当第一层for 执行第一次时，里面的for就是 n - 0 - 1 把当前数组的所有元素都逐一对比，把最大的推到末尾
				当第一层for执行第二次时，里面的for是 n - 1 - 1 就是把数组减掉最后一个(因为最后一个是最大值)去做逐一对比
				当第一层for执行第三次时，里面的for是 n - 2 - 1 就是把数组减掉最后两个(因为后面两个是数组汇总依次最大的两个值)去做逐一对比
				依次类推当for执行到 第 n - 1次时，里面的for是 n - (n - 2) - 1 == 1 就是剩下的最后一个最小的被排到第一位
				todo 所以是： (n - 0 - 1) * (n - 1 - 1) * (n - 2 - 1) * ... * (n - (n - 2) - 1) == n的阶乘
				 */
				arr[i+1], arr[i] = arr[i], arr[i+1]
			}
		}
	}

	/**
	下面这种是做浪费for次数的冒泡，todo 完全就是 n的n次方的运算
	 */
	//for i := 0; i < n - 1; i++ {
	//	for j := 0; j < n - 1; j++ {   // 因为 此时 数组 尾巴一部分已经在前面的 循环中排好序了, 大数推到后面,
	//								   // 而现在里面的 for如果继续把后面的元素作比较, 其实是在做 浪费的工作
	//		if arr[j+1] < arr[j] {
	//			arr[j], arr[j+1] = arr[j+1], arr[j]
	//		}
	//	}
	//}
}

//////////////////////////////////////////////////////////////////////////////// 选择排序 ////////////////////////////////////////////////////////////////////////
/**
选择排序
基本思想：在长度为N的无序数组中，第一次遍历n-1个数，找到最小的数值与第一个元素交换，
第二次遍历n-2个数，找到最小的数值与第二个元素交换。。。
第n-1次遍历，找到最小的数值与第n-1个元素交换，排序完成。


todo  从原理上看， 不会比冒泡好啊， 我博客中写错了还是比冒泡好 | 看最后的说明

todo  因为冒泡是逐个往后 顶的，在顶的过程中还会有 互换动作在这一步可能使得某些短序列变成有序了

todo  而 选择 只是逐个比较出最小的或者最大的， 在遍历过的元素没有做 互换动作，导致后面还是需要一遍又一遍的 做排序


todo  为什么说, 选择排序 会比  冒泡排序 好点？

todo  同样数据的情况下，2种算法的循环次数是一样的，但选择排序只有0到1次交换，而冒泡排序只有0到n次交换
 */


// todo 提取最小值 （用 这种  好理解）
func selectSort(arr []int) {
	var minIndex int
	for i := 0; i < len(arr) - 1; i++ { // 表示对数字做for遍历的次数
		minIndex = i // 初始化最小元素的下标为0  todo 中间变量
		/**
		在第一次对数组做for 的遍历中，从第二个元素开始逐一最小的那个对比，如果自己是最小的则把该元素的下标记住然后继续往后对比
		第二次对数组的for是从第三个开始，依次类推
		所以他也是 n的阶乘
		 */
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}


		if minIndex != i { //这里把最小的那个放到第i位，即第一次对数组的for是下标0，第二次的是下标1依次类推
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}


// todo 提取最大值
func selectSort2(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		maxIndex := 0
		/**
		第二种做法是，每一次对数组的for都需要把最大的书的下标记录下来，
		然后放置本次for的数组的末尾(本次for的数组是除去上次for之后的某尾数的元素)
		也是n的阶乘，todo 比第一种写法少了一个if判断，效率提升一些 ? 狗屁, 这个相差不大的
		 */
		for j := 1; j < length - i; j++ {
			if arr[j] > arr[maxIndex] {
				maxIndex = j
			}
		}
		arr[length - i - 1], arr[maxIndex] = arr[maxIndex], arr[length - i - 1]
	}
}


//////////////////////////////////////////////////////////////////////////////// 快速排序 ////////////////////////////////////////////////////////////////////////
/**
快排
基本思想：如果数组中只有一个元素，直接返回
如果有多个元素，先指定一个 【数轴】 (也就是一个基准数，开始第一遍的时候可以是最左边的数)
然后根据数轴把数分成 左右两部分，然后继续递归

todo  说白了就是 分片去各自调整  基准数的 位置， 各个递归小分片中各自调整好 各自基准数位置后， 大家组到一起 就是一个有序数组
 */
func quickSort1(nums []int) {
	reSort(nums, 0, len(nums)-1)
}
func reSort(arr []int, left, right int) {
	if left < right {
		pivot := arrAdjust(arr, left, right)  // todo 精髓
		reSort(arr, left, pivot-1)
		reSort(arr, pivot+1, right)
	}
}

// 返回调整后基准数的位置
func arrAdjust(arr []int, left, right int) int {

	x := arr[left] // 基准   中间变量

	for left < right {

		//从右向左找出小于 基准数的下标
		for left < right &&  x <= arr[right] {
			right-- // 当出现右边有数小于左边的数时，for停止进入下面if
		}

		// 如果 右边的下标已经 <= 左边的下标那么说明，上面的for已经找过头了，右边往左缩进缩过头了
		if left < right {
			// 用上面for找出来的下标的那个元素的值替换掉基准数所在的位置，且开始把左游标由基准数位置往右移 一位开始找左边的数来和基准数对比
			arr[left] = arr[right]
			left++
		}

		//从左向右找大于或等于x的数来填从 之前右侧找出来的那个值的位置，且把右游标往左移一位
		for left < right &&  x > arr[left]{
			left++
		}
		if left < right {
			arr[right] = arr[left]
			right--
		}
		/**
		然后重复 从新的右游标和左右表做上面的动作，直到本次左的所有数都小于基准数，右侧的数都大于基准数
		 */
	}

	arr[left] = x //退出前把基准放到最后一次左往右移时找出的数的那个位置，可以认为是 一个逻辑中间的位置
	/**
	todo
		到这里的操作步骤就是：
		先从右往左注意和基准数对比，把小于基准数的都移到左侧；
		再从左到右把大于等于基准数的都移到右边，不管是从左到右还是从右到左，的每一次for都是逐一往中心靠拢的游标
		最后把，基准数放置最后一次替换的那个位置，这样纸最终就达到了
		【以该基准数的新位置为新的左右分割点，形成左侧都是小于基准数，右侧都是大于基准数】
		最终把新的基准位置返回，并进入下次左右两个 数组的迭代上述重复动作，知道最终排完
	 */

	return left
}



// todo 第一种是做了个中间变量来实现数据互换的，这个是直接互换，写法简洁一点
func quickSort2(arr []int) {
	recursionSort(arr, 0, len(arr) - 1)
}

func recursionSort(arr []int, left int, right int) {
	if left < right { // 加一个强制限定 左边小于右边的下标，否则结束快排,这个也是结束递归的方式
		pivot := partition(arr, left, right)
		recursionSort(arr, left, pivot-1)
		recursionSort(arr, pivot+1, right)
	}
}
// todo 精华就在这个里面
func partition(arr []int, left int, right int) int {

	// 为什么要这么做 因为看里面的逻辑我们知道，left和right是逐渐往中间(不是二分中心点)靠拢的
	for left < right {

		//先逐个缩进遍历右侧的元素，由右到左，到发现有小于基准数的时候则记住当前下标 往下走
		for left < right && arr[left] <= arr[right] {
			right--
		}
		if left < right {
			// 互换基准数和 对比数的位置，注意现在 基准数是在 之前对比数的位置的，然后开始从之前基准数所在的下一个位置开始启动左边对比
			arr[left], arr[right] = arr[right], arr[left]
			left++
		}

		// 逐个缩进遍历新的左侧(从之前右侧对比的基准的下一个位置开始)，由左到右和基准数对比(注意次时基准数的下标是和之前右侧查找的数下标互换了)，直到发现有大于等于基准数的时候记住下标 往下走
		for left < right && arr[left] <= arr[right] {
			left++
		}
		if left < right {
			// 把基准数和对比数互换位置
			arr[left], arr[right] = arr[right], arr[left]
			right--
		}
	}
	/**
	直到最后 todo 基准数被调节到某个位置，且该位置左边的元素都小于基准数，右边的元素都大于等于基准数，
	 			  且返回该基准数最终的位置，把数组分为左右两部分各自进入新的递归。
	 */
	return left
}


// todo  三分单向快排
// 		之前的 都是 单轴快排
//		下面这种是一种 是 单轴的优化，  减少了当 元素相等时 也要做 互换的无脑动作
func quickSort3 (arr []int) {
	suffleThreeWay(arr, 0, len(arr) - 1)
}

func suffleThreeWay (arr []int, left, right int) {

	//	最开始:
	//
	//  | - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - |
	//
	//	left, k																			right
	//	pivot = arr[left]
	//
	//
	//  过程中 呈现的是:
	//
	//	| - - - - - - - - - | - - - - - - - - - | - - - - - - - - - - | - - - - - - - - - |
	//
	//	left                lt      			k					  gt				right
	//
	// 		   小于  pivot			等于 pivot			还没处理的			大于 pivot
	//
	//
	//	最后是:
	//
	//	| - - - - - - - - - - - - - | - - - - - - - - - - - | - - - - - - - - - - - - - - |
	//
	//	left						lt						gt							right
	//
	//			小于 pivot 的			等于 pivot 的					大于 pivot 的
	//
	//
	if left < right  {
		pivot := arr[left]
		lt := left
		gt := right
		k := left + 1  // 第一个元素是切分元素，所以指针i可以从lo+1开始
		for k <= gt {
			if arr[k] < pivot {   // 小于切分元素的放在lt左边，因此指针lt和指针i整体右移
				arr[lt], arr[k] = arr[k], arr[lt]
				lt++
				k++
			} else if arr[k] > pivot {  // 大于切分元素的放在gt右边，因此指针gt需要左移
				arr[gt], arr[k] = arr[k], arr[gt]
				gt--
			} else {
				k++
			}
		}

		// lt - gt 的元素已经排定，只需对 it左边 和 gt右边 的元素进行递归求解
		suffleThreeWay(arr, left, lt- 1)
		suffleThreeWay(arr, gt+ 1, right)
	}
}

// todo  三分 双向 快排
//		是 三分 单向 的一种优化， (亦是 单轴 的一种优化)
//func





// todo 双轴 快排 (最优 写法??)
//		双轴快速排序算法思路和三向切分快速排序算法的思路基本一致，
// 		双轴快速排序算法使用两个轴 （pivot），通常选取最左边的元素作为pivot1 和 最右边的元素作pivot2.
// 		首先要比较这两个轴的大小，如果pivot1 > pivot2，则交换最左边的元素和最右边的元素，已保证pivot1 <= pivot2,
// 		双轴快速排序同样使用i，j，k三个变量将数组分成四部分.

func quickSort4 (arr []int) {
	doublePivot(arr, 0, len(arr) - 1)
}

func doublePivot(arr []int, left, right int) {
	/**

	todo https://www.cnblogs.com/nullzx/p/5880191.html
	todo https://blog.csdn.net/Holmofy/article/details/71168530

	| pivot1, - - - - - - - - - - - - - - - | - - - - - - - - - - - - - - - - | - - - - - - - - - - - - - - - - - | - - - - - - - - - -  - - - - - - , pivot2 |

	left,  | - - - - - - - -  - - - - - - i | - - - - - - - - - - - - - - - - | k - - - - - - - - - - - - - - - - | j - - - - - - - - - - - - -  - - | right
				 x < pivot1                     pivot1  <= x  <= pivot2 				还未处理部分					pivot2 < x

	todo 步骤
		A[L+1, i]    	是小于pivot1的部分，
		A[i+1, k-1]		是大于等于pivot1且小于等于pivot2的部分，
		A[j, R]			是大于pivot2的部分，而A[k, j-1]是未知部分。

		todo 和三向切分的快速排序算法一样，初始化 i = L，k = L+1，j=R，k自   左向右   扫描直到 k与j 相交为止（k == j）。
			我们扫描的目的就是逐个减少未知元素，并将每个元素按照和pivot1和pivot2的大小关系放到不同的区间上去。

		todo 在k的扫描过程中我们可以对a[k]分为三种情况讨论（注意我们始终保持最左边和最右边的元素，即双轴，不发生交换）

			（1）a[k] < pivot1   i先自增，交换a[i]和a[k]，k自增1，k接着继续扫描

			（2）a[k] >= pivot1 && a[k] <= pivot2 k自增1，k接着继续扫描

			（3）a[k] > pivot2: 这个时候显然a[k]应该放到最右端大于pivot2的部分。
								但此时，我们不能直接将   a[k]与j的下一个位置a[--j]交换
								（可以认为 A[j]  与 pivot1 和 pivot2 的大小关系在上一次j自右向左的扫描过程中就已经确定了，这样做主要是 j 首次扫描时避免 pivot2 参与其中），
								因为目前 a[--j] 和 pivot1 以及 pivot2 的关系未知，所以我们这个时候应该从j的下一个位置（--j）自右向左扫描。
								而a[--j]与pivot1和pivot2的关系可以继续分为三种情况讨论.

			       3.1）a[--j] > pivot2 j接着继续扫描

			       3.2）a[--j] >= pivot1且a[j] <= pivot2 交换a[k]和a[j]，k自增1，k继续扫描（注意此时j的扫描就结束了）

			       3.3） a[--j] < pivot1 先将i自增1，此时我们注意到	a[j] < pivot1,  a[k] > pivot2,  pivot1 <= a[i] <=pivot2，
										那么我们只需要将 a[j] 放到 a[i] 上，a[k] 放到 a[j] 上，而 a[i] 放到 a[k] 上。
										k自增1，然后k继续扫描（此时j的扫描就结束了）



	todo 注意：

		1. 	pivot1 和 pivot2在始终不参与k，j扫描过程

		2. 	扫描结束时，A[i]表示了小于pivot1部分的最后一个元素，A[j]表示了大于pivot2的第一个元素，
		  	这时我们只需要交换pivot1（即A[L]）和A[i]，交换pivot2（即A[R]）与A[j]，
			同时我们可以确定A[i]和A[j]所在的位置在后续的排序过程中不会发生变化
			（这一步非常重要，否则可能引起无限递归导致的栈溢出），
			todo 最后我们只需要对   A[L, i-1]，   A[i+1, j-1]，  A[j+1, R] 这三个部分继续递归上述操作即可。
	 */

	if left < right {

		// 以 收尾 定位最开始的 pivot1 和  pivot2
		if arr[left] > arr[right] {
			arr[left], arr[right] = arr[right], arr[left] //保证pivot1 <= pivot2
		}
		pivot1 := arr[left]
		pivot2 := arr[right]

		i := left + 1
		k := left + 1
		j := right - 1

	OUT_LOOP:
		for k <= j {
			if arr[k] < pivot1 {
				arr[i], arr[k] = arr[k], arr[i]
				k++
				i++
			}else {

				if arr[k] >= pivot1 && arr[k] <= pivot2 {
					k++
				}else{

					for arr[j] > pivot2 {
						j--
						if j < k { // 当 k 和 j 错过
							break OUT_LOOP
						}
					}
					if arr[j] >= pivot1 && arr[j] <= pivot2 {
						arr[k], arr[j] = arr[j], arr[k]
						k++
						j--
					}else{ // arr[j] < pivot1
						arr[k], arr[j] = arr[j], arr[k] // 注意 k 不动
						j--
					}
				}
			}

		}
		i--
		j++

		// 最后  k 和 j 重合,  且 i  和 j 就是当前的 双轴, 以 轴为切点 分成三部分  x < pivot1 (arr[i]) | pivot1 <= y <= pivot2 | pivot2 (arr[j]) < z 继续递归 新的双轴

		// todo 最后
		arr[left], arr[i] = arr[i], arr[left]		//将pivot1交换到适当位置
		arr[right], arr[j] = arr[j], arr[right]		//将pivot2交换到适当位置

		//	一次双轴切分至少确定两个元素的位置，这两个元素将整个数组区间分成三份
		doublePivot(arr, left, i-1)					// x 的取值范围
		doublePivot(arr, i+1, j-1)					// y 的取值范围
		doublePivot(arr, j+1, right)				// z 的取值范围
	}


}

//////////////////////////////////////////////////////////////////////////////// 插入排序 ////////////////////////////////////////////////////////////////////////
/**
插入排序
基本思想：

在要排序的一组数中，todo 假定前n-1个数已经排好序，现在将第n个数插到前面的有序数列中，
使得这n个数也是排好顺序的。如此反复循环，直到全部排号顺序
 */
func insertSort1(arr []int) {

	// todo  就用这种吧

	for i := 0; i < len(arr) - 1; i++ {

		/**
		第一层for指的是对数组的位置做遍历，

		第二层for是对第一层for遍历到的位置的后一个元素往前做遍历，
				并逐一和前面的元素做对比，把小的元素往前推,这样来实现插入到某一个位置
		 */
		for j := i + 1; j > 0; j-- {

			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}else {
				break
			}
		}
	}
}


func insertSort2(arr []int) {  // 总之，不建议这种比较绕的写法
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			j := i - 1
			temp := arr[i]
			for j >= 0 && arr[j] > temp {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = temp
		}
	}
}


func insertSort3(arr []int) {
	var n = len(arr)
	for i := 1; i < n; i++ { // 从第二个开始遍历
		j := i
		for j > 0 { // 依次往前对比
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j = j - 1
		}
	}
}


//////////////////////////////////////////////////////////////////////////////// 希尔排序 ////////////////////////////////////////////////////////////////////////

/**
希尔排序

基本思想：

在要排序的一组数中，根据某一增量分为若干子序列，并对子序列分别进行插入排序。
然后逐渐将增量减小,并重复上述过程。直至增量为1,此时数据序列基本有序,最后进行插入排序
 */
func shellSort(arr []int) {  // 这个没有第二个好理解，但是和第二个的执行是一样的
	increment := len(arr) // 设置初初始步进为 数组的长度
	for {
		increment = increment / 2 // 第一次排序用 1/2 数组长度
		for i := 0; i < increment; i++ { // 每次外层遍历只 遍历至和步进一样的次数
			for j := i + increment; j < len(arr); j = j + increment { // 以 步进来取需要做 插入排序的数组段
				for k := j; k > i; k = k - increment { // 由起点往前做插入排序
					if arr[k] < arr[k-increment] {
						arr[k], arr[k-increment] = arr[k-increment], arr[k]
					} else {
						break
					}
				}
			}
		}
		if increment == 1 { // 当最终步进为 1 时，停止排序
			break
		}
	}
}
// 以下只拿第一次的来说，后面的依次类推
/**
	以下为下面代码中第一次排序的逻辑
	6 5 5 9 3 4 8 7 11 8 10 2
	|_________| 4 和 6 做排序 == 4 5 9 3 6 8 7 11 10 2
	|_________|_____________|  2 和 6 做排序，再和4做排序 == 2 5 9 3 4 8 7 11 8 10 6
	第二次是 1/4 数组长度
	2 5 5 9 3 4 8 7 11 8 10 6
	|___|  5 和 2 排序 == 2 5 5 9 3 4 8 7 11 8 10 6
	|___|_____| 4 和 5 做排序 再和 2 做排序 == 2 5 4 9 3 5 8 7 11 8 10 6
	|___|_____|______| 11 和 5 做排序再和4做排序再和2做排序 == 2 5 4 9 3 5 8 7 11 8 10 6
	|___|_____|______|______| 6和11做排序再和5做排序再和4做排序再和2做排序 == 2 5 4 9 3 5 8 7 6 8 10 11
	然后再次缩短步进，重复上面的操作，依次类推直到步进 == 1

 */
// todo  这种写的不错
func shellSort2(arr []int) {

	// 确定步长，第一次先为 1/2 的数组长度，往后每一次遍历都逐渐减半
	for gap := len(arr) / 2; gap >= 1; gap /= 2 {
		// 对每一个步长对应的数组进行插入排序，从从第一个步进处开始
		for i := gap; i < len(arr); i += gap {
			flag := i - gap //从0个开始
			temp := arr[i] // 步进处的值
			for flag >= 0 && arr[flag] > temp { // 如果第零个的值大于第一个步进处的值，则把第零个值置赋值给第一个步进处
				arr[flag+gap] = arr[flag]
				flag -= gap	// 从第零处往前减一个步进
			}
			arr[flag+gap] = temp // 目前这是第零个，把之前缓存的步进处的值赋给第零个，完成了置换
		}
	}
}


// todo  这种逼格更高
func shellSort3(items []int) {
	var (
		n    = len(items)
		gaps = []int{1}
		k    = 1
	)

	for {
		gap := pow(2, k) + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...) // 原先步进小的都往尾部增加，大的步进都在 切片前面
		k++
	}

	for _, gap := range gaps { // 这块逻辑和方法2 一致
		for i := gap; i < n; i += gap {
			j := i
			for j > 0 {
				if items[j-gap] > items[j] {
					items[j-gap], items[j] = items[j], items[j-gap]
				}
				j = j - gap
			}
		}
	}
}
// 求间隙
func pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}


//////////////////////////////////////////////////////////////////////////////// 归并排序 ////////////////////////////////////////////////////////////////////////

/**
归并排序
基本思路：

先递归分解数组，再合并数组。先考虑合并两个有序数组，基本思路是比较两个数组的最前面的数，【谁小就先取谁】，取了后相应的指针就往后移一位。
然后再比较，直至一个数组为空，最后把另一个数组的剩余部分复制过来即可。再考虑递归分解，基本思路是将数组分解成left和right，
如果这两个数组内部数据是有序的，那么就可以用上面合并数组的方法将这两个数组合并排序。如何让这两个数组内部是有序的？可以再二分，
直至分解出的小组只含有一个元素时为止，此时认为该小组内部已有序。然后合并排序相邻二个小组即可
 */
func mergeSort(arr []int) {
	length := len(arr)
	temp := make([]int, length) // 提前开辟一块内存空间存放临时数据
	mSort(arr, 0, length-1, temp)
}

func mSort(arr []int, left, right int, temp []int) {
	if left < right {
		mid := (left + right) / 2 // 先去中位数
		mSort(arr, left, mid, temp)
		mSort(arr, mid+1, right, temp)
		//两边的子序列都是有序的，
		//如果左边的最大的元素比右边最小的元素大才需要合并
		if arr[mid] > arr[mid+1] {
			merge(arr, left, mid, right, temp)
		}
	}
}

func merge(arr []int, left, mid, right int, temp []int) {
	i := left
	j := mid + 1
	t := 0 //临时slice的指针
	for i <= mid && j <= right { // 开始同时遍历左边和右边两部分数组 【该循环会一直执行到左右其中一个数组为空】
		if arr[i] <= arr[j] { // 左边数组的第N个和右边数组的第M个作比较， 当右边的 大于左边的时，开始把左边的放入 临时数组的第K位中
			temp[t] = arr[i]
			i++ // 取完后，相应的指针向后移一位
		} else {
			temp[t] = arr[j] // 同上
			j++
		}
		t++
	}
	// 等上述对比完之后， 临时数组中刚好存完且空出最后一个空位 (因为上面做了 t++)
	//将左序列剩余元素填充进temp中
	for i <= mid {
		temp[t] = arr[i]
		t++
		i++
	}
	//将右序列剩余元素填充进temp中
	for j <= right {
		temp[t] = arr[j]
		t++
		j++
	}
	t = 0
	//将temp中的元素全部拷贝到原数组中，【因为到这里位置，可能做也可能右的某个数组剩下的那部分根本还没排序】
	// 我们先从左到右把临时数组中的元素逐一放回原数组中，然后再次做递归
	for left <= right {
		arr[left] = temp[t]
		left++
		t++
	}
}

func mergeSort2(arr []int){
	// 把排好序且返回的临时数组,从新逐个加入原数组中
	for i, v := range mergeSortTemp(arr) {
		arr[i] = v
	}
}
// 这一种比第一种耗的内存多一点，临时数组起的太多了
func mergeSortTemp(arr []int) []int {
	var n = len(arr)

	if n == 1 {
		return arr
	}

	middle := int(n / 2)
	var (
		leftArr  = make([]int, middle)
		rightArr = make([]int, n-middle)
	)
	for i := 0; i < n; i++ {
		if i < middle { // 把原来数组中左右两部分分别拆分到 两个临时数组中
			leftArr[i] = arr[i]
		} else {
			rightArr[i-middle] = arr[i]
		}
	}

	return merge2(mergeSortTemp(leftArr), mergeSortTemp(rightArr)) // 因为一直这样递归的拆分，所以，最终导致最底层的临时数组只会有一个元素
}
// 合并左右两个临时数组【这个才是排序的精髓】
func merge2(leftArr, rightArr []int) (result []int) {
	result = make([]int, len(leftArr) + len(rightArr)) // 临时的回收数组

	i := 0 // result 的下标
	for len(leftArr) > 0 && len(rightArr) > 0 { // 一直对比到左右两个数组有一个比完为止
		if leftArr[0] < rightArr[0] { // 当左边第一个比右边第一个大时，把左边的加入回收数组中
			result[i] = leftArr[0]
			leftArr = leftArr[1:] // 且左数组截除第一位
		} else {
			result[i] = rightArr[0] // 右边的同上
			rightArr = rightArr[1:]
		}
		i++
	}

	for j := 0; j < len(leftArr); j++ { // 如果左边有剩余，则把剩下的追加到 回数数组末尾，以便开启新的一轮递归
		result[i] = leftArr[j]
		i++
	}
	for j := 0; j < len(rightArr); j++ { // 同理
		result[i] = rightArr[j]
		i++
	}
	return
}

/**
todo 建议 这种写法 ##############
 */
func mergeSort3 (arr []int) {
	// 把排好序且返回的临时数组, 重新逐个加入原数组中
	for i, v := range mergeSort3Temp(arr) {
		arr[i] = v
	}
}
func mergeSort3Temp(arr []int) []int{
	if len(arr) <= 1 {
		return arr
	}
	var middle int = len(arr)/2
	left := mergeSort3Temp(arr[:middle])
	right := mergeSort3Temp(arr[middle:])
	return merge3(left, right)
}

func merge3(left, right []int) []int {
	leftLen, rightLen := len(left), len(right)
	var result []int = make([]int, leftLen + rightLen) // 起一个临时的回收数组
	k := 0//数组切片result的下标
	i, j := 0, 0//a、b起始下标均未0
	for i < leftLen && j < rightLen  {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}
	// 左和右哪个有剩余，把剩余的追加到 回收数组的末尾以便之后的递归
	for i != leftLen {
		result[k] = left[i]
		k++
		i++
	}
	for j != rightLen {
		result[k] = right[j]
		k++
		j++
	}
	return result
}


//////////////////////////////////////////////////////////////////////////////// 堆排序 ////////////////////////////////////////////////////////////////////////

/**
堆排序
基本思想：

先把待排序的序列构造成最大化堆，然后把第一位和末尾一位互换，再把不包含末尾一位的剩余元素再次构造新的最大花堆，
再把新堆的首位和新堆的末尾位(也就是原数组的倒数第二位)互换。如此重复，知道最后的堆只有两个数然后 R0 与 R1互换。
 */
func heapSort(a []int) {
	length := len(a)
	if length == 0 {
		return
	}
	// 构造初始堆
	for i := length/2 - 1; i >= 0; i-- {
		heapAdjust(a, i, length-1)
	}
	// 取出最大的数，然后对剩下的再次调整成新的堆
	for j := length - 1; j >= 0; j-- {
		a[0], a[j] = a[j], a[0]
		heapAdjust(a, 0, j-1)
	}
}

// 调整堆 【是用了中间变量】
func heapAdjust(a []int, start, end int) {
	temp := a[start]

	for k := 2*start + 1; k <= end; k = 2*k + 1 { //从i结点的左子结点开始，也就是2i+1处开始
		//选择出左右孩子较大的下标
		if k < end && a[k] < a[k+1] {
			k++
		}
		//如果子节点大于父节点，将子节点值赋给父节点（不用进行交换）
		if a[k] > temp {
			a[start] = a[k]
			start = k
		} else {
			break
		}
	}
	a[start] = temp //插入正确的位置
}

func heapSort2(arr []int) {
	N := len(arr)
	var first int = N/2    //最后一个非叶子节点
	// 初始化 最大化堆
	for start := first; start > -1; start-- {    //构造大根堆
		max_heapify(arr, start, N-1)
	}
	// 堆排序，并调整新堆
	for end := N-1; end > 0; end-- {    //堆排，将大根堆转换成有序数组
		arr[end],arr[0] = arr[0],arr[end]
		max_heapify(arr, 0, end-1)
	}
}
func max_heapify(arr []int, start int, end int) {
	root := start
	for true {
		child := root*2 + 1    //调整节点的子节点
		if child > end {
			break
		}
		if child + 1 <= end && arr[child] < arr[child+1] {
			child = child + 1    //取较大的子节点
		}
		if arr[root] < arr[child] {
			arr[root], arr[child] = arr[child], arr[root]    //较大的子节点成为父节点
			root = child
		} else {
			break
		}
	}
}

/**
todo 建议用这个
 */
func heapSort3(arr []int) {
	// todo 从最后一个非叶子节点（也就是最左侧的那个非叶子结点）开始调整(len(s)/2-1)
	/**
	【1】 若根结点存在序号0处， i结点的父结点下标就为(i-1)/2。i结点的左右子结点下标分别为2i+1和2i+2
	【2】 只要从n/2-1开始，向前依次构造大根堆，这样就能保证，构造到某个节点时，它的左右子树都已经是大根堆
	比如现在有一个序列长度为15，序号0代表根节点则：
							0
						 /    \				3 层的 二叉树的节点数为  2^3-1   （深度为 k 的二叉树,最多有2^k-1个节点）
					    1      2
	 	             /   \    /   \			二叉树子树最多的节点的个数称为二叉树的度。度为2代表着深度 即该二叉树最多有三个节点
	               3     4    5     6    todo 那么我们就是从 6 = 15/2 - 1 开始构造堆
	             /  \   / \  / \   /  \
	  		    7    8 9 10  11 12 13  14  (下标 是 双数 结尾)

		“二叉树bai中的度“是指树中du最大的结点度，叶子结点是终zhi端结点，是度dao为 0 的结点。

	二叉树的度是指树中zhuan所以结点的度数的最大值。二叉树的度小于等于2，因为二叉树的定义要求二叉树中任意结点的度数（结点的分支数）小于等于2 ，并且两个子树有左右之分，顺序不可颠倒。

	叶子结点就是度为0的结点，也就是没有子结点的结点叶子。如n0表示度为0的结点数，n1表示度为1的结点，n2表示度为2的结点数。在二叉树中：n0=n2+1；N=n0+n1+n2（N是总结点）。

	todo 因为任一棵树中，结点总数 = 度数 * 该度数对应的结点数 + 1
	 */
	for i := len(arr)/2 - 1; i >= 0; i-- { // 先从 6 = 15/2 - 1 处开始构造最大化堆结构，然后依次从 5, 4, 3处构建最大化堆结构
		heapAdjust4(arr, i, len(arr))      // todo 绝不会有 元素下标 == len(arr)
	}
	// 经过上面的for初始化堆结构，则其实在 这个堆的二叉树上其实数字已经是有序的了
	//
	// 下面这个for是把最大的依次取出来，并重新调整堆结构
	for i := len(arr) - 1; i > 0; i-- {
		//将第一个和最后一个交换然后继续调整堆
		//
		// 如 14 和 0 置换后 (这时候14位置是个最大值，被取出来)
		arr[0], arr[i] = arr[i], arr[0]

		// 把剩下的 0 -> n-2  的数由0位置 (现在是14位置的值了)， 再次往下和各个子节点的值对比
		heapAdjust4(arr, 0, i)     // todo 绝不会有 元素下标 == i
	}
}

/*
// todo 看 heapAdjust4 的写法, 优雅简洁
func heapAdjust3 (arr []int, parent, end int) {

	// 记录最大的叶子结点的下标
	var maxLeaf int

	// 从 parent 节点往下 调整~
	for 2 * parent + 1 < end { // 确保  parent 是非叶子节点

		lchild := 2 * parent + 1 // 左儿子的下标
		rchild := lchild + 1	 // 右儿子的下标

		// 先比较两个叶子节点, 取最大的叶子的下标

		maxLeaf = lchild
		if rchild < end && arr[rchild] > arr[lchild] {
			maxLeaf = rchild
		}

		// 如果最大的叶子节点 大于 父节点则交换，否则推出循环

		//这样纸就会把 最大的数依次上顶到根
		if arr[maxLeaf] > arr[parent] {

			arr[parent], arr[maxLeaf] = arr[maxLeaf], arr[parent]

			parent = maxLeaf // 然后设置该位置为新的父亲,因为该位置和根置换了，所以对该位子原先下属的各个子节点的值又需要重新比较了
			// 就这样一路比下去
			//
			// todo 【其实这个只是针对上面除了6之外的非叶子节点才有用】
			//
			//同时注意，如果进到这里面来了，比方说：说明叶子节点上的最大的那个值已经和6位置的更换了，
			// 然后当parant == 2 时 6又是2的其中一个非叶子节点
			// 则，在进入这里时就是把之前6为根时的子节点上的最大值置换根后的值再次置换到2上面，
			// 依次类推最终形成0索引是最大的数，这样的最大化堆结构
		} else {
			break
		}
	}
}

*/

// todo 用这种写法, 最优雅
func heapAdjust4 (arr []int, parent, end int) {

	maxVal := parent
	for  {
		lchild := 2*parent + 1
		rchild := 2 *parent + 2
		if lchild < end && arr[lchild] > arr[maxVal] {
			maxVal = lchild
		}

		if rchild < end && arr[rchild] > arr[maxVal]{
			maxVal = rchild
		}
		// 互换位置 (parent 和 最大叶子互换)
		if parent != maxVal {
			arr[maxVal], arr[parent] = arr[parent], arr[maxVal]
			parent = maxVal  // 下一个 子树的 root 由被替换的节点担当
		}else {
			break
		}
	}
}


//////////////////////////////////////////// topk (孤岛算法) ////////////////////////////////////////////

func HeapSearchK (arr []int, topk int) {
	// 初始化原始最小堆
	smallHeapArr := buildSmallHeap(arr, topk)
	for i := topk; i < len(arr); i ++ {
		// 如果当前原始比最小堆的根元素大，那么替换根，且重新调整最小堆
		if arr[i] > smallHeapArr[0] {
			swapRoot(smallHeapArr, arr[i])
		}
	}
}
//建立小顶堆 
func buildSmallHeap(arr []int,topk int) []int{
	smallHeapArr := arr[:topk]
	for i := topk/2 - 1; i >= 0; i-- {
		adjustSmallHeap(smallHeapArr, i, topk)
	}
	return smallHeapArr
}
// 调整最小堆
func adjustSmallHeap(arr []int, parent, end int) {
	minVal := parent
	for  {
		lchild := 2*parent + 1
		rchild := 2 *parent + 2
		if lchild < end && arr[lchild] < arr[minVal] {
			minVal = lchild
		}
		if rchild < end && arr[rchild] < arr[minVal]{
			minVal = rchild
		}
		// 互换位置 (parent 和 最小叶子互换)
		if parent != minVal {
			arr[minVal], arr[parent] = arr[parent], arr[minVal]
			parent = minVal
		}else {
			break
		}
	}
}
// 替换根部，且重新构造最小堆
func swapRoot(arr []int, root int) {
	arr[0] = root // 新的根
	// 重新调整堆
	adjustSmallHeap(arr, 0, len(arr))
}





