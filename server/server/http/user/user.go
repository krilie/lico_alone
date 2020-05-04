package user

import (
	"github.com/krilie/lico_alone/service"
	"github.com/krilie/lico_alone/service/user-service"
)

type UserCtrl struct {
	AppUser *user_service.AppUser
}

func NewUserCtrl(app *service.App) *UserCtrl {
	return &UserCtrl{AppUser: app.User}
}
