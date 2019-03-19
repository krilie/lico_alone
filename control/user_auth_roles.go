package control

import (
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico-my-site-user/common/common_struct/errs"
	"github.com/lico603/lico-my-site-user/control/gin_util"
	"github.com/lico603/lico-my-site-user/user_auth"
)

// /user/auth/roles get
func UserAuthRoles(c *gin.Context) {
	ctx := gin_util.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	userId := ctx.GetUserIdOrEmpty()
	if userId == "" {
		gin_util.ReturnWithAppErr(ctx, c, errs.UnAuthorized.NewWithMsg("can not take login user id"))
		return
	}
	set, err := user_auth.UserAuthRoles(ctx, userId)
	if err != nil {
		gin_util.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, set.ToSlice())
		return
	}
}
