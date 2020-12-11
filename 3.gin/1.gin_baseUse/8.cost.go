package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// Gin中间件
/*
Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。
这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等。
*/
// Gin全局中间件
func GlobalCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().Nanosecond()
		c.Set("name", "yy")
		c.Next()
		end := time.Now().Nanosecond()
		log.Println("耗时", end-start)
	}
}

// 单个请求中间件
func BaseCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("hehe")
		c.Next()
	}
}

// 为路由组注册中间件
/*
写法1
shopGroup := r.Group("/shop", StatCost())
{
    shopGroup.GET("/index", func(c *gin.Context) {...})
    ...
}
写法2
shopGroup := r.Group("/shop")
shopGroup.Use(StatCost())
{
    shopGroup.GET("/index", func(c *gin.Context) {...})
    ...
}
*/
// 注意啊！！！！！！！！！
/*
gin默认中间件
gin.Default()默认使用了Logger和Recovery中间件，其中：

Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。


gin中间件中使用goroutine
当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。
*/
func main() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	// 注册一个全局中间件
	r.Use(GlobalCost())
	// 为某个路由单独注册中间件
	r.GET("/test", BaseCost(), func(c *gin.Context) {
		name := c.MustGet("name").(string) //从上下值取值
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{"msg": "ok"})
	})
	r.Run(":9001")
}
