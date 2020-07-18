package com_model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id        string         `gorm:"column:id;primary_key;type:char(36)" json:"id"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;not null"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at"`
}
