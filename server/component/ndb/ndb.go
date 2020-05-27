package ndb

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/krilie/lico_alone/common/config"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/nlog"
	"runtime/debug"
	"sync"
	"time"
)

const gormTransConDb = "gormTransConDb"

type NDb struct {
	cfg         config.DB
	log         *nlog.NLog
	onceStartDb sync.Once
	onceStopDb  sync.Once
	db          *gorm.DB
}

func (ndb *NDb) GetDb(ctx context.Context) *gorm.DB {
	nCtx := context2.GetContextOrNil(ctx)
	if nCtx == nil {
		return ndb.db
	} else {
		if nCtx.Tx == nil {
			return ndb.db
		} else {
			return nCtx.Tx.(*gorm.DB)
		}
	}
}

func (ndb *NDb) Start() {
	ndb.onceStartDb.Do(func() {
		connStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
			ndb.cfg.User,
			ndb.cfg.Password,
			ndb.cfg.Host,
			ndb.cfg.Port,
			ndb.cfg.DbName,
		)
		var err error
		if ndb.db, err = gorm.Open("mysql", connStr+"&loc=Asia%2FShanghai"); err != nil {
			fmt.Println(err.Error())
			ndb.log.Fatal(err, string(debug.Stack())) // 报错退出程序
			return
		} else {
			ndb.db.DB().SetMaxOpenConns(ndb.cfg.MaxOpenConn)
			ndb.db.DB().SetMaxIdleConns(ndb.cfg.MaxIdleConn)
			ndb.db.DB().SetConnMaxLifetime(time.Second * time.Duration(ndb.cfg.ConnMaxLeftTime))
			ndb.log.Info("db init done. params:", connStr+"&loc=Asia%2FShanghai") // 数据库初始化成功
			ndb.db = ndb.db.Debug()
			ndb.db.LogMode(true)
		}
	})
}

func (ndb *NDb) CloseDb() {
	ndb.onceStopDb.Do(func() {
		err := ndb.db.Close()
		if err != nil {
			ndb.log.Warn(err)
		} else {
			ndb.log.Info("db closed.")
		}
	})
}

func NewNDb(dbCfg *config.Config, log *nlog.NLog) (ndb *NDb) {
	log.Info("no ndb created")
	ndb = &NDb{log: log, cfg: dbCfg.DB}
	ndb.Start()
	return ndb
}

type Dao struct {
	*NDb
	log *nlog.NLog
}
