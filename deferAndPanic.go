package main

import "fmt"

//func main() {
//	defer func() {
//		if err := recover();nil != err {
//			fmt.Println("A + ", fmt.Sprint(err))
//		}
//	}()
//	defer func() {
//		if err := recover();nil != err {
//			fmt.Println("B + ", fmt.Sprint(err))
//		}
//	}()
//	defer func() {
//		if err := recover();nil != err {
//			fmt.Println("C + ", fmt.Sprint(err))
//		}
//	}()
//	defer_call()
//
//	}
//
//func defer_call() {
//	defer func() { fmt.Println("打印前") }()
//	defer func() { fmt.Println("打印中") }()
//	defer func() { fmt.Println("打印后") }()
//	panic("触发异常")
//}

//func main() {
//	println(DeferFunc1(1))
//	println(DeferFunc2(1))
//	println(DeferFunc3(1))
//}
//func DeferFunc1(i int) (t int) {
//	t = i
//	defer func() { t += 3 }()
//	return t }
//
//func DeferFunc2(i int) int     {
//	t := i
//	defer func() { t += 3 }()
//	return t }
//
//func DeferFunc3(i int) (t int) {
//	defer func() { t += i }()
//	return 2
//}



//func main() {
//	list := new([]int)
//	list = append(list, 1)
//	fmt.Println(list)
//}

//func Foo(x interface{}) {
//	if x == nil {
//		fmt.Println("empty interface")
//		return
//	}
//	fmt.Println("non-empty interface", fmt.Sprint(x))
//}
//func main() {
//	var x /**int*//* interface{}*/ *float32 = nil
//	Foo(x)
//}




//const cl  = 100
//
//var bl    = 123
//
//func main() {
//	println(&bl, bl)
//	println(&cl, cl)
//}

type stu struct {
	Name string
}

func (s *stu) ShowA(){
	s.Name = "A"
}


func (s stu) ShowB(){
	s.Name = "B"
}

func main() {
	s := stu{Name: "c"}
	s.ShowA()
	fmt.Println(s)
	s.ShowB()
	fmt.Println(s)
}