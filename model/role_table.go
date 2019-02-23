package model

type Role struct {
	ID          string `gorm:"primary_key;auto_increment;type:varchar(32)"`
	Name        string `gorm:"type:varchar(50);unique_index"`
	Description string `gorm:"type:varchar(100)"`
}

func (Role) TableName() string {
	return "role"
}
