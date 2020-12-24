package main

import (
	"github.com/gin-gonic/gin"
	"lean-go/3.gin/4.struct_demo/basic/config"
	"lean-go/3.gin/4.struct_demo/basic/plugins"
	"lean-go/3.gin/4.struct_demo/middlewares"
	"lean-go/3.gin/4.struct_demo/router"
)

func main() {
	config.InitEnv()
	// 注册前置组件
	plugins.UsePlugin(
		plugins.PluginDb, //初始化数据库
	)
	// 执行已注册组件动作
	plugins.ExecPlugin()

	// 协程通知处理
	//执行已注册组件动作 验证组件是否正常
	func() {
		// 启动定时任务
	}()
	r := gin.Default()
	r.Use(
		middlewares.MiddlewareErrorHandle, //错误信息拦截器
		middlewares.MyCors,
		middlewares.LoggerToFile(),
	)
	router.Register(r)
	r.Run(":9001")
}
