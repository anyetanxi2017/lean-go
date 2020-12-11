package blog

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/post", postHandler)
	e.GET("/commentHandler", commentHandler)
}
