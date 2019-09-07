package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// UserLogin 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags user_account
// @ID 用户注册
// @Produce  json
// @Param phone formData string true "用户手机号"
// @Param password formData string true "用户密码"
// @Success 200 {object} cmodel.StdReturn
// @Failure 400 {object} cmodel.ErrorReturn
// @Failure 404 {object} cmodel.ErrorReturn
// @Failure 500 {object} cmodel.ErrorReturn
// @Router /v1/user/register [post]
func (a *UserCtrl) UserRegister(c *gin.Context) {
	err := a.AppUser.UserRegister(ginutil.MustGetAppCtx(c), c.PostForm("phone"), c.PostForm("password"), "")
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnOk(c)
	return
}
