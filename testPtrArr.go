package main

import "fmt"

type Student struct {
	Name 	string
}

func main() {

	arr := []*Student{
		&Student{
			Name: "小明",
		},

		&Student{
			Name: 	"小紅",
		},
	}
	for i, st := range arr {
		n := (*st).Name
		(*st).Name = fmt.Sprint(i) + "_" + n
	}
	fmt.Println(arr)
	for _, st := range arr {
		fmt.Println(st.Name)
	}
}
