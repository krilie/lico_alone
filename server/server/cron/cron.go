package cron

import (
	"context"
	"github.com/krilie/lico_alone/service"
	"github.com/prometheus/common/log"
)

func InitAndStartCorn(ctx context.Context, app *service.App) (cronStop func()) {
	crone := NewCrone()
	//// 定时任务 * * 7 * * ?
	crone.mustAddCronFunc("0 0 7 * * *", func() {
		err := app.NotificationEmailService.SendGoodMorningEmail(ctx)
		if err != nil {
			log.Error(err)
		}
	})
	// stop 定时任务
	return func() {
		crone.StopAndWait()
	}
}
