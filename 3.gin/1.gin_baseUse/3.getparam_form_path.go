package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取path参数
/*
请求的数据通过form表单来提交，例如向/user/search发送一个POST请求，
*/
func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	// 获取path参数
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"address":  address,
		})
	})
	r.Run(":8080")
}
