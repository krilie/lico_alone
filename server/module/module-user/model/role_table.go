package model

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	CreatedAt          time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at;index;type:datetime(3)" json:"deleted_at"`
	Name               string         `gorm:"column:name;primaryKey;type:varchar(32)" json:"name"`
	Description        string         `gorm:"column:description;type:varchar(100);not null" json:"description"`
	MainPermissionName string         `gorm:"column:main_permission_name;type:varchar(32);not null;" json:"main_permission_name"`
}

func (Role) TableName() string {
	return "tb_auth_role"
}
