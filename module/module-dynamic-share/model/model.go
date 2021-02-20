package model

import com_model "github.com/krilie/lico_alone/common/com-model"

// DynamicShare 动态表
// 动态信息
type DynamicShare struct {
	com_model.Model
	Content string `json:"content" gorm:"column:content;type:text;not null"`    // 内容
	Sort    int    `json:"sort" gorm:"column:sort;type:int(11);not null;index"` // 大的在前
}

func (DynamicShare) TableName() string {
	return "tb_dynamic_share"
}

// DynamicShareLabel 标签表
type DynamicShareLabel struct {
	com_model.Model
	ShareId string `gorm:"column:id;primaryKey;type:char(36)" json:"id"`
	Label   string `gorm:"column:label,type:nvarchar(36);not null;index"`
}

func (DynamicShareLabel) TableName() string {
	return "tb_dynamic_share_label"
}

type CreateDynamicShareLabelModel struct {
	ShareId string `json:"share_id"`
	Label   string `json:"label"`
}

type CreateDynamicShareModel struct {
	Content string `json:"content" gorm:"column:content;type:text;not null"` // 内容
	Sort    int    `json:"sort" gorm:"column:sort;type:int(11);not null;index"`
}
type UpdateDynamicShareModel struct {
	Id string `gorm:"column:id;primaryKey;type:char(36)" json:"id"`
	CreateDynamicShareModel
}

type QueryDynamicShareModel struct {
	com_model.PageParams
	ContentLike string
}
type QueryDynamicShareResModel struct {
	TotalPage  int64
	TotalCount int64
	PageNum    int64
	PageSize   int64
	Data       []DynamicShare
}
