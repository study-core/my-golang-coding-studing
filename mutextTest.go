package main

import (
	"sync/atomic"
	"runtime/race"
	"unsafe"
)

func throw(string) // 被定义在 runtime 包中，src/runtime/panic.go 的 sync_throw 方法

// mutex 是一个互斥锁
// 零值是没有被上锁的互斥锁。
//
// 首次使用后，不得复制互斥锁。(意思是不能复制值，可以做成引用复制)
type Mutex struct {
	// 将一个32位整数拆分为
	// 当前阻塞的goroutine数目(29位)|饥饿状态(1位)|唤醒状态(1位)|锁状态(1位) 的形式，来简化字段设计
	state int32

	// 信号量
	sema uint32
}

// 锁接口
type Locker interface {
	Lock()
	Unlock()
}

const (
	// 定义锁的状态
	mutexLocked      = 1 << iota // 1 表示是否被锁定  0001 含义：用最后一位表示当前对象锁的状态，0-未锁住 1-已锁住
	mutexWoken                   // 2 表示是否被唤醒  0010 含义：用倒数第二位表示当前对象是否被唤醒 0- 未唤醒 1-唤醒  【注意： 未被唤醒并不是指 休眠，而是指为了让所能被设置 被唤醒的一个初始值】
	mutexStarving                // 4 表示是否饥饿   0100 含义：用倒数第三位表示当前对象是否为饥饿模式，0为正常模式，1为饥饿模式。
	mutexWaiterShift = iota      // 3 表示 从倒数第四位往前的bit位表示在排队等待的goroutine数目(共对于 32位中占用 29 位)

	//
	/** 互斥量可分为两种操作模式:正常和饥饿。

	【正常模式】，等待的goroutines按照FIFO（先进先出）顺序排队，但是goroutine被唤醒之后并不能立即得到mutex锁，它需要与新到达的goroutine争夺mutex锁。

	因为新到达的goroutine已经在CPU上运行了，所以被唤醒的goroutine很大概率是争夺mutex锁是失败的。出现这样的情况时候，被唤醒的goroutine需要排队在队列的前面。

	如果被唤醒的goroutine有超过1ms没有获取到mutex锁，那么它就会变为饥饿模式。

	在饥饿模式中，mutex锁直接从解锁的goroutine交给队列前面的goroutine。新达到的goroutine也不会去争夺mutex锁（即使没有锁，也不能去自旋），而是到等待队列尾部排队。

	【饥饿模式】，锁的所有权将从unlock的gorutine直接交给交给等待队列中的第一个。新来的goroutine将不会尝试去获得锁，即使锁看起来是unlock状态, 也不会去尝试自旋操作，而是放在等待队列的尾部。如果有一个等待的goroutine获取到mutex锁了，如果它满足下条件中的任意一个，mutex将会切换回去正常模式：

	1. 是等待队列中的最后一个goroutine

	2. 它的等待时间不超过1ms。

	正常模式：有更好的性能，因为goroutine可以连续多次获得mutex锁；

	饥饿模式：能阻止尾部延迟的现象，对于预防队列尾部goroutine一致无法获取mutex锁的问题。

	*/
	starvationThresholdNs = 1e6 // 1ms
)

