package model

type Role struct {
	DbHandler
	Name        string `gorm:"type:varchar(50);unique_index"`
	Description string `gorm:"type:varchar(100);not null"`
}

func (Role) TableName() string {
	return "tb_role"
}

// 主角色 user_admin app service user_normal
// 辅角色 user_normal_vip user_normal_super_vip
// 不分主辅，都是角色
