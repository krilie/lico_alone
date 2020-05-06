package ctl_user

import (
	"github.com/krilie/lico_alone/service/user-service"
)

type UserCtrl struct {
	userService *user_service.UserService
}

func NewUserCtrl(userService *user_service.UserService) *UserCtrl {
	return &UserCtrl{userService: userService}
}
