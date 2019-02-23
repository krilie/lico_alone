package model

type UserRole struct {
	UserID int64 `gorm:"primary_key;type:varchar(32)"`
	RoleID int64 `gorm:"primary_key;type:varchar(32)"`
}

func (UserRole) TableName() string {
	return "tb_user_role"
}
