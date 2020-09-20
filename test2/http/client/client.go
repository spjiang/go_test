package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//模拟一个post提交请求
	resp, err := http.Post("http://127.0.0.1:8888/hello", "application/x-www-form-urlencoded", strings.NewReader("id=1"))
	if err != nil {
		panic(err)
	}
	//关闭连接
	defer resp.Body.Close()
	//读取报文中所有内容
	body, err := ioutil.ReadAll(resp.Body)
	//输出内容
	fmt.Println(string(body))
}