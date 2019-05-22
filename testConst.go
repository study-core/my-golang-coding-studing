package main

import (
	"fmt"
	//"golang.org/x/net/html/atom"
)

func main() {
	//fmt.Println(MUTEXLOCKED, MUTEXWOKEN, MUTEXSTARVING, MUTEXWAITERSHIFT)
	//fmt.Println(PREVIOUS_C, CURRENT_C, NEXT_C)
	//fmt.Println(LT, GT, EQ)


	total := 250000000

	i := 1

	use := 25000000

	for total > 0 {

		if total == 100000012 {
			fmt.Println()
		}
		if total <= 0 {
			break
		}
		if i <= 5 {
			fmt.Println("第" + fmt.Sprint(i) + "年", "total", total, "use", use)
			total -= use
			i ++
			continue
		}

		if i%5 == 0 {

			use = use/2
			fmt.Println("第" + fmt.Sprint(i) + "年", "total", total, "use", use)
			total -= use
			i ++
			continue
		}
		i ++
	}
	i--
	fmt.Println(i)
}


const (
	MUTEXLOCKED      = 1 << iota			// 	0001
	MUTEXWOKEN								//	0010
	MUTEXSTARVING							//	0100
	MUTEXWAITERSHIFT = iota
)

//const (
//	PREVIOUS_C, LT = iota -1, iota -1
//	CURRENT_C, GT
//	NEXT_C, EQ
//)


