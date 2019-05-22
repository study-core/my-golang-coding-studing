package main
import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"fmt"
	"crypto/sha256"
)
func main() {
	// 先获取一个椭圆实例
	curve := elliptic.P256()
	//得到私钥
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	//产生公钥
	publicKey := privateKey.PublicKey
	fmt.Println("priKey", privateKey, "\npubKey", publicKey)

	strHash := sha256.New().Sum([]byte( "我是学生"))
	// 签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, strHash)
	if nil != err {
		log.Panic(err)
	}
	fmt.Println("r", r, "\ns", s)
	strHash2 := sha256.New().Sum([]byte("我是个程序员"))
	// 验签
	fmt.Println(ecdsa.Verify(&publicKey, strHash2, r, s))  	// false
	fmt.Println(ecdsa.Verify(&publicKey, strHash, r, s))	// true
}
