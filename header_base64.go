package main 

import(
	b64 "encoding/base64"
	"fmt"
)

func main() {
	headerStr := `{
    "channelId":1,
    "version":"4.8",
    "clientIp":"192.168.1.1",
    "appStoreId":"0-maizuo",
    "equipmentId":"807f45f139a70658f1a1d3a5fe915952",
    "userId":201069818,
    "cityId":10,
    "longitude":114.085947,
    "latitude":22.547101,
    "locationType":61,
    "carrier":1,
    "salesChannelId":54
}`

/**
ewogICAgImNoYW5uZWxJZCI6MSwKICAgICJ2ZXJzaW9uIjoiNC44IiwKICAgICJjbGllbnRJcCI6IjE5Mi4xNjguMS4xIiwKICAgICJhcHBTdG9yZUlkIjoiMC1tYWl6dW8iLAogICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLAogICAgInVzZXJJZCI6MjAxMDY5ODE4LAogICAgImNpdHlJZCI6MTAsCiAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LAogICAgImxhdGl0dWRlIjoyMi41NDcxMDEsCiAgICAibG9jYXRpb25UeXBlIjo2MSwKICAgICJjYXJyaWVyIjoxLAogICAgInNhbGVzQ2hhbm5lbElkIjo1NAp9
*/

	sDec:= b64.StdEncoding.EncodeToString([]byte(headerStr))
	fmt.Println("请求头编码结果:", string(sDec))


	//decode
	deStr := `ewogICAgImNoYW5uZWxJZCI6MSwKICAgICJ2ZXJzaW9uIjoiNC44IiwKICAgICJjbGllbnRJcCI6IjE5Mi4xNjguMS4xIiwKICAgICJhcHBTdG9yZUlkIjoiMC1tYWl6dW8iLAogICAgImVxdWlwbWVudElkIjoiODA3ZjQ1ZjEzOWE3MDY1OGYxYTFkM2E1ZmU5MTU5NTIiLAogICAgInVzZXJJZCI6MjAxMDY5ODE4LAogICAgImNpdHlJZCI6MTAsCiAgICAibG9uZ2l0dWRlIjoxMTQuMDg1OTQ3LAogICAgImxhdGl0dWRlIjoyMi41NDcxMDEsCiAgICAibG9jYXRpb25UeXBlIjo2MSwKICAgICJjYXJyaWVyIjoxLAogICAgInNhbGVzQ2hhbm5lbElkIjowCn0=
`
/**
{
    "channelId":1,
    "version":"4.8",
    "clientIp":"192.168.1.1",
    "appStoreId":"0-maizuo",
    "equipmentId":"807f45f139a70658f1a1d3a5fe915952",
    "userId":201069818,
    "cityId":10,
    "longitude":114.085947,
    "latitude":22.547101,
    "locationType":61,
    "carrier":1,
    "salesChannelId":0
}

*/

	sDec2, _ := b64.StdEncoding.DecodeString(deStr)
	fmt.Println("请求头解码结果:", string(sDec2))
}