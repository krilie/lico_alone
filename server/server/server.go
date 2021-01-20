package server

import (
	"context"
	"github.com/krilie/lico_alone/component/broker"
	cron2 "github.com/krilie/lico_alone/component/cron"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/component/nlog/logcat"
	"github.com/krilie/lico_alone/server/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.HttpService
	cfg        *ncfg.NConfig
	log        *nlog.NLog
	broker     *broker.Broker
	nCron      *cron2.NCron
	db         *ndb.NDb
}

func NewServer(httpServer *http.HttpService, cfg *ncfg.NConfig, log *nlog.NLog, broker *broker.Broker, nCron *cron2.NCron, db *ndb.NDb) *Server {
	return &Server{httpServer: httpServer, cfg: cfg, log: log, broker: broker, nCron: nCron, db: db}
}

func (a *Server) StartService(ctx context.Context) {

	var logCfg = a.cfg.GetLogCfg()

	closeLogLimit := logcat.BeginLogFileLimit(1024, logCfg.LogFile, time.Hour*8)
	defer closeLogLimit()
	defer a.log.CloseAndWait(ctx)
	defer a.db.CloseDb()
	defer a.nCron.StopAndWait(ctx)
	defer func() { a.broker.Close(); a.log.Get(ctx).Infof("消息队列退出") }()

	// 最后初始化为开启http服务
	shutDownApi := a.httpServer.InitAndStartHttpService(ctx)
	// 收尾工作
	WaitSignalAndExit(ctx, func() {
		err := shutDownApi(time.Second * 30)
		if err != nil {
			a.log.Get(ctx).WithError(err).Errorln("service is closed with err")
		} else {
			a.log.Get(ctx).Infoln("service is closed gracefully")
		}
		return
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

var DigServerProviderAll = []interface{}{
	http.NewHttpService,
	NewServer,
}
