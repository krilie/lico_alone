package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/module/module-statistic/model"
	"time"
)

func (a *StatisticDao) AddStatVisitorLogs(ctx context.Context, item *model.AddStatVisitorLogsModel) error {
	log := a.log.Get(ctx).WithFuncName("AddStatVisitorLogs")
	a.GetDb(ctx).Model(new(model.StatVisitorLogs)).Save(&model.StatVisitorLogs{
		Model:      com_model.Model{},
		AccessTime: time.Time{},
		Ip:         "",
		TraceId:    "",
	})
}
