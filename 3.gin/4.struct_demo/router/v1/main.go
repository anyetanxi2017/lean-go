package v1

import "github.com/gin-gonic/gin"

func InitRouterV1(r *gin.Engine) {
	r.GET("svcstatus", func(c *gin.Context) {
		c.String(200, "")
	})
	rGroup := r.Group("v1/app")
	{
		rUserAccount(rGroup)
	}
}
