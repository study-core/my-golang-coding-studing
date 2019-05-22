package main 

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"fmt"
)

//包装 类
var exits = &struct{
	sync.RWMutex
	funcs []func()
	signals chan os.Signal 
}{}

//自定义: 登记函数 (被登记的函数 将会被exit函数执行)
func atexit(f func()) {
	//先锁定
	exits.Lock()
	//释放锁
	defer exits.Unlock()
	//往队列中 登记需要被执行的函数 f 
	exits.funcs = append(exits.funcs, f)
}

func waitExit(){

	//先初始化信号处理工具实体
	if nil == exits.signals{
		//缓存接收到系统信号的 chan
		exits.signals = make(chan os.Signal)
		//声明需要接受的系统信号类型
		signal.Notify(exits.signals, syscall.SIGINT, syscall.SIGTERM)
	}

	//读锁
	exits.RLock()
	for _, f := range exits.funcs{
		//defer f()    //这里使用 defer 是因为即使 某些func 发生panic了,也能确保之后的func正常被执行
		f()
	}
	//释放读锁
	exits.RUnlock()

	//不接收到信号就一直阻塞
	<- exits.signals
}

func main() {
	
	//登记 需要被执行的func
	atexit(func(){fmt.Println("exit 1 ...")})
	atexit(func(){fmt.Println("exit 2 ...")})
	waitExit()
}