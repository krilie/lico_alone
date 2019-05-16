package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/utils"
	"github.com/krilie/lico_alone/module/userbase/auth"
)

// /user/auth/has_role get
// role 参数
//
func UserAuthHasRole(c *gin.Context) {
	ctx := common.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	userId := ctx.GetUserIdOrEmpty()
	if userId == "" {
		common.ReturnWithAppErr(ctx, c, errs.UnAuthorized.NewWithMsg("can not take login user id"))
		return
	}
	role := c.Query("role")
	if role == "" {
		common.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("no role param received"))
		return
	}
	//
	b, e := auth.UserAuthHasRole(ctx, userId, role)
	if e != nil {
		common.ReturnWithErr(ctx, c, e)
		return
	} else {
		if b {
			c.JSON(200, &common_struct.StdReturn{Code: "success", Message: "find"})
			return
		} else {
			c.JSON(404, &common_struct.StdReturn{Code: "noThisRole", Message: "not find"})
			return
		}
	}
}
