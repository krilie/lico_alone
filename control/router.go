package control

import (
	"github.com/gin-gonic/gin"
	mg "github.com/krilie/lico_alone/control/manager"
	"github.com/krilie/lico_alone/control/midfunc"
	"github.com/krilie/lico_alone/control/user_auth"
)

var LocalRouter *gin.Engine

func init() {
	LocalRouter = gin.Default()
	//api for client token need to check
	apis := LocalRouter.Group("/api")
	//check context and acc token
	apis.Use(midfunc.BuildContext())     //创建上下文
	apis.Use(midfunc.CheckClientToken()) //检查客户端的acc token
	{
		//要登录的接口
		needLogged := apis.Group("")
		needLogged.Use(midfunc.CheckAuthToken()) //检查用户的token是否登录了,即检查是否有基本准入门槛
		{
			//管理组
			manager := needLogged.Group("/manager")
			manager.Use(midfunc.NeedRoles("admin")) //是否有admin权限
			manager.POST("/client_user/new_acc_token", mg.ManagerClientUserNewAccToken)
			manager.POST("/permission/new_permission", mg.ManagerPermissionNewPermission)
			manager.POST("/role/add_permission", mg.ManagerRoleAddPermission)
			manager.POST("/role/new_role", mg.ManagerRoleNewRole)
			manager.POST("/user/add_role", mg.ManagerUserAddRole)
		}
		{
			//用户认证组
			userAuth := needLogged.Group("/user/auth")
			userAuth.GET("/client/acc_token", midfunc.NeedRoles("client", "admin"), user_auth.UserAuthClientAccToken)
			userAuth.POST("/client/has_acc_token", midfunc.NeedRoles("client", "admin"), user_auth.UserAuthClientHasAccToken)
			userAuth.POST("/client/new_acc_token", midfunc.NeedRoles("client"), mg.ManagerClientUserNewAccToken)
			userAuth.POST("/has_permission", user_auth.UserAuthHasPermission) //登录就可以调用的接口
			userAuth.POST("/has_role", user_auth.UserAuthHasRole)
			userAuth.GET("/permissions", user_auth.UserAuthPermissions)
			userAuth.GET("/roles", user_auth.UserAuthRoles)
		}
	}
	//用户基础
	userBase := apis.Group("/user/base")
	userBase.POST("/login", user_auth.UserBaseLogin)
	userBase.POST("/logout", user_auth.UserBaseLogout)
	userBase.GET("/valid", user_auth.UserBaseValid)                                // 不要登录，要有客户端的key
	userBase.GET("/valid_client_acc_token", user_auth.UserBaseValidClientAccToken) //不要权限的
}
