package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

func (m *GinMiddleware) NeedRoles(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ginWrap := ginutil.NewGinWrap(c, m.log)

		// check user get context
		userId := ginWrap.GetUserIdOrAbort()
		if userId == "" {
			return
		}
		//check user has permission
		b, err := m.IAuth.HasRole(ginWrap.MustGetAppContext(), userId, role)
		if err != nil {
			ginWrap.AbortWithErr(err)
			return
		}
		if !b {
			ginWrap.AbortWithErr(errs.NewNoPermission().WithMsg("没有权限"))
			return
		}
		c.Next()
	}
}
