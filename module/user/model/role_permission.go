package model

import "github.com/krilie/lico_alone/common/cmodel"

type RolePermission struct {
	cmodel.Model
	RoleName       string `gorm:"column:role_name;type:varchar(32)"`
	PermissionName string `gorm:"column:permission_name;type:varchar(32)"`
}

func (RolePermission) TableName() string {
	return "tb_role_permission"
}
