package cron

import (
	"context"
	"github.com/krilie/lico_alone/service"
	"github.com/prometheus/common/log"
	"github.com/robfig/cron/v3"
)

func mustAddCronFunc(cron *cron.Cron, spec string, f func()) {
	_, err := cron.AddFunc(spec, f)
	if err != nil {
		panic(err)
	}
}

func InitAndStartCorn(ctx context.Context, app *service.App) (cronStop func()) {
	crone := cron.New()
	//// 定时任务 * * 7 * * ?
	mustAddCronFunc(crone, "0 0 7 * * *", func() {
		err := app.NotificationEmailService.SendGoodMorningEmail(ctx)
		if err != nil {
			log.Error(err)
		}
	})
	// stop 定时任务
	return func() {
		crone.Stop()
	}
}
