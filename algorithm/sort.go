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

	heapSort3(items)
	fmt.Println(items[len(items) - 7:])


	HeapSearchK(items, 7)
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
	for swapped {
		swapped = false
		for i := 0; i < n - 1; i++ {
			/**
			因为外层是个死循环，内层为值遍历到数组的倒数第二个，然后遍历到的每一个都是用当前的和后面一个最对比
			如果发现只要有需要把大的往后推的就把标识位改为true，但是如果在遍历完本次之后发现就没有 i > i + 1 的，那么就属于数组全部都排序好的了
			如果该次数组的遍历有大的就把他逐个往后推到末尾，然后标识位为 true 且 在进入下一轮 遍历之前把下一轮遍历的数组元素遍历个数 减1，【因为最后一个已经是拍好的最大的了】
			这个做法比 方法2的好，因为增加了标识位的判断避免了多余的 for 的遍历次数
			而方法2是预定好一定会做 n - 1次对数组的for，导致真正整体循环的次数是 n的阶乘，而本方法根据实际情况(由标识位来控制)确定是否需要是n的完整阶乘还是一半额阶乘
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
			if arr[i] > arr[i+1] {  // 对数组的各个元素逐个对比，把大的值一直推到末尾
				/**
				对于 for 中的 n - j - 1 解释
				当第一层for 执行第一次时，里面的for就是 n - 0 - 1 把当前数组的所有元素都逐一对比，把最大的推到末尾
				当第一层for执行第二次时，里面的for是 n - 1 - 1 就是把数组减掉最后一个(因为最后一个是最大值)去做逐一对比
				当第一层for执行第三次时，里面的for是 n - 2 - 1 就是把数组减掉最后两个(因为后面两个是数组汇总依次最大的两个值)去做逐一对比
				依次类推当for执行到 第 n - 1次时，里面的for是 n - (n - 2) - 1 == 1 就是剩下的最后一个最小的被排到第一位
				所以是： (n - 0 - 1) * (n - 1 - 1) * (n - 2 - 1) * ... * (n - (n - 2) - 1) == n的阶乘
				 */
				arr[i+1], arr[i] = arr[i], arr[i+1]
			}
		}
	}

	/**
	这一种是做浪费for次数的冒泡，完全就是 n的n次方的运算
	 */
	//for i := 0; i < n - 1; i++ {
	//	for j := 0; j < n - 1; j++ {
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
 */
