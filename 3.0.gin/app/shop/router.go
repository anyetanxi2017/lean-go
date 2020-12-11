package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(e *gin.Engine) {
	e.GET("/goods", goodsHandler)
	e.GET("/checkout", checkoutHandler)
}

func checkoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 1})
}

func goodsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 1})
}
