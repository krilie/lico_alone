package control

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/gin_util"
	"github.com/krilie/lico_alone/user_base"
)

// /user/base/valid_client_acc_token POST
// token
// 无权限
func UserBaseValidClientAccToken(c *gin.Context) {
	ctx := gin_util.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	token := c.PostForm("token")
	if token == "" {
		gin_util.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("no find token param in form"))
		return
	}
	key, err := user_base.UserValidateClientAccToken(ctx, token)
	if err != nil {
		gin_util.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, key)
		return
	}
}
