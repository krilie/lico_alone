package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/string_util"
	"github.com/krilie/lico_alone/common/uuid_util"
	"github.com/krilie/lico_alone/control/utils"
	"time"
)

// 请求到来后第一个经过的中间件
// 从请中中构建context上下文的中间件
func BuildContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		context := &context_util.Context{}
		context.TraceId = string_util.EmptyOrDefault(c.GetHeader(common.HeaderTraceId), uuid_util.GetUuid())
		context.StartTime = time.Now()
		c.Set(common.GinKeyAppContext, context)
		c.Next()
		context.LastTime = time.Now()
	}
}
