package control

import (
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico-my-site-user/common/common_struct"
	"github.com/lico603/lico-my-site-user/common/common_struct/errs"
	"github.com/lico603/lico-my-site-user/control/gin_util"
	"github.com/lico603/lico-my-site-user/user_auth"
)

/**
 *  /user/auth/has_permission get
 *  是不是有那个权限
 *  要求role base
 *  userId, permissionName string
 *  permission  参数要查的权限
 */
func UserAuthHasPermission(c *gin.Context) {
	ctx := gin_util.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	userId := ctx.GetUserIdOrEmpty()
	if userId == "" {
		gin_util.ReturnWithAppErr(ctx, c, errs.UnAuthorized.NewWithMsg("can not take login user id"))
		return
	}
	permission := c.Query("permission")
	if permission == "" {
		gin_util.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("no permission param received"))
		return
	}
	//
	b, e := user_auth.UserAuthHasPermission(ctx, userId, permission)
	if e != nil {
		gin_util.ReturnWithErr(ctx, c, e)
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
