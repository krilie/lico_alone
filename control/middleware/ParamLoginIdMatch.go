package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/control/utils"
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
func ParamLoginIdMatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		//取用户信息
		userId := checkUserId(c)
		if userId == "" {
			log.Errorln("ParamLoginIdMatch", "no user_id find")
			c.AbortWithStatusJSON(400, errs.ErrParam.ToStdWithMsg("no user id find"))
			return
		}
		//取认证信息
		ctx := utils.GetAppCtxOrAbort(c)
		if ctx == nil {
			return
		} else if ctx.GetUserIdOrEmpty() != userId {
			log.Errorln("ParamLoginIdMatch", "user id and login user id not match.")
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
