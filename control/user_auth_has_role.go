package control

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/gin_util"
	"github.com/krilie/lico_alone/module_user_auth/user_auth"
)

// /user/auth/has_role get
// role 参数
//
func UserAuthHasRole(c *gin.Context) {
	ctx := gin_util.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	userId := ctx.GetUserIdOrEmpty()
	if userId == "" {
		gin_util.ReturnWithAppErr(ctx, c, errs.UnAuthorized.NewWithMsg("can not take login user id"))
		return
	}
	role := c.Query("role")
	if role == "" {
		gin_util.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("no role param received"))
		return
	}
	//
	b, e := user_auth.UserAuthHasRole(ctx, userId, role)
	if e != nil {
		gin_util.ReturnWithErr(ctx, c, e)
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
