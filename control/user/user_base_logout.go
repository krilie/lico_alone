package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct"
	"github.com/krilie/lico_alone/control/gin_util"
	"github.com/krilie/lico_alone/module_user_auth/user_base"
)

// /user/base/logout post
// jwtToken string
// 从登录信息中取jwttoken
func UserBaseLogout(c *gin.Context) {
	ctx := gin_util.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	logout := user_base.UserLogout(ctx, ctx.GetNowUserTokenOrEmpty())
	if logout != nil {
		gin_util.ReturnWithErr(ctx, c, logout)
		return
	} else {
		c.JSON(200, common_struct.StdSuccess)
		return
	}
}
