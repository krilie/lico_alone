package model

type Permission struct {
	DbHandler
	Name        string `gorm:"type:varchar(50);unique_index;not null"`
	Description string `gorm:"type:varchar(100);not null"`
}

func (Permission) TableName() string {
	return "tb_permission"
}
