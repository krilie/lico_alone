package control

import "github.com/gin-gonic/gin"

//先检查role 拦截与权限
func init() {
	authMap = make(validate, 4)
	//authMap.addRole("/user/base/login","")//	登录
	//authMap.addRole("/user/base/logout",)//	登出
	//authMap.addRole("/user/base/register","base")//		注册
	authMap.addRole("/user/base/info", "base")                         //		信息
	authMap.addRole("/user/base/valid", "base")                        //	检查用户的key是否正常
	authMap.addRole("/user/auth/has_role", "base")                     //		检查用户是否有role
	authMap.addRole("/user/auth/has_permission", "base")               //	检查用户是否有permission
	authMap.addRole("/user/auth/roles", "base")                        //	返回用户的所有role
	authMap.addRole("/user/auth/permissions", "base")                  //	返回用户的所有权限permission
	authMap.addRole("/user/auth/app_keys", "sys_user", "admin")        //		返回app用户的所有app_token
	authMap.addRole("/manager/role/new_role", "admin")                 //    给系统添加新角色
	authMap.addRole("/manager/role/new_permission", "admin")           //    给角色添加permission
	authMap.addRole("/manager/permission/new_permission", "admin")     // 给系统添加新权限
	authMap.addRole("/manager/user/new_role", "admin")                 //    给用户添加新角色
	authMap.addRole("/manager/app_user/new_keys", "sys_user", "admin") //    给app用户添加新的key
}

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
