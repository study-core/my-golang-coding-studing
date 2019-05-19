package main

import "fmt"

func main() {
	//var a interface{} = nil
	//Show(a)

	num := 15

	fmt.Println(num & 1)


}

func Show(n interface{}){
	if nil != n {
		fmt.Println("No nil")
		return
	}
	fmt.Println("nil")
}
