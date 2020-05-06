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

func (common *CommonCtrl) GetIcpInfo(c *gin.Context) {
	info := common.CommonService.GetIcpInfo(ginutil.MustGetAppCtx(c))
	ginutil.ReturnData(c, info)
}
