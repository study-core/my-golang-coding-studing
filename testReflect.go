package main

import (
	"reflect"
	//"fmt"
	//"time"
	//"log"
	"sync"
	"errors"
	"fmt"
)

//func main(){
//	var n int = 1
//	var chs = make([]chan int, n)
//
//	var worker = func(n int, c chan int){
//		for i:= 0; i < n; i++ {
//			c <- i
//		}
//		close(c)
//	}
//
//	//不定数量的channel数组
//	for i := 0; i < n; i++ {
//		chs[i] = make(chan int)
//		go worker(3 + i, chs[i])
//	}
//
//
//	var selectCase = make([]reflect.SelectCase, n)
//	//将channel绑定到SelectCase
//	for i := 0; i < n; i++ {
//		selectCase[i].Dir = reflect.SelectRecv //设置信道是接收,可以为下面值之一
//		/*
//		const (
//			SelectSend    SelectDir // case Chan <- Send
//			SelectRecv              // case <-Chan:
//			SelectDefault           // default
//		)
//		*/
//		selectCase[i].Chan = reflect.ValueOf(chs[i])
//	}
//
//	numDone := 0
//	//从所有channel中取出最先到达的N个值
//	for numDone < n {
//		chosen, recv, recvOk := reflect.Select(selectCase)
//		if recvOk {
//			fmt.Println(chosen, recv.Int(), recvOk)
//			numDone++
//		}else{
//			fmt.Println("recv error")
//		}
//	}
//}









//func main() {
//	c1 := make(chan int)
//	c2 := make(chan string)
//	c3 := make(chan bool)
//
//	go func() {
//		for i := 0; i < 5; i++ {
//			time.Sleep(50 * time.Millisecond)
//			c1 <- i
//			log.Println("c2 <- ", <-c2)
//			c3 <- true
//		}
//	}()
//
//
//	// 收集 chan 的反射类型
//	cases := []reflect.SelectCase{
//		//  对 c1 做接收
//		{
//			Dir:  reflect.SelectRecv,
//			Chan: reflect.ValueOf(c1),
//		},
//
//		// 对 c2 做发送
//		{
//			Dir:  reflect.SelectSend,
//			Chan: reflect.ValueOf(c2),
//			Send: reflect.ValueOf("Hello"),
//		},
//
//		// 对 c3 做接收
//		{
//			Dir:  reflect.SelectRecv,
//			Chan: reflect.ValueOf(c3),
//		},
//
//		//
//		{
//			Dir: reflect.SelectDefault,
//		},
//	}
//	for i := 0; i < 10; i++ {
//		time.Sleep(1 * time.Second)
//		switch index, value, recvOK := reflect.Select(cases); index {
//		case 0:
//			// c1 was selected, recv is the value received
//			// we get recvOK whether we need it or not
//			log.Println("c1接收到的:", value, recvOK)
//		case 1:
//			// c2 was selected, therefore "hello" was sent
//			// recv and recvOK are garbage
//			log.Println("c2有值可以接收", value, recvOK)
//		case 2:
//			// c3 was selected, value is good if recvOK == true
//			// recvOK is false if c3 is closed
//			log.Println("c3接收到的:", value, recvOK)
//		case 3:
//			// default case
//			// recv and recvOK are useless here
//			log.Println("reflect.SelectDefault:", value, recvOK)
//		}
//	}
//}

func main() {

	ch1 := make(chan int, 0)
	ch2 := make(chan int, 0)
	ch3 := make(chan int, 0)

	go func() {
		fmt.Println("ch1收到:", <- ch1)
	}()
	go func() {
		fmt.Println("ch2收到:", <- ch1)
	}()
	go func() {
		fmt.Println("ch3收到:", <- ch1)
	}()
	e := &Event{}
	/*sub1 :=*/ e.Subscribe(ch1)
	/*sub2 :=*/ e.Subscribe(ch2)
	/*sub3 :=*/ e.Subscribe(ch3)

	e.Send(15)



	/*sub1.Unsubscribe()
	sub2.Unsubscribe()
	sub3.Unsubscribe()*/


}

