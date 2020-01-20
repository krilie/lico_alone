package cdb

import (
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/config"
)

type Dao struct {
	Db *gorm.DB
}

func NewDao(cfg config.DB) *Dao {
	return &Dao{Db: GetDbByConfig(cfg)}
}
