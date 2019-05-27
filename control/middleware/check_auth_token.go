package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/jwt"
	"github.com/krilie/lico_alone/common/string_util"
	"github.com/krilie/lico_alone/control/utils"
)

// check user is login and auth token validation
func CheckAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get context
		ctx := utils.GetApplicationContextOrAbort(c)
		if ctx == nil {
			return
		}
		headerAuth := c.GetHeader(utils.HeaderAuthorization)

		var claims, userValidate = apiUser.Validate(ctx, headerAuth)
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
			ctx.NowUserToken = string_util.NewString(headerAuth)
			ctx.UserClaims = claims.(*jwt.UserClaims)
			c.Next()
			return
		}
	}
}
