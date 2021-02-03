package ctl_user

import (
	"github.com/gin-gonic/gin"
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
	ginWrap := ginutil.NewGinWrap(c, a.log)

	var data = a.userService.InitAppData(ginWrap.AppCtx)
	ginWrap.ReturnData(data)
	return
}
