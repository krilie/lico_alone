package control

import (
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico_user/common/common_struct/errs"
	"github.com/lico603/lico_user/control/gin_util"
	"github.com/lico603/lico_user/user_base"
)

// /user/base/valid POST
// token 用户的jwttoken
// 检查这个jwtToken是否有效，并返回有效载荷
func UserBaseValid(c *gin.Context) {
	ctx := gin_util.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	token := c.PostForm("token")
	if token == "" {
		gin_util.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("no find token param in form"))
		return
	}
	claims, e := user_base.UserValidate(ctx, token)
	if e != nil {
		gin_util.ReturnWithErr(ctx, c, e)
		return
	} else {
		c.JSON(200, claims)
		return
	}
}
