package user

import (
	"github.com/krilie/lico_alone/common/dig"
	user_service "github.com/krilie/lico_alone/service/user-service"
)

func init() {
	dig.Container.MustProvide(func(userService *user_service.UserService) *UserCtrl {
		return NewUserCtrl(userService)
	})
}
