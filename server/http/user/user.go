package user

import (
	"github.com/krilie/lico_alone/application"
	"github.com/krilie/lico_alone/application/user-api"
)

type UserCtrl struct {
	AppUser *user_api.AppUser
}

func NewUserCtrl(app *application.App) *UserCtrl {
	return &UserCtrl{AppUser: app.User}
}
