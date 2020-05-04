package user

import (
	"github.com/krilie/lico_alone/service"
	"github.com/krilie/lico_alone/service/user-api"
)

type UserCtrl struct {
	AppUser *user_api.AppUser
}

func NewUserCtrl(app *service.App) *UserCtrl {
	return &UserCtrl{AppUser: app.User}
}
