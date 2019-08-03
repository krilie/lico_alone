package main

import (
	"github.com/krilie/lico_alone/common/comlog"
	"github.com/krilie/lico_alone/common/config"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/db"
	"github.com/krilie/lico_alone/control"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var log = comlog.NewLog(context2.NewContext(), "lico.main", "main")

func main() {
	defer db.Close()
	// 开始
	srv := &http.Server{
		Addr:    ":" + config.GetString("service.port"),
		Handler: control.LocalRouter,
	}
	//是否有ssl.public_key ssl.private_key
	pubKey := config.GetString("ssl.public_key")
	priKey := config.GetString("ssl.private_key")
	if pubKey == "" || priKey == "" {
		go func() {
			if err := srv.ListenAndServe(); err != nil {
				log.Warningln(err)
				return
			}
		}()
	} else {
		go func() {
			if err := srv.ListenAndServeTLS(pubKey, priKey); err != nil {
				log.Warningln(err)
				return
			}
		}()
	}
	// 关闭服务器
	c := make(chan os.Signal, 0)
	signal.Notify(c)
	for {
		// Block until a signal is received.
		s := <-c
		log.Info("Got signal:", s) //Got signal: terminated
		if s == syscall.SIGINT || s == syscall.SIGTERM || s == syscall.SIGKILL {
			// shutdown
			shutdown := srv.Shutdown(context.Background())
			if shutdown != nil {
				log.Error(shutdown)
				return
			} else {
				log.Info("end of service...")
				return
			}
		}
	}

}
