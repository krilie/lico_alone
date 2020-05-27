package http

import (
	"context"
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/krilie/lico_alone/docs"
	"github.com/krilie/lico_alone/server/http/middleware"
	"github.com/krilie/lico_alone/service"
	"github.com/prometheus/common/log"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func InitAndStartHttpServer(ctx context.Context, app *service.App, ctrl *Controllers) (shutDown func(waitSec time.Duration) error) {
	// 设置gin mode
	gin.SetMode(app.Cfg.GinMode)
	// 路径设置 根路径
	RootRouter := gin.Default() // logger recover
	// 跨域
	//var theCors = cors.DefaultConfig()
	//theCors.AllowAllOrigins = true
	//theCors.AddAllowHeaders("Content-Type", "X-CSRF-Token", "Authorization", "Token")
	//theCors.AllowCredentials = true
	//RootRouter.Use(cors.New(theCors))
	RootRouter.Use(Cors())
	// 静态文件 图片等
	RootRouter.StaticFile("/files", app.Cfg.FileSave.LocalFileSaveDir)
	// swagger + gzip压缩
	if app.Cfg.EnableSwagger {
		RootRouter.GET("/swagger/*any", gzip.Gzip(gzip.DefaultCompression), ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// 健康检查
	RootRouter.GET("health/", ctrl.healthCheckCtrl.Hello)
	RootRouter.GET("health/ping", ctrl.healthCheckCtrl.Ping)
	// 版本号
	RootRouter.GET("/version", Version(app.RunEnv.Version, app.RunEnv.BuildTime, app.RunEnv.GitCommit, app.RunEnv.GoVersion))
	// web 网页
	// web 站点
	webRouter := RootRouter.Group("/")
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
			if info.Name() == "index.html" {
				webRouter.StaticFile("/", "./www/"+info.Name())
			}
		}
	}
	// api路由 + 中间件
	apiGroup := RootRouter.Group("/api")
	apiGroup.Use(middleware.BuildContext())

	// 不检查权限的分组
	noCheckToken := apiGroup.Group("")
	noCheckToken.POST("/user/login", ctrl.userCtrl.UserLogin)
	noCheckToken.POST("/user/register", ctrl.userCtrl.UserRegister)
	noCheckToken.POST("/user/send_sms", ctrl.userCtrl.UserSendSms)
	//检查权限的分组
	checkToken := apiGroup.Group("")
	checkToken.Use(middleware.CheckAuthToken(app.UserService.GetAuthFace()))
	checkToken.GET("/manage/setting/get_setting_all", ctrl.userCtrl.ManageGetConfigList)
	checkToken.POST("/manage/setting/update_config", ctrl.userCtrl.ManageUpdateConfig)

	// common 服务
	commonApi := apiGroup.Group("")
	commonApi.GET("/common/icp_info", ctrl.commonCtrl.GetIcpInfo)

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
