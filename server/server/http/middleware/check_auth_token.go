package middleware

import (
	"context"
	"errors"
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/com-model"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/jwt"
	"github.com/krilie/lico_alone/server/http/ginutil"
	"strings"
)

// 权限接口
type IAuth interface {
	HasUser(ctx context.Context, userId string) (bool, error)
	HasPermission(ctx context.Context, userId, method, path string) (bool, error)
	HasRole(ctx context.Context, userId, roleName string) (bool, error)
	CheckJwtToken(tokenStr string) (userClaims jwt.UserClaims, err error)
}

// check user is login and auth token validation
func (m *GinMiddleware) CheckAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get context
		ctx := m.GinUtil.GetAppCtxOrAbort(c)
		if ctx == nil {
			return
		}
		headerAuth := c.GetHeader(ginutil.HeaderAuthorization)
		headerAuth = strings.TrimPrefix(headerAuth, "Bearer ")

		var claims, err = m.IAuth.CheckJwtToken(headerAuth)
		if err != nil {
			if errors.As(err, &jwt2.ValidationError{}) {
				validateErr := err.(*jwt2.ValidationError)
				if errors.Is(validateErr.Inner, jwt.ErrIatTime) {
					ginutil.AbortWithErr(c, errs.NewInvalidToken().WithMsg("token format error"))
					return
				} else if errors.Is(validateErr.Inner, jwt.ErrTimeExp) {
					c.AbortWithStatusJSON(200, com_model.NewRetFromErr(errs.NewInvalidToken().WithMsg("token expired")))
					return
				}
				c.AbortWithStatusJSON(200, com_model.NewRetFromErr(errs.NewInvalidToken().WithMsg("token error")))
				return
			} else {
				c.AbortWithStatusJSON(200, com_model.NewRetFromErr(errs.NewInvalidToken().WithMsg("unknown token error")))
				return
			}
		} else {
			has, err := m.IAuth.HasUser(ctx, claims.UserId)
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
