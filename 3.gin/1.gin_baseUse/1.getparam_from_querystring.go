package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取querystring参数
func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	// Get 获取参数
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "yy")
		address := c.Query("address")
		c.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run()
}
