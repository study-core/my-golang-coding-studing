package main

import (
	"fmt"
	//"encoding/json"
)

func main() {
	m := make(map[string][]string, 0)
	m["a"] = []string{"a1", "a2", "a3", "a4"}
	fmt.Println(m)
	l := m["a"]
	//l = append(l[:1], l[2:]...)
	//fmt.Println(m)
	l = append(l, "d")
	fmt.Println(m)
	l = append(l[:1], l[2:]...)
	fmt.Println(m)
}
//func main() {
//
//	//m := make(map[string][]string, 0)
//	//
//	//m["a"] = []string{"a1", "a2", "a3", "a4"}
//	//m["b"] = []string{"b1", "b2", "b3", "b4"}
//	m := make(map[string]*[]string, 0)
//
//	m["a"] = &[]string{"a1", "a2", "a3", "a4"}
//	m["b"] = &[]string{"b1", "b2", "b3", "b4"}
//	str, _ := json.Marshal(m)
//	fmt.Println(string(str))
//	//fmt.Println(m)
//
//	l := m["a"]
//	//for i, v := range l {
//	for i, v := range *l {
//		if "a2" == v {
//			//m["a"] = append((m["a"])[:i],(m["a"])[i+1:]...)
//			//l = append(l[:i], l[i+1:]...)
//			*l = append((*l)[:i], (*l)[i+1:]...)
//		}
//	}
//	//fmt.Println(l)
//	str, _ = json.Marshal(m)
//	fmt.Println(string(str))
//	//fmt.Println(m)
//	//s1 :=[]string{"a", "b", "c"}
//	//
//	//s2 := s1
//	//fmt.Println(s1, s2)
//	//for i, v := range s2 {
//	//	if "b" == v {
//	//		s2 = append(s2[:i], s2[i+1:]...)
//	//	}
//	//}
//	//fmt.Println(s1, s2)
//	//arr := []string {"Hello", "World"}
//	/*fmt.Println(m)
//	//fmt.Println(arr)
//	//handleSlice(arr)
//	handleMap(m)
//
//	fmt.Println(m)
//	//fmt.Println(arr)*/
//}



func handleSlice (arr []string) {
	//arr = append(arr, "FFFX")
	arr[0], arr[1] = arr[1], arr[0]
}

func handleMap(m map[string][]string) {
	delete(m, "a")
	m["b"] = []string{"BBBBBBBBBBBBB"}
	m["ass"] = []string{"SDDDD"}
}