package main

import "sync/atomic"

// Once 是一个只执行一个动作的对象。
type Once struct {
	m    Mutex
	done uint32		// 初始值为0表示还未执行过，1表示已经执行过
}

// Once 的实现超级简单
// 用 互斥锁做线程安全控制
// 用uint32的done字段标识是否执行过
func (o *Once) Do(f func()) {
	// 每次一进来先读标识位 0 标识没有被执行过，1 标识已经被执行过
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// 施加互斥锁
	o.m.Lock()
	defer o.m.Unlock()
	// 如果之前未被执行过，则执行
	if o.done == 0 {
		// 先调用目标函数, 然后标识位更该去为 1
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

