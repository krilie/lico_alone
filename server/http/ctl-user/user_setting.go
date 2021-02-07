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
// @Param Authorization header string true "凭证token" default({{token}})
// @Param searchKey query string true "搜索值"
// @Success 200 {object} com_model.CommonReturn{data=[]model.Config}
// @Failure 500 {string} errInfo
// @Router /api/manage/setting/get_setting_all [get]
func (a *UserCtrl) ManageGetConfigList(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	searchKey := c.Query("searchKey")
	config, err := a.userService.GetAllConfig(ginWrap.AppCtx, searchKey)
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}
	ginWrap.ReturnData(config)
	return
}

// ManageUpdateConfig 更新配置项
// @Summary 更新配置项
// @Description 更新配置项
// @Tags 配置项
// @ID 更新配置项
// @Produce  json
// @Param Authorization header string true "凭证token" default({{token}})
// @Param name formData string true "配置名"
// @Param value formData string true "配置值"
// @Success 200 {object} com_model.CommonReturn{data=object}
// @Failure 500 {string} errInfo
// @Router /api/manage/setting/update_config [post]
func (a *UserCtrl) ManageUpdateConfig(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	name := c.PostForm("name")
	value := c.PostForm("value")
	err := a.userService.UpdateConfig(ginWrap.AppCtx, name, value)
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}
	ginWrap.ReturnOk()
	return
}

// ManageGetAMapKey 获取AMap配置项
// @Summary 获取AMap配置项
// @Description 获取AMap配置项
// @Tags 配置项
// @ID 获取AMap配置项
// @Produce  json
// @Param Authorization header string true "凭证token" default({{token}})
// @Success 200 {object} com_model.CommonReturn{data=object} => "data":{"a_map_key":"the a map key"}
// @Failure 500 {string} errInfo
// @Router /api/manage/setting/get_a_map_key [get]
func (a *UserCtrl) ManageGetAMapKey(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	key, err := a.userService.GetAMapKey(ginWrap.AppCtx)
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}
	ginWrap.ReturnData(gin.H{"a_map_key": key})
	return
}
