package main

import (
	"unsafe"
	"runtime/race"
	"sync/atomic"
	"runtime"
)

// 池是一组可以单独保存和检索的临时对象。
//
//
// 存储在池中的任何项目都可以随时自动删除，
// 并且不会被通知。如果池在发生这种情况时保留唯一引用，
// 则该项可能会被释放。
//
// Pool可以安全地同时使用多个goroutine。
//
// 池的目的是缓存已分配但未使用的 对象 以供以后重用，
// 从而减轻对 gc 的压力。也就是说，
// 它可以轻松构建高效，线程安全的 free 列表。
// 但是，它不适用于所有 free 列表。
//
// 池的适当使用是管理一组默认共享的临时项，
// 并且可能由包的并发独立客户端重用。
// 池提供了一种在许多客户端上分摊分配开销的方法。
//
//
// 很好地使用池的一个例子是fmt包，
// 它维护一个动态大小的临时输出缓冲区存储。
// 底层存储队列 在负载下（当许多goroutine正在积极打印时）
// 进行缩放，并在静止时收缩。
//
// 另一方面，作为短期对象的一部分维护的空闲列表不适合用于池，
// 因为在该场景中开销不能很好地摊销。
// 使这些对象实现自己的空闲列表更有效。
//
//
// 首次使用后不得复制池。
//

/**
pool 的两个特点
1、在本地私有池和本地共享池均获取 obj 失败时,
	则会从其他 p 偷一个 obj 返回给调用方。

2、obj 在池中的生命周期取决于垃圾回收任务的下一次执行时间,
	并且从池中获取到的值可能是 put 进去的其中一个值，
	也可能是 new fun 处 新生成的一个值，在应用时很容易入坑。

在多个goroutine之间使用同一个pool做到高效，是因为sync.pool为每个P都分配了一个子池，
当执行一个pool的get或者put操作的时候都会先把当前的goroutine固定到某个P的子池上面，
然后再对该子池进行操作。每个子池里面有一个私有对象和共享列表对象，
私有对象是只有对应的P能够访问，因为一个P同一时间只能执行一个goroutine，
【因此对私有对象存取操作是不需要加锁的】。
共享列表是和其他P分享的，因此操作共享列表是需要加锁的。
 */
type Pool struct {
	// 不允许复制,一个结构体,有一个Lock()方法,嵌入别的结构体中,表示不允许复制
	// noCopy对象，拥有一个Lock方法，使得Cond对象在进行go vet扫描的时候，能够被检测到是否被复制
	noCopy noCopy

	/** local 和 localSize 维护一个动态 poolLocal 数组 */
	// 每个固定大小的池， 真实类型是 [P]poolLocal
	// 其实就是一个[P]poolLocal 的指针地址
	local     unsafe.Pointer
	// local 数组的大小
	// typedef uint64   uintptr
	localSize uintptr

	// New 是一个回调函数指针
	// 即：当Get 获取到目标对象为 nil 时，需要调用 此处的回调函数
	// 用于生成 新的对象
	New func() interface{}
}

// 本地池的附录
// 也就是一些包装
type poolLocalInternal struct {
	// 只能由相应的P 使用
	// 私有缓冲区
	private interface{}
	// 可以由任意的P 使用
	// 公共缓冲区
	shared  []interface{}
	// 保护共享的互斥锁
	Mutex
}

/**
【注意】
因为poolLocal中的对象可能会被其他P偷走，
private域保证这个P不会被偷光，至少保留一个对象供自己用。
否则，如果这个P只剩一个对象，被偷走了，
那么当它本身需要对象时又要从别的P偷回来，造成了不必要的开销。
 */
type poolLocal struct {
	poolLocalInternal

	/**
	cache使用中常见的一个问题是false sharing。
	当不同的线程同时读写同一cache line上不同数据时就可能发生false sharing。
	false sharing会导致多核处理器上严重的系统性能下降。
	 */
	// 字节对齐，避免 false sharing （伪共享）
	pad [128 - unsafe.Sizeof(poolLocalInternal{})%128]byte
}

// 在 runtime 包中实现 /src/runtime/stubs.go 的 sync_fastrand()
func fastrand() uint32

var poolRaceHash [128]uint64

// poolRaceAddr returns an address to use as the synchronization point
// for race detector logic. We don't use the actual pointer stored in x
// directly, for fear of conflicting with other synchronization on that address.
// Instead, we hash the pointer to get an index into poolRaceHash.
// See discussion on golang.org/cl/31589.
func poolRaceAddr(x interface{}) unsafe.Pointer {
	ptr := uintptr((*[2]unsafe.Pointer)(unsafe.Pointer(&x))[1])
	h := uint32((uint64(uint32(ptr)) * 0x85ebca6b) >> 16)
	return unsafe.Pointer(&poolRaceHash[h%uint32(len(poolRaceHash))])
}

