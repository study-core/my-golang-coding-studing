package main

import "fmt"

// 散列表
func main() {
	npages := ^uintptr(0)
	npages2 := uintptr(0)
	fmt.Println(npages, npages2)
}

/*

和数组相比，数组的寻址比较容易，插入和删除困难（最差为O(n)），而链表寻址困难（O(n)），插入和删除比较容易，

todo  散列表除了不支持排序，散列表寻址容易，同时插入和删除也容易。散列表有以下特点：

1、访问速度快，最快为O(1)
2、需要额外空间，散列表需要更多的空间来避免散列冲突或聚集
3、无序，散列表不能实现元素排序
4、可能产生碰撞，散列碰撞会影响散列表的执行效率


todo  散列 函数

		1、




todo  填充 因子


 */

