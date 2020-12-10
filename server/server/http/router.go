package http

import (
	"context"
	"errors"
	"github.com/arl/statsviz"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/ncfg"
	_ "github.com/krilie/lico_alone/docs"
	"github.com/krilie/lico_alone/server/http/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"strconv"
	"time"
)

type HttpService struct {
	cfg        *ncfg.NConfig
	ctrl       *Controllers
	middleware *middleware.GinMiddleware
}

func NewHttpService(cfg *ncfg.NConfig, ctrl *Controllers, middleware *middleware.GinMiddleware) *HttpService {
	return &HttpService{cfg: cfg, ctrl: ctrl, middleware: middleware}
}

func (h *HttpService) InitAndStartHttpService(ctx context.Context) (shutDown func(waitSec time.Duration) error) {
	httpCfg := &h.cfg.Cfg.Http
	fileCfg := &h.cfg.Cfg.FileSave
	// 设置gin mode
	gin.SetMode(httpCfg.GinMode)
	// 路径设置 根路径
	rootRouter := gin.Default() // logger recover
	rootRouter.Use(h.middleware.MiddlewareRecovery())
	rootRouter.Use(middleware.RequestOpsLimit()) // 限流

	{
		// 性能
		pprof.Register(rootRouter, "pprof")
		// 性能
		mux := http.NewServeMux()
		statsviz.Register(mux)
		rootRouter.Any("statsviz", func(c *gin.Context) {
			mux.ServeHTTP(c.Writer, c.Request)
		})
		// 静态文件 图片等
		if fileCfg.Channel == "local" {
			rootRouter.StaticFile("/files", fileCfg.OssBucket)
		}
		// swagger + gzip压缩
		if httpCfg.EnableSwagger {
			rootRouter.GET("/swagger/*any", gzip.Gzip(gzip.DefaultCompression), ginSwagger.WrapHandler(swaggerFiles.Handler))
		}
	}

	{
		// 健康检查
		rootRouter.GET("health/", h.ctrl.healthCheckCtrl.Hello)
		rootRouter.GET("health/ping", h.ctrl.healthCheckCtrl.Ping)
		rootRouter.GET("health/panic", h.ctrl.healthCheckCtrl.Panic)
	}

	// api路由 + 中间件
	apiGroup := rootRouter.Group("/api")
	apiGroup.Use(h.middleware.BuildContext())

	{
		manageGroup := apiGroup.Group("")
		origin, err := h.middleware.CfgService.GetValueStr(context2.EmptyAppCtx(), "access-control-allow-origin-m")
		if err != nil {
			origin = new(string)
			*origin = "*"
		}
		manageGroup.Use(h.middleware.Cors(*origin)) // 管理页面跨域
		// 不检查权限的分组
		noCheckToken := manageGroup.Group("")
		noCheckToken.POST("/user/login", h.ctrl.userCtrl.UserLogin)
		noCheckToken.POST("/user/register", h.ctrl.userCtrl.UserRegister)
		noCheckToken.POST("/user/send_sms", h.ctrl.userCtrl.UserSendSms)

		//检查权限的分组
		checkToken := manageGroup.Group("")
		checkToken.Use(h.middleware.CheckAuthToken())
		checkToken.GET("/user/init_app", h.ctrl.userCtrl.InitApp)
		checkToken.GET("/manage/setting/get_setting_all", h.ctrl.userCtrl.ManageGetConfigList)
		checkToken.POST("/manage/setting/update_config", h.ctrl.userCtrl.ManageUpdateConfig)
		checkToken.GET("/manage/setting/get_a_map_key", h.ctrl.userCtrl.ManageGetAMapKey) // 高德地图 获取配置key
		checkToken.GET("/manage/article/query", h.ctrl.userCtrl.QueryArticle)
		checkToken.GET("/manage/article/get_by_id", h.ctrl.userCtrl.GetArticleById)
		checkToken.POST("/manage/article/update", h.ctrl.userCtrl.UpdateArticle)
		checkToken.POST("/manage/article/delete", h.ctrl.userCtrl.DeleteArticle)
		checkToken.POST("/manage/article/create", h.ctrl.userCtrl.CreateArticle)
		checkToken.POST("/manage/file/upload", middleware.OpsLimit(1), h.ctrl.userCtrl.UploadFile)
		checkToken.POST("/manage/file/delete", h.ctrl.userCtrl.DeleteFile)
		checkToken.GET("/manage/file/query", h.ctrl.userCtrl.QueryFile)
		checkToken.GET("/manage/carousel/query", h.ctrl.userCtrl.QueryCarousel)
		checkToken.POST("/manage/carousel/create", h.ctrl.userCtrl.CreateCarousel)
		checkToken.POST("/manage/carousel/update", h.ctrl.userCtrl.UpdateCarousel)
		checkToken.POST("/manage/carousel/delete_by_id", h.ctrl.userCtrl.DeleteCarouselById)
		checkToken.GET("/manage/statistic/get_visitor_points", h.ctrl.userCtrl.ManageGetVisitorPoints)
	}

	{
		commonGroup := apiGroup.Group("")
		origin, err := h.middleware.CfgService.GetValueStr(context2.EmptyAppCtx(), "access-control-allow-origin-c")
		if err != nil {
			origin = new(string)
			*origin = "*"
		}
		commonGroup.Use(h.middleware.Cors(*origin))
		// common 服务
		commonApi := commonGroup.Group("")
		commonApi.GET("/common/icp_info", h.ctrl.commonCtrl.GetIcpInfo)
		commonApi.GET("/common/article/query_sample", h.ctrl.commonCtrl.QueryArticleSample)
		commonApi.GET("/common/article/get_article", h.ctrl.commonCtrl.GetArticle)
		commonApi.POST("/common/article/mark/like", h.ctrl.commonCtrl.MarkArticleLike)
		commonApi.POST("/common/article/mark/dislike", h.ctrl.commonCtrl.MarkArticleDisLike)
		commonApi.POST("/common/article/mark/remove_like", h.ctrl.commonCtrl.RemoveMarkArticleLike)
		commonApi.POST("/common/article/mark/remove_dislike", h.ctrl.commonCtrl.RemoveMarkArticleDisLike)
		commonApi.GET("/common/carousel/query", h.ctrl.commonCtrl.QueryCarousel)
		commonApi.GET("/common/version", h.ctrl.commonCtrl.Version)     // 版本号
		commonApi.POST("/common/visited", h.ctrl.commonCtrl.WebVisited) // WebVisited
		commonApi.GET("/common/about_app", h.ctrl.commonCtrl.AboutApp)  // AboutApp
	}

	// 开始服务
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(httpCfg.Port),
		Handler: rootRouter,
	}
	//是否有ssl.public_key ssl.private_key
	pubKey := httpCfg.SslPub
	priKey := httpCfg.SslPri
	if pubKey == "" || priKey == "" {
		go func() {
			if err := srv.ListenAndServe(); err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					return
				} else {
					panic(err.Error())
				}
				return
			}
		}()
	} else {
		go func() {
			if err := srv.ListenAndServeTLS(pubKey, priKey); err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					return
				} else {
					panic(err.Error())
				}
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
			return err
		} else {
			return nil
		}
	}
}
