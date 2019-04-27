package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/common_struct/errs"
)

type Role struct {
	DbHandler
	Name        string `gorm:"type:varchar(50);unique_index;not null" json:"name"`
	Description string `gorm:"type:varchar(100);not null" json:"description"`
}

func (Role) TableName() string {
	return "tb_role"
}

// 主角色 user_admin app service user_normal
// 辅角色 user_normal_vip user_normal_super_vip
// 不分主辅，都是角色

func GetRoleIdByName(db *gorm.DB, name string) (id string, err error) {
	//根据roleName 取到role id 自动关rows
	err = db.Raw("select id from tb_role where name = ?;", name).Row().Scan(&id)
	if err != nil && err == sql.ErrNoRows {
		return "", errs.ErrNotFound
	} else if err != nil {
		return "", err
	} else {
		return id, nil
	}
}
