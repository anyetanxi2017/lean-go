package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello q1mi!",
	})
}

// 基本的路由注册
/*
下面最基础的gin路由注册方式，适用于路由条目比较少的简单项目或者项目demo。
*/
func main() {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
