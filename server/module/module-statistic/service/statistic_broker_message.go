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
	err := a.Dao.AddStatVisitorLogs(ctx, &model.AddStatVisitorLogsModel{
		AccessTime: msg.AccessTime,
		Ip:         msg.Ip,
		TraceId:    msg.TraceId,
	})
	if err != nil {
		a.log.Get(ctx).WithField("err", err).Error("add stat visitor error")
	}
}
