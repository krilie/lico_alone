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
	manager.POST("/client_user/new_acc_token", MCtrl.CreateNewAccToken)
	manager.POST("/permission/new_permission", MCtrl.CreateNewPermission)
	manager.POST("/role/add_permission", MCtrl.AddPermissionToRole)
	manager.POST("/role/new_role", MCtrl.CreateNewRole)
	manager.POST("/user/add_role", MCtrl.AddRoleToUser)
}

type ManagerCtrl struct{}

var MCtrl ManagerCtrl

type ManagerCtrler interface {
	AddPermissionToRole(c *gin.Context)
	CreateNewRole(c *gin.Context)
	CreateNewPermission(c *gin.Context)
	AddRoleToUser(c *gin.Context)
	CreateNewAccToken(c *gin.Context)
}
