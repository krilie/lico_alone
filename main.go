package main

import (
	"github.com/krilie/lico_alone/application"
	"github.com/krilie/lico_alone/common/ccontext"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/server/cron"
	"github.com/krilie/lico_alone/server/http"
	"os"
	"os/signal"
	"syscall"
)

// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
// @BasePath http:127.0.0.1
func main() {
	ctx := ccontext.NewContext()
	var log = clog.NewLog(ctx, "lico.main", "main")
	defer cdb.Close()
	app := application.NewApp(config.Cfg)
	// 初始化数据 权限账号等
	app.Init.InitData(ctx)
	// 加载所有权限
	app.User.UserService.AuthCacheLoadAll(ctx)
	// 初始化为开启http服务
	shutDown := http.InitAndStartHttpServer(app)
	// 初始化定时任务
	cronStop := cron.InitAndStartCorn(app)
	// 收到信号并关闭服务器
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	// Block until a signal is received.
	case s := <-c:
		log.Info("Got signal:", s) //Got signal: terminated
		if s == syscall.SIGINT || s == syscall.SIGTERM || s == syscall.SIGKILL || s == syscall.SIGHUP || s == syscall.SIGQUIT {
			// 关闭定时任务
			cronStop()
			log.Infoln("cron job end.")
			// shutdown
			err := shutDown(30)
			if err != nil {
				log.Errorln(err)
			} else {
				log.Infoln("service is closed normally")
			}
			log.Infoln("service is done.")
			return
		}
	}
}
