package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// UserLogin 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags 用户账号
// @ID 用户注册
// @Produce  json
// @Param phone formData string true "用户手机号"
// @Param password formData string true "用户密码"
// @Param valid_code formData string true "验证码"
// @Success 200 {object} cmodel.CommonReturn
// @Failure 400 {object} cmodel.CommonReturn
// @Failure 404 {object} cmodel.CommonReturn
// @Failure 500 {object} cmodel.CommonReturn
// @Router /api/user/register [post]
func (a *UserCtrl) UserRegister(c *gin.Context) {
	err := a.AppUser.UserRegister(ginutil.MustGetAppCtx(c), c.PostForm("phone"), c.PostForm("password"), c.PostForm("valid_code"), "")
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnOk(c)
	return
}
