package cron

import "github.com/robfig/cron/v3"

type CronLocal struct {
	*cron.Cron
}

func (c *CronLocal) StopAndWait() {
	if c != nil {
		stop := c.Stop()
		<-stop.Done()
	}
}

func NewCrone() *CronLocal {
	CronGlobal := cron.New(cron.WithParser(cron.NewParser(cron.Second|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.DowOptional|cron.Descriptor)),
		cron.WithChain(cron.Recover(cron.DefaultLogger)))
	CronGlobal.Start()
	return &CronLocal{Cron: CronGlobal}
}

func (c *CronLocal) mustAddCronFunc(spec string, f func()) {
	_, err := c.Cron.AddFunc(spec, f)
	if err != nil {
		panic(err)
	}
}
