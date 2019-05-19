package main

import (
	"fmt"
	"reflect"
)

/**
测试方法集
 */
func main() {
	m := my{}
	m.String()

	t := reflect.TypeOf(&m)
	me := t.Method(0)
	fmt.Println(me.Name, me.Type)
}


type my struct {
}



func (a *my)String(){
	fmt.Println("my")
}