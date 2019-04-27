package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/gin_util"
	"github.com/krilie/lico_alone/module_user_auth/user_auth"
)

// /user/auth/client/acc_token POST
// 取到app角色用户的所有keys admin client 如果没有admin,检查client与用户id的一致性
// appUserId string
// user_id 要查询的用户id
func UserAuthClientAccToken(c *gin.Context) {
	ctx := gin_util.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	userId := c.PostForm("user_id")
	if userId == "" {
		gin_util.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("user_id is empty"))
		return
	}
	//取list
	list, err := user_auth.UserAuthClientAccToken(ctx, userId)
	if err != nil {
		gin_util.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, list)
		return
	}
}
