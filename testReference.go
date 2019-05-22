package main
import (
	"fmt"
)
type AA struct {
	B interface{}
}

func geth(a, b int){
	fmt.Println(a + b)
}
func main() {
	a := AA{}
	a.B = geth
	(a.B.(func (int,int)))(1, 2)
}

