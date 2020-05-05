package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// UserSendSms 用户发短信
// @Summary 用户发短信
// @Description 用户发短信
// @Tags 用户发短信
// @ID 用户发短信
// @Produce  json
// @Param phone formData string true "用户手机号"
// @Param send_type formData string true "register login change_password"
// @Success 200 {object} com_model.CommonReturn
// @Failure 500 {object} com_model.CommonReturn
// @Router /api/user/send_sms [post]
func (a *UserCtrl) UserSendSms(c *gin.Context) {
	phone := c.PostForm("phone")
	if phone == "" {
		ginutil.ReturnWithErr(c, errs.NewParamError().WithMsg("手机号不正确"))
		return
	}
	switch c.PostForm("send_type") {
	case "register":
		err := a.userService.SendRegisterSms(ginutil.MustGetAppCtx(c), phone)
		if err != nil {
			ginutil.ReturnWithErr(c, err)
			return
		}
		ginutil.ReturnOk(c)
		return
	case "login", "change_password":
		ginutil.ReturnWithErr(c, errs.NewNormal().WithMsg("未实现"))
		return
	default:
		ginutil.ReturnWithErr(c, errs.NewParamError().WithMsg("未知发送类型"))
		return
	}
}
