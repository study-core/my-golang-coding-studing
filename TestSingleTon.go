package main

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"

	//"github.com/pangu/PlatON-Go/rlp"
	"math/big"
)

//import "sync"

func main() {


	c := &Candi{
		Amount: big.NewInt(12),
	}

	/*fmt.Println(c.Amount)

	c.Amount = new(big.Int).Sub(c.Amount, big.NewInt(13))

	fmt.Println(c.Amount)

	if by, err := rlp.EncodeToBytes(c); nil != err {
		fmt.Println("err", err.Error())
	}else {
		var ca *Candi
		rlp.DecodeBytes(by, ca)
		fmt.Println(ca)
	}*/

	i := include{
		a: c,
	}


	fmt.Println(GetAString(i.a))
	fmt.Println(GetBString(i.a))


	flag := common.HexToHash(common.Hash{}.String()) == (common.Hash{})
	fmt.Println(flag)
}


type Candi struct {

	Amount *big.Int
}

func (c *Candi) String() string {
	status := fmt.Sprintf("[")
	status += fmt.Sprintf("[%d]", c.Amount)
	status += fmt.Sprintf("]")
	return status
}


type include struct {
	a Ainterface
}

//
//type MyFeed struct {
//	once      sync.Once
//}
//
//
//
//
//func (m *MyFeed) GetInstance () *MyFeed  {
//
//}



type Ainterface interface {
	String() string
}


func GetAString(a Ainterface) string {
	fmt.Println("我是A接口")
	return a.String()
}


type Binterface interface {
	String() string
}


func GetBString(b Binterface) string {
	fmt.Println("我是B接口")
	return b.String()
}