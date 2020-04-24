package model

import "github.com/krilie/lico_alone/common/model"

type RolePermission struct {
	model.Model
	RoleName       string `gorm:"column:role_name;type:varchar(32)"`
	PermissionName string `gorm:"column:permission_name;type:varchar(32)"`
}

func (RolePermission) TableName() string {
	return "tb_auth_role_permission"
}
