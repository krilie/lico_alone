package ctl_user

import (
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/service/user-service"
)

type UserCtrl struct {
	userService *user_service.UserService
	log         *nlog.NLog
}

func NewUserCtrl(userService *user_service.UserService, log *nlog.NLog) *UserCtrl {
	fieldLog := log.WithField("ctrl", "NewUserCtrl")
	return &UserCtrl{userService: userService, log: fieldLog}
}
