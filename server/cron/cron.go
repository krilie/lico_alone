package cron

import (
	"github.com/krilie/lico_alone/application"
	"github.com/krilie/lico_alone/common/ccron"
	"github.com/robfig/cron/v3"
)

func mustAddCronFunc(cron *cron.Cron, spec string, f func()) {
	_, err := cron.AddFunc(spec, f)
	if err != nil {
		panic(err)
	}
}

func InitAndStartCorn(app *application.App) (cronStop func()) {
	crone := ccron.NewCrone()
	//// 定时任务
	//mustAddCronFunc(crone, "*/1 * * * * *", func() {
	//	fmt.Println("in crone")
	//})
	// stop 定时任务
	return func() {
		ccron.Stop(crone)
	}
}
