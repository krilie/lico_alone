package ctl_common

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/server/http/ginutil"
	common_service "github.com/krilie/lico_alone/service/common-service"
)

type CommonCtrl struct {
	CommonService *common_service.CommonService
}

func NewCommonCtrl(common *common_service.CommonService) *CommonCtrl {
	return &CommonCtrl{CommonService: common}
}

// Health Icp信息
// @Summary Icp信息
// @Description Icp信息
// @Tags Icp信息
// @ID Icp信息
// @Success 200 {string} model.IcpInfo "pong start time up time"
// @Router /api/common/icp_info [post]
func (common *CommonCtrl) GetIcpInfo(c *gin.Context) {
	info := common.CommonService.GetIcpInfo(ginutil.MustGetAppCtx(c))
	ginutil.ReturnData(c, info)
}