/** 总的来说，sync.Pool的定位不是做类似连接池的东西，它的用途仅仅是增加对象重用的几率，减少gc的负担，而开销方面也不是很便宜的 */

/**
归还对象的过程：

1）固定到某个P，如果私有对象为空则放到私有对象；

2）否则加入到该P子池的共享列表中（需要加锁）。

可以看到一次put操作最少0次加锁，最多1次加锁。
 */
func (p *Pool) Put(x interface{}) {
	/**
	如果放入的值为空，直接return.
	检查当前goroutine的是否设置对象池私有值，
	如果没有则将x赋值给其私有成员，并将x设置为nil。
	如果当前goroutine私有值已经被设置，那么将该值追加到共享列表。
 	*/

	if x == nil {
		return
	}

	if race.Enabled {
		if fastrand()%4 == 0 {
			// Randomly drop x on floor.
			return
		}
		race.ReleaseMerge(poolRaceAddr(x))
		race.Disable()
	}

	// 先获得当前P绑定的 localPool
	l := p.pin()
	// 如果当前 localPool中的 私有 缓冲区为nil
	// 则将 obj的值赋给 私有缓冲区，并将 obj 引用设为 nil (入参的 obj是个interface{})
	if l.private == nil {
		l.private = x
		x = nil
	}

	// 调用方必须在完成取值后调用 runtime_procUnpin() 来取消禁用抢占
	runtime_procUnpin()
	// 如果上面添加入私有缓冲不成功，则加入公共缓冲区
	if x != nil {
		l.Lock()
		l.shared = append(l.shared, x)
		l.Unlock()
	}
	if race.Enabled {
		race.Enable()
	}
}


/**
获取对象过程是：

1）固定到某个P，尝试从私有对象获取，如果私有对象非空则返回该对象，并把私有对象置空；

2）如果私有对象是空的时候，就去当前子池的共享列表获取（需要加锁）；

3）如果当前子池的共享列表也是空的，那么就尝试去其他P的子池的共享列表偷取一个（需要加锁）,并删除其他P的共享池中的该值(p.getSlow())；

4）如果其他子池都是空的，最后就用用户指定的New函数产生一个新的对象返回。注意这个分配的值不会被放入池中。

可以看到一次get操作最少0次加锁，最大N（N等于MAXPROCS）次加锁。
 */
func (p *Pool) Get() interface{} {

	/**
	尝试从本地P对应的那个本地池中获取一个对象值, 并从本地池冲删除该值。
	如果获取失败，那么从共享池中获取, 并从共享队列中删除该值。
	如果获取失败，那么从其他P的共享池中偷一个过来，并删除共享池中的该值(p.getSlow())。
	如果仍然失败，那么直接通过New()分配一个返回值，注意这个分配的值不会被放入池中。
	New()返回用户注册的New函数的值，如果用户未注册New，那么返回nil。
	 */

	if race.Enabled {
		race.Disable()
	}

	// l : *poolLocal
	// 即：与当前P绑定的 poolLocal
	l := p.pin()
	// 这里不需要加锁，因为 p.pin() 中将 goroutine 设为了不可抢占。
	x := l.private
	l.private = nil
	// 先取消禁用抢占
	runtime_procUnpin()
	// 如果私有缓冲取不到则，从公共缓冲区尾部拿
	if x == nil {
		l.Lock()
		last := len(l.shared) - 1
		if last >= 0 {
			x = l.shared[last]
			l.shared = l.shared[:last]
		}
		l.Unlock()
		// 如果还取不到，则去其他P的公共缓冲区，偷一个
		if x == nil {
			x = p.getSlow()
		}
	}
	if race.Enabled {
		race.Enable()
		if x != nil {
			race.Acquire(poolRaceAddr(x))
		}
	}
	// 如果 还是空，则 new一个
	if x == nil && p.New != nil {
		x = p.New()
	}
	return x
}

/**
在我们获取到了 poolLocal。
就回到了我们从中取值的过程。
在取对象的过程中，我们仍然会面临：
既不能从 private 取、也不能从 shared 中取得尴尬境地。
这时候就来到了 getSlow()。
 */
 // 从其他P的共享缓冲区偷取 obj
