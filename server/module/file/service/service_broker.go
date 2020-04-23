package service

import (
	"context"
	"github.com/krilie/lico_alone/component/nlog"
)

func (a *Service) RegisterBroker(ctx context.Context) {
	log := nlog.NewLog(ctx, "module/file/service/service_broker.go:5", "RegisterBroker")
	log.Infoln("not implement")
}
