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

	Id        string `gorm:"column:id;primaryKey;type:varchar(32)" json:"id"`      // 这是一个id
	UserId    string `gorm:"column:user_id;type:varchar(32)" json:"user_id"`       // 这是评论的用户的id
	CommentId string `gorm:"column:comment_id;type:varchar(32)" json:"comment_id"` // 这是子评论的父评论的id
	TargetId  string `gorm:"column:target_id;type:varchar(32)" json:"target_id"`   // 这是评论目标的id

	Content      string `gorm:"column:content;type:varchar(100);not null" json:"content"`        // 此为评论的内容
	LikeCount    int    `gorm:"column:like_count;type:int(11);not null" json:"like_count"`       // 喜欢数
	DislikeCount int    `gorm:"column:dislike_count;type:int(11);not null" json:"dislike_count"` // 不喜欢数
	IsCheck      bool   `gorm:"column:is_check;type:tinyint(1);not null" json:"is_check"`        // 是否审核
}

func (TbComment) TableName() string {
	return "tb_comment"
}
