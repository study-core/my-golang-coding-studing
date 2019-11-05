package main

import (
	"encoding/hex"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"github.com/PlatONnetwork/PlatON-Go/x/xutil"
)

//func main() {
//
//
//	callback(callbackB)
//}
//
//
//func callbackA (str string) {
//	fmt.Println(str)
//}
//
//func callbackB (str string) bool {
//	fmt.Println("a" == str)
//	return "a" == str
//}
//
//func callback(f func()) {
//
//	f()
//}

func main() {


	addr1, _ := xutil.NodeId2Addr(discover.MustHexID("ced880d4769331f47af07a8d1b79de1e40c95a37ea1890bb9d3f0da8349e1a7c0ea4cadbb9c5bf185b051061eef8e5eadca251c24e1db1d9faf0fb24cbd06f9a"))
	addr2, _ := xutil.NodeId2Addr(discover.MustHexID("65e2ab09161e32e6d07d82adaa416ee6d41d617c52db20e3145a4d1b7d396af38d095c87508ad5bb35df741513bdc4bf12fec215e58450e255f05d194d41d089"))
	addr3, _ := xutil.NodeId2Addr(discover.MustHexID("248af08a775ff63a47a5970e4928bcccd1a8cef984fd4142ea7f89cd13015bdab9ca4a8c5e1070dc00fa81a047542f53ca596f553c4acfb7abe75a8fb5019057"))
	addr4, _ := xutil.NodeId2Addr(discover.MustHexID("56d243db84a521cb204f582ee84bca7f4af29437dd447a6e36d17f4853888e05343844bd64294b99b835ca7f72ef5b1325ef1c89b0c5c2744154cdadf7c4e9fa"))

	addr5, _ := xutil.NodeId2Addr(discover.MustHexID("8796a6fcefd9037d8433e3a959ff8f3c4552a482ce727b00a90bfd1ec365ce2faa33e19aa6a172b5c186b51f5a875b5acd35063171f0d9501a9c8f1c98513825"))
	addr6, _ := xutil.NodeId2Addr(discover.MustHexID("547b876036165d66274ce31692165c8acb6f140a65cab0e0e12f1f09d1c7d8d53decf997830919e4f5cacb2df1adfe914c53d22e3ab284730b78f5c63a273b8c"))
	addr7, _ := xutil.NodeId2Addr(discover.MustHexID("9fdbeb873bea2557752eabd2c96419b8a700b680716081472601ddf7498f0db9b8a40797b677f2fac541031f742c2bbd110ff264ae3400bf177c456a76a93d42"))
	addr8, _ := xutil.NodeId2Addr(discover.MustHexID("8fb8a89ad4cbe24db75c1d0b4a371a89d01357f8b0fbd2aa108804541a1a0e3cba78346610b1b6796536b133f475b4cd66392d7e7fd42cef96e41ab358e86a93"))

	fmt.Println("addr1", addr1.String())
	fmt.Println("addr2", addr2.String())
	fmt.Println("addr3", addr3.String())
	fmt.Println("addr4", addr4.String())

	fmt.Println("addr5", addr5.String())
	fmt.Println("addr6", addr6.String())
	fmt.Println("addr7", addr7.String())
	fmt.Println("addr8", addr8.String())



	str := "我们啊啊啊ada4897"  // 中文每个char长 3， 英文每个 char长 1
	fmt.Println(len(str))


	event := common.BytesToHash(crypto.Keccak256([]byte("1001")))
	s := hex.EncodeToString(event[:])
	fmt.Println(s)  // 0x6171fdca598d5b5b0972f826968da3c55dbf3423f194b07e62a2d999bd0d14e0
					//   6171fdca598d5b5b0972f826968da3c55dbf3423f194b07e62a2d999bd0d14e0


	newVersion  := uint32(1<<16 | 2<<8 | 0)

	fmt.Println(newVersion)

}
