package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lean-go/3.gin/service"
	"log"
	"net/http"
	"time"
)

// Gin是一个用Go语言编写的web框架。
/*
https://gin-gonic.com/zh-cn/docs/

 安装 go get -u github.com/gin-gonic/gin
*/
func demo1() {
	r := gin.Default()
	r.Use(StartCost()) // 注册全局中间件

	demoBaseController(r)
	demoGetParam(r)
	demoUploadFile(r)
	demoRedirect(r)
	demoRouter(r)

	r.Run(":9002")
}

// 上传文件
func demoUploadFile(r *gin.Engine) {
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		} else {
			log.Println(file.Filename)
			dst := fmt.Sprintf("C:/%s", file.Filename)
			err := c.SaveUploadedFile(file, dst)
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"msg": "ok"})
		}
	})
}

// 重定向
func demoRedirect(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
}

// 路由组
func demoRouter(r *gin.Engine) {
	walletGroup := r.Group("/wa")
	{
		walletGroup.GET("/info", func(c *gin.Context) {
			info := walletService.GetInfo()
			c.JSON(http.StatusOK, info)
		})
		// TestCost() 为单个路由注册中间件
		walletGroup.GET("/list", TestCost(), func(c *gin.Context) {
			var res = map[string]interface{}{}
			wallet := walletService.ListWallet()
			res["data"] = wallet
			res["code"] = 0
			res["msg"] = "ok"
			c.JSON(http.StatusOK, res)
		})
	}
}

// Gin中间件
func StartCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().Nanosecond()
		c.Set("name", "yy")
		c.Next()
		end := time.Now().Nanosecond()
		log.Println("耗时", end-start)
	}
}
func TestCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("hehe")
		c.Next()
	}
}

func demoGetParam(r *gin.Engine) {
	// Get 获取参数
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "yy")
		address := c.Query("address")
		c.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"address":  address,
		})
	})
	// 获取Form参数
	r.POST("/user/search", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.PostForm("address")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"address":  address,
		})
	})
	// 获取path参数
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"address":  address,
		})
	})
	// 参数绑定
	r.POST("/login", func(c *gin.Context) {
		var login Login
		if err := c.BindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"username": login.Username,
				"package":  login.Password,
			})
		}
	})
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func demoBaseController(r *gin.Engine) {
	// Get 请求响应 方式1
	r.GET("/hello", func(c *gin.Context) {
		res := map[string]interface{}{
			"name": "yy",
			"age":  12,
		}

		c.JSON(http.StatusOK, res)
	})

	// Get 请求响应 方式2
	r.GET("/hello2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "yy",
			"age":  22,
		})
	})

	// Get 请求响应 方式3
	r.GET("/json", func(c *gin.Context) {
		var msg struct {
			Name string
			Msg  string
			Age  int
		}
		msg.Name = "yy"
		msg.Msg = "hello"
		msg.Age = 20
		c.JSON(http.StatusOK, msg)
	})
}
func main() {
	demo1()
}
