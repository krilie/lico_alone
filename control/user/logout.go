package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/comstruct"
	"github.com/krilie/lico_alone/control/utils"
)

// /user/base/logout post
// jwtToken string
// 从登录信息中取jwttoken
func (UserCtrl) Logout(c *gin.Context) {
	ctx := utils.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	logout := appUser.Logout(ctx, ctx.GetNowUserTokenOrEmpty())
	if logout != nil {
		utils.ReturnWithErr(ctx, c, logout)
		return
	} else {
		c.JSON(200, comstruct.StdSuccess)
		return
	}
}
