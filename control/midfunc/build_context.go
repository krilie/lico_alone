package midfunc

import (
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico_user/common/context_util"
	"github.com/lico603/lico_user/common/string_util"
	"github.com/lico603/lico_user/common/uuid_util"
	"github.com/lico603/lico_user/control/gin_util"
	"time"
)

// 请求到来后第一个经过的中间件
// 从请中中构建context上下文的中间件
func BuildContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		context := &context_util.Context{}
		context.TraceId = string_util.EmptyOrDefault(c.GetHeader(gin_util.HeaderTraceId), uuid_util.GetUuid())
		context.StartTime = time.Now()
		c.Set(gin_util.GinKeyAppContext, context)
		c.Next()
	}
}
