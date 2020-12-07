package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/arl/statsviz"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/component/ncfg"
	_ "github.com/krilie/lico_alone/docs"
	"github.com/krilie/lico_alone/server/http/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"strconv"
	"strings"
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
	// 性能
	pprof.Register(rootRouter, "pprof")
	// 性能
	mux := http.NewServeMux()
	statsviz.Register(mux)
	rootRouter.Any("statsviz", func(c *gin.Context) {
		mux.ServeHTTP(c.Writer, c.Request)
	})

	// 跨域
	rootRouter.Use(Cors())
	// 静态文件 图片等
	if fileCfg.Channel == "local" {
		rootRouter.StaticFile("/files", fileCfg.OssBucket)
	}
	// swagger + gzip压缩
	if httpCfg.EnableSwagger {
		rootRouter.GET("/swagger/*any", gzip.Gzip(gzip.DefaultCompression), ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 健康检查
	rootRouter.GET("health/", h.ctrl.healthCheckCtrl.Hello)
	rootRouter.GET("health/ping", h.ctrl.healthCheckCtrl.Ping)
	rootRouter.GET("health/panic", h.ctrl.healthCheckCtrl.Panic)

	// api路由 + 中间件
	apiGroup := rootRouter.Group("/api")
	apiGroup.Use(middleware.BuildContext())

	// 不检查权限的分组
	noCheckToken := apiGroup.Group("")
	noCheckToken.POST("/user/login", h.ctrl.userCtrl.UserLogin)
	noCheckToken.POST("/user/register", h.ctrl.userCtrl.UserRegister)
	noCheckToken.POST("/user/send_sms", h.ctrl.userCtrl.UserSendSms)

	//检查权限的分组
	checkToken := apiGroup.Group("")
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

	// common 服务
	commonApi := apiGroup.Group("")
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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//				允许跨域设置																										可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //	跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //	处理请求
	}
}
