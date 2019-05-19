package main

import (
	"runtime/race"
	"sync/atomic"
	"unsafe"
)

type RWMutex struct {
	w           Mutex  // 互斥锁
	writerSem   uint32 // 写锁信号量
	readerSem   uint32 // 读锁信号量
	readerCount int32  // 读锁计数器
	readerWait  int32  // 获取写锁时需要等待的读锁释放数量
}

const rwmutexMaxReaders = 1 << 30    // 支持最多2^30个读锁

// 读锁锁定：
//
// 它不应该用于递归读锁定;
func (rw *RWMutex) RLock() {

	// 竞态检测
	if race.Enabled {
		_ = rw.w.state
		race.Disable()
	}

	// 每次goroutine获取读锁时，readerCount+1
	// 如果写锁已经被获取，那么readerCount在 - rwmutexMaxReaders与 0 之间，这时挂起获取读锁的goroutine，
	// 如果写锁没有被获取，那么readerCount>=0，获取读锁,不阻塞
	// 通过readerCount的正负判断读锁与写锁互斥,如果有写锁存在就挂起读锁的goroutine,多个读锁可以并行

	if atomic.AddInt32(&rw.readerCount, 1) < 0 {
		// 将goroutine排到G队列的后面,挂起goroutine, 监听readerSem信号量
		runtime_Semacquire(&rw.readerSem)
	}

	// 竞态检测
	if race.Enabled {
		race.Enable()
		race.Acquire(unsafe.Pointer(&rw.readerSem))
	}
}

// 释放读锁

// 读锁不会影响其他读操作
// 如果在进入RUnlock时没有锁没有被施加读锁的话，则会出现运行时错误。
func (rw *RWMutex) RUnlock() {

	// 竞态检测
	if race.Enabled {
		_ = rw.w.state
		race.ReleaseMerge(unsafe.Pointer(&rw.writerSem))
		race.Disable()
	}
	// 读锁计数器 -1
	// 有四种情况，其中后面三种都会进这个 if
	// 【一】有读锁，单没有写锁被挂起
	// 【二】有读锁，且也有写锁被挂起
	// 【三】没有读锁且没有写锁被挂起的时候， r+1 == 0
	// 【四】没有读锁但是有写锁被挂起，则 r+1 == -(1 << 30)
	if r := atomic.AddInt32(&rw.readerCount, -1); r < 0 {
		// 读锁早就被没有了，那么在此 -1 是需要抛异常的
		// 这里只有当读锁没有的时候才会出现的两种极端情况
		// 【一】没有读锁且没有写锁被挂起的时候， r+1 == 0
		// 【二】没有读锁但是有写锁被挂起，则 r+1 == -(1 << 30)
		if r+1 == 0 || r+1 == -rwmutexMaxReaders {
			race.Enable()
			throw("sync: RUnlock of unlocked RWMutex")
		}
		// 否则，就属于 有读锁，且也有写锁被挂起
		// 如果获取写锁时的goroutine被阻塞，这时需要获取读锁的goroutine全部都释放，才会被唤醒
		// 更新需要释放的 写锁的等待读锁释放数目
		// 最后一个读锁解除时，写锁的阻塞才会被解除.
		if atomic.AddInt32(&rw.readerWait, -1) == 0 {
			// 更新信号量，通知被挂起的写锁去获取锁
			runtime_Semrelease(&rw.writerSem, false)
		}
	}
	if race.Enabled {
		race.Enable()
	}
}

// 对一个已经lock的rw上锁会被阻塞
// 如果锁已经锁定以进行读取或写入，则锁定将被阻塞，直到锁定可用。
func (rw *RWMutex) Lock() {
	if race.Enabled {
		_ = rw.w.state
		race.Disable()
	}
	// 先获取一把互斥锁
	// 首先，获取互斥锁，与其他来获取写锁的goroutine 互斥
	rw.w.Lock()
	// 告诉其他来获取读锁操作的goroutine，现在有人获取了写锁
	// 减去最大的读锁数量，用0 -负数 来表示写锁已经被获取
	r := atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders) + rwmutexMaxReaders

	// 设置需要等待释放的读锁数量，如果有，则挂起获取 竞争写锁 goroutine
	if r != 0 && atomic.AddInt32(&rw.readerWait, r) != 0 {
		// 挂起，监控写锁信号量
		runtime_Semacquire(&rw.writerSem)
	}
	if race.Enabled {
		race.Enable()
		race.Acquire(unsafe.Pointer(&rw.readerSem))
		race.Acquire(unsafe.Pointer(&rw.writerSem))
	}
}

// Unlock 已经Unlock的锁会被阻塞.
// 如果在写锁时，rw没有被解锁，则会出现运行时错误。
//
// 与互斥锁一样，锁定的RWMutex与特定的goroutine无关。
// 一个goroutine可以RLock（锁定）RWMutex然后安排另一个goroutine到RUnlock（解锁）它。
func (rw *RWMutex) Unlock() {
	if race.Enabled {
		_ = rw.w.state
		race.Release(unsafe.Pointer(&rw.readerSem))
		race.Release(unsafe.Pointer(&rw.writerSem))
		race.Disable()
	}

	// 向 读锁的goroutine发出通知，现在已经没有写锁了
	// 还原加锁时减去的那一部分readerCount
	r := atomic.AddInt32(&rw.readerCount, rwmutexMaxReaders)
	// 读锁数目超过了 最大允许数
	if r >= rwmutexMaxReaders {
		race.Enable()
		throw("sync: Unlock of unlocked RWMutex")
	}
	// 唤醒获取读锁期间所有被阻塞的goroutine
	for i := 0; i < int(r); i++ {
		runtime_Semrelease(&rw.readerSem, false)
	}
	// 释放互斥锁资源
	rw.w.Unlock()
	if race.Enabled {
		race.Enable()
	}
}


// RLocker返回一个Locker接口的实现
// 通过调用rw.RLock和rw.RUnlock来锁定和解锁方法。
func (rw *RWMutex) RLocker() Locker {
	return (*rlocker)(rw)
}

type rlocker RWMutex

func (r *rlocker) Lock()   { (*RWMutex)(r).RLock() }
func (r *rlocker) Unlock() { (*RWMutex)(r).RUnlock() }

