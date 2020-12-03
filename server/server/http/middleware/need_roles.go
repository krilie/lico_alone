package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

func (m *GinMiddleware) NeedRoles(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// check user get context
		userId := m.GinUtil.GetUserIdOrAbort(c)
		if userId == "" {
			return
		}
		//check user has permission
		b, err := m.IAuth.HasRole(m.GinUtil.MustGetAppContext(c), userId, role)
		if err != nil {
			ginutil.AbortWithErr(c, err)
			return
		}
		if !b {
			ginutil.AbortWithErr(c, errs.NewNoPermission().WithMsg("没有权限"))
			return
		}
		c.Next()
	}
}
