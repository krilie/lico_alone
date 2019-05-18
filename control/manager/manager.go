package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/control/middleware"
	"github.com/krilie/lico_alone/module/userbase/auth"
)

var apiManagerUser auth.Manage

func Init(group *gin.RouterGroup) {
	//管理组
	manager := group.Group("/manager")
	manager.Use(middleware.NeedRoles("admin")) //是否有admin权限
	manager.POST("/client_user/new_acc_token", ManagerClientUserNewAccToken)
	manager.POST("/permission/new_permission", ManagerPermissionNewPermission)
	manager.POST("/role/add_permission", ManagerRoleAddPermission)
	manager.POST("/role/new_role", ManagerRoleNewRole)
	manager.POST("/user/add_role", ManagerUserAddRole)
}