func (p *Pool) getSlow() (x interface{}) {
	size := atomic.LoadUintptr(&p.localSize) // 获取当前 poolLocal 的大小
	local := p.local                         // 获取当前 poolLocal

	// 先固定 当前P，并取得当前的 P.id 来从其他 P 中偷值 (obj)，
	// 那么我们需要先获取到其他 P 对应的 poolLocal。
	// 假设 size 为数组的大小，local 为 p.local，那么尝试遍历其他所有 P：

	// 获取 P.id
	// 从其他 proc (poolLocal) 偷 一个对象
	pid := runtime_procPin()
	runtime_procUnpin()
	for i := 0; i < int(size); i++ {
		// 获取目标 poolLocal, 引入 pid 保证不是自身
		l := indexLocal(local, (pid+i+1)%int(size))
		// 对目标 poolLocal 加锁，用于访问 share 区域
		l.Lock()
		// steal 一个缓存对象
		last := len(l.shared) - 1
		if last >= 0 {
			x = l.shared[last]
			l.shared = l.shared[:last]
			l.Unlock()
			break
		}
		l.Unlock()
	}
	return x
}

// pin 会将当前 goroutine 订到 P 上, 禁止抢占(preemption) 并从 poolLocal 池中返回 P 对应的 poolLocal
// 调用方必须在完成取值后调用 runtime_procUnpin() 来取消禁止抢占。
func (p *Pool) pin() *poolLocal {

	/***
	pin() 首先会调用运行时实现获得当前 P 的 id，
	将 P 设置为禁止抢占。然后检查 pid 与 p.localSize 的值
	来确保从 p.local 中取值不会发生越界。如果不会发生，
	则调用 indexLocal() 完成取值。否则还需要继续调用 pinSlow() 。
	 */

	// 返回当前  P.id   PID
	// 并将 P设置为 禁止抢占
	pid := runtime_procPin()

	// 在 pinSlow 中会存储 localSize 后再存储 local，因此这里反过来读取
	// 因为我们已经禁用了抢占，这时不会发生 GC
	// 因此，我们必须观察 local 和 localSize 是否对应
	// 观察到一个全新或很大的的 local 是正常行为
	s := atomic.LoadUintptr(&p.localSize) // 获取当前 poolLocal 的大小
	l := p.local                          // 获取当前 poolLocal

	// 因为可能存在动态的 P（运行时调整 P 的个数）procresize/GOMAXPROCS
	// 如果 P.id 没有越界，则直接返回   PID
	/**
	具体的逻辑就是首先拿到当前的pid，
	然后以pid作为index找到local中的poolLocal，
	但是如果pid大于了localsize，
	说明当前线程的poollocal不存在,就会新创建一个poolLocal
	 */
	if uintptr(pid) < s {
		// 说明空间已分配好，直接返回
		return indexLocal(l, pid)
	}

	// 没有结果时，涉及全局加锁
	// 例如重新分配数组内存，添加到全局列表
	return p.pinSlow()
}


func (p *Pool) pinSlow() *poolLocal {

	/**
	因为需要对全局进行加锁，pinSlow() 会首先取消 P 的不可抢占，然后使用 allPoolsMu 进行加锁
	 */

	// 在互斥锁下重试。
	// 固定时无法锁定互斥锁。
	// 这时取消 P 的禁止抢占，因为使用 mutex 时候 P 必须可抢占
	runtime_procUnpin()

	// 加锁
	allPoolsMu.Lock()
	defer allPoolsMu.Unlock()

	// 当锁住后，再次固定 P 取其 id
	pid := runtime_procPin()

	// 并再次检查是否符合条件，因为可能中途已被其他线程调用
	// 当再次固定 P 时 poolCleanup 不会被调用
	s := p.localSize
	l := p.local

	/**
	具体的逻辑就是首先拿到当前的pid，
	然后以pid作为index找到local中的poolLocal，
	但是如果pid大于了localsize，
	说明当前线程的poollocal不存在,就会新创建一个poolLocal
	 */
	if uintptr(pid) < s {
		return indexLocal(l, pid)
	}

	// 如果数组为空，新建
	// 将其添加到 allPools，垃圾回收器从这里获取所有 Pool 实例
	if p.local == nil {
		allPools = append(allPools, p)
	}

	// 根据 P 数量创建 slice，如果 GOMAXPROCS 在 GC 间发生变化
	// 我们重新分配此数组并丢弃旧的
	size := runtime.GOMAXPROCS(0)
	local := make([]poolLocal, size)

	// 将底层数组起始指针保存到 p.local，并设置 p.localSize
	atomic.StorePointer(&p.local, unsafe.Pointer(&local[0])) // store-release
	atomic.StoreUintptr(&p.localSize, uintptr(size))         // store-release

	// 返回所需的 pollLocal
	return &local[pid]
}


