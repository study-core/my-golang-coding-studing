package main
//import (
//	"sync"
//	"fmt"
//	"net"
//	"runtime"
//)
//
////sync.Pool是一个可以存或取的临时对象集合
////sync.Pool可以安全被多个线程同时使用，保证线程安全
////注意、注意、注意，sync.Pool中保存的任何项都可能随时不做通知的释放掉，所以不适合用于像socket长连接或数据库连接池。
////sync.Pool主要用途是增加临时对象的重用率，减少GC负担。
//
//func testTcpConnPool() {
//	sp2 := sync.Pool{
//		New: func() interface{} {
//			conn, err := net.Dial("tcp", ":80")
//			if err != nil {
//				return nil
//			}
//			return conn
//		},
//	}
//	buf := make([]byte, 1024)
//	//获取对象
//	conn := sp2.Get().(net.Conn)
//	//使用对象
//	conn.Write([]byte("GET / HTTP/1.1 \r\n\r\n"))
//	n, _ := conn.Read(buf)
//	fmt.Println("conn read : ", string(buf[:n]))
//	//打印conn的地址
//	fmt.Println(conn)
//	//把对象放回池中
//	sp2.Put(conn)
//	//我们人为的进行一次垃圾回收
//	runtime.GC()
//	//再次获取池中的对象
//	conn2 := sp2.Get().(net.Conn)
//	//这时发现conn2的地址与上面的conn的地址不一样了
//	//说明池中我们之前放回的对象被全部清除了，显然这并不是我们想看到的
//	//所以sync.Pool不适合用于scoket长连接或数据库连接池
//	fmt.Println(conn2)
//}
//
//func main() {
//	//我们创建一个Pool，并实现New()函数
//	sp := sync.Pool{
//		//New()函数的作用是当我们从Pool中Get()对象时，如果Pool为空，则先通过New创建一个对象，插入Pool中，然后返回对象。
//		New: func() interface{} {
//			return make([]int, 16)
//		},
//	}
//	item := sp.Get()
//	//打印可以看到，我们通过New返回的大小为16的[]int
//	fmt.Println("item : ", item)
//
//	//然后我们对item进行操作
//	//New()返回的是interface{}，我们需要通过类型断言来转换
//	for i := 0; i < len(item.([]int)); i++ {
//		item.([]int)[i] = i
//	}
//	fmt.Println("item : ", item)
//
//	//使用完后，我们把item放回池中，让对象可以重用
//	sp.Put(item)
//
//	//再次从池中获取对象
//	item2 := sp.Get()
//	//注意这里获取的对象就是上面我们放回池中的对象
//	fmt.Println("item2 : ", item2)
//	//我们再次获取对象
//	item3 := sp.Get()
//	//因为池中的对象已经没有了，所以又重新通过New()创建一个新对象，放入池中，然后返回
//	//所以item3是大小为16的空[]int
//	fmt.Println("item3 : ", item3)
//
//	//测试sync.Pool保存socket长连接池
//	//testTcpConnPool()
//}



import (
	"runtime/debug"
	"sync/atomic"
	"sync"
	"fmt"
	"runtime"
)

func main() {
	defer debug.SetGCPercent(debug.SetGCPercent(-1))

	var count int32
	newfun := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}

	pool := sync.Pool{New: newfun}

	v1 := pool.Get()
	fmt.Printf("v1 :%v\n", v1)

	pool.Put(9)
	pool.Put(10)
	pool.Put(11)
	pool.Put(12)

	v2 := pool.Get()
	fmt.Printf("v2 :%v\n", v2)

	debug.SetGCPercent(100)
	runtime.GC()

	v3 := pool.Get()
	fmt.Printf("v3 :%v\n", v3)

	pool.New = nil

	v4 := pool.Get()
	fmt.Printf("v4 :%v\n", v4)
}

