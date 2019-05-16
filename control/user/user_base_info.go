package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/utils"
	"github.com/krilie/lico_alone/module/userbase/user"
)

// /user/base/info get
// get info
func UserBaseInfo(c *gin.Context) {
	ctx := utils.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	userId := ctx.GetUserIdOrEmpty()
	if userId == "" {
		common.ReturnWithAppErr(ctx, c, errs.UnAuthorized.NewWithMsg("can not take login user id"))
		return
	}
	info, err := user.UserBaseGetInfo(ctx, userId)
	if err != nil {
		common.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, info)
		return
	}
}
