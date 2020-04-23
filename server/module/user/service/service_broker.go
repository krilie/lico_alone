package service

import (
	"context"
	"github.com/krilie/lico_alone/component/clog"
)

// 注册处理事件 有错误panic
func (s *Service) RegisterBroker(ctx context.Context) {
	log := clog.NewLog(ctx, "module/user/service/service_broker.go", "RegisterBroker")
	log.Infoln("not implement")
}
