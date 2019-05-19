package main 

import(
	"fmt"
	"errors"
	"github.com/parnurzeal/gorequest"
	"github.com/bitly/go-simplejson"
)



func HttpGet(url string, params interface{}) (string, error) {
	fmt.Println("get请求链接地址:", url, params)
	request := gorequest.New()
	_, body, errs := request.Get(url).Query(params).End()
	for _, err := range errs {
		if err != nil {
			fmt.Println("get请求链接地址:", url, params, "请求出错: err=", errs)
			return "", errors.New("请求第三方网络发生错误")
		}
	}
	fmt.Println("请求链接地址:", url, params, "返回的结果是: ", body)
	return body, nil
}


func HttpPostJson(url string, params interface{}) (string, error) {
	fmt.Println("postJson请求链接地址:", url, params)
	request := gorequest.New()
	_, body, errs := request.Post(url).Send(params).End()
	for _, err := range errs {
		if err != nil {
			fmt.Println("postJson请求链接地址:", url, params, "请求出错: err=", errs)
			return "", errors.New("请求第三方网络发生错误")
		}
	}
	fmt.Println("请求链接地址:", url, params, "返回的结果是: ", body)
	return body, nil
}

func main(){
	//获取access_token: 用httpClinet发一个 get请求 调用这个接口  https://oapi.dingtalk.com/gettoken?corpid=id&corpsecret=secrect
	access_token_url := "https://oapi.dingtalk.com/gettoken?corpid=ding09b9a7dcac3d504835c2f4657eb6378f&corpsecret=bwpDywU-Cxq87v-2EvG5FDIaGoZ7Dmos2k3xZyVA8b3EMV7OkMRqap-HCo686s5D"
	result, err := HttpGet(access_token_url, nil)
	if err != nil {
		fmt.Println("请求异常:" + err.Error())
	}

	fmt.Println("获取access_token返回的结果是: ", result)

	//处理返回内容 反序列化 string -> json
	json, err := simplejson.NewJson([]byte(result))
	if nil != err {
		fmt.Println("json解析异常")
	}
	//从json中拿到 access_token的值
	access_token := json.Get("access_token").MustString()

	fmt.Println("\n\n")

	//获取部门Id列表
	departmentId_url := "https://oapi.dingtalk.com/department/list_ids"
	departmentId_params := make(map[string]interface{}, 0)
	departmentId_params["access_token"] = access_token
	departmentId_params["id"] = "1"
	
	
	departmentId, err := HttpGet(departmentId_url, departmentId_params)
	if err != nil {
		fmt.Println("请求异常:" + err.Error())
	}
	fmt.Println("获取部门Id列表返回的结果是: ", departmentId)

	fmt.Println("\n\n")

	//获取部门列表
	departmentInfo_url := "https://oapi.dingtalk.com/department/list"
	departmentInfo_params := make(map[string]interface{}, 0)
	departmentInfo_params["access_token"] = access_token
	departmentInfo_params["id"] = "1"
	
	
	departmentInfo, err := HttpGet(departmentInfo_url, departmentInfo_params)
	if err != nil {
		fmt.Println("请求异常:" + err.Error())
	}
	fmt.Println("获取部门列表返回的结果是: ", departmentInfo)
	fmt.Println("\n\n")


	//获取部门成员（详情）
	userInfo_url := "https://oapi.dingtalk.com/user/list"
	params := make(map[string]interface{}, 0)
	params["access_token"] = access_token
	params["department_id"] = "1"
	params["offset"] = "0"
	params["size"] = "100"
	
	
	userInfo, err := HttpGet(userInfo_url, params)
	if err != nil {
		fmt.Println("请求异常:" + err.Error())
	}
	fmt.Println("获取部门成员（详情返回的结果是: ", userInfo)
	fmt.Println("\n\n")


	//获取部门成员
	simplelist_url := "https://oapi.dingtalk.com/user/simplelist"
	simplelist_params := make(map[string]interface{}, 0)
	simplelist_params["access_token"] = access_token
	simplelist_params["department_id"] = "1"
	
	
	simplelist, err := HttpGet(simplelist_url, simplelist_params)
	if err != nil {
		fmt.Println("请求异常:" + err.Error())
	}
	fmt.Println("获取部门成员返回的结果是: ", simplelist)
	fmt.Println("\n\n")



	//发送钉钉消息
	sendMsg_url := "https://oapi.dingtalk.com/message/send?access_token=" + access_token
	args := make(map[string]interface{}, 0)
	args["touser"] = "manager8500" //用户Id
	args["toparty"] = "1"    //部门Id
	args["agentid"] = "163934651" //微应用Id
	args["msgtype"] = "text"    //发送的消息类型
	content := make(map[string]interface{}, 0)
	content["content"] = "啦啦啦 wetetetet"  //对应类型的消息内容
	args["text"] = content
	
	sendResult, err := HttpPostJson(sendMsg_url, args)
	if err != nil {
		fmt.Println("请求异常:" + err.Error())
	}
	fmt.Println("发送钉钉消息返回的结果是: ", sendResult)
}



