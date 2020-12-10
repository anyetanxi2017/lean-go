package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Go语言内置的net/http包十分的优秀，提供了HTTP客户端和服务端的实现。

func demo70() {
	// GET请求
	//res, err := http.Get("http://www.baidu.com")
	//if err != nil {
	//	fmt.Printf("get failed,err:%v\n", err)
	//}
	//defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Printf("read from res.Body fialed.err:%v\n", err)
	//}
	//fmt.Print(string(body))

	// GET 带参数请求
	apiurl := "http://m.kaileizhengu5.com/v1/lottery/outer/getStockOpenMessage"
	data := url.Values{}
	data.Set("lotteryCodes", "1801,1841")
	u, err := url.ParseRequestURI(apiurl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
		return
	}
	u.RawQuery = data.Encode()
	fmt.Println(u.String())

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed,err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read res body failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))

	// Post请求
	demoPost()
}

func demoPost() {
	url := "http://127.0.0.1:9090/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=小王子&age=18"
	// json
	contentType := "application/json"
	data := `{"name":"小王子","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
// http server

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello 沙河！")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
