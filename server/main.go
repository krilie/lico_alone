package main

import (
	"context"
	"fmt"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	run_env "github.com/krilie/lico_alone/common/run_env"
	"github.com/krilie/lico_alone/component"
	"github.com/krilie/lico_alone/component/broker"
	cron2 "github.com/krilie/lico_alone/component/cron"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/component/nlog/logcat"
	"github.com/krilie/lico_alone/module/module"
	"github.com/krilie/lico_alone/module/module-user/service"
	service2 "github.com/krilie/lico_alone/module/service"
	"github.com/krilie/lico_alone/server/http"
	_ "go.uber.org/automaxprocs"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:generate swag init -g ./main.go
//go:generate go test -run Auto -v ./...

// @title lizo_alone
// @version 1.0.0
// @description  api docs for lizo_alone
// @license.name all right

// @host localhost:80
// @BasePath /

func main() {
	// dig config
	component.DigProvider()
	module.DigProviderModule()
	service2.DigProviderService()
	http.DigProviderController()
	// begin service
	dig.Container.MustInvoke(
		func(log *nlog.NLog,
			broker *broker.Broker,
			cfg *ncfg.NConfig,
			auth *service.UserModule,
			nCron *cron2.NCron,
			db *ndb.NDb,
			ctrl *http.Controllers) {

			ctx := context2.NewContext()
			ctx.Module = "main"
			ctx.Function = "main"

			closeLogLimit := logcat.BeginLogFileLimit(1024, cfg.Cfg.Log.LogFile, time.Hour*8)
			defer closeLogLimit()
			defer log.CloseAndWait(ctx)
			defer db.CloseDb()
			defer nCron.StopAndWait(ctx)
			defer func() { broker.Close(); log.Get(ctx).Infof("消息队列退出") }()

			// 最后初始化为开启http服务
			shutDownApi := http.InitAndStartHttpServer(ctx, cfg, auth, ctrl)
			// 收尾工作
			WaitSignalAndExit(ctx, func() {
				err := shutDownApi(time.Second * 30)
				if err != nil {
					log.Get(ctx).WithError(err).Errorln("service is closed with err")
				} else {
					log.Get(ctx).Infoln("service is closed gracefully")
				}
				return
			})
		})
}

// 接收信号和退出
func WaitSignalAndExit(ctx context.Context, exit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL:
			exit()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

// 命令行处理 没用
func CmdForInfo() (isHit bool) { // 命令行 命令
	if len(os.Args) >= 2 {
		cmd := os.Args[1]
		switch cmd {
		case "version", "--version", "-version":
			fmt.Println(run_env.VERSION)
			return true
		case "git-commit", "--git-commit", "-git-commit":
			fmt.Println(run_env.GIT_COMMIT)
			return true
		case "go-version", "-go-version", "--go-version":
			fmt.Println(run_env.GO_VERSION)
			return true
		case "build-time", "-build-time", "--build-time":
			fmt.Println(run_env.BUILD_TIME)
			return true
		default:
			return false
		}
	} else {
		return false
	}
}
