package service

import (
	"context"
	"github.com/krilie/lico_alone/common/clog"
)

func (s *Service) RegisterBroker(ctx context.Context) {
	log := clog.NewLog(ctx, "module/account/service/service_broker.go:8", "RegisterBroker")
	log.Infoln("not implement")
}
