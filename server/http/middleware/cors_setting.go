package middleware

import (
	"github.com/gin-gonic/gin"
	common_model "github.com/krilie/lico_alone/common/com-model"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/server/http/ginutil"
	"net/http"
	"strings"
)

func (m *GinMiddleware) Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		m.log.WithField("path", c.Request.RequestURI).
			WithField("origin", "origin").
			WithField("method", c.Request.Method).
			Info("in cors")

		var originSettingName = "access-control-allow-origin-c"
		if strings.HasPrefix(c.Request.RequestURI, "/api/manage") {
			originSettingName = "access-control-allow-origin-m"
		} else if strings.HasPrefix(c.Request.RequestURI, "/api/common") {
			originSettingName = "access-control-allow-origin-c"
		}
		origin, err := m.CfgService.GetValueStr(context2.EmptyAppCtx(), originSettingName)
		if err != nil {
			origin = new(string)
			*origin = "*"
		}
		c.Header("Access-Control-Allow-Origin", *origin) // 这是允许访问所有域

		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
		c.Header("Access-Control-Allow-Headers", "Authorization,Content-Length,X-CSRF-Token,Token,session,X_Requested_With,Accept,Origin,Host,Connection,Accept-Encoding,Accept-Language,DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Pragma")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
		c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                         // 缓存请求信息 单位为秒
		c.Header("Access-Control-Allow-Credentials", "true")                                                                                                                                                 //	跨域请求是否需要带cookie信息 默认设置为true
		//c.Set("content-type", "application/json")                                                                                                                                                            // 设置返回格式是json

		//放行所有OPTIONS方法
		//允许跨域设置	可以返回其他子段
		//header的类型
		method := c.Request.Method //请求方法
		if method == "OPTIONS" {
			c.AbortWithStatusJSON(http.StatusOK, common_model.StdSuccess)
			return
		}
		// 处理请求
		c.Next() //	处理请求
		m.log.Get(ginutil.NewGinWrap(c, m.log).AppCtx).Info("end of cors")
	}
}
