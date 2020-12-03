package middleware

import (
	context2 "context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/common/utils/str_util"
	"github.com/krilie/lico_alone/server/http/ginutil"
	http_common "github.com/krilie/lico_alone/server/http/http-common"
	"net/http"
	"time"
)

// 请求到来后第一个经过的中间件
// 从请中中构建context上下文的中间件
func BuildContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		parentCtx := context2.Background()
		values := context.NewAppCtxValues()
		values.TraceId = str_util.EmptyOrDefault(c.GetHeader(ginutil.HeaderTraceId), id_util.GetUuid())
		values.StartTime = time.Now()
		values.RemoteIp = c.ClientIP()
		c.Set(ginutil.GinKeyAppContext, context.NewAppCtx(parentCtx, values))
		// cookie trace id
		traceId, err := c.Cookie(http_common.CookieCustomerTraceId)
		if errors.Is(err, http.ErrNoCookie) {
			traceId = id_util.GetUuid()
			c.SetCookie(http_common.CookieCustomerTraceId, traceId, 3600*24*365*10, "", "", true, true)
		}
		values.CustomerTraceId = traceId
		c.Next()
		values.LastTime = time.Now()
	}
}
