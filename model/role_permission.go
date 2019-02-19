package model

type RolePermission struct {
	RoleID       int64 `gorm:"primary_key"`
	PermissionID int64 `gorm:"primary_key"`
}

func (RolePermission) TableName() string {
	return "role_permission"
}
