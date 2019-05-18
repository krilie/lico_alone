package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct"
	"github.com/krilie/lico_alone/control/utils"
)

// /user/base/logout post
// jwtToken string
// 从登录信息中取jwttoken
func UserBaseLogout(c *gin.Context) {
	ctx := utils.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	logout := apiUser.Logout(ctx, ctx.GetNowUserTokenOrEmpty())
	if logout != nil {
		utils.ReturnWithErr(ctx, c, logout)
		return
	} else {
		c.JSON(200, common_struct.StdSuccess)
		return
	}
}
