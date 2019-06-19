package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/control/utils"
	"time"
)

// 请求到来后第一个经过的中间件
// 从请中中构建context上下文的中间件
func BuildContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &context.Context{}
		ctx.TraceId = str_util.EmptyOrDefault(c.GetHeader(utils.HeaderTraceId), id_util.GetUuid())
		ctx.StartTime = time.Now()
		c.Set(utils.GinKeyAppContext, ctx)
		c.Next()
		ctx.LastTime = time.Now()
	}
}
