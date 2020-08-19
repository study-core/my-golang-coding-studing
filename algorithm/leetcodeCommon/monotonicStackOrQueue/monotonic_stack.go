package main

import (
	"fmt"
)

// todo 单调栈
//
//      从栈顶到栈底的元素是严格递增 (or 递减)
//
// 规则:
//
// 如果栈中是5 4 3 2 1，如果压入3怎么办？
// 原来我们只需要添加到栈尾即可，现在则需要将3 2 1弹出，再压入3，栈变成5 4 3
// 注：弹出的元素我们直接舍弃掉。
func main() {
	arr := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(arr))
}


/**
todo 单调栈有什么应用？
	`单调递增栈` 能表示入栈元素左边第一个比它 `大` 的元素
	`单调递减栈` 能表示入栈元素左边第一个比它 `小` 的元素
 */

/**
todo  题 1

	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

	输入: [0,1,0,2,1,0,1,3,2,1,2,1]

	输出: 6



todo  解法：
     使用单调递增栈，我们知道栈有两个操作，入栈和出栈。单调栈的出入栈表示：

todo
	入栈：  表明本身比栈顶小，说明在下台阶，下台阶是 形成不了积水的。
	出栈：  表明本身比栈顶大，肯定会形成积水。


每次计算积水肯定是在pop的时候计算。
而且栈里最少有两个元素的时候才会形成，
因为最小的积水也是有两个边和一个坑组成的，最少也是栈里两个加上刚来的一个。

这里我们可以想象成一个木桶，todo 根据木桶理论，容量由最低的那块木板决定，所以桶的容量需要由 长木板+桶底+短木板决定

遍历到数组下标 3 的时候。栈：[1, 0]，来了一个2，2比0大，0要出栈，
这个时候就可以知道1和2中间夹了一个0，找1和2最小值，短木板是1，桶底是0，
所以宽度是1，高度是1，得到面积是1.


遍历到数组下标 6 的时候。栈：[2，1，0]，来了1，所以1和1中间夹了0，同样得到面积是1，得到的是浅蓝色的部分。


继续往后遍历来到7，来了一个3，栈：[2,1] (看入栈的逻辑，栈里也可以是[2,1,1]，相同的1可以入可以不入)。
假设是[2,1]，先弹1，短木板是1，长木板是3，但是桶底也是1，所以能装的水是0。
1弹出之后栈里还剩[2]，短木板是2，长木板是3，桶底是1，水深度为1，宽度index = 6-3 = 3，所以面积是3.


 */


func trap(height []int) int {

	if len(height) == 0 {
		return 0
	}

	// 记录累计和
	var sumArea int

	// 初始化一个装 木板的栈, 最坏的打算是 木板全装
	stack := NewMonotonicStack(len(height))

	// todo  算法精髓

	// 遍历 arr 的右边界
	for right := 0; right < len(height); right++ {

		// 如果当前 栈顶的元素 < 当前数组元素
		for !stack.IsEmpty() && height[stack.Peek()] <= height[right] {
			if stack.Size() >= 2 {
				j := stack.Pop()
				left := stack.Peek()
				waterHeight := MinFn(height[right], height[left]) - height[j] // 就像一个木桶：得到最低的木板减去底得到能装水的高度
				waterLength := right - left - 1
				curArea := waterHeight * waterLength
				sumArea += curArea
			} else {
				stack.Pop()
			}
		}
		stack.Push(right)
	}
	return sumArea
}


type monotoniNode struct {
	pre *monotoniNode
	value int
}
type monotonicIncreaseStack struct {
	capacity int
	size int
	curr *monotoniNode
}

func NewMonotonicStack(capacity int) *monotonicIncreaseStack {
	return &monotonicIncreaseStack{
		capacity: capacity,
		size:     0,
		curr:     nil,
	}
}
func (self *monotonicIncreaseStack) IsEmpty() bool {
	return self.size == 0
}
func (self *monotonicIncreaseStack) Size() int {
	return self.size
}
func (self *monotonicIncreaseStack) Capacity() int {
	return self.capacity
}
func (self *monotonicIncreaseStack) Push(value int) bool {

	if self.size == self.capacity {
		return false
	}
	newNode := &monotoniNode{
		pre:   self.curr,
		value: value,
	}
	self.curr = newNode
	self.size ++
	return true
}
func (self *monotonicIncreaseStack) Peek() int {
	return self.curr.value
}
func (self *monotonicIncreaseStack) Pop() int {
	if self.size == 0 {
		return -1
	}
	node := self.curr
	self.curr = node.pre
	self.size --
	return node.value
}

func MinFn(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxFn(a, b int) int {
	if a < b {
		return b
	}
	return a
}