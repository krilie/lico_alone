package middleware

import (
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-user/service"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

type GinMiddleware struct {
	log     *nlog.NLog
	IAuth   IAuth
	GinUtil *ginutil.GinUtils
}

func NewGinMiddleware(ginUtil *ginutil.GinUtils, auth *service.UserModule, log *nlog.NLog) *GinMiddleware {
	return &GinMiddleware{
		log:     log,
		IAuth:   auth,
		GinUtil: ginUtil,
	}
}
