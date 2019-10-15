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

func (d *Dao) Begin(db *gorm.DB) (*Dao, error) {
	if db == nil {
		if d.IsInTx() {
			return d, nil
		}
		db = d.Db.Begin()
		if err := db.Error; err != nil {
			return nil, db.Error
		}
	}
	return &Dao{
		Db: db,
	}, nil
}

func (d *Dao) IsInTx() bool {
	return IsInTx(d.Db)
}

func IsInTx(db *gorm.DB) bool {
	if db.DB() == nil {
		return true
	} else {
		return false
	}
}
