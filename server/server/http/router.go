package http

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/nlog"
	_ "github.com/krilie/lico_alone/docs"
	ctl_common "github.com/krilie/lico_alone/server/http/ctl-common"
	"github.com/krilie/lico_alone/server/http/ctl-health-check"
	"github.com/krilie/lico_alone/server/http/ctl-user"
	"github.com/krilie/lico_alone/server/http/middleware"
	"github.com/krilie/lico_alone/service"
	"github.com/prometheus/common/log"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"strconv"
	"time"
)

func InitAndStartHttpServer(ctx context.Context, app *service.App) (shutDown func(waitSec time.Duration) error) {
	// 设置gin mode
	gin.SetMode(app.Cfg.GinMode)
	// 路径设置 根路径
	RootRouter := gin.Default() // logger recover
	// 跨域
	RootRouter.Use(cors.Default())
	// 静态文件 图片等
	RootRouter.StaticFile("/files", app.Cfg.FileSave.LocalFileSaveDir)
	// swagger + gzip压缩
	if app.Cfg.EnableSwagger {
		RootRouter.GET("/swagger/*any", gzip.Gzip(gzip.DefaultCompression), ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// 健康检查
	ctl_health_check.Init(RootRouter)
	// 版本号
	RootRouter.GET("/version", Version(app.Version, app.BuildTime, app.GitCommit, app.GoVersion))
	// api路由 + 中间件
	apiGroup := RootRouter.Group("/api")
	apiGroup.Use(middleware.BuildContext())

	// 不检查权限的分组
	noCheckToken := apiGroup.Group("")
	userCtrl := ctl_user.NewUserCtrl(app.UserService)
	noCheckToken.POST("/user/login", userCtrl.UserLogin)
	noCheckToken.POST("/user/register", userCtrl.UserRegister)
	noCheckToken.POST("/user/send_sms", userCtrl.UserSendSms)

	//检查权限的分组
	checkToken := apiGroup.Group("")
	checkToken.Use(middleware.CheckAuthToken(app.UnionService.ModuleUser))
	checkToken.GET("/manage/setting/get_setting_all", userCtrl.ManageGetConfigList)

	// common 服务
	dig.Container.MustInvoke(func(commonCtl *ctl_common.CommonCtrl) {
		commonApi := apiGroup.Group("")
		commonApi.GET("/common/icp_info", commonCtl.GetIcpInfo)
	})

	// 开始服务
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(app.Cfg.HttpPort),
		Handler: RootRouter,
	}
	//是否有ssl.public_key ssl.private_key
	pubKey := app.Cfg.SslPub
	priKey := app.Cfg.SslPri
	if pubKey == "" || priKey == "" {
		go func() {
			if err := srv.ListenAndServe(); err != nil {
				log.Warnln(err)
				return
			}
		}()
	} else {
		go func() {
			if err := srv.ListenAndServeTLS(pubKey, priKey); err != nil {
				log.Warnln(err)
				return
			}
		}()
	}
	return func(waitDuration time.Duration) error {
		ctxTimeout, cancelFunc := context.WithTimeout(ctx, waitDuration)
		defer cancelFunc()
		// shutdown
		err := srv.Shutdown(ctxTimeout)
		if err != nil {
			log.Error(err)
			return err
		} else {
			log.Info("end of service...")
			return nil
		}
	}
}

func InitAndStartStaticWebServer(ctx context.Context, cfg config.Config) (shutDown func(waitSec time.Duration) error) {
	log := nlog.Log.NewLog(ctx, "controller.router", "InitAndStartStaticWebServer")
	// 设置gin mode
	gin.SetMode(cfg.GinMode)
	// 路径设置 根路径
	RootRouter := gin.New()
	RootRouter.Use(gin.Logger(), gin.Recovery())
	// web 站点
	webRouter := RootRouter.Group("/")
	webRouter.Use(gzip.Gzip(gzip.DefaultCompression)) // 开启gzip压缩
	webRouter.Static("/", "./www")
	// 开始服务
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.WebPort),
		Handler: RootRouter,
	}
	//是否有ssl.public_key ssl.private_key
	pubKey := cfg.SslPub
	priKey := cfg.SslPri
	if pubKey == "" || priKey == "" {
		go func() {
			if err := srv.ListenAndServe(); err != nil {
				log.Warnln(err)
				return
			}
		}()
	} else {
		go func() {
			if err := srv.ListenAndServeTLS(pubKey, priKey); err != nil {
				log.Warnln(err)
				return
			}
		}()
	}
	return func(waitSec time.Duration) error {
		ctxTimeout, cancelFunc := context.WithTimeout(ctx, waitSec)
		defer cancelFunc()
		// shutdown
		err := srv.Shutdown(ctxTimeout)
		if err != nil {
			log.Error(err)
			return err
		} else {
			log.Info("end of service...")
			return nil
		}
	}
}

// UserLogin Version
// @Summary Version
// @Description Version
// @Tags 基本信息
// @ID Version
// @Success 200 {string} string "version build_time git_commit go_version"
// @Failure 500 {string} string ""
// @Router /version [get]
func Version(version, buildTime, gitCommit, goVersion string) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version":    version,
			"build_time": buildTime,
			"git_commit": gitCommit,
			"go_version": goVersion,
		})
	}
}
