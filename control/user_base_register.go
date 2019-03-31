package control

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lico603/lico_user/common/common_struct"
	"github.com/lico603/lico_user/common/common_struct/errs"
	"github.com/lico603/lico_user/control/gin_util"
	"github.com/lico603/lico_user/user_base"
)

// /user/base/register POST
// loginName string, password string
// login_name 登录名
// password 密码
func UserBaseRegister(c *gin.Context) {
	ctx := gin_util.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	req := struct {
		LoginName string `form:"login_name" binding:"required"`
		Password  string `form:"password" binding:"required"`
	}{}
	if e := c.ShouldBindWith(req, binding.FormPost); e != nil {
		gin_util.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg(e.Error()))
		return
	}
	//开始注册
	err := user_base.UserBaseRegister(ctx, req.LoginName, req.Password)
	if err != nil {
		gin_util.ReturnWithErr(ctx, c, err)
		return
	} else {
		// 可以带一次登录过程
		// jwtString, err := user_base.UserLogin(ctx, req.LoginName, req.Password)
		c.JSON(200, common_struct.StdSuccess)
		return
	}
}
