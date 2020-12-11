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
	r.Run()
}
