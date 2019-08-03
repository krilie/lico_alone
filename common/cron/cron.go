package cron

import (
	"github.com/robfig/cron/v3"
)

//import "github.com/robfig/cron/v3"

var CronGlobal *cron.Cron

// 定时调度任务 隔离
func init() {
	CronGlobal = cron.New(cron.WithParser(cron.NewParser(cron.Second|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.DowOptional|cron.Descriptor)),
		cron.WithChain(cron.Recover(cron.DefaultLogger)))
	CronGlobal.Start()
}

var Close = Stop

func Stop() {
	if CronGlobal != nil {
		stop := CronGlobal.Stop()
		<-stop.Done()
	}
}
