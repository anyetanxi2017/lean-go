package main

import (
	"fmt"
	"lean-go/3.gin/3.1.gin_routing/3.routing_for_multiple_app/app/blog"
	"lean-go/3.gin/3.1.gin_routing/3.routing_for_multiple_app/app/shop"
	"lean-go/3.gin/3.1.gin_routing/3.routing_for_multiple_app/routers"
)

// 路由拆分到不同的APP
/*
有时候项目规模实在太大，那么我们就更倾向于把业务拆分的更详细一些，例如把不同的业务代码拆分成不同的APP。

因此我们在项目目录下单独定义一个app目录，用来存放我们不同业务线的代码文件，这样就很容易进行横向扩展。
*/
func main() {
	routers.Include(shop.Routers, blog.Routers)
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed,err:%v\n", err)
	}
}

// gin框架是一个非常容易扩展的web框架，本文是我在日常编码中总结的一点点经验，
// 因为世界上不可能有完全相同的项目，每个人也都有自己的编程习惯，关于gin框架路由注册的方式我就在此抛砖引玉了。
