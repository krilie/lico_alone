package service

import (
	"context"
	"github.com/krilie/lico_alone/component/nlog"
)

// 注册处理事件 有错误panic
func (s *Service) RegisterBroker(ctx context.Context) {
	log := nlog.NewLog(ctx, "module/user/service/service_broker.go", "RegisterBroker")
	log.Infoln("not implement")
}
