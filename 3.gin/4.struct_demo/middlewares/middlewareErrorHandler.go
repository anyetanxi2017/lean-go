package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"lean-go/3.gin/4.struct_demo/basic/utils"
	"net/http"
	"runtime/debug"
)


// 错误中间件处理器
func MiddlewareErrorHandle(c *gin.Context) {

	//b:= response.NewBaseResponse()
	// 最后执行，发现程序有 panic 出现时 将其拯救出来
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			logrus.Errorln("panic errMsg: ", err)
			logrus.Errorln("panic errStack:", string(debug.Stack()))
			switch err.(type) {
			case utils.RuntimeException:
				c.JSON(http.StatusOK, err.(utils.RuntimeException).Result)
			case error:
				c.JSON(http.StatusOK, err.(error).Error())
			case utils.SQLError:
				c.JSON(http.StatusOK, err.(utils.SQLError).Result)
			default:
				result := utils.NewResult()
				result.SetCode(-1).SetMsg("System Exception").SetData(nil)
				c.JSON(http.StatusInternalServerError, result)
			}
		}
	}()
	c.Next()
}
