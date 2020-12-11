package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由组
func main() {
	r := gin.Default()
	walletGroup := r.Group("/wa")
	{
		walletGroup.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"code": 1})
		})
		walletGroup.GET("/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"code": 1})
		})
	}
}
