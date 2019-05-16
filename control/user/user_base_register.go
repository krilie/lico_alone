package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/krilie/lico_alone/common/common_struct"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/control/utils"
	"github.com/krilie/lico_alone/module/userbase/user"
)

// /user/base/register POST
// loginName string, password string
// login_name 登录名
// password 密码
func UserBaseRegister(c *gin.Context) {
	ctx := common.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	req := struct {
		LoginName string `form:"login_name" binding:"required"`
		Password  string `form:"password" binding:"required"`
	}{}
	if e := c.ShouldBindWith(req, binding.FormPost); e != nil {
		common.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg(e.Error()))
		return
	}
	//开始注册
	err := user.UserBaseRegister(ctx, req.LoginName, req.Password)
	if err != nil {
		common.ReturnWithErr(ctx, c, err)
		return
	} else {
		// 可以带一次登录过程
		// jwtString, err := user_base.UserLogin(ctx, req.LoginName, req.Password)
		c.JSON(200, common_struct.StdSuccess)
		return
	}
}
