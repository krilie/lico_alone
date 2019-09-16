package model

import (
	"time"
)

type Permission struct {
	Name        string    `gorm:"column:name;primary_key;type:varchar(32)" json:"name"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	Description string    `gorm:"column:description;type:varchar(100);not null" json:"description"`
	RefMethod   string    `gorm:"column:ref_method;" json:"ref_method"`
	RefPath     string    `gorm:"column:ref_path;" json:"ref_path"`
	Sort        int       `gorm:"column:sort;" json:"sort"`
}

func (Permission) TableName() string {
	return "tb_permission"
}