// 如果锁已经在使用中，则调用goroutine 直到互斥锁可用为止。
/**
在此之前我们必须先说下 四个重要的方法；
【runtime_canSpin】，【runtime_doSpin】，【runtime_SemacquireMutex】，【runtime_Semrelease】
【runtime_canSpin】： 在 src/runtime/proc.go 中被实现 sync_runtime_canSpin； 表示 比较保守的自旋，
					golang中自旋锁并不会一直自旋下去，在runtime包中runtime_canSpin方法做了一些限制,
					传递过来的iter大等于4或者cpu核数小等于1，最大逻辑处理器大于1，至少有个本地的P队列，
					并且本地的P队列可运行G队列为空。
【runtime_doSpin】： 在src/runtime/proc.go 中被实现 sync_runtime_doSpin；表示 会调用procyield函数，
					该函数也是汇编语言实现。函数内部循环调用PAUSE指令。PAUSE指令什么都不做，
					但是会消耗CPU时间，在执行PAUSE指令时，CPU不会对它做不必要的优化。
【runtime_SemacquireMutex】：在 src/runtime/sema.go 中被实现 sync_runtime_SemacquireMutex；表示通过信号量 阻塞当前协程
【runtime_Semrelease】: 在src/runtime/sema.go 中被实现 sync_runtime_Semrelease
*/
func (m *Mutex) Lock() {
	// 如果m.state为 0，说明当前的对象还没有被锁住，进行原子性赋值操作设置为mutexLocked状态，CompareAnSwapInt32返回true
	// 否则说明对象已被其他goroutine锁住，不会进行原子赋值操作设置，CopareAndSwapInt32返回false
	/**
	如果mutext的state没有被锁，也没有等待/唤醒的goroutine, 锁处于正常状态，那么获得锁，返回.
    比如锁第一次被goroutine请求时，就是这种状态。或者锁处于空闲的时候，也是这种状态
	 */
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		if race.Enabled {
			race.Acquire(unsafe.Pointer(m))
		}
		return
	}
	/** 在 锁定没有成功的时候，才会往下面走 */

	// 首先判断是否已经加锁并处于 正常模式，
	// 将原先锁的state & (1 和 4 | 的结果，目的就是为了检验 state 是处于 1 还是 4 状态， 还是两者都是.
	// 如果与1相等，则说明此时处于 正常模式并且已经加锁，而后判断当前协程是否可以自旋。
	// 如果可以自旋，则通过右移三位判断是否还有协程正在等待这个锁，
	// 如果有，并通过 低2位 判断是否该所处于被唤醒状态，
	// 如果并没有，则将其状态量设为被唤醒的状态，之后进行自旋，直到该协程自旋数量达到上限，
	// 或者当前锁被解锁，
	// 或者当前锁已经处于 饥饿模式


	// 标记本goroutine的等待时间
	// 开始等待时间戳
	var waitStartTime int64

	// 本goroutine是否已经处于饥饿状态
	// 饥饿模式标识 true: 饥饿  false: 未饥饿
	starving := false

	// 本goroutine是否已唤醒
	// 被唤醒标识  true: 被唤醒   flase: 未被唤醒
	awoke := false

	// 自旋次数
	iter := 0

	// 保存当前对象锁状态，做对比用
	old := m.state

	// for 来实现 CAS(Compare and Swap) 非阻塞同步算法 (对比交换)
	for {
		// 不要在饥饿模式下自旋，将锁的控制权交给阻塞任务，否则无论如何 当前goroutine都无法获得互斥锁。

		// 相当于xxxx...x0xx & 0101 = 01，当前对象锁被使用
		// old & (是否锁定|是否饥饿) == 是否锁定
		// runtime_canSpin() 表示 是否可以自旋。runtime_canSpin返回true，可以自旋。即： 判断当前goroutine是否可以进入自旋锁
		/**
		第一个条件：是state已被锁，但是不是饥饿状态。如果时饥饿状态，自旋时没有用的，锁的拥有权直接交给了等待队列的第一个。
        第二个条件：是还可以自旋，多核、压力不大并且在一定次数内可以自旋， 具体的条件可以参考`sync_runtime_canSpin`的实现。
        如果满足这两个条件，不断自旋来等待锁被释放、或者进入饥饿状态、或者不能再自旋。
		 */
		if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {


			// 主动旋转是有意义的。试着设置 mutexWoken （锁唤醒）标志，告知解锁，不唤醒其他阻塞的goroutines。
			// old&mutexWoken == 0 再次确定是否被唤醒： xxxx...xx0x & 0010 = 0
			// old>>mutexWaiterShift != 0 查看是否有goroution在排队
			// tomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) 将对象锁在老状态上追加唤醒状态：xxxx...xx0x | 0010 = xxxx...xx1x

			// 如果当前标识位 awoke为 未被唤醒 && （old 也为 未被唤醒） && 有正在等待的 goroutine && 则修改 old 为 被唤醒
			// 且修改标识位 awoke 为 true 被唤醒
			/**
			自旋的过程中如果发现state还没有设置woken标识，则设置它的woken标识， 并标记自己为被唤醒。
			 */
			if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 && atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {

				// 更改标识位为 唤醒true
				awoke = true
			}

			// 否则： 进入自旋
			// 进入自旋锁后当前goroutine并不挂起，仍然在占用cpu资源，所以重试一定次数后，不会再进入自旋锁逻辑
			runtime_doSpin()

			// 累加自旋次数
			iter++

			// 更新中转变量
			// 保存mutex对象即将被设置成的状态
			old = m.state
			continue
		}


		// 以下代码是不使用**自旋**的情况

		/**
		 到了这一步， state的状态可能是：
         1. 锁还没有被释放，锁处于正常状态
         2. 锁还没有被释放， 锁处于饥饿状态
         3. 锁已经被释放， 锁处于正常状态
         4. 锁已经被释放， 锁处于饥饿状态

         并且本gorutine的 awoke可能是true, 也可能是false (其它goutine已经设置了state的woken标识)
         new 复制 state的当前状态， 用来设置新的状态
         old 是锁当前的状态
		 */

		new := old


		/** 下面的几个 if 分别是并列语句，来判断如给 设置 state 的new 状态 */

		/**
		如果old state状态不是饥饿状态, new state 设置锁， 尝试通过CAS获取锁,
        如果old state状态是饥饿状态, 则不设置new state的锁，因为饥饿状态下锁直接转给等待队列的第一个.
		 */
		// 不要试图获得饥饿goroutine的互斥锁，新来的goroutines必须排队。
		// 对象锁饥饿位被改变 为 1 ，说明处于饥饿模式
		// xxxx...x0xx & 0100 = 0xxxx...x0xx



		/**【一】如果是正常状态 （如果是正常，则可以竞争到锁） */
		if old&mutexStarving == 0 {

			// xxxx...x0xx | 0001 = xxxx...x0x1，将标识对象锁被锁住
			new |= mutexLocked
		}

		/** 【二】处于饥饿且锁被占用 状态下  */
		// xxxx...x1x1 & (0001 | 0100) => xxxx...x1x1 & 0101 != 0;当前mutex处于饥饿模式并且锁已被占用，新加入进来的goroutine放到队列后面，所以 等待者数目 +1
		if old&(mutexLocked|mutexStarving) != 0 {
			// 更新阻塞goroutine的数量,表示mutex的等待goroutine数目加1
			// 首先，如果此时还是由于别的协程的占用 无法获得锁 或者 处于 饥饿模式，都在其state加8表示有新的协程正在处于等待状态
			new += 1 << mutexWaiterShift
		}

		/**
		如果之前由于自旋而将该锁唤醒，那么此时将其低二位的状态量重置为0 (即 未被唤醒)。
		之后判断starving是否为true，如果为true说明在上一次的循环中，
		锁需要被定义为 饥饿模式，那么在这里就将相应的状态量低3位设置为1表示进入饥饿模式
		*/
		/***
		【三】
		如果当前goroutine已经处于饥饿状态 （表示当前 goroutine 的饥饿标识位 starving）， 并且old state的已被加锁,
        将new state的状态标记为饥饿状态, 将锁转变为饥饿状态.
		 */
		// 当前的goroutine将互斥锁转换为饥饿模式。但是，如果互斥锁当前没有解锁，就不要打开开关,设置mutex状态为饥饿模式。Unlock预期有饥饿的goroutine
		// old&mutexLocked != 0  xxxx...xxx1 & 0001 != 0；锁已经被占用
		// 如果 饥饿且已被锁定
		if starving && old&mutexLocked != 0 {

			// 【追加】饥饿状态
			new |= mutexStarving
		}


		/**
		【四】
		如果本goroutine已经设置为唤醒状态, 需要清除new state的唤醒标记, 因为本goroutine要么获得了锁，要么进入休眠，
        总之state的新状态不再是woken状态.
		 */
		// 如果 goroutine已经被唤醒，因此需要在两种情况下重设标志
		if awoke {
			// xxxx...xx0x & 0010 == 0,如果唤醒标志为与awoke的值不相协调就panic
			// 即 state 为 未被唤醒
			if new&mutexWoken == 0 {
				panic("sync: inconsistent mutex state")
			}

			// new & (^mutexWoken) => xxxx...xxxx & (^0010) => xxxx...xxxx & 1101 = xxxx...xx0x
			// 设置唤醒状态位0,被  未唤醒【只是为了， 下次被可被设置为i被唤醒的 初识化标识，而不是指休眠】
			new &^= mutexWoken
		}

		/**
		之后尝试通过cas将 new 的state状态量赋值给state，
		如果失败，则重新获得其 state在下一步循环重新重复上述的操作。
		如果成功，首先判断已经阻塞时间 (通过 标记本goroutine的等待时间 waitStartTime )，如果为零，则从现在开始记录
		*/

		// 将新的状态赋值给 state
		// 注意new的锁标记不一定是true, 也可能只是标记一下锁的state是饥饿状态
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			/**
			 如果old state的状态是未被锁状态，并且锁不处于饥饿状态,
             那么当前goroutine已经获取了锁的拥有权，返回
			 */
			// xxxx...x0x0 & 0101 = 0，表示可以获取对象锁 （即 还是判断之前的状态，锁不是饥饿 也不是被被锁定， 所已经可用了）
			if old&(mutexLocked|mutexStarving) == 0 {
				break // 结束cas
			}
			// 以下的操作都是为了判断是否从【饥饿模式】中恢复为【正常模式】
			// 判断处于FIFO还是LIFO模式
			// 如果等待时间不为0 那么就是 LIFO
			// 在正常模式下，等待的goroutines按照FIFO（先进先出）顺序排队

			/**
			设置/计算本goroutine的等待时间
			 */
			queueLifo := waitStartTime != 0
			if waitStartTime == 0 {
				// 更新等待时间
				waitStartTime = runtime_nanotime()
			}


			// 通过runtime_SemacquireMutex()通过信号量将当前协程阻塞
			// 函数 runtime_SemacquireMutex 定义在 sema.go
			/**
			既然未能获取到锁， 那么就使用 [sleep原语] 阻塞本goroutine
            如果是新来的goroutine,queueLifo=false, 加入到等待队列的尾部，耐心等待
            如果是唤醒的goroutine, queueLifo=true, 加入到等待队列的头部
			 */
			runtime_SemacquireMutex(&m.sema, queueLifo)

			// 当之前调用 runtime_SemacquireMutex 方法将当前新进来争夺锁的协程挂起后，如果协程被唤醒，那么就会继续下面的流程
			// 如果当前 饥饿状态标识为 饥饿 || 当前时间 - 开始等待时间 > 1 ms 则 都切换为饥饿状态标识
			/**
			使用 [sleep原语] 之后，此goroutine被唤醒
            计算当前goroutine是否已经处于饥饿状态.
			 */
			starving = starving || runtime_nanotime()-waitStartTime > starvationThresholdNs
			// 刷新下 中转变量
			/** 得到当前的锁状态 */
			old = m.state


			/**
			如果当前的state已经是饥饿状态
            那么锁应该处于Unlock状态，那么应该是锁被直接交给了本goroutine
			 */
			// xxxx...x1xx & 0100 != 0  处于 饥饿状态
			if old&mutexStarving != 0 {


				 /**
				 如果当前的state已被锁，或者已标记为唤醒， 或者等待的队列中为空,
                 那么state是一个非法状态
				  */
				// xxxx...xx11 & 0011 != 0 又可能是被锁定，又可能是被唤醒 或者 没有等待的goroutine
				if old&(mutexLocked|mutexWoken) != 0 || old>>mutexWaiterShift == 0 {
					panic("sync: inconsistent mutex state")
				}

				// delta 表示当前状态下的等待数
				// 否则下一次的循环中将该锁这是为 饥饿模式。
				// 如果已经是这个模式，那么就会将 状态量的等待数 减1
				/**
				当前goroutine用来设置锁，并将等待的goroutine数减1.
				lock状态 -一个gorotine数，表示 状态 delta == (lock + （减去一个等待goroutine数）)
				 */
				delta := int32(mutexLocked - 1<<mutexWaiterShift)
				// 并判断当前如果已经没有等待的协程，就没有必要继续维持 饥饿模式，同时也没必要继续执行该循环（当前只有一个协程在占用锁）
				/**
				如果本goroutine并不处于饥饿状态，或者它是最后一个等待者，
                那么我们需要把锁的state状态设置为正常模式.
				 */
				if !starving || old>>mutexWaiterShift == 1 {

					// 退出饥饿模式。
					// 在这里做到并考虑等待时间至关重要。
					// 饥饿模式是如此低效，一旦将互斥锁切换到饥饿模式，两个goroutine就可以无限锁定。
					delta -= mutexStarving
				}
				// 设置新state, 因为已经获得了锁，退出、返回
				atomic.AddInt32(&m.state, delta)
				break
			}
			// 修改为 本goroutine 是否被唤醒标识位
			/**
			如果当前的锁是正常模式，本goroutine被唤醒，自旋次数清零，从for循环开始处重新开始
			 */
			awoke = true
			// 自旋计数 初始化
			iter = 0
		} else {
			// 如果CAS不成功，重新获取锁的state, 从for循环开始处重新开始 继续上述动作
			old = m.state
		}
	}

	if race.Enabled {
		race.Acquire(unsafe.Pointer(m))
	}
}

