package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"lean-go/3.gin/4.struct_demo/basic/config"
	"time"
)

//func main() {
//	r := gin.Default()
//	r.Use(LoggerToFile()) // 暂时注释，测试稳定后再放开
//	r.GET("/hello", func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "hello"})
//	})
//	r.Run()
//}

// 日志记录中间件
func LoggerToFile() gin.HandlerFunc {
	// 例項化
	logger := logrus.New()
	// 設定日誌級別
	logger.SetLevel(logrus.DebugLevel)
	// 設定日誌格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	lfHook := config.GetLfsHook("visit")
	// 新增 Hook
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		// 開始時間
		startTime := time.Now()

		// 處理請求
		c.Next()

		// 結束時間
		endTime := time.Now()

		// 執行時間
		latencyTime := endTime.Sub(startTime).Seconds()

		// 請求方式
		reqMethod := c.Request.Method

		// 請求路由
		reqUri := c.Request.RequestURI

		// 请求参数
		param := c.Request.Form.Encode()
		//param := c.Request.Header()
		// 狀態碼
		statusCode := c.Writer.Status()

		// 請求IP
		clientIP := c.ClientIP()

		// 日誌格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"param":        param,
			"@timestamp":   time.Now().Format("2006-01-02 15:04:05"),
		}).Info()
	}
}
