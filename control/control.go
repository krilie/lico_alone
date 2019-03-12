package control

//https gin 控制层
/*
 *   /user/base/login				登录
 *   /user/base/logout				登出
 *   /user/base/info				信息
 *   /user/base/valid				检查用户的key是否正常
 *   /user/base/valid_client_acc_token   检查acckey是否有效过期
 *   /user/base/register			注册
 *
 *   /user/auth/has_role			检查用户是否有role
 *   /user/auth/has_permission		检查用户是否有permission
 *   /user/auth/roles				返回用户的所有role
 *   /user/auth/permissions			返回用户的所有权限permission
 *   /user/auth/client_acc_token			返回app用户的所有app_token
 *
 *   /manager/role/new_role             给系统添加新角色
 *   /manager/role/new_permission       给角色添加permission
 *	 /manager/permission/new_permission 给系统添加新权限
 *   /manager/user/new_role             给用户添加新角色
 *   /manager/client_user/new_acc_token         给app用户添加新的key
 */
