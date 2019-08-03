package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/module/user/model"
	"time"
)

var Db *gorm.DB
var log = clog.NewLog(context.NewContext(), "alone.module.userbase.model", "init")

func init() {
	connStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		config.GetString("db.user"),
		config.GetString("db.password"),
		config.GetString("db.ip"),
		config.GetInt("db.port"),
		config.GetString("db.database"),
	)

	if db, err := gorm.Open("mysql", connStr); err != nil {
		log.Panicln(err)
		return
	} else {
		Db = db

		if !Db.HasTable(&model.User{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&model.User{}).Error; err != nil {
				log.Panicln(err)
			}
		}
		if !Db.HasTable(&model.Role{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&model.Role{}).Error; err != nil {
				log.Panicln(err)
			}
		}
		if !Db.HasTable(&model.Permission{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&model.Permission{}).Error; err != nil {
				log.Panicln(err)
			}
		}
		if !Db.HasTable(&model.RolePermission{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&model.RolePermission{}).Error; err != nil {
				log.Panicln(err)
			}
		}
		if !Db.HasTable(&model.UserRole{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&model.UserRole{}).Error; err != nil {
				log.Panicln(err)
			}
		}
		if !Db.HasTable(&model.ClientUserAccessToken{}) {
			if err := Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&model.ClientUserAccessToken{}).Error; err != nil {
				log.Panicln(err)
			}
		}

		Db.DB().SetMaxOpenConns(config.GetInt("db.max_open_conn"))
		Db.DB().SetMaxIdleConns(config.GetInt("db.max_idle_conn"))
		Db.DB().SetConnMaxLifetime(time.Second * time.Duration(config.GetInt("db.conn_max_left_time")))
	}
}
