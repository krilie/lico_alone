package service

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/broker/messages"
	"github.com/krilie/lico_alone/module/module-statistic/model"
)

func (a *StatisticService) HandleBrokerWebStationVisited(msg *messages.WebStationVisitedMessage) {
	ctx := context.NewContext()
	ctx.Module = "StatisticService"
	ctx.Function = "HandleBrokerWebStationVisited"
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
