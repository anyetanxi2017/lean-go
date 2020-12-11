package routers

// 路由拆分-购物相关
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理shop 相关的业务
func LoadShop(e *gin.Engine) {
	e.GET("/shop", shop)
	e.GET("/goods", goodsList)
}

func goodsList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"mgs": "ok", "code": 1, "data": "dats"})
}

func shop(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "ok", "code": 1})
}
