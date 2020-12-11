package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向
func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	r.Run()
}
