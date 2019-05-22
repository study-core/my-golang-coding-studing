// package main 

// import (
// 	"fmt"
// )



// func main() {
// 	a := 5
// 	fmt.Println(fmt.Sprint(a) + "的1的个数:" + fmt.Sprint(CalcOneNum(a)))
// 	fmt.Println(fmt.Sprint(a) + "交换符号:" + fmt.Sprint(ExcangeSymbol(a)))
// }

package main

import (
	"fmt"
	"time"
)

var battle = make(chan string)

func warrior(name string, done chan struct{}) {
    select {
    case opponent := <-battle:
        fmt.Printf("%s beat %s\n", name, opponent)
    case battle <- name:
    }
    done <- struct{}{}
}

// func main() {
//     done := make(chan struct{})
//     langs := []string{"Go", "C", "C++", "Java", "Perl", "Python"}
//     for _, l := range langs { go warrior(l, done) }
//     for _ = range langs { <-done }
// }


type Work struct {}

func (w *Work)Do(){
	fmt.Println("I am Do Func")
}

func (w *Work)Refuse (){
	fmt.Println("I am Refuse Func")
}

func worker(i int, ch chan Work, quit chan struct{}) {
    for {
        select {
        case w := <-ch:
            if quit == nil {
                w.Refuse(); fmt.Println("worker", i, "refused", w)
                break
            }
            w.Do(); fmt.Println("worker", i, "processed", w)
        case <-quit:
            fmt.Println("worker", i, "quitting")
            quit = nil
        }
    }
}

func main() {
    ch, quit := make(chan Work), make(chan struct{})
    go makeWork(ch)
    for i := 0; i < 4; i++ { go worker(i, ch, quit) }
    time.Sleep(5 * time.Second)
    close(quit)
    time.Sleep(2 * time.Second)
}





func CalcOneNum(a int) int {
    a = ((a & 0xAAAA) >> 1) + (a & 0x5555)
    a = ((a & 0xCCCC) >> 2) + (a & 0x3333)
    a = ((a & 0xF0F0) >> 4) + (a & 0x0F0F)
    a = ((a & 0xFF00) >> 8) + (a & 0x00FF)
    return a
}

func ExcangeSymbol(a int) (ret int) {
	ret = ^a + 1
	return
}