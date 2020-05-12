package ctl_user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// UserLogin 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags 用户
// @ID 用户注册
// @Produce  json
// @Param searchKey query string true "验证码"
// @Success 200 {object} com_model.CommonReturn{data=model.config}
// @Failure 500 {string} errInfo
// @Router /api/manage/get_setting_all [post]
func (a *UserCtrl) ManageGetConfigList(c *gin.Context) {
	searchKey := c.Query("searchKey")
	config, err := a.userService.GetAllConfig(ginutil.MustGetAppCtx(c), searchKey)
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnData(c, config)
	return
}
