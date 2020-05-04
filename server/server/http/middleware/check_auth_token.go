package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/jwt"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// 权限接口
type IAuth interface {
	HasUser(userId string) (bool, error)
	HasPermission(userId, permission string) (bool, error)
	HasRole(userId, roleId string) (bool, error)
}

// check user is login and auth token validation
func CheckAuthToken(auth IAuth) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get context
		ctx := ginutil.GetAppCtxOrAbort(c)
		if ctx == nil {
			return
		}
		headerAuth := c.GetHeader(ginutil.HeaderAuthorization)

		var claims, err = jwt.CheckJwtToken(headerAuth)
		if err != nil {
			if errors.Is(err, jwt.ErrIatTime) {
				ginutil.AbortWithErr(c, errs.NewInvalidToken().WithMsg("token format error"))
				return
			} else if errors.Is(err, jwt.ErrTimeExp) {
				c.AbortWithStatusJSON(200, com_model.NewRetFromErr(errs.NewInvalidToken().WithMsg("token expired")))
				return
			} else {
				c.AbortWithStatusJSON(500, com_model.NewRetFromErr(errs.NewInternal().WithMsg(err.Error())))
				return
			}
		} else {
			b, err := auth.HasUser(claims.UserId)
			if err != nil {
				ginutil.AbortWithErr(c, err)
				return
			}
			if !b {
				ginutil.AbortWithAppErr(c, errs.NewInvalidToken())
				return
			}
			ctx.SetUserId(claims.UserId)
			c.Next()
			return
		}
	}
}
