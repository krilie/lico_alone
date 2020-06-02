package ctl_common

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/run_env"
	common_service "github.com/krilie/lico_alone/module/service-common"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

type CommonCtrl struct {
	CommonService *common_service.CommonService
	runEnv        *run_env.RunEnv
}

func NewCommonCtrl(common *common_service.CommonService, env *run_env.RunEnv) *CommonCtrl {
	return &CommonCtrl{CommonService: common, runEnv: env}
}

// Health Icp信息
// @Summary Icp信息
// @Description Icp信息
// @Tags 公共接口
// @ID Icp信息
// @Success 200 {object} model.IcpInfo
// @Success 400 {string} errInfo
// @Success 500 {string} errInfo
// @Router /api/common/icp_info [get]
func (common *CommonCtrl) GetIcpInfo(c *gin.Context) {
	info := common.CommonService.GetIcpInfo(ginutil.MustGetAppCtx(c))
	ginutil.ReturnData(c, info)
}

// UserLogin Version
// @Summary Version
// @Description Version
// @Tags 公共接口
// @ID Version
// @Success 200 {string} string "version build_time git_commit go_version"
// @Failure 500 {string} string ""
// @Router /version [get]
func (common *CommonCtrl) Version(c *gin.Context) {
	c.JSON(200, gin.H{
		"version":    common.runEnv.Version,
		"build_time": common.runEnv.BuildTime,
		"git_commit": common.runEnv.GitCommit,
		"go_version": common.runEnv.GoVersion,
	})
}
