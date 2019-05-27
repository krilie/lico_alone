package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/application/user"
	"github.com/krilie/lico_alone/control/middleware"
)

var appUser user.AppUser

type UserCtrl struct{}
type UserCtrler interface {
	Login(c *gin.Context)
	GetInfo(c *gin.Context)
	Logout(c *gin.Context)
	ValidClientAccToken(c *gin.Context)
	Valid(c *gin.Context)
	Register(c *gin.Context)
}

func Init(group *gin.RouterGroup) {
	var userCtrl UserCtrl
	//用户基础
	userBase := group.Group("/user")
	userBase.POST("/login", userCtrl.Login)
	userBase.POST("/logout", middleware.CheckAuthToken(), userCtrl.Logout) // 要token有效
	userBase.POST("/register", userCtrl.Register)
	userBase.GET("/valid", userCtrl.Valid)                                // 不要登录，要有客户端的key
	userBase.GET("/valid_client_acc_token", userCtrl.ValidClientAccToken) //不要权限的
}
