package model

type Role struct {
	ID          int64  `gorm:"primary_key;auto_increment"`
	Name        string `gorm:"type:varchar(50);unique_index"`
	Description string `gorm:"type:varchar(100)"`
}

func (Role) TableName() string {
	return "role"
}
