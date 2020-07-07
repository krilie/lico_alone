package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-statistic/model"
	"time"
)

func (a *StatisticDao) AddStatVisitorLogs(ctx context.Context, item *model.AddStatVisitorLogsModel) error {
	log := a.log.Get(ctx).WithFuncName("AddStatVisitorLogs")
	err := a.GetDb(ctx).Model(new(model.StatVisitorLogs)).Save(&model.StatVisitorLogs{
		Model: com_model.Model{
			Id:        id_util.GetUuid(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		AccessTime: item.AccessTime,
		Ip:         item.Ip,
		TraceId:    item.TraceId,
		RegionName: item.RegionName,
		City:       item.CityName,
		Memo:       item.Memo,
	}).Error
	if err != nil {
		log.WithField("err", err).Error("err on save stat visitor logs")
		return err
	}
	return nil
}
