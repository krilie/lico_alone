package manager

import (
	"github.com/krilie/lico_alone/module/file/service"
	"github.com/krilie/lico_alone/module/userbase/auth"
)

type AppManager struct {
	auth.UserAuth
	auth.UserManage
	service.FileOp
}
