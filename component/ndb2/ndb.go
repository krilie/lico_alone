package ndb

import (
	"context"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sync"
	"time"
)

// IDb interface that wrap of db operation
type IDb interface {
	sqlx.ExtContext
	sqlx.Ext
}

// MustToSqlxTx convert IDb to *sqlx.Tx if not will panic
func MustToSqlxTx(db IDb) *sqlx.Tx {
	tx, ok := db.(*sqlx.Tx)
	if !ok {
		panic(errors.New("db is not a *sqlx.tx"))
	}
	return tx
}

// MustToSqlxDb convert IDb to *sqlx.DB if not will panic
func MustToSqlxDb(db IDb) *sqlx.DB {
	tx, ok := db.(*sqlx.DB)
	if !ok {
		panic(errors.New("db is not a *sqlx.tx"))
	}
	return tx
}

// NDb struct wrap of sql db config and start close
type NDb struct {
	cfg struct {
		ConnStr         string
		MaxOpenConn     int
		MaxIdleConn     int
		ConnMaxLeftTime int
	}
	onceStartDb sync.Once
	onceStopDb  sync.Once
	sqlxDb      IDb // *sqlx.DB
}

func NewNDb(
	MaxOpenConn int, MaxIdleConn int,
	ConnMaxLeftTime int, ConnStr string,
) *NDb {
	db := &NDb{}
	db.cfg.ConnStr = ConnStr
	db.cfg.ConnMaxLeftTime = ConnMaxLeftTime
	db.cfg.MaxIdleConn = MaxIdleConn
	db.cfg.MaxOpenConn = MaxOpenConn
	db.Start()
	return db
}

func (ndb *NDb) Ping() error {
	return ndb.sqlxDb.(*sqlx.DB).Ping()
}

func (ndb *NDb) Start() {
	ndb.onceStartDb.Do(func() {
		db, err := sqlx.Connect("mysql", ndb.cfg.ConnStr)
		if err != nil {
			panic(err)
		}
		db.SetConnMaxIdleTime(time.Second * time.Duration(ndb.cfg.ConnMaxLeftTime))
		db.SetConnMaxLifetime(time.Hour * 6)
		db.SetMaxIdleConns(ndb.cfg.MaxIdleConn)
		db.SetMaxOpenConns(ndb.cfg.MaxOpenConn)
		err = db.Ping()
		if err != nil {
			panic(err)
		}
		ndb.sqlxDb = db
	})
}

func (ndb *NDb) CloseDb() {
	ndb.onceStopDb.Do(func() {
		err := ndb.sqlxDb.(*sqlx.DB).Close()
		if err != nil {
			panic(err)
		}
	})
}

// GetDb get db before use
// if in transaction this function will return tx set on context or. return sqlx.db
func (ndb *NDb) GetDb(ctx context.Context) IDb {
	txDb := GetTxDbFromCtx(ctx)
	if txDb != nil {
		return txDb
	}
	return ndb.sqlxDb
}

// GetGlobalDb get db before use
// get ori db of global which can begin a new trans
func (ndb *NDb) GetGlobalDb(ctx context.Context) IDb {
	return ndb.sqlxDb
}

// WithTrans start trans with db on context
func WithTrans(ctx context.Context, trans func(ctx context.Context) error, onNewTrans ...bool) (err error) {
	// 环境变量
	isOnNewTrans := len(onNewTrans) >= 1 && onNewTrans[0]
	globalDb := MustGetGlobalDbFromCtx(ctx)
	oldTransDb := GetTxDbFromCtx(ctx)
	// 执行新事务
	var doTransOnNewSession = func() error {
		// 准备环境
		newTransDb := MustToSqlxDb(globalDb).MustBeginTx(ctx, nil)    // 新的事务对象
		newTransCtx := ForceNewDbValuesCtx(ctx, globalDb, newTransDb) // 新的上下文 覆盖旧的上下文
		// panic or err 回滚
		panicked := true
		defer func() {
			if panicked || err != nil {
				err = errors.New("err or panic on trans " + err.Error())
				err2 := newTransDb.Rollback()
				if err2 != nil {
					err = errors.New("err on trans rollback " + err.Error() + err2.Error())
				}
			} else {
				err := newTransDb.Commit()
				if err != nil {
					panic(errors.New("err on commit " + err.Error()))
				}
			}
		}()
		// 执行事务代码并返回
		err = trans(newTransCtx)
		panicked = false
		return err
	}
	// 执行代码
	if isOnNewTrans {
		return doTransOnNewSession()
	} else {
		if oldTransDb == nil {
			return doTransOnNewSession()
		} else {
			// panic or err 回滚
			panicked := true
			defer func() {
				if panicked || err != nil {
					err = errors.New("err or panic on trans (inner trans no commit) " + err.Error())
				}
			}()
			// 执行事务代码并返回
			err = trans(ctx)
			panicked = false
			return err
		}
	}
}
