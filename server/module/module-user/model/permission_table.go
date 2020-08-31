package model

import (
	"gorm.io/gorm"
	"time"
)

type Permission struct {
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `sql:"index" json:"deleted_at"`
	Name        string         `gorm:"column:name;primary_key;type:varchar(32)" json:"name"`
	Description string         `gorm:"column:description;type:varchar(100);not null" json:"description"`
	RefMethod   string         `gorm:"column:ref_method;not null" json:"ref_method"`
	RefPath     string         `gorm:"column:ref_path;not null" json:"ref_path"`
	Sort        int            `gorm:"column:sort;not null" json:"sort"`
}

func (Permission) TableName() string {
	return "tb_auth_permission"
}
