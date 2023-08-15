package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.oioweb.cn/api/weather/GetWeather"
	// 建立http客户端
	client := &http.Client{}
	// 传入参数 : 请求方式、url、请求体body
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("api requests failed ", err)
		return
	}
	// 发送请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("requests failed ", err)
		return
	}
	defer res.Body.Close()
	// 响应的头部信息
	//fmt.Println(res.Header.Get("Content-Type"))
	// 响应码
	//fmt.Println(res.StatusCode)
	// cookies
	// fmt.Println(res.Cookies())
	// 响应数据
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read fiiled ", err)
		return
	}
	//fmt.Println(string(body))
	// 创建map, 将返回的数据通过json.Unmarshal()反序列化到map
	newmap := make(map[string]interface{})
	err = json.Unmarshal([]byte(body), &newmap)
	if err != nil {
		fmt.Println("转换为map类型失败", err)
		return
	}
	//  打印map数据
	//  fmt.Println(newmap)
	// map 长度
	//fmt.Println(len(newmap))
	// 获取指定数据
	// 由于map的values 为空接口类型 需要使用类型断言才能进行for 循环取值
	data := newmap["result"].(map[string]interface{})["city"] //断言
	for k, v := range data.(map[string]interface{}) {
		fmt.Println(k, "=", v)

	}
	//fmt.Println("##############################################################")
	data1 := newmap["result"].(map[string]interface{})["forecast"]
	for _, v := range data1.([]interface{}) {
		for key, values := range v.(map[string]interface{}) {
			fmt.Println(key, ":", values)
		}
		fmt.Println("#################################################################")
	}

}
