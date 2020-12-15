package main

import (
	"github.com/gin-gonic/gin"
	"lean-go/3.gin/4.struct_demo/middlewares"
	"lean-go/3.gin/4.struct_demo/router"
)

func main() {
	// 注册前置组件
	// 协程通知处理
	//执行已注册组件动作 验证组件是否正常
	func() {
		// 启动定时任务
	}()
	r := gin.Default()
	r.Use(
		middlewares.MiddlewareErrorHandle,
		middlewares.MyCors,
		middlewares.LoggerToFile(),
	)
	router.Register(r)
	//GoMicroService.Handle("/", http.DefaultServeMux)
	r.Run("9001")
}
