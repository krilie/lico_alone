package user_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/gin_util"
	"github.com/krilie/lico_alone/module/user_auth/user_auth"
)

// /user/auth/client/has_acc_token GET
// userId, accTokenStr string
// token 要传的token
// 用户是否有client acctoken 要有client权限，使用登录用户的id
func UserAuthClientHasAccToken(c *gin.Context) {
	ctx := gin_util.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	userId := ctx.GetUserIdOrEmpty()
	if userId == "" {
		gin_util.ReturnWithAppErr(ctx, c, errs.UnAuthorized.NewWithMsg("can not take login user id"))
		return
	}
	token := c.Query("token")
	if token == "" {
		gin_util.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("no token param received"))
		return
	}
	//调用接口方法取结果
	accessToken, err := user_auth.UserAuthClientHasAccToken(ctx, userId, token)
	if err != nil {
		gin_util.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, accessToken)
		return
	}
}
