package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/utils"
	"github.com/krilie/lico_alone/module/userbase/auth"
)

// /user/auth/client/acc_token POST
// 取到app角色用户的所有keys admin client 如果没有admin,检查client与用户id的一致性
// appUserId string
// user_id 要查询的用户id
func UserAuthClientAccToken(c *gin.Context) {
	ctx := common.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	userId := c.PostForm("user_id")
	if userId == "" {
		common.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("user_id is empty"))
		return
	}
	//取list
	list, err := auth.UserAuthClientAccToken(ctx, userId)
	if err != nil {
		common.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, list)
		return
	}
}
