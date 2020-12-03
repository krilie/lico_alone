package ctl_common

import (
	"github.com/gin-gonic/gin"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/run_env"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/nlog"
	common_service "github.com/krilie/lico_alone/module/service-common"
	"github.com/krilie/lico_alone/server/http/ginutil"
	"github.com/krilie/lico_alone/server/http/middleware"
)

type CommonCtrl struct {
	CommonService *common_service.CommonService
	runEnv        *run_env.RunEnv
	log           *nlog.NLog
	middleware    *middleware.GinMiddleware
	ginUtil       *ginutil.GinUtils
}

func NewCommonCtrl(
	log *nlog.NLog,
	common *common_service.CommonService,
	middleware *middleware.GinMiddleware,
	ginUtil *ginutil.GinUtils,
	cfg *ncfg.NConfig) *CommonCtrl {

	log = log.WithField(context_enum.Module.Str(), "common controller")
	return &CommonCtrl{
		CommonService: common,
		runEnv:        cfg.RunEnv,
		log:           log,
		middleware:    middleware,
		ginUtil:       ginUtil,
	}
}

// Health Icp信息
// @Summary Icp信息
// @Description Icp信息
// @Tags 公共接口
// @ID Icp信息
// @Success 200 {object} com_model.CommonReturn{data=model.IcpInfo}
// @Success 500 {object} com_model.CommonReturn
// @Router /api/common/icp_info [get]
func (con *CommonCtrl) GetIcpInfo(c *gin.Context) {
	info := con.CommonService.GetIcpInfo(con.ginUtil.MustGetAppContext(c))
	ginutil.ReturnData(c, info)
}

// UserLogin Version
// @Summary Version
// @Description Version
// @Tags 公共接口
// @ID Version
// @Success 200 {string} string "version build_time git_commit go_version"
// @Failure 500 {string} string ""
// @Router /api/common/version [get]
func (con *CommonCtrl) Version(c *gin.Context) {
	c.JSON(200, gin.H{
		"version":    con.runEnv.Version,
		"build_time": con.runEnv.BuildTime,
		"git_commit": con.runEnv.GitCommit,
		"go_version": con.runEnv.GoVersion,
		"host":       con.runEnv.AppHost,
	})
}

// WebVisited WebVisited
// @Summary WebVisited
// @Description WebVisited
// @Tags 公共接口
// @ID WebVisited
// @Success 200 {object} com_model.CommonReturn
// @Success 500 {object} com_model.CommonReturn
// @Router /api/common/visited [post]
func (con *CommonCtrl) WebVisited(c *gin.Context) {
	ctx := con.ginUtil.MustGetAppContext(c)
	values := context.MustGetAppValues(ctx)
	con.CommonService.WebVisited(ctx, values.RemoteIp, values.CustomerTraceId)
	ginutil.ReturnOk(c)
}

// AboutApp AboutApp
// @Summary AboutApp
// @Description AboutApp
// @Tags 公共接口
// @ID AboutApp
// @Success 200 {object} com_model.CommonReturn{data=string}
// @Success 500 {object} com_model.CommonReturn
// @Router /api/common/about_app [get]
func (con *CommonCtrl) AboutApp(c *gin.Context) {
	ctx := con.ginUtil.MustGetAppContext(c)
	app, err := con.CommonService.GetAboutApp(ctx)
	ginutil.HandlerErrorOrReturnData(c, err, app)
}
