package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// UserLogin 用户登录
// @Summary 用户登录
// @Description 用户用密码登录
// @Tags user_account
// @ID 用户登录
// @Produce  json
// @Param phone formData string true "用户手机号"
// @Param password formData string true "用户密码"
// @Success 200 {string} string "{jwt:"asb"}"
// @Failure 400 {object} cmodel.ErrorReturn
// @Failure 404 {object} cmodel.ErrorReturn
// @Failure 500 {object} cmodel.ErrorReturn
// @Router /v1/user/login [post]
func (a *UserCtrl) UserLogin(c *gin.Context) {
	jwt, err := a.AppUser.UserLogin(ginutil.MustGetAppCtx(c), c.PostForm("phone"), c.PostForm("password"), "")
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	c.JSON(200, gin.H{"jwt": jwt})
	return
}
