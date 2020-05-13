package main

import (
	"fmt"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/dig"
	"github.com/krilie/lico_alone/component/broker"
	"github.com/krilie/lico_alone/component/nlog"
	run_env "github.com/krilie/lico_alone/run_env"
	broker2 "github.com/krilie/lico_alone/server/broker"
	"github.com/krilie/lico_alone/server/cron"
	"github.com/krilie/lico_alone/server/http"
	"github.com/krilie/lico_alone/service"
	"os"
	"os/signal"
	"syscall"
)

//go:generate swag init -g ./main.go

// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
func main() {
	// 命令行 命令
	if len(os.Args) >= 2 {
		cmd := os.Args[1]
		switch cmd {
		case "version", "--version", "-version":
			fmt.Println(run_env.VERSION)
			return
		case "git-commit", "--git-commit", "-git-commit":
			fmt.Println(run_env.GIT_COMMIT)
			return
		case "go-version", "-go-version", "--go-version":
			fmt.Println(run_env.GO_VERSION)
			return
		case "build-time", "-build-time", "--build-time":
			fmt.Println(run_env.BUILD_TIME)
			return
		default:
			break
		}
	}
	// 开始服务
	dig.Container.MustInvoke(func(log *nlog.NLog, app *service.App) {
		ctx := context.NewContext()
		// 初始化日志文件
		defer func() {
			broker.Smq.Close()
			log.Infof("消息队列退出")
		}()
		// 初始化数据
		app.InitService.InitData(ctx)
		// 注册所有消息处理句柄
		broker2.RegisterHandler(ctx, app)
		// 初始化定时任务
		cronStop := cron.InitAndStartCorn(ctx, app)
		// 最后初始化为开启http服务
		shutDownApi := http.InitAndStartHttpServer(ctx, app)
		// 收尾工作
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		for {
			s := <-c
			log.Info("get a signal %s", s.String())
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL:
				err := shutDownApi(30)
				if err != nil {
					log.Errorln(err)
				} else {
					log.Infoln("service is closed normally")
				}
				// 关闭定时任务
				cronStop()
				log.Infoln("cron job end.")
				log.Infoln("service is done.")
				return
			case syscall.SIGHUP:
			default:
				return
			}
		}
	})
}
