package main

import (
	"github.com/krilie/lico_alone/application"
	"github.com/krilie/lico_alone/common/broker"
	"github.com/krilie/lico_alone/common/ccontext"
	"github.com/krilie/lico_alone/common/cdb"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/utils/time_util"
	broker2 "github.com/krilie/lico_alone/server/broker"
	"github.com/krilie/lico_alone/server/cron"
	"github.com/krilie/lico_alone/server/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:generate swag init -g ./main.go

// @title Swagger Example API
// @version 0.0.1
// @description  This is a sample server Petstore server.
func main() {
	ctx := ccontext.NewContext()
	clog.SetUpLogFile(config.Cfg.LogFile)
	var log = clog.NewLog(ctx, "lico.main", "main")
	cdb.StartDb(config.Cfg.DB)
	defer cdb.CloseDb()                                        // 最后关闭数据库
	defer func() { broker.Smq.Close(); log.Infof("消息队列退出") }() // 关闭消息队列
	app := application.NewApp(config.Cfg)
	// 初始化数据 权限账号等
	app.Init.InitData(ctx)
	// 加载所有权限
	app.User.UserService.MustAuthCacheLoadAll(ctx)
	// 注册所有消息处理句柄
	broker2.RegisterHandler(ctx, app)
	// 初始化定时任务
	cronStop := cron.InitAndStartCorn(ctx, app)
	// 最后初始化为开启http服务
	shutDown := http.InitAndStartHttpServer(app)
	// 发送上线邮件
	err := app.All.Message.SendEmail(ctx, "1197829331@qq.com", "app-server", "启动成功"+time.Now().Format(time_util.DefaultFormat))
	if err != nil {
		log.Error(err)
	}
	// 收到信号并关闭服务器
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	// Block until a signal is received.
	case s := <-c:
		log.Info("Got signal:", s) //Got signal: terminated
		if s == syscall.SIGINT || s == syscall.SIGTERM || s == syscall.SIGKILL || s == syscall.SIGHUP || s == syscall.SIGQUIT {
			// shutdown
			err := shutDown(30)
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
		}
	}
}