// 解锁一个未被锁定的互斥锁时，是会报错
// 锁定的互斥锁与特定的goroutine无关。
// 允许一个goroutine锁定Mutex然后
// 安排另一个goroutine解锁它。
func (m *Mutex) Unlock() {
	if race.Enabled {
		_ = m.state
		race.Release(unsafe.Pointer(m))
	}
	/** 如果state不是处于锁的状态, 那么就是Unlock根本没有加锁的mutex, panic  */
	// state -1 标识解锁 (移除锁定标记)
	new := atomic.AddInt32(&m.state, -mutexLocked)

	/**
	释放了锁，还得需要通知其它等待者

	被通知的 goroutine 会去做下面的事情
    锁如果处于饥饿状态，直接交给等待队列的第一个, 唤醒它，让它去获取锁
    锁如果处于正常状态，则需要唤醒对头的 goroutine 让它和新来的goroutine去竞争锁，当然极大几率为失败，
		这时候 被唤醒的goroutine需要排队在队列的前面 (然后自旋)。如果被唤醒的goroutine有超过1ms没有获取到mutex锁，那么它就会变为饥饿模式
	 */


	// 再次校验下 标识，new state如果是正常状态, 验证锁状态是否符合
	if (new+mutexLocked)&mutexLocked == 0 {
		panic("sync: unlock of unlocked mutex")
	}
	// xxxx...x0xx & 0100 = 0 ;判断是否处于正常模式
	if new&mutexStarving == 0 {

		// 记录缓存值
		old := new
		for {

			// 如果没有等待的goroutine或goroutine不处于空闲，则无需唤醒任何人
			// 在饥饿模式下，锁的所有权直接从解锁goroutine交给下一个 正在等待的goroutine (等待队列中的第一个)。
			// 注意： old&(mutexLocked|mutexWoken|mutexStarving) 中，因为在最上面已经 -mutexLocked 并且进入了 if new&mutexStarving == 0
			// 说明目前 只有在还有goroutine 或者 被唤醒的情况下才会 old&(mutexLocked|mutexWoken|mutexStarving) != 0
			// 即：当休眠队列内的等待计数为 0  或者 是正常但是 处于被唤醒或者被锁定状态，退出
			// old&(mutexLocked|mutexWoken|mutexStarving) != 0     xxxx...x0xx & (0001 | 0010 | 0100) => xxxx...x0xx & 0111 != 0
			/**
			 如果没有等待的goroutine, 或者锁不处于空闲的状态，直接返回.
			 */
			if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
				return
			}

			// 减少等待goroutine个数，并添加 唤醒标识
			new = (old - 1<<mutexWaiterShift) | mutexWoken

			/** 设置新的state, 这里通过 信号量 去唤醒一个阻塞的goroutine去获取锁. */
			if atomic.CompareAndSwapInt32(&m.state, old, new) {

				// 释放锁,发送释放信号 (解除 阻塞信号量)
				runtime_Semrelease(&m.sema, false)
				return
			}
			// 赋值给中转变量，然后启动下一轮
			old = m.state
		}
	} else {

		/**
		饥饿模式下:
		直接将锁的拥有权传给等待队列中的第一个.
        注意:
		此时state的mutexLocked还没有加锁，唤醒的goroutine会设置它。
        在此期间，如果有新的goroutine来请求锁， 因为mutex处于饥饿状态， mutex还是被认为处于锁状态，
        新来的goroutine不会把锁抢过去.
		 */
		runtime_Semrelease(&m.sema, true)
	}
}
