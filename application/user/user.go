package user

import (
	apiAuth "github.com/krilie/lico_alone/module/userbase/auth"
)

type AppUser struct {
	apiAuth.UserAuth
	apiAuth.UserManage
}