const firstSubSendCase = 1 // 记录 sendCases 中的起始可用下标

var (
	chanSendErr = errors.New("Subscribe argument must have sendable channel type")
)

type eventTypeError struct {
	got, want reflect.Type
	op        string
}

type caseList []reflect.SelectCase    // send chan的反射实例集

type Event struct {
	once      	sync.Once        	// 确保只会被初始化一次
	sendLock  	chan struct{}    	// 用来对 sendCases 做锁
	removeSub 	chan interface{} 	// 用来加收需要删除 sendCases 队列中的chan
	sendCases 	caseList         	// 保存所有监听的 send chan


	mu     		sync.Mutex			// Event 对象互斥锁
	inbox  		caseList			// 用来缓存新注册进来的 send chan
	etype  		reflect.Type		// 记录当前Event实例的 允许注册的sned chan 的 Element type
}

type EventSub struct {
	e    		*Event			// 因为需要操作Event的方法，故注入
	channel 	reflect.Value	// 封装注入监听的那个chan
	errOnce 	sync.Once		// 用来保证关闭chan 只会一次操作
}

func (e *Event)Subscribe (channel interface{}) *EventSub {
	// 初始化event 实例中的相关字段，仅一次
	e.once.Do(e.init)
	cvalue := reflect.ValueOf(channel)
	ctype := reflect.TypeOf(channel)

	// 如果入参的类型不是chan 或者 chan的方向不是 send
	if ctype.Kind() != reflect.Chan || ctype.ChanDir()&reflect.SendDir == 0 {
		panic(chanSendErr)
	}
	// 检查通道中元素的类型
	if !e.checktype(ctype.Elem()) {
		panic(eventTypeError{op: "Subscribe", got: ctype, want: reflect.ChanOf(reflect.SendDir, e.etype)})
	}
	// 组装订阅信息
	sub := &EventSub{e: e, channel: cvalue}

	e.mu.Lock()
	defer e.mu.Unlock()
	// 把新注册进来的  send chan 追加到 inbox 中
	cas := reflect.SelectCase{Dir: reflect.SelectSend, Chan: cvalue}
	e.inbox = append(e.inbox, cas)
	return sub
}

