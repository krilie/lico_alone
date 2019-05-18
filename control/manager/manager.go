package manager

import (
	"github.com/krilie/lico_alone/control"
	"github.com/krilie/lico_alone/control/middleware"
	"github.com/krilie/lico_alone/module/userbase/auth"
)

var apiManagerUser auth.Manage

func init() {
	//管理组
	manager := control.NeedLogin.Group("/manager")
	manager.Use(middleware.NeedRoles("admin")) //是否有admin权限
	manager.POST("/client_user/new_acc_token", ManagerClientUserNewAccToken)
	manager.POST("/permission/new_permission", ManagerPermissionNewPermission)
	manager.POST("/role/add_permission", ManagerRoleAddPermission)
	manager.POST("/role/new_role", ManagerRoleNewRole)
	manager.POST("/user/add_role", ManagerUserAddRole)
}
