package model

import (
	"gorm.io/gorm"
	"time"
)

type Permission struct {
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index;type:datetime(3)" json:"deleted_at"`
	Name        string         `gorm:"column:name;primaryKey;type:varchar(32)" json:"name"`
	Description string         `gorm:"column:description;type:varchar(100);not null" json:"description"`
	RefMethod   string         `gorm:"column:ref_method;not null;type:varchar(255)" json:"ref_method"`
	RefPath     string         `gorm:"column:ref_path;not null;type:varchar(255)" json:"ref_path"`
	Sort        int            `gorm:"column:sort;not null;type:int(11)" json:"sort"`
}

func (Permission) TableName() string {
	return "tb_auth_permission"
}
