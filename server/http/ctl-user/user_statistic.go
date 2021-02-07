package ctl_user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// ManageGetVisitorPoints 获取所有访问地点
// @Summary 获取所有访问地点
// @Description 获取所有访问地点
// @Tags 配置项
// @ID 获取所有访问地点
// @Produce  json
// @Param Authorization header string true "凭证token" default({{token}})
// @Success 200 {object} com_model.CommonReturn{data=[]model.VisitorLonlatModel}
// @Failure 500 {string} errInfo
// @Router /api/manage/statistic/get_visitor_points [get]
func (a *UserCtrl) ManageGetVisitorPoints(c *gin.Context) {
	ginWrap := ginutil.NewGinWrap(c, a.log)

	point, err := a.userService.GetAllVisitorPoint(ginWrap.AppCtx)
	if err != nil {
		ginWrap.ReturnWithErr(err)
		return
	}
	ginWrap.ReturnData(point)
	return
}
