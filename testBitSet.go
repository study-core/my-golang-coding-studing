package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// 初始化 Bitset ，首先在1 2 10 99 这几个bit 上初始化为1的 bitset
	set := NewBitSet(1, 2, 10, 99)
	// 使用 Visit 遍历所有bit，并打印出该bit的index
	set.Visit(func(n int) bool {
		fmt.Println(n)
		return false
	})
}

const (
	shift = 6    // 2^n = 64 的 n
	mask  = 0x3f // n=6，即 2^n - 1 = 63，即 0x3f
)

// 表示一个长度可变的大位图
// 状态标志位 元素个数固定 一个整数的就足以表示左右的状态
// 集合场景 元素个数不确定 一个整数不足以表示
type BitSet struct {
	data []uint64 // uint64每个bit上表示 一个状态标识
	size int      // 用于存放集合元素的个数
}

// 该函数 获取 某个元素的bit处于动态位图的 data 中的第几个 int64 上
func index(n int) int {
	return n >> shift
}

// 相对于标志位使用场景中某个标志的值
// 该函数 获取 某个元素处于位图的某个int64的第几个 bit，且将索引值转化成uint64的十进制返回
func posVal(n int) uint64 {
	// value & 0x3f ==  value % 64
	return 1 << uint(n&mask)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// 创建一个 BitSet
//
// ns 是一个 int 类型的变长参数，用于设置集合中的初始bit上的1
func NewBitSet(ns ...int) *BitSet {
	if len(ns) == 0 {
		return new(BitSet)
	}

	// 确定给 bitset 分配多大存放集合元素
	// 循环所有入参的 ns 看看哪个数是大数，用最大数决定最大分配
	max := ns[0]
	for _, n := range ns {
		if n > max {
			max = n
		}
	}

	if max < 0 {
		return new(BitSet)
	}

	// 通过 max >> shift+1 计算最大值 max 所在 index
	// 而 index + 1 即为要开辟的空间
	s := &BitSet{
		data: make([]uint64, index(max)+1),
	}

	// 在对应的 bit 上初始化 1
	for _, n := range ns {
		if n >= 0 {
			// 这个骚操作

			// e >> shift 获取索引位置，即行，一般叫 index
			// e&mask 获取所在列，一般叫 pos，F1 0 F2 1
			s.data[index(n)] |= posVal(n)
			// 增加元素个数
			s.size++
		}
	}

	return s
}

func (set *BitSet) Contains(n int) bool {
	// 应该存在什么位置
	i := index(n)
	if i >= len(set.data) {
		return false
	}

	// 通过它判断是否存在指定元素
	return set.data[i]&posVal(n) != 0
}

func (set *BitSet) Clear(n int) *BitSet {
	if n < 0 {
		return set
	}

	i := index(n)
	if i >= len(set.data) {
		return set
	}

	// 通过它判断是否存在指定元素
	if set.data[i]&posVal(n) != 0 {

		// 清除对应bit上的 1
		set.data[i] &^= posVal(n)
		// 递减计数
		set.size--
	}

	// 重新决定是否收缩 bitset
	set.trim()

	return set
}

func (set *BitSet) Add(n int) *BitSet {
	if n < 0 {
		return set
	}

	i := index(n)
	if i >= len(set.data) {
		// 扩展 bitset 容量
		ndata := make([]uint64, i+1)
		copy(ndata, set.data)
		set.data = ndata
	}

	// 如果不存在，则在对应的bit上添加1
	if set.data[i]&posVal(n) == 0 {
		set.data[i] |= posVal(n)
		set.size++
	}

	return set
}

func (set *BitSet) Size() int {
	return set.size
}

func (set *BitSet) Intersect(other *BitSet) *BitSet {

	// 求交集，获取最小的容量
	minLen := min(len(set.data), len(other.data))

	// 创建一个存放交集的bitset
	intersectSet := &BitSet{
		data: make([]uint64, minLen),
	}

	// 根据最小的 索引，去遍历爽bitset，并将各个元素的交集存起来
	for i := 0; i < minLen; i++ {
		intersectSet.data[i] = set.data[i] & other.data[i]
	}

	intersectSet.size = set.computeSize()

	return intersectSet
}

func (set *BitSet) Union(other *BitSet) *BitSet {
	var maxSet, minSet *BitSet
	if len(set.data) > len(other.data) {
		maxSet, minSet = set, other
	} else {
		maxSet, minSet = other, set
	}

	unionSet := &BitSet{
		data: make([]uint64, len(maxSet.data)),
	}

	// 先把大对于小的 多余出来的那部分，直接追加到并集的末尾
	minLen := len(minSet.data)
	copy(unionSet.data[minLen:], maxSet.data[minLen:])

	for i := 0; i < minLen; i++ {
		unionSet.data[i] = set.data[i] | other.data[i]
	}
	unionSet.size = unionSet.computeSize()
	return unionSet
}

// 差集
// A和B求差集，说的是 A -(A∩B)
func (set *BitSet) Difference(other *BitSet) *BitSet {
	setLen := len(set.data)
	otherLen := len(other.data)

	differenceSet := &BitSet{
		data: make([]uint64, setLen),
	}

	minLen := setLen
	if setLen > otherLen {
		copy(differenceSet.data[otherLen:], set.data[otherLen:])
		minLen = otherLen
	}

	for i := 0; i < minLen; i++ {
		// 使用 A &^ B == A - (A∩B)
		differenceSet.data[i] = set.data[i] &^ other.data[i]
	}

	differenceSet.size = differenceSet.computeSize()

	return differenceSet
}

// 该函数逐个遍历bitset上所有的 bit
// 入参的为一个 接收 元素 索引的 func
func (set *BitSet) Visit(do func(int) (skip bool)) (aborted bool) {

	d := set.data

	// 逐个遍历 bitset中的 uint64
	for i, len := 0, len(d); i < len; i++ {

		w := d[i]
		if w == 0 { // 如果该 uint64 == 0 说明该 uint64 的bit 上没有元素，跳过
			continue
		}

		// n： 计算出该 uint64 所代表的 bit范围
		// 这小段代码可以理解为从元素值到 index 的逆运算，
		// 只不过得到的值是诸如 0、64、128 的第一个位置的值。
		// 0 << 6，还是 0，1 << 6 就是 64，2 << 6 的就是 128
		//  0 ~ 63 : 64bit, index: 1
		//  64 ~ 127 : 64bit, index: 2
		//  128 ~ 192 : 64bit, index: 3
		//  ... 以此类推
		n := i << shift

		for w != 0 {
			// 000.....000100 64~127 的话，表示 66 index，即 64 + 2，这个 2 可以由结尾 0 的个数确定
			// 那怎么获取结果 0 的个数呢？可以使用 bits.TrailingZeros64 函数
			b := bits.TrailingZeros64(w) // 这个得到 某个数值转化成二进制之后末尾有几个0
			if do(n + b) {
				return true
			}
			// 将已经检查的位清零
			// 为了保证尾部 0 的个数能代表元素的值
			w &^= 1 << uint64(b)
		}
	}
	return false
}

// 计算元素的个数
func (set *BitSet) computeSize() int {
	d := set.data
	n := 0
	for i, l := 0, len(d); i < l; i++ {
		if w := d[i]; w != 0 {
			n += bits.OnesCount64(w)
		}
	}

	return n
}

// 清楚后面为 0 的 bit (uint64e)
func (set *BitSet) trim() {
	d := set.data
	n := len(d) - 1
	for n >= 0 && d[n] == 0 {
		n--
	}
	set.data = d[:n+1]
}
