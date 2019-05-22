package main 

import(

	"fmt"
	"sync"
	"time"
)

func main() {
	getFilter()
}


func getFilter ()([]string){
	var wg  sync.WaitGroup

	wg.Add(280000)

	filterChan := make(chan map[string]interface{}, 280000)
	startTime := time.Now().UnixNano()
	for i := 1; i <= 280000; i ++ {
		fmt.Println("这是第" + fmt.Sprint(i) + "次循环")

		go func (a int) {
			time.Sleep(time.Second * 2)
			filterMap := make(map[string]interface{}, 0)
			filterMap["resCode"] = a 
			filterMap["filterInfo"] = "这是第" + fmt.Sprint(a) + "个 goroutine"

			filterChan <- filterMap
			wg.Done()
			return

		}(i)
	}
	

	wg.Wait()
	close(filterChan)
	endTime := time.Now().UnixNano()
	
	count := 0 
	countArr := make([]string, 0)

	timeLength := (endTime - startTime)/1000000000
	for{
		select {
		case x, ok := <- filterChan:
			if !ok {
				fmt.Println("最终打印了:" + fmt.Sprint(count) + "个,总共消耗了:" + fmt.Sprint(timeLength) + "s")
				// filterChan = nil
				return countArr
			}
			if i, isOk := x["resCode"]; isOk {
				fmt.Println("取出" + fmt.Sprint(i.(int)) + ",的filter:=" + x["filterInfo"].(string))
				countArr = append(countArr, fmt.Sprint(i.(int)))
				count ++
			}

		}
	}
}