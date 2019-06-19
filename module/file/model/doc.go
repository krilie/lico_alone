package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/comlog"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/context"
	"time"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

var Db *gorm.DB
var log = comlog.NewLog(context.NewContext(), "alone.module.file.model")

func init() {
	connStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		config.GetString("db.user"),
		config.GetString("db.password"),
		config.GetString("db.ip"),
		config.GetInt("db.port"),
		config.GetString("db.database"),
	)

	if Db, err := gorm.Open("mysql", connStr); err != nil {
		log.Panicln(err)
		return
	} else {
		if !Db.HasTable(&File{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&File{}).Error; err != nil {
				log.Panicln(err)
			}
		}

		Db.DB().SetMaxOpenConns(config.GetInt("db.max_open_conn"))
		Db.DB().SetMaxIdleConns(config.GetInt("db.max_idle_conn"))
		Db.DB().SetConnMaxLifetime(time.Second * time.Duration(config.GetInt("db.conn_max_left_time")))
	}
}
