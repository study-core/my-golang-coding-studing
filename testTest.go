package main

import (
	"time"
	"encoding/json"
	"fmt"
)

type Testjson struct {
	UpdateTime  time.Time  `json:"updateTime"`
}

func main() {
	//jsonStr := `{
	//	"updateTime":"2018-05-10T10:36:04+08:00"
	//}`
	jsonStr := `{
		"updateTime":"2017-12-12T11:55:09+08:00"
	}`
	var test Testjson
	if err := json.Unmarshal([]byte(jsonStr), &test); nil != err {
		fmt.Println("err:=" + err.Error())
		return
	}
	fmt.Println("time:=" + test.UpdateTime.Format("2006-01-02 15:04:05"))


	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	thatDay := yesterday.AddDate(0, 0, -30)
	fmt.Println("今天:" + today.Format("2006-01-02") + ",昨天:" + yesterday.Format("2006-01-02") + ",昨天的钱30天:" + thatDay.Format("2006-01-02") + ",相差:" + fmt.Sprint(CalcAbs(TimeSub(thatDay, yesterday))))
}

func TimeSub(t1, t2 time.Time) int {
	t1 = t1.UTC().Truncate(24 * time.Hour)
	t2 = t2.UTC().Truncate(24 * time.Hour)
	return int(t1.Sub(t2).Hours() / 24)
}

func CalcAbs(a int) (ret int) {
	ret = (a ^ a>>31) - a>>31
	return
}