package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/module/userbase/user"
)

var apiUser user.User

func Init(group *gin.RouterGroup) {
	//用户基础
	userBase := group.Group("/user/base")
	userBase.POST("/login", UserBaseLogin)
	userBase.POST("/logout", UserBaseLogout)
	userBase.GET("/valid", UserBaseValid)                                // 不要登录，要有客户端的key
	userBase.GET("/valid_client_acc_token", UserBaseValidClientAccToken) //不要权限的
}
