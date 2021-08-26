

package main

import (
	"fmt"
)

func main() {
	a := uint64(1000000000000000000)
	b := uint64(10*1024*1024*1024*1024*1024)

	if a < b {
		fmt.Println("true")
	} else {
		fmt.Println("false", a, b)
		//1000000000000000000
		//11258999068426240
	}
}

//func main() {
//
//	//m := make(map[string]struct{}, 3)
//	//
//	//m["a"] = struct{}{}
//	//
//	//fmt.Println(len(m))
//
//	type FighterTaskReadyGoReqContractCfg struct {
//		PartyId     string          `json:"party_id"`
//		DataParty struct{
//			InputFile  string      `json:"input_file"`
//			IdColumnName string    `json:"id_column_name"`
//		}  `json:"data_party"`
//		DynamicParameter   map[string]interface{}  `json:"dynamic_parameter"`
//	}
//
//	req := &FighterTaskReadyGoReqContractCfg{
//		PartyId: "p2",
//		DataParty: struct {
//			InputFile  string    `json:"input_file"`
//			IdColumnName string    `json:"id_column_name"`
//		}{
//			InputFile: "./root/kally/gavin/emma",
//			IdColumnName: "xxx",
//		},
//	}
//
//
//
//	m := make(map[string]interface{})
//
//	im := make(map[string]interface{})
//	im["epochs"] = 10
//	im["genter"] = "female"
//	m["data_party"] = im
//	m["name"] = "kally"
//
//	req.DynamicParameter = m
//
//	b, err := json.Marshal(req)
//	if nil != err {
//		fmt.Errorf("Failed to json Marshal, %s", err)
//		return
//	}
//
//	str := string(b)
//	fmt.Println(string(b))
//
//
//	var  res map[string]interface{}
//
//	if err := json.Unmarshal([]byte(str), &res); nil != err {
//		fmt.Errorf("Failed to json Unmarshal, %s", err)
//		return
//	}
//	a := res
//	_ = a
//}

//import (
//	"fmt"
//	"reflect"
//	"sync"
//)
//
//func main() {
//	testChan := make(chan interface{}, 3)
//	for  i := 0; i < 3; i++ {
//		testChan <- "i:=" + fmt.Sprint(i)
//	}
//	close(testChan)
//	la:
//	for {
//		select {
//			//case <- testChan:
//			//	fmt.Println("aaa")
//			//	break la
//		case i, ok := <- testChan:
//			if !ok {
//				//testChan = nil
//				fmt.Println("chan is empty")
//				//break
//				break la
//			}
//			fmt.Println(i)
//		}
//	}
//	fmt.Println("ss")
//}
//
//func Test(a *[]int, t *TypeMux){
//	*a =  append(*a, 4, 5, 6)
//	fmt.Printf("%p\n", &(*a))
//}
//
//
// //Deprecated: 加这个注释 就会表示 已过期
//type TypeMux struct {
//	mutex   sync.RWMutex
//	subm    map[reflect.Type][]*int
//	stopped bool
//}




/*

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func test(x int) {
	cond.L.Lock() //获取锁
	cond.Wait()//等待通知  暂时阻塞
	fmt.Println(x)
	time.Sleep(time.Second * 1)
	cond.L.Unlock()//释放锁
}
func main() {
	for i := 0; i < 40; i++ {
		go test(i)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 3)
	fmt.Println("broadcast")
	cond.Signal()   // 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	cond.Signal()// 3秒之后 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	cond.Broadcast()//3秒之后 下发广播给所有等待的goroutine
	time.Sleep(time.Second * 60)
}*/


