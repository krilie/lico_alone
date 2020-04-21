package user

import (
	"github.com/gin-gonic/gin"
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
// @Success 200 {string} string "{jwt:"asb"}"
// @Failure 400 {object} cmodel.CommonReturn
// @Failure 404 {object} cmodel.CommonReturn
// @Failure 500 {object} cmodel.CommonReturn
// @Router /api/user/login [post]
func (a *UserCtrl) UserLogin(c *gin.Context) {
	jwt, err := a.AppUser.UserLogin(ginutil.MustGetAppCtx(c), c.PostForm("phone"), c.PostForm("password"), "")
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	c.JSON(200, gin.H{"jwt": jwt})
	return
}
