package ccron

import (
	"github.com/robfig/cron/v3"
)

func Stop(cron *cron.Cron) {
	if cron != nil {
		stop := cron.Stop()
		<-stop.Done()
	}
}

func NewCrone() *cron.Cron {
	CronGlobal := cron.New(cron.WithParser(cron.NewParser(cron.Second|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.DowOptional|cron.Descriptor)),
		cron.WithChain(cron.Recover(cron.DefaultLogger)))
	CronGlobal.Start()
	return CronGlobal
}
