package cdb

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/ccontext"
	"github.com/krilie/lico_alone/common/clog"
	"github.com/krilie/lico_alone/common/config"
	"github.com/sirupsen/logrus"
	"runtime/debug"
	"sync"
	"time"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

var Db *gorm.DB
var log *logrus.Entry
var onceStartDb sync.Once
var onceStopDb sync.Once

func StartDb(cfg config.DB) {
	onceStartDb.Do(func() {
		log = clog.NewLog(ccontext.NewContext(), "br_go.common.db", "init")

		connStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.DbName,
		)
		var err error
		if Db, err = gorm.Open("mysql", connStr+"&loc=Asia%2FShanghai"); err != nil {
			log.Fatal(err, string(debug.Stack())) // 报错退出程序
			return
		} else {
			Db.DB().SetMaxOpenConns(cfg.MaxOpenConn)
			Db.DB().SetMaxIdleConns(cfg.MaxIdleConn)
			Db.DB().SetConnMaxLifetime(time.Second * time.Duration(cfg.ConnMaxLeftTime))
			log.Info("db init done. params:", connStr+"&loc=Asia%2FShanghai") // 数据库初始化成功
			Db = Db.Debug()
		}
	})
}

func GetDbByConfig(cfg config.DB) *gorm.DB {
	StartDb(cfg)
	return Db
}

func CloseDb() {
	onceStopDb.Do(func() {
		err := Db.Close()
		if err != nil {
			log.Warn(err)
		} else {
			log.Info("db closed.")
		}
	})
}
