package middle_funcs

import (
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico-my-site-user/common/errs"
	"github.com/lico603/lico-my-site-user/common/log"
	"github.com/lico603/lico-my-site-user/control/gin_util"
)

// 中间件
// 参数中的用户名，必需与登录的一致
// query: user_id
// form: user_id
// handler: user_id
// context: key[context]
// 按顺序检查这些信息
// 根据需要选择使用
// 从context中取到有用的值
func UserIdMustMatchLogged() gin.HandlerFunc {
	return func(c *gin.Context) {
		//取用户信息
		userId := checkUserId(c)
		if userId == "" {
			log.Errorln("UserIdMustMatchLogged", "no user_id find")
			c.AbortWithStatusJSON(400, errs.ErrParam.ToStdWithMsg("no user id find"))
			return
		}
		//取认证信息
		ctx := gin_util.GetApplicationContextOrAbort(c)
		if ctx == nil {
			return
		} else if ctx.GetUserIdOrEmpty() != userId {
			log.Errorln("UserIdMustMatchLogged", "user id and login user id not match.")
			c.AbortWithStatusJSON(400, errs.ErrParam.ToStdWithMsg("user id and login user id not match."))
			return
		}
		c.Next()
	}
}

//获取用户名，如果没有用户名，则返回空
func checkUserId(c *gin.Context) string {
	//query user_id
	userId, ok := c.GetQuery("user_id")
	if ok {
		return userId
	}
	//form user_id
	userId, ok = c.GetPostForm("user_id")
	if ok {
		return userId
	}
	// handler user_id
	// _, ok := c.Request.Header["user.id"]
	userId = c.GetHeader("user_id")
	return userId
}
