package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/module/userbase/user"
)

var apiUser user.User

func Init(group *gin.RouterGroup) {
	//用户基础
	userBase := group.Group("/user/base")
	userBase.POST("/login", userCtrl.Login)
	userBase.POST("/logout", userCtrl.Logout)
	userBase.GET("/valid", userCtrl.Valid)                                // 不要登录，要有客户端的key
	userBase.GET("/valid_client_acc_token", userCtrl.ValidClientAccToken) //不要权限的
}

var userCtrl UserCtrl

type UserCtrl struct{}

type UserCtrler interface {
	Login(c *gin.Context)
	GetInfo(c *gin.Context)
	Logout(c *gin.Context)
	ValidClientAccToken(c *gin.Context)
	Valid(c *gin.Context)
	Register(c *gin.Context)
}
