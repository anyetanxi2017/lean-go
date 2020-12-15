package user

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"lean-go/3.gin/4.struct_demo/web/models/request"
	"lean-go/3.gin/4.struct_demo/web/models/response"
	"lean-go/3.gin/4.struct_demo/web/service/user"
)

type Controller struct {
	request.BaseController
	response.Response
}

func (v *Controller) Login(c *gin.Context) {
	resUsername := c.PostForm("username")
	resPwd := c.PostForm("password")
	// 参数验证
	valid := validation.Validation{}
	valid.Required(resUsername, "username").Message("用户名不能为空")
	valid.Required(resPwd, "password").Message("密码不能为空")
	if valid.HasErrors() {
		v.Response.Fail(c, valid.Errors[0].Message)
		return
	}
	svc := user.NewServiceUser()
	res, err := svc.Login(resUsername, resPwd)
	if err != nil {
		v.Response.Fail(c, err.Error())
		return
	}
	v.Response.OkWithData(res, "登录成功", c)
}
