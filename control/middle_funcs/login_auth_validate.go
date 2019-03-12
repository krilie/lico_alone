package middle_funcs

import "github.com/gin-gonic/gin"

//先检查role 拦截与权限
func init() {
	authMap = make(validate, 4)

	//authMap.addRole("/user/base/login","")//	登录
	//authMap.addRole("/user/base/logout",)//	登出
	//authMap.addRole("/user/base/register","base")//		注册

	//查询用户的基本信息 [base] base
	authMap.addRole("/user/base/info", "base")       //		信息
	authMap.addPermission("/user/base/info", "base") //		信息

	//用户jwt是否有效 [base] base
	authMap.addRole("/user/base/valid", "base")       //	检查用户的key是否正常
	authMap.addPermission("/user/base/valid", "base") //	检查用户的key是否正常

	//查用户是否有这个角色 [base] base
	authMap.addRole("/user/auth/has_role", "base")       //		检查用户是否有role
	authMap.addPermission("/user/auth/has_role", "base") //		检查用户是否有role

	//查用户是否有这个权限 [base] base
	authMap.addRole("/user/auth/has_permission", "base")       //	检查用户是否有permission
	authMap.addPermission("/user/auth/has_permission", "base") //	检查用户是否有permission

	//查这个用户的所有角色 [base] base
	authMap.addRole("/user/auth/roles", "base")       //	返回用户的所有role
	authMap.addPermission("/user/auth/roles", "base") //	返回用户的所有role

	//查这个用户的所有权限 [base] base
	authMap.addRole("/user/auth/permissions", "base")       //	返回用户的所有权限permission
	authMap.addPermission("/user/auth/permissions", "base") //	返回用户的所有权限permission

	//查这个用户的所有 client_acc_token [admin,client] -
	authMap.addRole("/user/auth/client_acc_token", "admin", "client") //		返回app用户的所有app_token
	//authMap.addPermission("/user/auth/client_acc_token",  "admin")        //		返回app用户的所有app_token

	//查询这个用户是否有某个acc_token [admin,client] -
	authMap.addRole("/user/auth/has_client_acc_token", "admin") //		返回app用户的所有app_token
	//authMap.addPermission("/user/auth/has_client_acc_token",  "admin")        //		返回app用户的所有app_token

	//给role表添加一个新角色 [admin] -
	authMap.addRole("/manager/role/new_role", "admin") //    给系统添加新角色
	//authMap.addPermission("/manager/role/new_role", "admin")                 //    给系统添加新角色

	//给role添加一条新权限 [admin] -
	authMap.addRole("/manager/role/new_permission", "admin") //    给角色添加permission
	//authMap.addPermission("/manager/role/new_permission", "admin")           //    给角色添加permission

	//给权限表添加一条新权限 [admin] -
	authMap.addRole("/manager/permission/new_permission", "admin") // 给系统添加新权限
	//authMap.addPermission("/manager/permission/new_permission", "admin")     // 给系统添加新权限

	//给用户添加一条新角色 [admin] -
	authMap.addRole("/manager/user/new_role", "admin") //    给用户添加新角色
	//authMap.addPermission("/manager/user/new_role", "admin")                 //    给用户添加新角色

	//给client添加一个新的acc_token [admin,client] - 如果是client，则检查是否和要生成的是同一个用户
	authMap.addRole("/manager/client_user/new_acc_token", "admin", "client") //    给app用户添加新的key
	//authMap.addPermission("/manager/client_user/new_acc_token", "admin","client") //    给app用户添加新的key
}
