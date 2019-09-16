package middleware

// 权限接口
type IAuth interface {
	HasUser(id string) (bool, error)
	HasPermission(id, permission string) (bool, error)
	HasRole(userId, roleId string) (bool, error)
}
