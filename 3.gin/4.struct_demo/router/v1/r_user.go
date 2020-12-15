package v1

import (
	"github.com/gin-gonic/gin"
	"lean-go/3.gin/4.struct_demo/web/controller/v1/user"
)

func rUserAccount(g *gin.RouterGroup) {
	// 这是原代码 ，这里 response.NewBaseResponse() 拿不到不清楚是什么
	//cSet := setting.Controller{BaseResponse: response.NewBaseResponse()}
	cUser := user.Controller{}
	rUser := g.Group("/user")
	{
		rUser.POST("/login", cUser.Login)
	}
}
