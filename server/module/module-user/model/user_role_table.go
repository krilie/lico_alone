package model

import "github.com/krilie/lico_alone/common/model"

type UserRole struct {
	model.Model
	RoleName string `gorm:"column:role_name;type:varchar(32)"`
	UserId   string `gorm:"column:user_id;type:char(36)"`
}

func (UserRole) TableName() string {
	return "tb_auth_user_role"
}
