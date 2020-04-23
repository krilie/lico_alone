package ndb

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/clog"
	"github.com/prometheus/common/log"
	"runtime/debug"
	"sync"
	"time"
)

type Ndb struct {
	onceStartDb sync.Once
	onceStopDb  sync.Once
	Db          *gorm.DB
}

func (ndb *Ndb) InitAndStart(dbCfg config.DB) {
	ndb.onceStartDb.Do(func() {
		log := clog.NewLog(context.NewContext(), "br_go.common.db", "init")

		connStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
			dbCfg.User,
			dbCfg.Password,
			dbCfg.Host,
			dbCfg.Port,
			dbCfg.DbName,
		)
		var err error
		if ndb.Db, err = gorm.Open("mysql", connStr+"&loc=Asia%2FShanghai"); err != nil {
			fmt.Println(err.Error())
			log.Fatal(err, string(debug.Stack())) // 报错退出程序
			return
		} else {
			ndb.Db.DB().SetMaxOpenConns(dbCfg.MaxOpenConn)
			ndb.Db.DB().SetMaxIdleConns(dbCfg.MaxIdleConn)
			ndb.Db.DB().SetConnMaxLifetime(time.Second * time.Duration(dbCfg.ConnMaxLeftTime))
			log.Info("db init done. params:", connStr+"&loc=Asia%2FShanghai") // 数据库初始化成功
			ndb.Db = ndb.Db.Debug()
		}
	})
}

func (ndb *Ndb) CloseDb() {
	ndb.onceStopDb.Do(func() {
		err := ndb.Db.Close()
		if err != nil {
			log.Warn(err)
		} else {
			log.Info("db closed.")
		}
	})
}

func NewNdb(dbCfg config.DB) (closer func(), ndb *Ndb) {
	var ndb = &Ndb{}
	onceStartDb.Do(func() {
		log = clog.NewLog(context.NewContext(), "br_go.common.db", "init")

		connStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.DbName,
		)
		var err error
		if Db, err = gorm.Open("mysql", connStr+"&loc=Asia%2FShanghai"); err != nil {
			fmt.Println(err.Error())
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
