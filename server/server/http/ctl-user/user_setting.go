package ctl_user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// ManageGetConfigList 获取所有配置项
// @Summary 获取所有配置项
// @Description 获取所有配置项
// @Tags 配置项
// @ID 获取所有配置项
// @Produce  json
// @Param searchKey query string true "搜索值"
// @Success 200 {object} com_model.CommonReturn{data=[]model.config}
// @Failure 500 {string} errInfo
// @Router /api/manage/setting/get_setting_all [post]
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

// ManageUpdateConfig 更新配置项
// @Summary 更新配置项
// @Description 更新配置项
// @Tags 配置项
// @ID 更新配置项
// @Produce  json
// @Param name formData string true "配置名"
// @Param value formData string true "配置值"
// @Success 200 {object} com_model.CommonReturn
// @Failure 500 {string} errInfo
// @Router /api/manage/setting/update_config [post]
func (a *UserCtrl) ManageUpdateConfig(c *gin.Context) {
	name := c.Query("name")
	value := c.Query("value")
	err := a.userService.UpdateConfig(ginutil.MustGetAppCtx(c), name, value)
	if err != nil {
		ginutil.ReturnWithErr(c, err)
		return
	}
	ginutil.ReturnOk(c)
	return
}
