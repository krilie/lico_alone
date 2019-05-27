package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/control/utils"
)

// /user/base/valid POST
// token 用户的jwttoken
// 检查这个jwtToken是否有效，并返回有效载荷
func (UserCtrl) Valid(c *gin.Context) {
	ctx := utils.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	token := c.PostForm("token")
	if token == "" {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("no find token param in form"))
		return
	}
	claims, e := apiUser.Validate(ctx, token)
	if e != nil {
		utils.ReturnWithErr(ctx, c, e)
		return
	} else {
		c.JSON(200, claims)
		return
	}
}
