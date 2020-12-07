package module_like_dislike

import com_model "github.com/krilie/lico_alone/common/com-model"

// LikeDisLikeModel 赞踩
type LikeDisLikeModel struct {
	com_model.Model
	UserId       string `json:"user_id" gorm:"column:user_id;type:char(36);not null;index"`
	BusinessType string `json:"business_type" gorm:"column:business_type;type:nvarchar(16);not null;index"` // article slider comment
	BusinessId   string `json:"business_id" gorm:"column:business_id;type:char(36);not null;index"`         // 唯一的
	GiveType     string `json:"give_type" gorm:"column:give_type;type:nvarchar(16);not null;index"`         // like dislike shock
}

func (LikeDisLikeModel) TableName() string {
	return "tb_like_dislike"
}

type LikeDisLikeModelParams struct {
	UserId       string `json:"user_id"`
	BusinessType string `json:"business_type"`
	BusinessId   string `json:"business_id"`
	GiveType     string `json:"give_type"`
}

type LikeDisLikeModelResult struct {
	BusinessType string `json:"business_type"`
	BusinessId   string `json:"business_id"`
	GiveType     string `json:"give_type"`
	Count        int64  `json:"count"`
}
