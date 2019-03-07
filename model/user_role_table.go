package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lico603/lico-my-site-user/common/log"
)

type UserRole struct {
	UserID int64 `gorm:"primary_key;type:varchar(32)"`
	RoleID int64 `gorm:"primary_key;type:varchar(32)"`
}

func (UserRole) TableName() string {
	return "tb_user_role"
}

//取到这个用户下的所有角色，用逗号串起来
func GetAllRolesByUserId(db *gorm.DB, userId string) (roles string, err error) {
	//执行语句
	rows, err := db.Raw("select b.name as name from tb_user_role a inner join tb_role b on a.role_id = b.id and a.user_id = ?", userId).
		Rows()
	if err != nil {
		return "", err
	}
	defer func() {
		if e := rows.Close(); e != nil {
			log.Error(e)
		}
	}()
	//取字符串
	var name string
	var total string
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			return "", err
		}
		total += "," + name
	}
	//去掉最前一个逗号
	if len(total) > 0 {
		total = total[1:]
	}
	return total, nil
}
