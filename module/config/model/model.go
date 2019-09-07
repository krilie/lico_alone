package model

import "github.com/krilie/lico_alone/common/cmodel"

type Config struct {
	cmodel.Model
	Name  string `gorm:"column:name;size:255"`
	Value string `gorm:"column:value;size:5000"`
}

func (Config) TableName() string {
	return "tb_config"
}
