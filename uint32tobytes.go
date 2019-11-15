package main

//
//import (
//	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
//
//	//"fmt"
//	//"github.com/PlatONnetwork/PlatON-Go/common"
//	//"github.com/PlatONnetwork/PlatON-Go/x/xutil"
//	"fmt"
//	//"github.com/pangu/PlatON-Go/rlp"
//
//	//"github.com/PlatONnetwork/PlatON-Go/core/state"
//	//"github.com/PlatONnetwork/PlatON-Go/core/vm"
//	//"github.com/PlatONnetwork/PlatON-Go/x/xcom"
//	"math"
//)
//
//func main() {
//
//	//a := uint32(1792)
//	//b := uint32(65536)
//	//c := xutil.CalcVersion(b)
//	//fmt.Println("a", common.Uint32ToBytes(a))
//	//fmt.Println("b", common.Uint32ToBytes(b))
//	//fmt.Println("c", common.Uint32ToBytes(c))
//
//
//	//fmt.Println("清除掉[x, y)范围bit上的1,得到的值:", clear(1221892080809121, 10, 63))
//	//
//	//// xxxx.1111  == 15
//	//fmt.Println(clear(15, 2, 3))
//
//	m := make(map[string]string, 0)
//	m["B"] = "b"
//	m["A"] = "a"
//	m["C"] = "c"
//	m["D"] = "d"
//
//	for k, v := range m {
//		fmt.Println("Key", k, "Value", v)
//	}
//
//
//
//	//stateDB := &state.StateDB{}
//	//
//	//var vms vm.StateDB
//	//
//	//var xcoms xcom.StateDB
//	//
//	//
//	//vms = stateDB
//	//
//	//
//	//xcoms = stateDB
//	//
//	//fmt.Printf("%p : %p : %p", stateDB, vms, xcoms)
//
//
//	//var str string
//
//	rlpByte, err := hexutil.Decode("0x7b22537461747573223a747275652c2244617461223a227b5c224e6f646549645c223a5c2238396134343039616265316163653862373763343439376332303733613861323034366462646162623538633862623538666537333932366262646335373266623834386437333962316432643039646430373936616263633165643864396133336262336566306136633265313036653430383039306466313739623034315c222c5c225374616b696e67416464726573735c223a5c223078343933333031373132363731616461353036626136636137383931663433366432393138353832315c222c5c2242656e65666974416464726573735c223a5c223078313030303030303030303030303030303030303030303030303030303030303030303030303030335c222c5c225374616b696e675478496e6465785c223a342c5c2250726f6772616d56657273696f6e5c223a313739322c5c225374617475735c223a302c5c225374616b696e6745706f63685c223a312c5c225374616b696e67426c6f636b4e756d5c223a302c5c225368617265735c223a333030303030303030303030303030303030303030303030302c5c2252656c65617365645c223a313030303030303030303030303030303030303030303030302c5c2252656c65617365644865735c223a323030303030303030303030303030303030303030303030302c5c225265737472696374696e67506c616e5c223a302c5c225265737472696374696e67506c616e4865735c223a302c5c2245787465726e616c49645c223a5c225c222c5c224e6f64654e616d655c223a5c22706c61746f6e2e6e6f64652e345c222c5c22576562736974655c223a5c227777772e706c61746f6e2e6e6574776f726b5c222c5c2244657461696c735c223a5c2254686520506c61744f4e204e6f64655c227d222c224572724d7367223a226f6b227d")
//
//	if nil != err {
//		fmt.Println("rlp byte err", err)
//		return
//	}
//
//	//err = rlp.DecodeBytes(rlpByte, &str)
//	//if nil != err {
//	//	fmt.Println("rlp decode err", err)
//	//	return
//	//}
//
//	fmt.Println(string(rlpByte))
//
//
//
//
//}
//
//
//// 清除掉[x, y)范围bit上的1
//func clear(n uint64, i, j uint8) uint64 {
//	return (math.MaxUint64<<j | ((1 << i) - 1)) & n
//}

import "fmt"
import "unicode/utf8"

func main() {
	fmt.Println("Hello, 世界", len("世界"), utf8.RuneCountInString("世界"))
	fmt.Println("Hello, 世界", len("Hello"), utf8.RuneCountInString("Hello"))
	fmt.Println("Hello, 世界", len("Hello, 世界"), utf8.RuneCountInString("Hello, 世界"))
}
