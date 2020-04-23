package cron

import (
	"context"
	"github.com/krilie/lico_alone/application"
	"github.com/krilie/lico_alone/component/cron"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/robfig/cron/v3"
)

func mustAddCronFunc(cron *cron.Cron, spec string, f func()) {
	_, err := cron.AddFunc(spec, f)
	if err != nil {
		panic(err)
	}
}

func InitAndStartCorn(ctx context.Context, app *application.App) (cronStop func()) {
	crone := cron.NewCrone()
	//// 定时任务 * * 7 * * ?
	mustAddCronFunc(crone, "0 0 7 * * *", func() {
		log := nlog.NewLog(ctx, "定时任务", "早上好")
		err := app.All.SendGoodMorningEmail(ctx)
		if err != nil {
			log.Error(err)
		}
	})
	// stop 定时任务
	return func() {
		cron.Stop(crone)
	}
}