// 对所有的注册 send chan 进行广播
func (e *Event) Send(val interface{}) int {
	// 初始化event 实例中的相关字段，仅一次
	e.once.Do(e.init)
	// 获取值的反射类型
	rval := reflect.ValueOf(val)
	// 上锁
	<-e.sendLock
	// 在 sendLock 之后操作 inbox
	e.mu.Lock() // 加锁是双重保险？
	e.sendCases = append(e.sendCases, e.inbox...)
	e.inbox = nil
	if !e.checktype(rval.Type()) {
		e.sendLock <- struct{}{}
		panic(eventTypeError{op: "Subscribe", got: rval.Type(), want: reflect.ChanOf(reflect.SendDir, e.etype)})
	}
	e.mu.Unlock()
	// 把需要广播的 值，逐个发给订阅了该时间的chan
	for i := firstSubSendCase; i < len(e.sendCases); i++ {
		e.sendCases[i].Send = rval  // 设置需要 发送给 send chan 的值
	}

	// 赋值给中间 cache 做操作
	cases := e.sendCases
	// 记录成功发送了多少个 send chan 的计数器
	var nsent int
	for {

		// 先逐个的 尝试性 发送一下 (因为每个 send chan 都是 阻塞的，所以send不一定会直接成功，所以用了 TrySend)
		for i := firstSubSendCase; i < len(cases); i++ {
			// 如果尝试性的发送，且成功发送出去了，则需要将 该send chan 移至 队列的尾部，并把尾部的send chan置换到该位置
			if cases[i].Chan.TrySend(rval) {
				// 累加计数
				nsent++
				// 将当前 send chan 和队列中的尾部 send chan 对换位置
				cases = cases.deactivate(i)
				// 将 遍历索引回退一个，目的是为了下一次遍历又从当前索引开始, 原因是当前索引的send chan 已经是置换后的尾部的send chan了，故需要对之 TrySend
				i--
			}
		}
		// 如果 缓存的 send chan 队列中的chan 只剩一个时(这一个是init队列时，置入的 removeSub)
		// 则，表示send chan 都被处理完了，结束循环
		if len(cases) == firstSubSendCase {
			break
		}

		/** 为毛这里老是会提示死锁啊 */
		// 这里是真正的调用可用的 chan 去做操作
		chosen, recv, _ := reflect.Select(cases)
		// 如果是第0下标的 chan 发生了处理，则表示 removeSub 中有值被处理
		// 即：接收到了 删除 send chan 中的某个chan的处理
		if chosen == 0 /* <-f.removeSub */ {
			// 找到该 send chan 在send chan队列中的索引
			index := e.sendCases.find(recv.Interface())
			e.sendCases = e.sendCases.delete(index)
			if index >= 0 && index < len(cases) {
				// 将新的 send chan 队列赋值给 缓存队列 (原因， 操作TrySend 或者 reflect.Select 都是对 缓存队上做处理的)
				cases = e.sendCases[:len(cases)-1]
			}
		} else {
			/** 和TrySend中的处理一样 */
			// 否则，置换位置
			cases = cases.deactivate(chosen)
			// 累加计数
			nsent++
		}
	}

	// 在完全处理完队列汇总的send chan 后，需要把chan中需要发送的value 清空(因为已经发出去了)
	for i := firstSubSendCase; i < len(e.sendCases); i++ {
		e.sendCases[i].Send = reflect.Value{}
	}
	// 释放 锁
	e.sendLock <- struct{}{}
	return nsent
}


// 检查注册的send chan element type
func (e *Event)checktype(typ reflect.Type) bool {
	if nil == e.etype {
		e.etype = typ
		return true
	}
	return e.etype == typ
}


func (e *Event) init() {
	e.removeSub = make(chan interface{})
	e.sendLock = make(chan struct{}, 1)
	e.sendLock <- struct{}{}
	// 在初始化的时候，事先放置 用来加收需要删除 sendCases 队列中的chan
	// 所以在 sendCases 队列中的chan 只能从下标为 1 处开始启用
	e.sendCases = caseList{{Chan: reflect.ValueOf(e.removeSub), Dir: reflect.SelectRecv}}
}

// 移除对应的 send chan
func (e *Event) remove(sub *EventSub) {
	// 获取需要删除的 chan 引用
	ch := sub.channel.Interface()
	e.mu.Lock()
	// 在inbox队列中找到这个引用
	index := e.inbox.find(ch)
	if index != -1 {
		// 找到直接删除，return
		e.inbox = e.inbox.delete(index)
		e.mu.Unlock()
		return
	}
	e.mu.Unlock()

	select {
	// inbox 找不到时，那么就是已经被转移到 send chan 队列中了, 需要由 removeSub 接收，在Send 方法中删除
	case e.removeSub <- ch:
	case <-e.sendLock:
		// 或者直接去 send chan 队列中找到，直接删除
		e.sendCases = e.sendCases.delete(e.sendCases.find(ch))
		e.sendLock <- struct{}{}
	}
}

// 取消订阅
func (sub *EventSub) Unsubscribe() {
	sub.errOnce.Do(func() {
		sub.e.remove(sub)
	})
}


// 置换队列中两个send chan的位置
func (cs caseList) deactivate(index int) caseList {
	last := len(cs) - 1
	cs[index], cs[last] = cs[last], cs[index]
	return cs[:last]
}

// 根据入参的 chan 找到其在队列中的索引
func (cs caseList) find(channel interface{}) int {
	for i, cas := range cs {
		if cas.Chan.Interface() == channel {
			return i
		}
	}
	return -1
}

// 根据索引删除 队列中对应的send chan
func (cs caseList) delete(index int) caseList {
	return append(cs[:index], cs[index+1:]...)
}