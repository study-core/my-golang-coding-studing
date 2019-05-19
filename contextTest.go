package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type favContextKey string

func main() {
	wg := &sync.WaitGroup{}
	values := []string{"https://www.baidu.com/", "https://www.zhihu.com/"}
	/**
	首先调用context.Background()生成根节点，然后调用withCancel方法，传入根节点，
	得到新的子Context以及根节点的cancel方法（通知所有子节点结束运行），
	【注意】：该WithCancel 方法也返回了一个Context，这是一个新的子节点，
	与初始传入的根节点不是同一个实例了，但是每一个子节点里会保存从
	最初的根节点到本节点的链路信息 ，才能实现链式
	 */
	ctx, cancelFunc := context.WithCancel(context.Background())

	for _, url := range values {
		wg.Add(1)

		/**
		程序的reqURL方法接收一个url，然后通过http请求该url获得response，
		然后在当前goroutine里再启动一个子groutine把response打印出来，


		然后从ReqURL开始Context树往下衍生叶子节点（每一个链式调用新产生的ctx）,
		中间每个ctx都可以通过WithValue方式传值（实现通信），

		而每一个子goroutine都能通过Value方法从父goroutine取值，实现协程间的通信，
		每个子ctx可以调用Done方法检测是否有父节点调用cancel方法通知子节点退出运行，
		根节点的cancel调用会沿着链路通知到每一个子节点，因此实现了强并发控制
		 */
		subCtx := context.WithValue(ctx, favContextKey("url"), url)
		go reqURL(subCtx, wg)
	}

	go func() {
		time.Sleep(time.Second * 3)
		cancelFunc()
	}()

	wg.Wait()
	fmt.Println("exit main goroutine")
}

func reqURL(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	url := ctx.Value(favContextKey("url")).(string)
	for {
			select {
			case <-ctx.Done():
				fmt.Printf("stop getting url:%s\n", url)
				return
			default:
				r, err := http.Get(url)
				if r.StatusCode == http.StatusOK && err == nil {
					body, _ := ioutil.ReadAll(r.Body)
					subCtx := context.WithValue(ctx, favContextKey("resp"), fmt.Sprintf("%s%x", url, md5.Sum(body)))
					wg.Add(1)
					go showResp(subCtx, wg)
					//fmt.Println("printing", fmt.Sprintf("%s%x", url, md5.Sum(body)))
				}
				r.Body.Close()
				//启动子goroutine是为了不阻塞当前goroutine，这里在实际场景中可以去执行其他逻辑，这里为了方便直接sleep一秒
				// doSometing()
				time.Sleep(time.Second * 1)
			}
	}
}

func showResp(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop showing resp ---------------------->", ctx.Value(favContextKey("resp")))
			return
		default:
			//子goroutine里一般会处理一些IO任务，如读写数据库或者rpc调用，这里为了方便直接把数据打印
			fmt.Println("printing ", ctx.Value(favContextKey("resp")))
			time.Sleep(time.Second * 1)
		}
	}
}