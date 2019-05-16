package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/utils"
	"github.com/krilie/lico_alone/module/userbase/auth"
)

/**
 *  /user/auth/has_permission get
 *  是不是有那个权限
 *  要求role base
 *  userId, permissionName string
 *  permission  参数要查的权限
 */
func UserAuthHasPermission(c *gin.Context) {
	ctx := common.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	userId := ctx.GetUserIdOrEmpty()
	if userId == "" {
		common.ReturnWithAppErr(ctx, c, errs.UnAuthorized.NewWithMsg("can not take login user id"))
		return
	}
	permission := c.Query("permission")
	if permission == "" {
		common.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("no permission param received"))
		return
	}
	//
	b, e := auth.UserAuthHasPermission(ctx, userId, permission)
	if e != nil {
		common.ReturnWithErr(ctx, c, e)
		return
	} else {
		if b {
			c.JSON(200, &common_struct.StdReturn{Code: "success", Message: "find"})
			return
		} else {
			c.JSON(404, &common_struct.StdReturn{Code: "noThisPermission", Message: "not find"})
			return
		}
	}
}
