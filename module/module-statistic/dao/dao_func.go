package dao

import (
	"context"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/module-statistic/model"
	"gorm.io/gorm"
	"time"
)

func (a *StatisticDao) AddStatVisitorLogs(ctx context.Context, item *model.AddStatVisitorLogsModel) error {
	log := a.log.Get(ctx).WithFuncName("AddStatVisitorLogs")
	err := a.GetDb(ctx).Model(new(model.StatVisitorLogs)).Create(&model.StatVisitorLogs{
		Model: com_model.Model{
			Id:        id_util.GetUuid(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
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

func (a *StatisticDao) AddStatArticleVisitorLogs(ctx context.Context, item *model.AddStatArticleVisitorModel) error {
	log := a.log.Get(ctx).WithFuncName("AddStatArticleVisitorLogs")
	err := a.GetDb(ctx).Model(new(model.StatArticleVisitorLogs)).Create(&model.StatArticleVisitorLogs{
		Model: com_model.Model{
			Id:        id_util.GetUuid(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt{},
		},
		AccessTime:      item.AccessTime,
		Ip:              item.Ip,
		CustomerTraceId: item.CustomerTraceId,
		ArticleId:       item.ArticleId,
		ArticleTitle:    item.ArticleTitle,
		RegionName:      item.RegionName,
		City:            item.CityName,
		Memo:            item.Memo,
	}).Error
	if err != nil {
		log.WithField("err", err).Error("err on save stat visitor logs")
		return err
	}
	return nil
}
