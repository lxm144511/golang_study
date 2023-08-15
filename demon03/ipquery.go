package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.oioweb.cn/api/ip/ipaddress"
	clinet := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("requests failed ", err)
		return
	}
	//url := "https://api.oioweb.cn/api/ip/ipaddress?ip=220.197.198.41"
	// 发送请求参数
	query := req.URL.Query()
	query.Add("ip", "220.197.198.41") //query.Add() 可以添加多个
	req.URL.RawQuery = query.Encode()
	res, err := clinet.Do(req)
	if err != nil {
		fmt.Println("requests failed", err)
		return
	}
	defer res.Body.Close()
	results, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(results))

}
