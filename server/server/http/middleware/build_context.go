package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/server/http/ginutil"
	"time"
)

// 请求到来后第一个经过的中间件
// 从请中中构建context上下文的中间件
func BuildContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.NewContext()
		ctx.SetTraceId(str_util.EmptyOrDefault(c.GetHeader(ginutil.HeaderTraceId), id_util.GetUuid()))
		ctx.SetStartTime(time.Now())
		ctx.SetRemoteIp(c.ClientIP())
		c.Set(ginutil.GinKeyAppContext, ctx)
		c.Next()
		ctx.SetLastTime(time.Now())
	}
}
