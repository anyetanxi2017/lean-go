package main

import (
	"github.com/gin-gonic/gin"
	utils2 "lean-go/3.gin/4.struct_demo/basic/utils"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()
	r.POST("/auth", authHandler)
	// 为 /home请求添加 中间件
	r.GET("/home", JWTAuthMiddleware(), homeHandler)
	r.Run()
}

// 资源
func homeHandler(c *gin.Context) {
	mc := c.MustGet("username").(*utils2.MyClaims)
	c.JSON(http.StatusOK, gin.H{"code": 2000, "msg": "success", "data": mc})
	return
}

type UserInfo struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// 获取token
func authHandler(c *gin.Context) {
	var user UserInfo
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 2001, "msg": " 无效的参数"})
		return
	}
	if user.Username == "yy" && user.Password == "123123" {
		token, _ := utils2.GenToken(user.Username)
		c.JSON(http.StatusOK, gin.H{"code": 2000, "msg": "success", "data": gin.H{"token": token}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 2002, "msg": "鉴权失败"})
	return
}

/*
用户通过上面的接口获取Token之后，后续就会携带着Token再来请求我们的其他接口，
这个时候就需要对这些请求的Token进行校验操作了，很显然我们应该实现一个检验Token的中间件，具体实现如下：
*/

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中token为空",
			})
			c.Abort() //停止下面的执行
			return
		}
		// 控空格分割
		parts := strings.SplitN(token, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中token格式有误",
			})
			c.Abort()
			return
		}
		info, err := utils2.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 2005, "msg": "无效的Token"})
			c.Abort()
			return
		}
		//  将当前请求的username信息保存到请求的上下文c上
		c.Set("username", info)
		c.Next() //后续的处理函数可以通过 c.Get("username")来获取当前请求的用户信息
	}
}

// 如果不想自己实现上述功能，你也可以使用Github上别人封装好的包 https://github.com/appleboy/gin-jwt
