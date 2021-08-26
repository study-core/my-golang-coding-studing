package main

import (
	"fmt"
	//"strconv"

	//"github.com/PlatONnetwork/PlatON-Go/rlp"
	//
	//"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto/vrf"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"encoding/hex"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"bytes"
	"encoding/binary"
	//"sync"
)

func adder() func(int) int {
     sum := 0
     return func(x int) int {
          sum += x
          return sum
     }
}

func main() {
     // pos, neg := adder(), adder()
     // for i := 0; i < 10; i++ {
     //      fmt.Println(
     //           pos(i),
     //           neg(-2*i),
     //      )
     // }
	//unicodeStr := strconv.QuoteToASCII("∫ζ")
	//fmt.Println(unicodeStr)


	/*a := ""

	aa, _ := rlp.EncodeToBytes(a)
	fmt.Println(aa)

	b := uint64(360)

	bb, _ := rlp.EncodeToBytes(b)
	fmt.Println(bb)

	//i := 127
	//c1:=rune(i)
	//fmt.Println( strconv.Itoa(i) + " convert to:"+string(c1))

	////c :=  string(c1)
	//c := ""
	//cc, _ := rlp.EncodeToBytes(c)
	//fmt.Println(cc)

	fmt.Println(common.Uint64ToBytes(360))*/

	// todo 【最多拉 128 个headers, 每间隔 191 个blockNumber 拉取一个header】

	// 私钥:  64 字符
	// 1191dc5317d5930beb77848f416ee023921fa4452f4d783384f35352409c0ad0
	//
	// 公钥：就是整个 NodeId:   128 字符
	// 5a942bc607d970259e203f5110887d6105cc787f7433c16ce28390fb39f1e67897b0fb445710cc836b89ed7f951c57a1f26a0940ca308d630448b5bd391a8aa6
	//
	// 地址: 40 字符
	// 0x1000000000000000000000000000000000000004
	//
	// bech32 地址:  42 字符
	// lax16pqmt742fdepysd92yrlarceecd68elae9unre
	//
	// Hash: 64 字符
	// 0x000000000000000000000000000000005249b59609286f2fa91a2abc8555e887
	//
	// BlsPrivateKey:  192 字符
	// 5d0f8a399533b3f9b3a7198282c4b7b8b414529c66861d7958ebf908664707e5e6b353630b94ac5c1173c36e889fb403208ff73d233c12865d9e32256bbb988b931d41fda48e450b992fa5ec67790081e730965f548120b6d9fdc6156d66a614
	//
	// BlsProof: 128 字符
	// db18af9be2af9dff2347c3d06db4b1bada0598d099a210275251b68fa7b5a863d47fcdd382cc4b3ea01e5b55e9dd0bdbce654133b7f58928ce74629d5e68b974
	//
	// VersionSign: 130 字符
	// db18af9be2af9dff2347c3d06db4b1bada0598d099a210275251b68fa7b5a863d47fcdd382cc4b3ea01e5b55e9dd0bdbce654133b7f58928ce74629d5e68b97412

	str := "db18af9be2af9dff2347c3d06db4b1bada0598d099a210275251b68fa7b5a863d47fcdd382cc4b3ea01e5b55e9dd0bdbce654133b7f58928ce74629d5e68b97412"

	rr := []rune(str)
	fmt.Println(len(rr))

	// vrf 每次生成的是一样的 ...

	privateKey := crypto.HexMustToECDSA("1191dc5317d5930beb77848f416ee023921fa4452f4d783384f35352409c0ad0")

	source := []byte(str)

	value, _ := vrf.Prove(privateKey, source)
	fmt.Println(hex.EncodeToString(value))
	bb, _ := vrf.Prove(privateKey, source)
	fmt.Println(hex.EncodeToString(bb))

	publicKey := privateKey.PublicKey

	flag, _ := vrf.Verify(&publicKey, value, source)
	fmt.Println(flag)

	status, _ := vrf.Verify(&publicKey, bb, source)
	fmt.Println(status)


	hash := common.HexToHash("3b198bfd5d2907285af009e9ae84a0ecd63677110d89d7e030251acb87f6487e")

	sig1, _ := crypto.Sign(hash.Bytes(), privateKey)
	sig2, _ := crypto.Sign(hash.Bytes(), privateKey)

	fmt.Println("sig1:", sig1, "\nsig2:",sig2,  "\nis equals?", bytes.Compare(sig1, sig2))


	var val uint32 = 257
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, val)
	fmt.Println("大端编码:", buf) // [0 0 0 12]   高位数字, 放在低地址    uint32:   低地址 ->  高地址  (左到右)   (257 分为  0, 0, 256, 1)

	lif := make([]byte, 4)
	binary.LittleEndian.PutUint32(lif, val)
	fmt.Println("小端编码:", lif) // [0 0 0 12]   高位数字, 放在低地址  (反着放, 257 分为   1 , 256, 0, 0)

	//ReInLock()

	var slice []int
	slice[0] = 0

	fmt.Println(slice)
}

//// 不可重入锁  [可重入锁指同一个线程可以再次获得之前已经获得的锁，避免产生死锁]
//type MyLocker struct {
//	L sync.Mutex
//
//}
//
//func (m *MyLocker) ReInLock () {
//	fmt.Println("尝试里面函数加锁")
//	m.L.Lock()
//	fmt.Println("里面函数获得锁")
//	m.L.Unlock()
//	fmt.Println("尝试里面函数解锁")
//}
//
//
//func ReInLock() {
//	m := &MyLocker{}
//	fmt.Println("尝试外面函数加锁")
//	m.L.Lock()
//	fmt.Println("外面函数获得锁")
//	m.ReInLock()
//	m.L.Unlock()
//	fmt.Println("尝试外面函数解锁")
//}