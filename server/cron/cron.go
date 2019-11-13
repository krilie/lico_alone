package cron

import (
	"context"
	"github.com/krilie/lico_alone/application"
	"github.com/krilie/lico_alone/common/ccron"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/utils/time_util"
	"github.com/robfig/cron/v3"
	"time"
)

func mustAddCronFunc(cron *cron.Cron, spec string, f func()) {
	_, err := cron.AddFunc(spec, f)
	if err != nil {
		panic(err)
	}
}

func InitAndStartCorn(ctx context.Context, app *application.App) (cronStop func()) {
	crone := ccron.NewCrone()
	//// 定时任务 * * 7 * * ?
	mustAddCronFunc(crone, "0 0 7 * * *", func() {
		log := clog.NewLog(ctx, "定时任务", "早上好")
		err := app.All.Message.SendEmail(ctx, "1197829331@qq.com", "早上好", "早上好"+time.Now().Format(time_util.DefaultFormat))
		if err != nil {
			log.Error(err)
		}
	})
	// stop 定时任务
	return func() {
		ccron.Stop(crone)
	}
}
