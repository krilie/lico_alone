package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// check user has some permission request by used url
func NeedPermission(auth IAuth, perms string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get user id from context
		userId := ginutil.GetUserIdOrAbort(c)
		if userId == "" {
			return
		}
		//check user has permission
		b, err := auth.HasPermission(userId, perms)
		if err != nil {
			ginutil.AbortWithErr(c, err)
			return
		}
		if !b {
			ginutil.AbortWithErr(c, errs.NewUnauthorized().WithMsg("无权限"))
			return
		}
		c.Next()
	}
}
