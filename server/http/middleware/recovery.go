package middleware

import (
	"github.com/gin-gonic/gin"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/server/http/ginutil"
)

func (m *GinMiddleware) MiddlewareRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			ginWrap := ginutil.NewGinWrap(c, m.log)

			if err := recover(); err != nil {
				ctx := ginWrap.GetAppContext()
				if ctx == nil {
					ctx = context2.EmptyAppCtx()
				}
				m.log.Get(ctx).WithError(err).Error("internal err")
				ginWrap.AbortWithErr(errs.NewInternal().WithMsg("internal err by recovery"))
				return
			}
		}()
		c.Next()
	}
}
