package model

type User struct {
	ID       int64  `gorm:"primary_key;auto_increment"`
	NickName string `gorm:"type:varchar(20);unique_index"`
	Phone    string `gorm:"type:varchar(20);unique_index"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:varchar(64)"`
}

func (User) TableName() string {
	return "user"
}
