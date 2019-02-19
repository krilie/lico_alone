package model

type UserRole struct {
	UserID int64 `gorm:"primary_key"`
	RoleID int64 `gorm:"primary_key"`
}

func (UserRole) TableName() string {
	return "user_role"
}
