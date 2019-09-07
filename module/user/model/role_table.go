package model

import (
	"time"
)

type Role struct {
	Name        string    `gorm:"column:name;primary_key;type:varchar(32)" json:"name"`
	ParentName  string    `gorm:"column:parent_name;type:varchar(32)" json:"parent_name"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	Description string    `gorm:"column:description;type:varchar(100);not null" json:"description"`
}

func (Role) TableName() string {
	return "tb_role"
}
