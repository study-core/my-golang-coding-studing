package main

import (
	"time"
	"strconv"
	"math/rand"
	"fmt"
	"strings"
	"encoding/hex"
	"crypto/md5"
	// "github.com/hugozhu/godingtalk"
	// "encoding/json"
	"github.com/shopspring/decimal"
	// "sort"
	"sync"
	// "encoding/json"
	// "net/url"
	// "errors"
)

func main() {
	// for i:=0; i< 10; i++ {
	// 	fmt.Println("time: ", time.Now().UnixNano())
	// 	fmt.Println(getAOrderId())
	// }
	/*showMap := map[string]string{}
	showMap["a"] = "aa"
	showMap["b"] = "bb"
	showMap["c"] = "cc"
	showMap["d"] = "dd"
	showLen := len(showMap)
	
	fmt.Println("所得长度为:" + fmt.Sprint(showLen))*/

	// randdomStr := RandString(int(6))
	// fmt.Println("随机生成的字符串为:" + randdomStr)

	// str := "15813846803"
	// subStr := SubString(str, int(8), int(11))
	// fmt.Println("被截取后的字符串为:" + subStr)

	// var aa uint32 
	// aa = uint32(-1)
	// fmt.Println("复制:" + fmt.Sprint(aa))

	// Md5("654321")

	//切片
	// s := [...]uint32{1, 2, 56}
	// p := &s
	// p0 := &s[0]
	// // fmt.Println("p:=" + fmt.Sprint(p), "p0:=" + fmt.Sprint(p0))
	// println(p, p0)

	// //数组
	// a := [...]uint32{1, 2}

	// println(&a, &a[0])



	// periodTime := uint32(10)
	// //获取当前时间 (时间戳: s 秒)
	// now := time.Now().Unix()

	// periodSecond :=  int64(periodTime) * 60
	// tm := time.Unix(now - periodSecond, 0)

	// then := tm.Format("2006-01-02 15:04:05")
	// fmt.Println("当前时间为:" + time.Unix(now, 0).Format("2006-01-02 15:04:05") + ",往前推:" + fmt.Sprint(periodTime) + 
	// 	"分钟后时间为:" + then)



	// a := "1"

	// b := "2"

	// fmt.Println("交换前a:=" + fmt.Sprint(a) + ",b:=" + fmt.Sprint(b))

	// // a ^= b
	// // b ^= a
	// // a ^= b
	// c := a

	// a = b

	// b = c

	// fmt.Println("交换后a:=" + fmt.Sprint(a) + ",b:=" + fmt.Sprint(b))


	// a := "2017-11-02 14:00:12"

	// b := "2017-14-02 12:00"

	// c := "12:08"

	// d := "11:11:14"

	// fmt.Println(a < b) 

	// fmt.Println(b < c)

	// fmt.Println(b > d)

	// fmt.Println(c > d)

	// arr := []int{1, 4, 5}

	// arrLL := append([]int{}, arr[1:]...)
	// arr = append(arr[:1], 6)
	// arr = append(arr, arrLL...)
	// fmt.Println(fmt.Sprint(arr))



	
		// for i := 0; i < 3; i ++ {
		// 	select{
		// 	//5秒之后
		// 	case <- time.After(time.Second * 5):
		// 		fmt.Println("timeout 5 second...现在是第" + fmt.Sprint(i + 1) + "次,当前年时间为:" + time.Now().Format("2006-01-02 15:04:05"))
		// 	}
		// }

	// str := "您好，您设置的单图“singleView”在时间：time，针对“dimensionName”的指标“quotaName”发生了“alarmRule”，请及时查看！"
	// now := time.Now().Format("01月02日 15时04分")
	
	// str2 := strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(str,
	// 	"singleView", "城市的卖座卡销量", -1),
	// 	"time", now, -1), "dimensionName", "城市", -1),
	// 	"quotaName", "卖座卡销量", -1), "alarmRule", "卖座卡销量>1000", -1)
	// fmt.Println("替换前:" + str + "\n替换后:" + str2)



	/** 测试钉钉 报警 */
	// c := godingtalk.NewDingTalkClient("ding67079df27e72cec9", "nFBY4HjVdZ-61m4NVqwdpK5nWirMHeKiG57dUfR3sC6fPhIw_ANWJjYbsQBbIya1")
 //    c.RefreshAccessToken()


		// tableId, _ := strconv.Atoi("")
		// fmt.Println("转化为：" + fmt.Sprint(tableId))


	// jsonStr := `[{"id":178,"name":"影片名称去重数量","tableId":12227,"mode":"","value":"","chartType":1}]`

	// jsonStr2 := `[{"id":178,"name":"影片名称去重数量","tableId":"44444","mode":"","value":"","chartType":1}]`

	// singleQuotaFrontList := []*SingleQuotas{}
	// err := json.Unmarshal([]byte(jsonStr), &singleQuotaFrontList)
	// if nil != err {
	// 	fmt.Println("数组1有误err:=" + err.Error())
	// }

	// singleQuotaFrontList2 := []*SingleQuotas{}
	// err = json.Unmarshal([]byte(jsonStr2), &singleQuotaFrontList2)
	// if nil != err {
	// 	fmt.Println("数组2有误err:=" + err.Error())
	// }

	// fmt.Println("最终序列化为:集合1:" + fmt.Sprint(singleQuotaFrontList) + ",集合2:" + fmt.Sprint(singleQuotaFrontList2))



	// quota := &SingleQuotas{
	// 	Id:		 	1,
	// 	Name:		"拉卡",
	// 	TableId:	10,
	// 	Mode:		"cc",
	// 	Value:		"XX",
	// 	IsMoney:	1,
	// 	ShowType:	0,
	// 	ChartType:	2,
	// }
	// // list := []SingleQuotas{}
	// // list = append(list, *quota)
	// var aa interface{}   = *quota
	// fmt.Println("反序列为:" + fmt.Sprint(aa))

	// ss := []int{14, 45, 78}
	// fmt.Println("最后一个为:" + fmt.Sprint(ss[len(ss) - 1]))



	// fmt.Println("等于:" + fmt.Sprint(float64(2.45)/float64(2)))
	
// 	averageValue := decimal.NewFromFloat(3).Sub(decimal.NewFromFloat(float64(2)))
// 	percentValue := averageValue.Div(decimal.NewFromFloat(1))	
// 	percentValueABS := percentValue.Abs().Add(decimal.NewFromFloat(0.005))
// 	percentValueABS2, _  := percentValueABS.Float64()
// 	percentValueABS = percentValueABS.Round(2)
// 	floatValue, _  := percentValueABS.Float64()
// 	//     -1 if d <  d2
// //      0 if d == d2
// //     +1 if d >  d2
// 	// averageValue := float64(0.451) - float64(0.450)
// 	fmt.Println(fmt.Sprint(averageValue) + "   " + fmt.Sprint(percentValue) + " " + fmt.Sprint(percentValueABS) + "  "+ fmt.Sprint(percentValueABS2)  + " " + fmt.Sprint(floatValue))


	// strArr := []string{"11.01", "1", "11.1", "12"}
	// 		sort.Strings(strArr)
	// 		fmt.Println("升序返回的数组为:" + fmt.Sprint(strArr))
 // 	sort.Sort(sort.Reverse(sort.StringSlice(strArr)))
 // 	fmt.Println("降序返回的数组222为:" + fmt.Sprint(strArr))
		// map1 := make(map[string]interface{}, 0)
		// map2 := make(map[string]int, 0)
		// map1["Թ"] = 14
		// map2["Թ"] = 14
		// map1["׷"] = 45
		// map2["׷"] = 45
		// map1[""] = 6
		// map2[""] = 6
		// fmt.Println("map1:=" + fmt.Sprint(map1) + ",map2:=" + fmt.Sprint(map2))

 	// fmt.Println("空数组:" + len(nil))

	// /*filterArr :=*/ getFilter()
	// strByte, _ := json.Marshal(filterArr)
	
	// fmt.Println("最终的数组是:" + string(strByte))

	// tmp := make([]string, 0)

	// arr := []string{"12", "45"}
	// tmp = append(tmp, "58")
	// arr = append(tmp, arr...)

	// fmt.Println("新数组:" + fmt.Sprint(arr))
	// var a int
	// arr := []string{"12", "45", "87"}
	// for i, name := range arr {
	// 	if name == "45" {
	// 		// tmp := arr[0:i]
	// 		arr = append(arr[:i], arr[i + 1 :]...)
	// 		a = i
	// 	}
	// }
	// fmt.Println("新数组:" + fmt.Sprint(arr) + ",索引为:" + fmt.Sprint(a) + ",根据新索引拿到新值:" + fmt.Sprint(arr[a]))

	
	// negativeArr := make([]float64, 0)
	// negative1, _ := decimal.NewFromFloat(-2).Div(decimal.NewFromFloat(10)).Float64()
	// negativeArr = append(negativeArr, negative1)
	// negativeArr = append(negativeArr, -10)
	// negativeArr = append(negativeArr, -45)

	// positiveArr := make([]float64, 0)
	// positive1, _ := decimal.NewFromFloat(8).Div(decimal.NewFromFloat(10)).Float64()
	// positiveArr = append(positiveArr, positive1)
	// positiveArr = append(positiveArr, 78)
	// positiveArr = append(positiveArr, 48)

	// fmt.Println("排序前：正数:" + fmt.Sprint(positiveArr) + ",负数:" + fmt.Sprint(negativeArr))

	// sort.Float64s(negativeArr)
	// sort.Float64s(positiveArr)
	// fmt.Println("排序后：正数:" + fmt.Sprint(positiveArr) + ",负数:" + fmt.Sprint(negativeArr))



	// arr1 := []string{"s1|s2|s3", "w1|w2|w3", "a1|a2|a3"}

	// var str, valStr string

	
	// i := 0
	// for _, arrStr := range arr1 {
	// 	arrStrArr := strings.Split(arrStr, "|")
	// 	for _, val := range arrStrArr {
	// 		i ++ 
	// 		if val == "w2" {
	// 			str = val
	// 			goto index
	// 		}
	// 	}
	// } 
	// index:
	// if str != "" {
	// 	valStr = str
	// }
	// fmt.Println("最终的val是:" + valStr + ",拿到的是:" + str +",执行了:" + fmt.Sprint(i) + "次循环")

	// var inter interface{} = 201702012359591918.00
	// strDecimal, _ := decimal.NewFromString(fmt.Sprint(inter))
	// fmt.Println("转之前:" + fmt.Sprint(inter) + ",最后输出:" + strDecimal.String())

	// urlStr := url.QueryEscape("https://glaucus.maizuo.com/#/singleview/369")
	// fmt.Println("编码完之后的url:" + urlStr)

	// fmt.Println(errors.New("result is empty") == errors.New("result is empty"))

	fmt.Println(decimal.NewFromFloat(112.7).Round(0).String())

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


type SingleQuotas struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	TableId  int    `json:"tableId"`  //单个值
	Mode     string `json:"mode"`     //公式
	Value    string `json:"value"`    //值
	IsMoney  int    `json:"isMoney"`  //指标是否含有金额，0否，1是
	ShowType int    `json:"showType"` //指标值的展示形式（0：数字；1：百分比）

	ChartType int 	`json:"chartType"` // 当前指标所需要渲染的图类型 1:折线图，2:柱状图(如果chartType 不等于9 混合图时,无需理会该字段)
}

