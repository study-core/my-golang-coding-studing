package main 

import (
	"fmt"
	"strings"
)


func main() {
	
	// inp := []string{}

	// if takes(inp,0, "a") == "add" {
	//    fmt.Println("you typed add")
	// } else if takes(inp,0, "b") == "" {
	//    fmt.Println("you didn't type add")
	// }else {
	// 	fmt.Println("......" + takes(inp,0, "c"))
	// }


	// firstSlice :=  []string{}       

	// secondsSlice := make([]string, 0)

 // 	 fmt.Println("firstSlice:=" + fmt.Sprint(firstSlice == nil) + 
 // 	 	",secondsSlice:=" + fmt.Sprint(secondsSlice != nil) + 
 // 	 	",1的长度:" + fmt.Sprint(len(firstSlice)) + ",2的长度:" + fmt.Sprint(len(secondsSlice)))

	 // a := []int{1,2,3,4}
  // b := a
  // a[1] = 5
  // fmt.Println(b[1])

	// a := [...]int{1,2,3,4}
 //  b := a
 //  a[1] = 5
 //  fmt.Println(b[1])


	// a := [4]int{1,2,3,4}
 //  b := 
 //  a[1] = 5
 //  fmt.Println(b[1])

// 	var a = [...]int{1: 5, 6: 8}
// 	fmt.Println(a[1])



	// const a uint64 = 4

	// fmt.Println(&a)

	// s := "sss真的哦"
	// args := []rune(s)
	
	// fmt.Println("入参解析成json, args:=" + fmt.Sprint(args))

	// var a []int

	// b := []int{}

	// fmt.Printf("a := %+v ,b := %+v ", a, b)
	// fmt.Println("a:=" + fmt.Sprint(a) + ",b:=" + fmt.Sprint(b))

	// var a, b = uint64(1), int32(4)
	// fmt.Println("a:=" + fmt.Sprint(a) + ",b:=" + fmt.Sprint(b))

	// type A struct{
	// 	Name string
	// 	Age  int
	// }
	// a := A{
	// 	Name: "xx",
	// 	Age: 1,
	// }

	// for k, v := range a {
	// 	fmt.Println("当前的k=" + k + ",v=" + fmt.Sprint(v))
	// }

	

	

	// a := N{}

	// b := &N{}

	// b.ToString()

	// a.Test()

	// const(
	// 	X uint32 = 4
	// 	Y uint32 = 5
	// 	Z uint32 = 2
	// 	A = uint32(iota)
	// 	B 
	// 	C 
	// 	)

	// fmt.Println("A:=" + fmt.Sprint(A) + ",B:=" + fmt.Sprint(B))
	userId := uint64(240000004)
	rs := []rune(fmt.Sprint(userId))
	length := len(rs)

	var start, end int
	//当大于10的Id只取后10位
	if length > 10 {
		start = length - 9
		end = length
	}else{  //等于 或小于10的取完
		start = 1
		end = length
	}

	//tirmUserId 经过截取后的userId
	tirmUserId := string(rs[(start - int(1)):end])
	tirmLen := len([]rune(tirmUserId))
	//往userId前拼接N个"0"使之达到10的长度
	newPwd := strings.TrimSpace(strings.Repeat("0",(int(10) - tirmLen)) + tirmUserId)
	fmt.Println(newPwd)
}



func (*N) Test() {
		fmt.Println("I am Test")
}

func (N) ToString() {
		fmt.Println("I am ToString")
}

type N struct{}


func takes(s []string, i int, str string) string {
	    defer func(str string) {
	        if err := recover(); err != nil {
	        	fmt.Println( str + " err :=" + fmt.Sprint(err))
	           return 
	        }
	    }(str)
	    return s[i]
}



func SubString(str string, size int) string {
	rs := []rune(str)
	length := len(rs)

	var start, end int
	if length > 10 {
		start = length - 9
		end = length 
	}else if length <= 0 {
		return ""
	}else{
		start = 1
		end = length
	}
		
	return strings.TrimSpace(string(rs[(start - int(1)):end]))
}
