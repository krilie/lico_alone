package model

import (
	"time"
)

// 客戶端也是用戶的一種，可以登錄后臺，請請一個access_token，
// 用于調用普通用戶登錄接口時認證，或是調用其他需要的接口時加上

//這個就不是jwt了，并且有效時間長一些，可以控制是否有效，可以吊銷等
//為簡間，只在部份接口上要求要有accesstoken 比如登錄，
//web 瀏覽器沒有token，要通過其它方式認證
type ClientUserAccessToken struct {
	Token          string    `gorm:"primary_key;type:varchar(32)" json:"token"` // token本身，可以是數字字母等
	CreateTime     time.Time `gorm:"type:DATETIME;not null" json:"create_time"`
	ExpirationTime time.Time `gorm:"type:DATETIME;not null" json:"expiration_time"`
	UserId         string    `gorm:"type:varchar(32);not null" json:"user_id"` //與之對應的userid ，就是app user的id
	CreateBy       string    `gorm:"type:varchar(32)" json:"create_by"`        //由哪個用戶創建
	IsValid        bool      `gorm:"type:boolean" json:"is_valid"`
	Description    string    `gorm:"type:varchar(128)" json:"description"` //给这个token加个描述，表明是那个干什么的
}

//一个app用户可以有多个token
func (ClientUserAccessToken) TableName() string {
	return "tb_client_user_acc_token"
}
