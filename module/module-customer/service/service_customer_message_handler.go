package service

import (
	context2 "context"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/broker/messages"
)

func (a *CustomerModule) HandleBrokerWebStationVisited(msg *messages.WebStationVisitedMessage) {
	values := context.NewAppCtxValues()
	values.Module = "CustomerModule"
	values.Function = "HandleBrokerWebStationVisited"
	ctx := context.NewAppCtx(context2.Background(), values)
	err := a.IncreaseCustomerAccessTimesByTraceId(ctx, msg.TraceId, msg.Ip)
	if err != nil {
		a.log.Get(ctx).WithField("err", err.Error()).Error("has err")
	}
}
