package com_model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id        string         `gorm:"column:id;primaryKey;type:char(36)" json:"id"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;not null;type:datetime(3)"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;not null;type:datetime(3)"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;type:datetime(3)" json:"deleted_at"`
}

type ModelVo struct {
	Id        string     `gorm:"column:id;primaryKey;type:char(36)" json:"id"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;not null;type:datetime(3)"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;not null;type:datetime(3)"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;type:datetime(3)" json:"deleted_at"`
}
