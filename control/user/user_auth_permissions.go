package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/utils"
	"github.com/krilie/lico_alone/module/userbase/auth"
)

// /user/auth/permissions
//
func UserAuthPermissions(c *gin.Context) {
	ctx := common.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	userId := ctx.GetUserIdOrEmpty()
	if userId == "" {
		common.ReturnWithAppErr(ctx, c, errs.UnAuthorized.NewWithMsg("can not take login user id"))
		return
	}
	set, err := auth.UserAuthPermissions(ctx, userId)
	if err != nil {
		common.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, set.ToSlice())
		return
	}
}
