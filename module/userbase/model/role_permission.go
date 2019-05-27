package model

import (
	"github.com/deckarep/golang-set"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/log"
)

type RolePermission struct {
	RoleID       string `gorm:"primary_key;type:varchar(32)"`
	PermissionID string `gorm:"primary_key;type:varchar(32)"`
}

func (RolePermission) TableName() string {
	return "tb_role_permission"
}

// 获取用户所有的权限
func GetAllPermissionByUserId(db *gorm.DB, userId string) (set mapset.Set, err error) {

	roles, err := GetAllRolesByUserId(Db, userId)
	if err != nil {
		return nil, err
	}
	//查询rows
	rows, err := db.Raw("select name from tb_role_permission where id in (?);", roles.ToSlice()).Rows()
	if err != nil {
		return nil, err
	}
	defer func() {
		if e := rows.Close(); e != nil {
			log.Error("GetAllPermissionByUserId", e)
		}
	}()
	//todo 可能有问题 temp name new
	var tempName string
	set = mapset.NewThreadUnsafeSet()
	for rows.Next() {
		if err := rows.Scan(&tempName); err != nil {
			return nil, err
		}
		if !set.Add(tempName) {
			log.Infoln("GetAllPermissionByUserId", "name duplicate:", tempName)
		}
	}
	return set, nil
}

//某个用户是否有某个权限
func IsPermissionExistsWithUser(db *gorm.DB, userId string, permissionName string) (bool, error) {
	row := db.Raw("select count(u.id) from "+
		"tb_user u inner join tb_user_role ur on u.id=? and u.id = ur.user_id"+
		"inner join tb_role_permission rp on ur.id = rp.role_id"+
		"inner join tb_permission p on p.id = rp.permission_id and p.name=?;",
		userId, permissionName).Row()
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, errs.ErrInternal
	}
	if count == 0 {
		return false, nil
	} else if count >= 1 {
		return true, nil
	} else {
		return false, errs.ErrInternal
	}
}
