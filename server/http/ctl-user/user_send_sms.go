package ctl_user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// UserSendSms 用户发短信
// @Summary 用户发短信
// @Description 用户发短信
// @Tags 用户
// @ID 用户发短信
// @Produce  json
// @Param phone formData string true "用户手机号"
// @Param send_type formData string true "register login change_password"
// @Success 200 {object} com_model.CommonReturn{data=object}
// @Failure 500 {string} errInfo
// @Router /api/user/send_sms [post]
func (a *UserCtrl) UserSendSms(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	phone := c.PostForm("phone")
	if phone == "" {
		ginWrap.ReturnWithErr(errs.NewParamError().WithMsg("手机号不正确"))
		return
	}
	switch c.PostForm("send_type") {
	case "register":
		err := a.userService.SendRegisterSms(ginWrap.AppCtx, phone)
		if err != nil {
			ginWrap.ReturnWithErr(err)
			return
		}
		ginWrap.ReturnOk()
		return
	case "login", "change_password":
		ginWrap.ReturnWithErr(errs.NewNormal().WithMsg("未实现"))
		return
	default:
		ginWrap.ReturnWithErr(errs.NewParamError().WithMsg("未知发送类型"))
		return
	}
}
