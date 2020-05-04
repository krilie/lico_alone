package user

import (
	"github.com/gin-gonic/gin"
	com_model "github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// UserLogin 用户登录
// @Summary 用户登录
// @Description 用户用密码登录
// @Tags 用户账号
// @ID 用户登录
// @Produce  json
// @Param phone formData string true "用户手机号"
// @Param password formData string true "用户密码"
// @Success 200 {string} com_model.CommonReturn "2000 {token:"asb"}"
// @Failure 500 {object} com_model.CommonReturn
// @Router /api/user/login [post]
func (a *UserCtrl) UserLogin(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	jwt, err := a.userService.UserLogin(ginutil.MustGetAppCtx(c), phone, password, "")
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	c.JSON(200, com_model.NewSuccess(gin.H{"token": jwt}))
	return
}
