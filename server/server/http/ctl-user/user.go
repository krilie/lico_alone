package ctl_user

import (
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/service-user"
)

type UserCtrl struct {
	userService *service_user.UserService
	log         *nlog.NLog
}

func NewUserCtrl(userService *service_user.UserService, log *nlog.NLog) *UserCtrl {
	fieldLog := log.WithField("ctrl", "NewUserCtrl")
	return &UserCtrl{userService: userService, log: fieldLog}
}
