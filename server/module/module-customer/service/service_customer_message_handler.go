package service

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/broker/messages"
)

func (a *CustomerModule) HandleBrokerWebStationVisited(msg *messages.WebStationVisitedMessage) {
	ctx := context.NewContext()
	ctx.Module = "CustomerModule"
	ctx.Function = "HandleBrokerWebStationVisited"
	err := a.IncreaseCustomerAccessTimesByTraceId(ctx, msg.TraceId, msg.Ip)
	if err != nil {
		a.log.Get(ctx).WithField("err", err.Error()).Error("has err")
	}
}
