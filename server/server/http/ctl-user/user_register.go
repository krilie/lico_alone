package ctl_user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// UserLogin 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags 用户
// @ID 用户注册
// @Produce  json
// @Param phone formData string true "用户手机号"
// @Param password formData string true "用户密码"
// @Param valid_code formData string true "验证码"
// @Success 200 {object} com_model.CommonReturn
// @Failure 500 {string} errInfo
// @Router /api/user/register [post]
func (a *UserCtrl) UserRegister(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	validCode := c.PostForm("valid_code")
	err := a.userService.UserRegister(ginutil.MustGetAppCtx(c), phone, password, validCode, "")
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnOk(c)
	return
}
