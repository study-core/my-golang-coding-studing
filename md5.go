package main 

import(
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

func main() {
	
	id := "0000000029"

	pwd := "670b14728ad9902aecba32e22fa4f6bd"

	fmt.Println("userId:=" + id + "\npwd:=" + pwd)
	Md5(id + pwd)
}



func Md5(data string) string {
	hash := md5.New()          //初始化一个MD5对象
	hash.Write([]byte(data))   // 需要加密的字符串
	cipherStr := hash.Sum(nil) //计算出校验和
	fmt.Println("md5前的信息: ", data)
	fmt.Println("md5后的信息: ", hex.EncodeToString(cipherStr))
	return hex.EncodeToString(cipherStr) // 输出加密结果
}