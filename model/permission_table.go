package model

type Permission struct {
	ID          string `gorm:"type:varchar(32);primary_key"`
	Name        string `gorm:"type:varchar(50);unique_index"`
	Description string `gorm:"type:varchar(100)"`
}

func (Permission) TableName() string {
	return "permission"
}
