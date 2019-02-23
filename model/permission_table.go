package model

type Permission struct {
	DbHandler
	Name        string `gorm:"type:varchar(50);unique_index"`
	Description string `gorm:"type:varchar(100)"`
}

func (Permission) TableName() string {
	return "tb_permission"
}