//生成一个订单号
func getAOrderId() string {
	orderId := time.Now().Format("20060102150405")
	rand.Seed(time.Now().UnixNano())
	orderId = orderId + strconv.Itoa(rand.Intn(90000) + 10000)
	fmt.Println("1: ", rand.Intn(90000))
	fmt.Println("2: ", rand.Intn(90000))
	return orderId
}


// func GetRandomString(len int) string{
//    str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//    bytes := []byte(str)
//    result := []byte{}
//    r := rand.New(rand.NewSource(time.Now().UnixNano()))
//    for i := 0; i < len; i++ {
//       result = append(result, bytes[r.Intn(len(bytes))])
//    }
//    return string(result)
// }


func RandString(length int) string {
    rand.Seed(time.Now().UnixNano())
    rs := make([]string, length)
    for start := 0; start < length; start++ {
        t := rand.Intn(3)
        if t == 0 {
            rs = append(rs, strconv.Itoa(rand.Intn(10)))
        } else if t == 1 {
            rs = append(rs, string(rand.Intn(26)+65))
        } else {
            rs = append(rs, string(rand.Intn(26)+97))
        }
    }
    return strings.Join(rs, "")
}



func SubString(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[(start - int(1)):end])
}


func Md5(data string) string {
	hash := md5.New()          //初始化一个MD5对象
	hash.Write([]byte(data))   // 需要加密的字符串
	cipherStr := hash.Sum(nil) //计算出校验和
	fmt.Println("md5前的信息: ", data)
	fmt.Println("md5后的信息: ", hex.EncodeToString(cipherStr))
	return hex.EncodeToString(cipherStr) // 输出加密结果
}
