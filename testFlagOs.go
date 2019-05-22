package main

import (
	//"os"
	"fmt"
	"flag"
	"os"
	"strings"
)

// 测试 os 和 flags

func main() {

	//if len(os.Args) != 0 {
	//	for i, arg := range os.Args {
	//		fmt.Println("i:", i, "v:", arg)
	//	}
	//}



	var ss bool
	var sr string
	b := flag.String("b", "我是b", "this is b !!")
	s := flag.Bool("s", false, "This bool")
	flag.BoolVar(&ss, "s2", false, "This is s2")
	flag.StringVar(&sr, "sr", "", "send `signal` to a master process: stop, quit, reopen, reload")

	// 自定义 flag的类型
	var languages []string
	flag.Var(newSliceValue([]string{}, &languages), "slice", "I like programming `languages`")

	// 替换Usage说明
	flag.Usage = usage
	flag.Parse()
	fmt.Println("b:=", *b)
	fmt.Println("s:=", *s)
	fmt.Println("s2:=", ss)
	fmt.Println("sr:=", sr)
	fmt.Println("slice:=", languages)
	for _, arg := range flag.Args() {
		fmt.Println(arg)
	}

}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]
 
Options:
`)
	flag.PrintDefaults()
}


// 自定义 flag 类型
// 需要实现 flag.Value 接口
// type Value interface {
//    String() string
//    Set(string) error
// }
type sliceValue []string

func newSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p)
}

func (s *sliceValue) Set(val string) error {
	*s = sliceValue(strings.Split(val, ","))
	return nil
}

func (s *sliceValue) String() string { return strings.Join([]string(*s), ",") }