package main

import (
	"fmt"
	"crypto/md5"
	//"encoding/hex"
	"time"
	"io"
)

func main() {


	//sub := float64(30)
	//div := sub/60
	//fmt.Println(div)
	//
	//fmt.Println(0.4 < div)

	//hash := md5.New()        // init one MD5 instance
	//hash.Write([]byte("")) //  the data need to MD5
	//cipherStr := hash.Sum(nil)
	//fmt.Println(hex.EncodeToString(cipherStr))

	fmt.Println(createPasswd())

}

// create random passwd
func createPasswd() string {
	t := time.Now()
	h := md5.New()
	io.WriteString(h, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	io.WriteString(h, t.String())
	passwd := fmt.Sprintf("%x", h.Sum(nil))
	return passwd
}

// a3231af7dc4752862177a0168288cc32
// 67e0e413282edaf556cfe339c48eba74
// 39152b4d3d5bd7fc0ff1f9ea4ce420c4
