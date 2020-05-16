package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/jwt"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// 权限接口
type IAuth interface {
	HasUser(ctx context.Context, userId string) (bool, error)
	HasPermission(ctx context.Context, userId, method, path string) (bool, error)
	HasRole(ctx context.Context, userId, roleName string) (bool, error)
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
			has, err := auth.HasUser(ctx, claims.UserId)
			if err != nil {
				ginutil.AbortWithErr(c, err)
				return
			}
			if !has {
				ginutil.AbortWithAppErr(c, errs.NewInvalidToken())
				return
			}
			ctx.SetUserId(claims.UserId)
			c.Next()
			return
		}
	}
}
