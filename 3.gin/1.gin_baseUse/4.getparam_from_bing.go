package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 参数绑定
/*
为了能够更方便的获取请求相关参数，提高开发效率，我们可以基于请求的Content-Type识别请求数据类型并利用反射机制自动提取请求中QueryString、
form表单、JSON、XML等参数到结构体中。
下面的示例代码演示了.ShouldBind()强大的功能，它能够基于请求自动提取JSON、form表单和QueryString类型的数据，
并把值绑定到指定的结构体对象。
*/

// Binding from JSON
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// 绑定QueryString示例 (/loginQuery?username=q1mi&password=123456)
	router.GET("/loginQuery", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"username": login.Username,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	// 绑定form表单示例 (username=yy&password=123456)
	router.POST("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		// application/json
		// application-context/x-www-form-urlencoded
		// application-context/form-data
		// ...
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"username": login.Username,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
