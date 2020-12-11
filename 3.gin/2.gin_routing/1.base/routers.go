package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", hello2Handler)
	return r
}

func hello2Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "1"})
}
