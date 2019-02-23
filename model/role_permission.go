package model

type RolePermission struct {
	RoleID       string `gorm:"primary_key;type:varchar(32)"`
	PermissionID string `gorm:"primary_key;type:varchar(32)"`
}

func (RolePermission) TableName() string {
	return "tb_role_permission"
}
