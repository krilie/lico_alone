package control

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/krilie/lico_alone/common/log"
	myvalid "github.com/krilie/lico_alone/common/validator"
	"github.com/krilie/lico_alone/control/common"
	"github.com/krilie/lico_alone/control/health"
	"github.com/krilie/lico_alone/control/manager"
	"github.com/krilie/lico_alone/control/middleware"
	"github.com/krilie/lico_alone/control/user"
	"gopkg.in/go-playground/validator.v8"
)

var LocalRouter *gin.Engine
var ApiGroup *gin.RouterGroup
var NeedLogin *gin.RouterGroup

func init() {
	LocalRouter = gin.Default()
	// 数据校验
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if e := v.RegisterValidation("user_id", myvalid.UserIdValid); e != nil {
			log.Panicln("valid error:", e)
		}
	}

	//api for client token need to check
	ApiGroup = LocalRouter.Group("/api")
	//check context and acc token
	ApiGroup.Use(middleware.BuildContext())     //创建上下文
	ApiGroup.Use(middleware.CheckClientToken()) //检查客户端的acc token
	//要登录的接口
	NeedLogin = ApiGroup.Group("")
	NeedLogin.Use(middleware.CheckAuthToken()) //检查用户的token是否登录了,即检查是否有基本准入门槛
	// init
	manager.Init(NeedLogin)
	user.Init(ApiGroup)
	health.Init(LocalRouter)
	common.Init(NeedLogin)
}
