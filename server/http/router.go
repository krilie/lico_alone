package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/application"
	"github.com/krilie/lico_alone/common/ccontext"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/config"
	_ "github.com/krilie/lico_alone/docs"
	"github.com/krilie/lico_alone/server/http/health"
	"github.com/krilie/lico_alone/server/http/middleware"
	"github.com/krilie/lico_alone/server/http/user"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"strconv"
	"time"
)

var RootRouter *gin.Engine

func InitAndStartHttpServer(app *application.App) (shutDown func(waitSec time.Duration) error) {
	ctx := ccontext.NewContext()
	log := clog.NewLog(ctx, "controller.router", "InitHttpServer")

	// 路径设置
	RootRouter = gin.Default() // logger recover
	// 静态文件
	RootRouter.StaticFile("/static", config.Cfg.FileSave.LocalFileSaveDir)
	// swagger
	if config.Cfg.EnableSwagger {
		RootRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// 健康检查
	health.Init(RootRouter)
	// 全局中间件
	RootRouter.Use(middleware.BuildContext())

	noCheckToken := RootRouter.Group("")
	//checkToken :=RootRouter.Group("").Use(middleware.CheckAuthToken(app.User))

	userCtrl := user.NewUserCtrl(app)
	noCheckToken.POST("/v1/user/login", userCtrl.UserLogin)
	noCheckToken.POST("/v1/user/register", userCtrl.UserRegister)

	// 开始服务
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.Cfg.HttpPort),
		Handler: RootRouter,
	}
	//是否有ssl.public_key ssl.private_key
	pubKey := config.Cfg.SslPub
	priKey := config.Cfg.SslPri
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
				log.Panicln(err)
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
