package model

import "github.com/krilie/lico_alone/common/com-model"

type RolePermission struct {
	com_model.Model
	RoleName       string `gorm:"column:role_name;type:varchar(32);not null"`
	PermissionName string `gorm:"column:permission_name;type:varchar(32);not null"`
}

func (RolePermission) TableName() string {
	return "tb_auth_role_permission"
}
