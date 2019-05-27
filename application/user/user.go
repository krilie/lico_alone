package user

import (
	apiAuth "github.com/krilie/lico_alone/module/userbase/auth"
	"github.com/krilie/lico_alone/module/userbase/user"
)

type AppUser struct {
	apiAuth.UserAuth
	apiAuth.UserManage
	user.User
}
