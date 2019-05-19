package main

import (
	//"strings"
	//"fmt"
	//"bytes"
	//"github.com/ethereum/go-ethereum/rlp"
	"fmt"
	//"time"
)

func main() {
	//s := "248 88 179 84 104 101 32 108 101 110 103 116 104 32 111 102 32 116 104 105 115 32 115 101 110 116 101 110 99 101 32 105 115 32 109 111 114 101 32 116 104 97 110 32 53 53 32 98 121 116 101 115 44 32 163 73 32 107 110 111 119 32 105 116 32 98 101 99 97 117 115 101 32 73 32 112 114 101 45 100 101 115 105 103 110 101 100 32 105 116"
	//sArr := strings.Split(s, " ")
	//fmt.Println(sArr)
	//fmt.Println(len(sArr))
	//sr := strings.Join(sArr, ",")
	//fmt.Println(sr)


	//sArr := []string{"The length of this sentence is more than 55 bytes, ", "I know it because I pre-designed it"}
	//for _, v := range  sArr {
	//	fmt.Println(len([]byte(v)))
	//}

	//type Student struct{
	//	Name string
	//	Sex string
	//}
	//s := Student{Name:"icattlecoder",Sex:"male"}
	//buff := bytes.Buffer{}
	//rlp.Encode(&buff, &s)
	//fmt.Println(buff.Bytes())
	// [210 140 105 99 97 116 116 108 101 99 111 100 101 114 132 109 97 108 101]

	//var m1 map[string]string = nil
	// //m2 := make(map[string]string, 0)
	// var m2 map[string]string = nil
	////m["s"] = "4"
	//
	////var m []string
	//
	////m = append(m, "a")
	//fmt.Println(m1 == m2)

	// Send the filter channel to the fetcher
	//filter := make(chan string)
	//
	//
	//// Request the filtering of the header list
	//select {
	//case filter <- "lala":
	//	case
	//}
	//// Retrieve the headers remaining after filtering
	//select {
	//case task := <-filter:
	//fmt.Println(task)
	//}


	//time.Tick(time.Second)
	//
	//time.After(time.Second)
	//time.NewTimer(time.Second).C

	//select {
	//	go func() {
	//		fmt.Println("a")
	//	}()
	//}
	//
	//fmt.Println("x")


	//round := uint64(751)/ 250   // round = 3
	//mod := uint64(751) % 250 	// mod = 1
	//fmt.Println(round, mod)
	//fmt.Println(round * 250)




	//round := uint64(750)/ 250   // round = 3
	//mod := uint64(750) % 250 	// mod = 1
	//fmt.Println(round, mod)
	//fmt.Println(round * 250)



	//round := uint64(200)/ 250   // round = 3
	//mod := uint64(200) % 250 	// mod = 1
	//fmt.Println(round, mod)
	//fmt.Println(round * 250)

	//fmt.Println(calcurround(1001))

	/**
	有 1 得 1
	0010
	0100
	----
	0110
	*/
	fmt.Println(TWO|FOUR)   // 0010|0100 == 0110
	// 0110 & 0010 == 0010
	// 0110 & 0100 == 0100

	/**
	1 1 得 1
	0010
	0100
	----
	0000
	 */
	fmt.Println(TWO&FOUR) 	// 0010&0100 == 0000

	/**
	1 1 得 1
	0010
	0010
	----
	0010
	*/
	fmt.Println(TWO&TWO)


	//fmt.Println(TWO)
	//fmt.Println(FOUR)



}

const (


	ONE	 = 1 << iota	// 0001
	TWO					// 0010
	FOUR				// 0100
	ZERO = 0   			// 0000

)

func calcurround (blocknumber uint64) uint64 {
	// current num
	var round uint64
	div := blocknumber/ 250
	mod := blocknumber% 250
	if (div == 0 && mod == 0) || (div == 0 && mod > 0 && mod < 250) { // first round
		round = 1
	}else if div > 0 && mod == 0 {
		round = div
	} else if div > 0 && mod > 0 && mod < 250 {
		round = div + 1
	}
	return round
}