package ctl_user

import (
	"github.com/gin-gonic/gin"
	service_user "github.com/krilie/lico_alone/module/service-user"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// InitApp 用户初始化
// @Summary 用户初始化
// @Description 用户初始化数据
// @Tags 用户
// @ID 用户初始化
// @Produce  json
// @Param Authorization header string true "凭证token" default({{token}})
// @Success 200 {object} com_model.CommonReturn{data=service_user.InitAppData}
// @Failure 500 {string} errInfo
// @Router /api/user/init_app [get]
func (a *UserCtrl) InitApp(c *gin.Context) {
	var data *service_user.InitAppData = a.userService.InitAppData(a.ginUtil.MustGetAppValues(c))
	ginutil.ReturnData(c, data)
	return
}
