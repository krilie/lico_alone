package user

import (
	apiAuth "github.com/krilie/lico_alone/module/user/auth"
	"github.com/krilie/lico_alone/module/user/info"
)

type AppUser struct {
	apiAuth.UserAuth
	apiAuth.UserManage
	info.User
	Account
}
