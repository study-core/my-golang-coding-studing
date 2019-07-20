package main

import (
	"unsafe"
	"fmt"
	"sync/atomic"
)
/**
保证 check 在第一次被使用后，不能被复制
 */
func main() {
	h := Human{
		Name: "A",
	}
	fmt.Println("第一次使用")
	h.C.checkFunc()
	fmt.Println("第二次使用,但没有被复制")
	h.C.checkFunc()

	hCopy := h // 这里相当于把 h 复制给了 hCopy
	fmt.Println("第三次使用 ... ")
	hCopy.C.checkFunc()

}

type checkCopy uintptr

func (c *checkCopy) checkFunc() {
	fmt.Println(uintptr(*c))
	fmt.Println(uintptr(unsafe.Pointer(c)))


	/**
	因为 check 的底层类型为 uintptr
	那么 这里的 *c其实就是 check 类型本身，然后强转成uintptr
	和拿着 c 也就是 check 的指针去求 uintptr，理论上要想等
	即：内存地址为一样，则表示没有被复制
	*/
	// 下述做法是：
	// 其实 check 中存储的对象地址就是 check 对象自身的地址
	// 先把 check 处存储的对象地址 和 自己通过 unsafe.Pointer求出来的对象地址作比较，
	// 如果发现不相等，那么就尝试的替换，由于使用的 old是0，
	// 则表示c还没有开辟内存空间，也就是说，只有是首次开辟地址才会替换成功
	// 如果替换不成功，则表示 check 处所存储的地址和 unsafe计算出来的不一致
	// 则表示对象是被复制了

	if uintptr(*c) != uintptr(unsafe.Pointer(c)) &&
		!atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c))) &&
		uintptr(*c) != uintptr(unsafe.Pointer(c)) {
		//panic("obj is copied")
		fmt.Println("obj is copied")
	}
}

type Human struct {
	C		checkCopy
	Name 	string
}

