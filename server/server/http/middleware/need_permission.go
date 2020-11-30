package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

// check user has some permission request by used url
func (m *GinMiddleware) NeedPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path
		m.log.Infof("%v %v", method, path)
		// get user id from context
		userId := m.GinUtil.GetUserIdOrAbort(c)
		if userId == "" {
			return
		}
		//check user has permission
		has, err := m.IAuth.HasPermission(m.GinUtil.MustGetAppCtx(c), userId, method, path)
		if err != nil {
			ginutil.AbortWithErr(c, err)
			return
		}
		if !has {
			ginutil.AbortWithErr(c, errs.NewNoPermission().WithMsg("无权限"))
			return
		}
		c.Next()
	}
}
