package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func commentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0})
}

func postHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0})
}

