package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func checkoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 1})
}

func goodsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 1})
}
