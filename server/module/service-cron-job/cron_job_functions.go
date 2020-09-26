package service_cron_job

import "context"

func (cron *CronJobService) InitCronJobAndFunctions(ctx context.Context) error {
	log := cron.log.Get(ctx).WithFuncName("InitCronJobAndFunctions")
	log.WithField("cron_status", "begin").Trace("cron job init")
	defer log.WithField("cron_status", "end").Trace("cron job init")
	// 定时发邮件的
	_, err := cron.cron.AddFunc("@every 1m", func() {
		err := cron.moduleMessage.SendEmail(ctx, "", "hello", "hello corn job")
		if err != nil {
			log.WithField("err", err.Error()).Error("send hello email err")
		}
	})
	if err != nil {
		return err
	}
	// 其它信息

	return nil
}
