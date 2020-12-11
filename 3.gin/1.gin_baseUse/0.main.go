package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Gin官网 https://gin-gonic.com/zh-cn/docs/

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// Get 请求响应 方式1 使用 interface 空接口
	r.GET("/hello", func(c *gin.Context) {
		res := map[string]interface{}{
			"name": "yy",
			"age":  12,
		}
		c.JSON(http.StatusOK, res)
	})
	// Get 请求响应 方式2 使用 gin.H{}(返回就是一个 空接口)
	r.GET("/hello2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "yy",
			"age":  22,
		})
	})
	// Get 请求响应 方式3
	r.GET("/json", func(c *gin.Context) {
		msg := UserInfo{"yy", "hello", 20}
		c.JSON(http.StatusOK, msg)
	})
	r.Run()
}

type UserInfo struct {
	Name string
	Msg  string
	Age  int
}
