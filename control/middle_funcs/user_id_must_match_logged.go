package middle_funcs

import "github.com/gin-gonic/gin"

// 中间件
// 参数中的用户名，必需与登录的一致
// query: user_id
// form: user_id
// handler: user_id
// 按顺序检查这些信息
// 根据需要选择使用
func UserIdMustMatchLogged() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
