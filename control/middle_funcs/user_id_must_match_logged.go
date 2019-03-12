package middle_funcs

import (
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/errs"
	"github.com/lico603/lico-my-site-user/common/log"
)

// 中间件
// 参数中的用户名，必需与登录的一致
// query: user_id
// form: user_id
// handler: user_id
// context: key[context]
// 按顺序检查这些信息
// 根据需要选择使用
func UserIdMustMatchLogged() gin.HandlerFunc {
	return func(c *gin.Context) {
		//取用户信息
		userId := checkUserId(c)
		if userId == "" {
			log.Errorln("UserIdMustMatchLogged", "no user_id find")
			c.AbortWithStatusJSON(400, errs.NewErr(errs.ErrParam, "no user id find"))
			return
		}
		//取认证信息
		value, exists := c.Get("context")
		if !exists {
			log.Errorln("UserIdMustMatchLogged", "context not set,no context find")
			c.AbortWithStatusJSON(500, errs.NewErr(errs.ErrInternal, "internal err.please upload this issue to us"))
			return
		}
		ctx := context_util.GetContextOrNil(value)
		if ctx == nil {
			log.Errorln("UserIdMustMatchLogged", "context not set,no context find")
			c.AbortWithStatusJSON(500, errs.NewErr(errs.ErrInternal, "internal err.please upload this issue to us"))
			return
		} else if ctx.GetUserIdOrEmpty() != userId {
			log.Errorln("UserIdMustMatchLogged", "user id and login user id not match.")
			c.AbortWithStatusJSON(400, errs.NewErr(errs.ErrParam, "no user id find"))
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
