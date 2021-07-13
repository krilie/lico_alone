package model

import (
	"gorm.io/gorm"
	"time"
)

// TbComment
// 1c'f
type TbComment struct {
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;type:datetime(3)" json:"deleted_at"`

	Id        string `gorm:"column:id;primaryKey;type:varchar(32)" json:"id"`
	UserId    string `gorm:"column:user_id;type:varchar(32)" json:"user_id"`
	CommentId string `gorm:"column:comment_id;type:varchar(32)" json:"comment_id"`
	TargetId  string `gorm:"column:target_id;type:varchar(32)" json:"target_id"`

	Content string `gorm:"column:content;type:varchar(100);not null" json:"content"`
}

func (TbComment) TableName() string {
	return "tb_comment"
}
