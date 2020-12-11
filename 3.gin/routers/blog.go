package routers

// 路由拆分-博客
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadBlog(e *gin.Engine) {
	e.GET("/listLog", listLog)
	e.GET("/saveLog", saveLog)
	e.GET("/delete/:id", deleteLog)
}

func deleteLog(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"code": 1, "'msg": fmt.Sprintf("已经删除id为%s的博客", id)})
}

func saveLog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "ok"})
}

func listLog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "ok", "data": "datas"})
}
