package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取form参数
/*
请求的数据通过form表单来提交，例如向/user/search发送一个POST请求，获取请求数据的方式如下：
*/
func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	// 获取Form参数
	r.POST("/user/search", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.PostForm("address")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"address":  address,
		})
	})
	r.Run(":8080")
}
