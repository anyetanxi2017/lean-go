package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 上传文件
func main() {
	r := gin.Default()
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

	// 多个文件上传
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/uploadMultipartForm", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})
	r.Run()
}
