package ndb

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/ncfg"
	"github.com/krilie/lico_alone/component/nlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"runtime/debug"
	"sync"
	"time"
)

const gormTransConDb = "gormTransConDb"

type NDb struct {
	cfg struct {
		ConnStr         string
		MaxOpenConn     int
		MaxIdleConn     int
		ConnMaxLeftTime int
	}
	log         *nlog.NLog
	onceStartDb sync.Once
	onceStopDb  sync.Once
	db          *gorm.DB
}

func (ndb *NDb) GetDb(ctx context.Context) *gorm.DB {
	values := context2.GetAppValues(ctx)
	if values == nil {
		return ndb.db
	}
	if values.Tx == nil {
		return ndb.db
	} else {
		return values.Tx.(*gorm.DB)
	}
}

// GetSessionDb 获取上下文中的数据库连接 可以为nil
func (ndb *NDb) GetSessionDb(ctx context.Context) *gorm.DB {
	values := context2.GetAppValues(ctx)
	if values == nil {
		return nil
	}
	if values.Tx == nil {
		return nil
	} else {
		return values.Tx.(*gorm.DB)
	}
}

func (ndb *NDb) Ping() error {
	db, err := ndb.db.DB()
	if err != nil {
		panic(err)
	}
	return db.Ping()
}

func (ndb *NDb) Start() {
	ndb.onceStartDb.Do(func() {
		var err error
		if ndb.db, err = gorm.Open(mysql.Open(ndb.cfg.ConnStr), &gorm.Config{}); err != nil {
			fmt.Println(err.Error())
			ndb.log.Fatal(err, string(debug.Stack())) // 报错退出程序
			return
		} else {
			db, err := ndb.db.DB()
			if err != nil {
				panic(err)
				return
			}
			db.SetMaxOpenConns(ndb.cfg.MaxOpenConn)
			db.SetMaxIdleConns(ndb.cfg.MaxIdleConn)
			db.SetConnMaxLifetime(time.Second * time.Duration(ndb.cfg.ConnMaxLeftTime))
			ndb.log.Info("db init done. params:", "connect string") // 数据库初始化成功
			ndb.db.Logger = &ndbLogger{NLog: ndb.log.WithField("gorm", "gorm-inner")}
			ndb.db.Logger.LogMode(logger.Info)
			ndb.db = ndb.db.Debug()
		}
	})
}

func (ndb *NDb) CloseDb() {
	ndb.onceStopDb.Do(func() {
		db, err2 := ndb.db.DB()
		if err2 != nil {
			panic(err2)
		}
		err := db.Close()
		if err != nil {
			ndb.log.Warn(err)
		} else {
			ndb.log.Info("db closed.")
		}
	})
}

func NewNDb(dbCfg *ncfg.NConfig, log *nlog.NLog) (ndb *NDb) {
	values := context2.NewAppCtxValues()
	values.Module = "ndb"
	values.Function = "NewNDb"
	ctx := context2.NewAppCtx(context.Background(), values)
	log = log.Get(ctx)
	log.Info("no ndb created")
	ndb = &NDb{log: log}
	ndb.cfg.ConnStr = dbCfg.Cfg.DB.ConnStr
	ndb.cfg.MaxOpenConn = dbCfg.Cfg.DB.MaxOpenConn
	ndb.cfg.MaxIdleConn = dbCfg.Cfg.DB.MaxIdleConn
	ndb.cfg.ConnMaxLeftTime = dbCfg.Cfg.DB.ConnMaxLeftTime
	ndb.Start()
	return ndb
}
