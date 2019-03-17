package middle_funcs

import (
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico-my-site-user/common/common_struct/errs"
	"github.com/lico603/lico-my-site-user/common/jwt"
	"github.com/lico603/lico-my-site-user/control/gin_util"
	"github.com/lico603/lico-my-site-user/user_base"
)

// check user is login and auth token validation
func CheckUserAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get context
		ctx := gin_util.GetApplicationContextOrAbort(c)
		if ctx == nil {
			return
		}
		headerAuth := c.GetHeader(gin_util.HeaderAuthorization)

		userValidate := user_base.UserValidate(ctx, headerAuth)
		if userValidate != nil {
			if userValidate == jwt.ErrIatTime {
				c.AbortWithStatusJSON(401, errs.UnAuthorized.ToStdWithMsg("token format error"))
				return
			} else if userValidate == jwt.ErrTimeExp {
				c.AbortWithStatusJSON(401, errs.UnAuthorized.ToStdWithMsg("token expired"))
				return
			} else {
				c.AbortWithStatusJSON(500, errs.ErrInternal.ToStdAppendMsg(userValidate.Error()))
				return
			}
		} else {
			c.Next()
			return
		}
	}
}
