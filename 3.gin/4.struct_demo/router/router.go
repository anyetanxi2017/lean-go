package router

import (
	"github.com/gin-gonic/gin"
	v1 "lean-go/3.gin/4.struct_demo/router/v1"
)

// 路由注册demo

type FuncRouter func(r *gin.RouterGroup)

var FuncRouters = make([]FuncRouter, 0)

func Register(r *gin.Engine) *gin.Engine {
	rr := r.Group("")

	//gin.SetMode(gin.ReleaseMode)
	for _, funcHandle := range FuncRouters {
		funcHandle(rr)
	}
	//初始化v1路由
	v1.InitRouterV1(r)
	//初始化v1路由
	//初始化第三方开发接口
	//vc.InitRouterVc(r)
	return r
}
