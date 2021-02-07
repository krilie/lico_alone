package middleware

import (
	"github.com/krilie/lico_alone/component/nlog"
	service2 "github.com/krilie/lico_alone/module/module-config/service"
	"github.com/krilie/lico_alone/module/module-user/service"
)

type GinMiddleware struct {
	log        *nlog.NLog
	IAuth      IAuth
	CfgService *service2.ConfigModule
}

func NewGinMiddleware(auth *service.UserModule, log *nlog.NLog, cfg *service2.ConfigModule) *GinMiddleware {
	return &GinMiddleware{
		log:        log,
		IAuth:      auth,
		CfgService: cfg,
	}
}
