package shop

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/goods", goodsHandler)
	e.GET("/checkout", checkoutHandler)
}

