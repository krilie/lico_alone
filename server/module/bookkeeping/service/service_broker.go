package service

import (
	"context"
	"github.com/krilie/lico_alone/component/nlog"
)

func (s *Service) RegisterBroker(ctx context.Context) {
	log := nlog.NewLog(ctx, "module/account/service/service_broker.go:8", "RegisterBroker")
	log.Infoln("not implement")
}
