package model

import (
	"time"
)

type Permission struct {
	ID          string    `gorm:"primary_key;type:varchar(32)" json:"id"` // 用户id uuid
	CreateTime  time.Time `gorm:"type:DATETIME;not null" json:"create_time"`
	Name        string    `gorm:"type:varchar(50);unique_index;not null"`
	Description string    `gorm:"type:varchar(100);not null"`
}

func (Permission) TableName() string {
	return "tb_permission"
}
