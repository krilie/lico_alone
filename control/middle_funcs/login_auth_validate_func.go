package middle_funcs

import "github.com/gin-gonic/gin"

//权限，登录中间件
func LoginAuthValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		//取appKey和userKey，没有就是返回了""
		clientKey := c.GetHeader("ClientKey")
		userKey := c.GetHeader("Authorization")
		url := c.Request.URL.String()
		ip := c.ClientIP()

		//检查clientKey即appkey

		//检查当前url是否在拦截表里,不在则放行，放行不要appKey
		if !authMap.hasUrl(url) {
			c.Next()
			return
		}
		//如果在,就是需要appKey了,注册不要任何权限，包括appkey
		_ = clientKey
		_ = userKey
		_ = ip

	}
}
