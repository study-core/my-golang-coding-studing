package main

import (
    "fmt"
    "regexp"
)



func main(){
    fieldInfo := "external_pay_amount+FindAllString"
    //初始化
    reg :=regexp.MustCompile(`[-\+\*/\(\)]+`)  //匹配运算符
    // reg :=regexp.MustCompilePOSIX()(`[-\+\*/\(\)]+`) 
    // reg := regexp.MustCompile(`[\PP]+`)
    fieldInfoArr := reg.Split(fieldInfo, -1)
    fmt.Println(fieldInfoArr)

  // reg :=regexp.MustCompile(`[-\+\*/\(\)]+`)  //匹配运算符
  // submatch := reg.Split(fieldInfo, -1)
  // fmt.Println(submatch)
  // for _, arr := range submatch {
  //   fmt.Println(string(arr))
  // }
}