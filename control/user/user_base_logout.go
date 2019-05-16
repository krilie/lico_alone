package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct"
	"github.com/krilie/lico_alone/control/utils"
	"github.com/krilie/lico_alone/module/userbase/user"
)

// /user/base/logout post
// jwtToken string
// 从登录信息中取jwttoken
func UserBaseLogout(c *gin.Context) {
	ctx := common.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	logout := user.UserLogout(ctx, ctx.GetNowUserTokenOrEmpty())
	if logout != nil {
		common.ReturnWithErr(ctx, c, logout)
		return
	} else {
		c.JSON(200, common_struct.StdSuccess)
		return
	}
}
