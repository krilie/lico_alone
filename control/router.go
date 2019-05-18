package control

import (
	"github.com/gin-gonic/gin"
	_ "github.com/krilie/lico_alone/control/manager"
	"github.com/krilie/lico_alone/control/middleware"
	_ "github.com/krilie/lico_alone/control/user"
)

var LocalRouter *gin.Engine
var ApiGroup *gin.RouterGroup
var NeedLogin *gin.RouterGroup

func init() {
	LocalRouter = gin.Default()
	//api for client token need to check
	ApiGroup = LocalRouter.Group("/api")
	//check context and acc token
	ApiGroup.Use(middleware.BuildContext())     //创建上下文
	ApiGroup.Use(middleware.CheckClientToken()) //检查客户端的acc token
	//要登录的接口
	needLogged := ApiGroup.Group("")
	needLogged.Use(middleware.CheckAuthToken()) //检查用户的token是否登录了,即检查是否有基本准入门槛
}
