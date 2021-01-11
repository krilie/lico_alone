package ctl_user

import (
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/service-user"
	"github.com/krilie/lico_alone/server/http/ginutil"
	"github.com/krilie/lico_alone/server/http/middleware"
)

type UserCtrl struct {
	userService *service_user.UserService
	log         *nlog.NLog
	middleware  *middleware.GinMiddleware
	ginUtil     *ginutil.GinUtils
}

func NewUserCtrl(
	userService *service_user.UserService,
	middleware *middleware.GinMiddleware,
	ginUtil *ginutil.GinUtils,
	log *nlog.NLog) *UserCtrl {

	log = log.WithField(context_enum.Module.Str(), "user controller")
	return &UserCtrl{
		userService: userService,
		log:         log,
		middleware:  middleware,
		ginUtil:     ginUtil,
	}
}
