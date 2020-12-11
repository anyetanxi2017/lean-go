package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lean-go/3.gin/app/blog"
	"lean-go/3.gin/app/shop"
	"lean-go/3.gin/routers"
)

//gin框架路由拆分与注册
/*
当项目的规模增大后就不太适合继续在项目的main.go文件中去实现路由注册相关逻辑了，我们会倾向于把路由部分的代码都拆分出来，形成一个单独的文件或包：
我们在routers.go文件中定义并注册路由信息：
*/
func main() {
	// 路由拆分 示例
	//splitRoutingDemo()

	// 路由拆分到不同的APP
	splitRoutingWhenApp()
}

func splitRoutingWhenApp() {
	// 有时候项目规模实在太大，那么我们就更倾向于把业务拆分的更详细一些，例如把不同的业务代码拆分成不同的APP。
	routers.Include(shop.Routers, blog.Routers)
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("starup service failed,err:%v\n", err)
	}
	// gin框架是一个非常容易扩展的web框架，本文是我在日常编码中总结的一点点经验，
	// 因为世界上不可能有完全相同的项目，每个人也都有自己的编程习惯，
	// 关于gin框架路由注册的方式我就在此抛砖引玉了。
}

func splitRoutingDemo() {
	r := gin.Default()
	//注册到r上下文
	routers.LoadBlog(r)
	routers.LoadShop(r)
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