// 实现缓存清理
// 当 stop the world  (STW) 来临，在 GC 之前会调用该函数
func poolCleanup() {

	// 该函数会注册到运行时 GC 阶段(前)，此时为 STW 状态，不需要加锁
	// 它必须不处理分配且不调用任何运行时函数，防御性的将一切归零，有以下两点原因:
	// 1. 防止整个 Pool 的 false retention
	// 2. 如果 GC 发生在当有 goroutine 与 l.shared 进行 Put/Get 时，它会保留整个 Pool.
	//    那么下个 GC 周期的内存消耗将会翻倍。
	// 遍历所有 Pool 实例，接触相关引用，交由 GC 进行回收
	for i, p := range allPools {

		// 解除引用
		allPools[i] = nil
		// 遍历 p.localSize 数组
		for i := 0; i < int(p.localSize); i++ {
			// 获取 poolLocal
			l := indexLocal(p.local, i)
			// 清理 private 和 shared 区域
			l.private = nil
			for j := range l.shared {
				l.shared[j] = nil
			}
			l.shared = nil
		}
		// 设置 p.local = nil 除解引用之外的数组空间
		// 同时 p.pinSlow 方法会将其重新添加到 allPool
		p.local = nil
		p.localSize = 0
	}
	// 重置 allPools，需要所有 p.pinSlow 重新添加
	allPools = []*Pool{}
}

var (
	// allPools 的锁?
	allPoolsMu Mutex
	// 所有P的 pool 队列?
	allPools   []*Pool
)

/**
pool创建的时候是不能指定大小的，
所有sync.Pool的缓存对象数量是没有限制的（只受限于内存），
因此使用sync.pool是没办法做到控制缓存对象数量的个数的。
另外sync.pool缓存对象的期限是很诡异的，
先看一下src/pkg/sync/pool.go里面的一段实现代码。
 */
func init() {
	/***
	可以看到pool包在init的时候注册了一个poolCleanup函数，
	它会清除所有的pool里面的所有缓存的对象，
	该函数注册进去之后会在每次gc之前都会调用，
	因此sync.Pool缓存的期限只是两次gc之间这段时间。
	例如我们把上面的例子改成下面这样之后，输出的结果将是0 0。

	a := p.Get().(int)
    p.Put(1)
    runtime.GC()
    b := p.Get().(int)
    fmt.Println(a, b)


	正因gc的时候会清掉缓存对象，也不用担心pool会无限增大的问题。
	 */
	runtime_registerPoolCleanup(poolCleanup)
	/**
	可以看到在init的时候注册了一个PoolCleanup函数，
	他会清除掉sync.Pool中的所有的缓存的对象，
	这个注册函数会在每次GC的时候运行，
	所以【sync.Pool中的值只在两次GC中间的时段有效】。

	通过以上的解读，我们可以看到，Get方法并不会对获取到的对象值做任何的保证，
	因为放入本地池中的值有可能会在任何时候被删除，
	但是不通知调用者。放入共享池中的值也有可能被其他的goroutine偷走。
	所以对象池比较适合用来存储一些临时切状态无关的数据，
	但是不适合用来存储数据库连接的实例，以及 net conn 等，
	因为存入对象池重的值有可能会在垃圾回收时被删除掉，
	这违反了数据库连接池建立的初衷。
	 */
}

/** 根据当前 pid 作为索引，和localPool的头指针，去load localPool */
func indexLocal(l unsafe.Pointer, i int) *poolLocal {
	/**
	在这个过程中我们可以看到在运行时调整 P 的大小的代价。
	如果此时 P 被调大，而没有对应的 poolLocal 时，
	必须在取之前创建好，从而必须依赖全局加锁，
	这对于以性能著称的池化概念是比较致命的，因此这也是 pinSlow 的由来。
	*/

	// 简单的通过 p.local 的头指针与索引来第 i 个 pooLocal
	lp := unsafe.Pointer(uintptr(l) + uintptr(i)*unsafe.Sizeof(poolLocal{}))
	return (*poolLocal)(lp)
}

// Implemented in runtime.

// 定义在 runtime 包中， src/runtime/mgc.go 的 sync_runtime_registerPoolCleanup()
func runtime_registerPoolCleanup(cleanup func())

// 定义在 runtime 包中， src/runtime/proc.go 的 sync_runtime_procPin()
func runtime_procPin() int

// 定义在 runtime包中， src/runtime/proc.go 的 sync_runtime_procUnpin()
func runtime_procUnpin()
