package service

import (
	context2 "context"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/broker/messages"
	"github.com/krilie/lico_alone/module/module-statistic/model"
)

func (a *StatisticService) HandleBrokerWebStationVisited(msg *messages.WebStationVisitedMessage) {
	values := context.NewAppCtxValues()
	values.Module = "StatisticService"
	values.Function = "HandleBrokerWebStationVisited"
	ctx := context.NewAppCtx(context2.Background(), values)
	var vLogs = &model.AddStatVisitorLogsModel{
		AccessTime: msg.AccessTime,
		Ip:         msg.Ip,
		TraceId:    msg.TraceId,
		RegionName: "",
		CityName:   "",
		Memo:       "",
	}
	info, err2 := a.ipInfoApi.GetIpInfo(ctx, msg.Ip)
	if err2 == nil {
		vLogs.RegionName = info.RegionName
		vLogs.CityName = info.City
		vLogs.Memo = info.RawResponse
	}
	err := a.Dao.AddStatVisitorLogs(ctx, vLogs)
	if err != nil {
		a.log.Get(ctx).WithField("err", err).Error("add stat visitor error")
	}
}

func (a *StatisticService) HandleBrokerArticleVisitorMessage(msg *messages.BlogArticleVisitedMessage) {
	values := context.NewAppCtxValues()
	values.Module = "StatisticService"
	values.Function = "HandleBrokerArticleVisitorMessage"
	ctx := context.NewAppCtx(context2.Background(), values)
	var vLogs = &model.AddStatArticleVisitorModel{
		AccessTime:      msg.VisitedTime,
		Ip:              msg.VisitorIp,
		CustomerTraceId: msg.CustomerTraceId,
		ArticleId:       msg.ArticleId,
		ArticleTitle:    msg.ArticleTitle,
		RegionName:      "",
		CityName:        "",
		Memo:            "",
	}
	info, err2 := a.ipInfoApi.GetIpInfo(ctx, msg.VisitorIp)
	if err2 == nil {
		vLogs.RegionName = info.RegionName
		vLogs.CityName = info.City
		vLogs.Memo = info.RawResponse
	}
	err := a.Dao.AddStatArticleVisitorLogs(ctx, vLogs)
	if err != nil {
		a.log.Get(ctx).WithField("err", err).Error("add stat article visitor error")
	}
}
