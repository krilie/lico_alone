package model

import (
	"github.com/deckarep/golang-set"
	"github.com/jinzhu/gorm"
	"github.com/lico603/lico-my-site-user/common/errs"
	"github.com/lico603/lico-my-site-user/common/log"
)

type UserRole struct {
	UserID string `gorm:"primary_key;type:varchar(32)"`
	RoleID string `gorm:"primary_key;type:varchar(32)"`
}

func (UserRole) TableName() string {
	return "tb_user_role"
}

//取到这个用户下的所有角色，用逗号串起来
func GetAllRolesByUserId(db *gorm.DB, userId string) (roles mapset.Set, err error) {
	//执行语句
	rows, err := db.Raw("select b.name as name from tb_user_role a inner join tb_role b on a.role_id = b.id and a.user_id = ?", userId).
		Rows()
	if err != nil {
		return nil, err
	}
	defer func() {
		if e := rows.Close(); e != nil {
			log.Error(e)
		}
	}()
	//取字符串
	var name string
	var nameMap = mapset.NewThreadUnsafeSet()
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		if !nameMap.Add(name) {
			log.Infoln("GetAllRolesByUserId", "roles duplicate:", name)
		}
	}
	return nameMap, nil
}

//检查这个角色 用户 是否有关联
func IsUserRoleRelationExist(db *gorm.DB, userId, roleId string) (b bool, err error) {
	//根据roleName 取到role id
	var count int
	err = db.Raw("select count(*) from tb_user_role where user_id = ? and role_id = ?;",
		userId, roleId).Row().Scan(&count)
	if err != nil {
		return false, err
	}
	//count 值
	if count == 1 {
		return true, nil
	} else if count == 0 {
		return false, nil
	} else {
		return false, errs.ErrInternal
	}
}