func selectSort(arr []int) {
	var minIndex int
	for i := 0; i < len(arr) - 1; i++ { // 表示对数字做for遍历的次数
		minIndex = i // 初始化下标为0
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

func selectSort2(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		maxIndex := 0
		/**
		第二种做法是，每一次对数组的for都需要把最大的书的下标记录下来，然后放置本次for的数组的末尾(本次for的数组是除去上次for之后的某尾数的元素)
		也是n的阶乘，比第一种写法少了一个if判断，效率提升一些
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
如果有多个元素，先指定一个数轴(也就是一个基准数，开始第一遍的时候可以是最左边的数)
然后根据数轴把数分成 左右两部分，然后继续递归
 */
func quickSort1(nums []int) {
	reSort(nums, 0, len(nums)-1)
}
func reSort(arr []int, left, right int) {
	if left < right {
		i := arrAdjust(arr, left, right)
		reSort(arr, left, i-1)
		reSort(arr, i+1, right)
	}
}

//返回调整后基准数的位置
func arrAdjust(arr []int, left, right int) int {
	x := arr[left] //基准
	for left < right {
		//从右向左找出小于 基准数的下标
		for left < right &&  x <= arr[right] {
			right-- //当出现右边有数小于左边的数时，for停止进入下面if
		}
		if left < right { // 如果 右边的下标已经 <= 左边的下标那么说明，上面的for已经找过头了，右边往左缩进缩过头了
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
	到这里的操作步骤就是：先从右往左注意和基准数对比，把小于基准数的都移到左侧；
	再从左到右把大于等于基准数的都移到右边，不管是从左到右还是从右到左，的每一次for都是逐一往中心靠拢的游标
	最后把，基准数放置最后一次替换的那个位置，这样纸最终就达到了 以该基准数的新位置为新的左右分割点，形成左侧都是小于基准数，右侧都是大于基准数
	最终把新的基准位置返回，并进入下次左右两个 数组的迭代上述重复动作，知道最终排完
	 */

	return left
}



// 第一种是做了个中间变量来实现数据互换的，这个是直接互换，写法简洁一点
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
// 精华就在这个里面
func partition(arr []int, left int, right int) int {
	for left < right { // 为什么要这么做 因为看里面的逻辑我们知道，left和right是逐渐往中间(不是二分中心点)靠拢的
		for left < right && arr[left] <= arr[right] { //先逐个缩进遍历右侧的元素，由右到左，到发现有小于基准数的时候则记住当前下标 往下走
			right--
		}
		if left < right {
			// 互换基准数和 对比数的位置，注意现在 基准数是在 之前对比数的位置的，然后开始从之前基准数所在的下一个位置开始启动左边对比
			arr[left], arr[right] = arr[right], arr[left]
			left++
		}

		for left < right && arr[left] <= arr[right] { // 逐个缩进遍历新的左侧(从之前右侧对比的基准的下一个位置开始)，由左到右和基准数对比(注意次时基准数的下标是和之前右侧查找的数下标互换了)，直到发现有大于等于基准数的时候记住下标 往下走
			left++
		}
		if left < right {
			// 把基准数和对比数互换位置
			arr[left], arr[right] = arr[right], arr[left]
			right--
		}
	}
	/**
	直到最后 基准数被调节到某个位置，且该位置左边的元素都小于基准数，右边的元素都大于等于基准数，且返回该基准数最终的位置，把数组分为左右两部分各自进入新的递归
	 */
	return left
}

//////////////////////////////////////////////////////////////////////////////// 插入排序 ////////////////////////////////////////////////////////////////////////
/**
插入排序
基本思想：在要排序的一组数中，假定前n-1个数已经排好序，现在将第n个数插到前面的有序数列中，使得这n个数也是排好顺序的。如此反复循环，直到全部排号顺序
 */
func insertSort1(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {
		/**
		第一层for指的是对数组的位置做遍历，第二层for是对第一层for遍历到的位置的后一个元素往前做遍历，并逐一和前面的元素做对比，把小的元素往前推,这样来实现插入到某一个位置
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
基本思想：在要排序的一组数中，根据某一增量分为若干子序列，并对子序列分别进行插入排序。
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
	|_________|_____________|  2 和 6 做排序，再和4做排序 == 2 5 5 9 3 4 8 7 11 8 10 6
	第二次是 1/4 数组长度
	2 5 5 9 3 4 8 7 11 8 10 6
	|___|  5 和 2 排序 == 2 5 5 9 3 4 8 7 11 8 10 6
	|___|_____| 4 和 5 做排序 再和 2 做排序 == 2 5 4 9 3 5 8 7 11 8 10 6
	|___|_____|______| 11 和 5 做排序再和4做排序再和2做排序 == 2 5 4 9 3 5 8 7 11 8 10 6
	|___|_____|______|______| 6和11做排序再和5做排序再和4做排序再和2做排序 == 2 5 4 9 3 5 8 7 6 8 10 11
	然后再次缩短步进，重复上面的操作，依次类推直到步进 == 1

 */
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
基本思路：先递归分解数组，再合并数组。先考虑合并两个有序数组，基本思路是比较两个数组的最前面的数，【谁小就先取谁】，取了后相应的指针就往后移一位。
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
建议 这种写法
 */
func mergeSort3 (arr []int) {
	// 把排好序且返回的临时数组,从新逐个加入原数组中
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
基本思想：先把待排序的序列构造成最大化堆，然后把第一位和末尾一位互换，再把不包含末尾一位的剩余元素再次构造新的最大花堆，
再把新堆的首位和新堆的末尾位(也就是原数组的倒数第二位)互换。如此重复，知道最后的堆只有两个数然后 R0 与 R1互换。
 */
func heapSort(a []int) {
	length := len(a)
	if length == 0 {
		return
	}
	//构造初始堆
	for i := length/2 - 1; i >= 0; i-- {
		heapAdjust(a, i, length-1)
	}
	// 取出最大的数，然后对剩下的再次调整成新的堆
	for j := length - 1; j >= 0; j-- {
		a[0], a[j] = a[j], a[0]
		heapAdjust(a, 0, j-1)
	}
}

//调整堆 【是用了中间变量】
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
建议用这个
 */
func heapSort3(arr []int) {
	//从最后一个非叶子节点（也就是最左侧的那个非叶子结点）开始调整(len(s)/2-1)
	/**
	【1】 若根结点存在序号0处， i结点的父结点下标就为(i-1)/2。i结点的左右子结点下标分别为2i+1和2i+2
	【2】 只要从n/2-1开始，向前依次构造大根堆，这样就能保证，构造到某个节点时，它的左右子树都已经是大根堆
	比如现在有一个序列长度为15，序号0代表根节点则：
							0
						 /    \
					    1      2
	 	             /   \    /   \
	               3     4    5     6   那么我们就是从 6 = 15/2 - 1 开始构造堆
	             /  \   / \  / \   /  \
	  		    7    8 9 10  11 12 13  14
	 */
	for i := len(arr)/2 - 1; i >= 0; i-- { // 先从 6 = 15/2 - 1 处开始构造最大化堆结构，然后依次从 5, 4, 3处构建最大化堆结构
		heapAdjust3(arr, i, len(arr))
	}
	// 经过上面的for初始化堆结构，则其实在 这个堆的二叉树上其实数字已经是有序的了
	// 下面这个for是把最大的依次取出来，并重新调整堆结构
	for i := len(arr) - 1; i > 0; i-- {
		//将第一个和最后一个交换然后继续调整堆
		arr[0], arr[i] = arr[i], arr[0] // 如 14 和 0 置换后(这时候14位置是个最大值，被取出来)
		heapAdjust3(arr, 0, i)  // 把剩下的 0 -> n-2  的数由0位置(现在是14位置的值了)，再次往下和各个子节点的值对比
	}
}

func heapAdjust3 (arr []int, parent, len int) {
	var i int
	for 2 * parent + 1 < len { //确保是非叶子节点

		lchild := 2 * parent + 1 // 左儿子的下标
		rchild := lchild + 1	 // 右儿子的下标
		i = lchild
		//取出两个叶子节点中最大的一个
		if rchild < len && arr[rchild] > arr[lchild] {
			i = rchild
		}
		//如果最大的叶子节点大于父节点则交换，否则推出循环
		if arr[i] > arr[parent] { //这样纸就会把 最大的数依次上顶到根
			arr[parent], arr[i] = arr[i], arr[parent]
			parent = i // 然后设置该位置为新的父亲,因为该位置和根置换了，所以对该位子原先下属的各个子节点的值又需要重新比较了
			// 就这样一路比下去
			// 【其实这个只是针对上面除了6之外的非叶子节点才有用】
			//同时注意，如果进到这里面来了，比方说：说明叶子节点上的最大的那个值已经和6位置的更换了，
			// 然后当parant == 2 时 6又是2的其中一个非叶子节点
			// 则，在进入这里时就是把之前6为根时的子节点上的最大值置换根后的值再次置换到2上面，
			// 依次类推最终形成0索引是最大的数，这样的最大化堆结构
		} else {
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
	// 最大的K个数
	fmt.Println(smallHeapArr)
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
func adjustSmallHeap(arr []int, parent, length int) {
	i := parent
	for  {
		lchild := 2*parent + 1
		rchild := 2 *parent + 2
		if lchild < length && arr[lchild] < arr[i] {
			i = lchild
		}
		// 右节点和根
		if rchild < length && arr[rchild] < arr[i]{
			i = rchild
		}
		// 互换位置
		if parent != i {
			arr[i], arr[parent] = arr[parent], arr[i]
			parent = i
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





func quickSort(s []int) []int {
	if len(s) < 2 {
		return  s
	}
	v := s[0]
	var left, right []int
	for _, e := range s[1:] {
		if e <= v {
			left = append(left, e)
		} else {
			right = append(right, e)
		}
	}
	return  append(append(quickSort(left), v), quickSort(right)...)
}