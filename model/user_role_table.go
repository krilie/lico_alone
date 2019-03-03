package model

import "github.com/jinzhu/gorm"

type UserRole struct {
	UserID int64 `gorm:"primary_key;type:varchar(32)"`
	RoleID int64 `gorm:"primary_key;type:varchar(32)"`
}

func (UserRole) TableName() string {
	return "tb_user_role"
}

//取到这个用户下的所有角色，用逗号串起来
func GetAllRolesByUserId(db *gorm.DB, userId string) (roles string, err error) {
	var rolesS []string
	err = db.Exec("select b.name from tb_user_role a inner join tb_role b on a.role_id = b.id and a.user_id = ?", userId).
		Find(rolesS).Error
	if err != nil {
		return "", err
	} else {
		for _, v := range rolesS {
			roles = roles + "," + v
		}
		return roles, nil
	}
}
