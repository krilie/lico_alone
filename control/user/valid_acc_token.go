package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/control/utils"
)

// /user/base/valid_client_acc_token POST
// token
// 无权限
func (UserCtrl) ValidClientAccToken(c *gin.Context) {
	ctx := utils.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	token := c.PostForm("token")
	if token == "" {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg("no find token param in form"))
		return
	}
	key, err := appUser.ValidateClientAccToken(ctx, token)
	if err != nil {
		utils.ReturnWithErr(ctx, c, err)
		return
	} else {
		c.JSON(200, key)
		return
	}
}
