package model

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `sql:"index" json:"deleted_at"`
	Name               string         `gorm:"column:name;primary_key;type:varchar(32)" json:"name"`
	Description        string         `gorm:"column:description;type:varchar(100);not null" json:"description"`
	MainPermissionName string         `gorm:"column:main_permission_name;type:varchar(32);not null;" json:"main_permission_name"`
}

func (Role) TableName() string {
	return "tb_auth_role"
}
