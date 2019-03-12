package middle_funcs

import (
	"github.com/gin-gonic/gin"
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/string_util"
	"github.com/lico603/lico-my-site-user/common/uuid_util"
	"time"
)

// 请求到来后第一个经过的中间件
// 从请中中构建context上下文的中间件
func BuildContextFunc() gin.HandlerFunc {
	return func(c *gin.Context) {

		context := &context_util.Context{}
		context.TraceId = string_util.EmptyDefault(c.GetHeader("TraceId"), uuid_util.GetUuid())
		context.StartTime = time.Now()
		c.Set("context", context)
	}
}
