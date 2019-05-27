package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/krilie/lico_alone/common/comstruct"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/control/utils"
)

// /user/base/register POST
// loginName string, password string
// login_name 登录名
// password 密码
func (UserCtrl) Register(c *gin.Context) {
	ctx := utils.GetApplicationContextOrReturn(c)
	if ctx == nil {
		return
	}
	req := struct {
		LoginName string `form:"login_name" binding:"required"`
		Password  string `form:"password" binding:"required"`
	}{}
	if e := c.ShouldBindWith(&req, binding.FormPost); e != nil {
		utils.ReturnWithAppErr(ctx, c, errs.ErrParam.NewWithMsg(e.Error()))
		return
	}
	//开始注册
	err := apiUser.Register(ctx, req.LoginName, req.Password)
	if err != nil {
		utils.ReturnWithErr(ctx, c, err)
		return
	} else {
		// 可以带一次登录过程
		// jwtString, err := user_base.Login(ctx, req.LoginName, req.Password)
		c.JSON(200, comstruct.StdSuccess)
		return
	}
}
