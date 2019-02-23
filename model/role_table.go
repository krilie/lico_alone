package model

type Role struct {
	DbHandler
	Name        string `gorm:"type:varchar(50);unique_index"`
	Description string `gorm:"type:varchar(100)"`
}

func (Role) TableName() string {
	return "tb_role"
}
