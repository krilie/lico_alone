package control

/**

权限体系

用户们： admin[root] normal sys_user[app service] base 管理员用户 普通用户 终端用户 基础用户[默认]
权限：   admin:  user_add user_delete user_update user_query app_key_delete app_key_add
		normal: normal_update_info,normal_change_nick_name,
		app:	user_normal,key_add,key_delete,key_set_invalid

base 基础用户, 一般面向所有注册用户的接口权限




*/
