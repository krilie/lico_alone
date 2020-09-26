package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/component/ncfg"
	_ "github.com/krilie/lico_alone/docs"
	"github.com/krilie/lico_alone/server/http/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func InitAndStartHttpServer(ctx context.Context, cfg *ncfg.NConfig, auth middleware.IAuth, ctrl *Controllers) (shutDown func(waitSec time.Duration) error) {
	httpCfg := &cfg.Cfg.Http
	fileCfg := &cfg.Cfg.FileSave
	// 设置gin mode
	gin.SetMode(httpCfg.GinMode)
	// 路径设置 根路径
	rootRouter := gin.Default()                  // logger recover
	rootRouter.Use(middleware.RequestOpsLimit()) // 限流

	pprof.Register(rootRouter, "pprof") // 性能
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
	rootRouter.GET("health/", ctrl.healthCheckCtrl.Hello)
	rootRouter.GET("health/ping", ctrl.healthCheckCtrl.Ping)

	// web 网页
	InitStaticWeb(ctx, rootRouter)

	// api路由 + 中间件
	apiGroup := rootRouter.Group("/api")
	apiGroup.Use(middleware.BuildContext())

	// 不检查权限的分组
	noCheckToken := apiGroup.Group("")
	noCheckToken.POST("/user/login", ctrl.userCtrl.UserLogin)
	noCheckToken.POST("/user/register", ctrl.userCtrl.UserRegister)
	noCheckToken.POST("/user/send_sms", ctrl.userCtrl.UserSendSms)

	//检查权限的分组
	checkToken := apiGroup.Group("")
	checkToken.Use(middleware.CheckAuthToken(auth))
	checkToken.GET("/user/init_app", ctrl.userCtrl.InitApp)
	checkToken.GET("/manage/setting/get_setting_all", ctrl.userCtrl.ManageGetConfigList)
	checkToken.POST("/manage/setting/update_config", ctrl.userCtrl.ManageUpdateConfig)
	checkToken.GET("/manage/setting/get_a_map_key", ctrl.userCtrl.ManageGetAMapKey) // 高德地图 获取配置key
	checkToken.GET("/manage/article/query", ctrl.userCtrl.QueryArticle)
	checkToken.GET("/manage/article/get_by_id", ctrl.userCtrl.GetArticleById)
	checkToken.POST("/manage/article/update", ctrl.userCtrl.UpdateArticle)
	checkToken.POST("/manage/article/delete", ctrl.userCtrl.DeleteArticle)
	checkToken.POST("/manage/article/create", ctrl.userCtrl.CreateArticle)
	checkToken.POST("/manage/file/upload", middleware.OpsLimit(1), ctrl.userCtrl.UploadFile)
	checkToken.POST("/manage/file/delete", ctrl.userCtrl.DeleteFile)
	checkToken.GET("/manage/file/query", ctrl.userCtrl.QueryFile)
	checkToken.GET("/manage/carousel/query", ctrl.userCtrl.QueryCarousel)
	checkToken.POST("/manage/carousel/create", ctrl.userCtrl.CreateCarousel)
	checkToken.POST("/manage/carousel/update", ctrl.userCtrl.UpdateCarousel)
	checkToken.POST("/manage/carousel/delete_by_id", ctrl.userCtrl.DeleteCarouselById)
	checkToken.GET("/manage/statistic/get_visitor_points", ctrl.userCtrl.ManageGetVisitorPoints)

	// common 服务
	commonApi := apiGroup.Group("")
	commonApi.GET("/common/icp_info", ctrl.commonCtrl.GetIcpInfo)
	commonApi.GET("/common/article/query_sample", ctrl.commonCtrl.QueryArticleSample)
	commonApi.GET("/common/article/get_article", ctrl.commonCtrl.GetArticle)
	commonApi.GET("/common/carousel/query", ctrl.commonCtrl.QueryCarousel)
	commonApi.GET("/common/version", ctrl.commonCtrl.Version)     // 版本号
	commonApi.POST("/common/visited", ctrl.commonCtrl.WebVisited) // 版本号

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

func InitStaticWeb(ctx context.Context, rootRouter *gin.Engine) {
	// web 网页
	webRouter := rootRouter.Group("/")
	webRouter.Use(gzip.Gzip(gzip.DefaultCompression)) // 开启gzip压缩
	dir, err := ioutil.ReadDir("./www")
	if err != nil {
		panic(err)
	}
	for _, info := range dir {
		if info.IsDir() {
			webRouter.Static("/"+info.Name(), "./www/"+info.Name())
		} else {
			webRouter.StaticFile("/"+info.Name(), "./www/"+info.Name())
			// 主页
			if info.Name() == "index.html" {
				webRouter.StaticFile("/", "./www/"+info.Name())
			}
			// 文章详情
			if info.Name() == "article_detail.html" {
				webRouter.StaticFile("/article_detail", "./www/"+info.Name())
			}
			// 管理
			if info.Name() == "management.html" {
				webRouter.StaticFile("/management", "./www/"+info.Name())
			}
		}
	}
	// 重定向
	rootRouter.NoRoute(func(c *gin.Context) {
		if c.Request.Method != "GET" {
			c.String(404, "page not found")
			return
		}
		path := c.Request.URL.Path
		prefix := []string{"/api", "/files", "/swagger", "/health", "/version"}
		for i := range prefix {
			if strings.HasPrefix(path, prefix[i]) {
				c.String(404, "page not found")
				return
			}
		}
		if strings.HasPrefix(path, "/article_detail") {
			c.Request.URL.Path = "/article_detail"
		} else if strings.HasPrefix(path, "/management") {
			c.Request.URL.Path = "/management"
		} else {
			c.Request.URL.Path = "/"
		}
		rootRouter.HandleContext(c)
		return
	})
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
