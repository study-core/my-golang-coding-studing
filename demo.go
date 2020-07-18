package main

import "fmt"

func main() {
	//var reqId uint64 = 1
	//u := &User{
	//	AccountId: 1,
	//	Desctiption: "没修改之前",
	//	Roles: []string{"全部修改A", "全部修改B"},
	//}
	//UpdateUser(reqId, u)

	//a := 2
	//b := 100
	//
	//fmt.Println(a%b, a/b)



	aa := 12

	var ai interface{}
	ai = aa
	switch ai.(type) {
	case int:
		fmt.Println("我是int", ai)
	case float64, float32:
		fmt.Println("我是float", ai)
	default:
		fmt.Println("我啥都不是", ai)
	}


}

type User struct{
	AccountId uint64
	Desctiption string
	Roles []string
}

func UpdateUser(reqId uint64, user *User) {
	if reqId == user.AccountId {
		user.Desctiption = "修改描述"
	} else {
		user.Desctiption = "全部修改"
		user.Roles = []string{"全部修改A", "全部修改B"}
	}
	// todo 使用 gorm update 当前 user
}