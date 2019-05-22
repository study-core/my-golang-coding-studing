package main 

import (
	"fmt"
	"strings"
	_ "errors"
	"encoding/json"
)


func main() {
	str := ""

	strArr := strings.Split(str, ",")

	bodyByte, err := json.Marshal(strArr)

	if nil != err {
		fmt.Println("转换错误~~")
	}

	fmt.Println("最终转换得到:", string(bodyByte), "len:=", len(strArr))
}


