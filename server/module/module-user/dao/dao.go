package dao

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-user/model"
)

type UserDao struct {
	log *nlog.NLog
	*ndb.NDb
}

func NewUserDao(db *ndb.NDb, log *nlog.NLog) *UserDao {

	db.GetDb(context.NewContext()).AutoMigrate(
		&model.Permission{},
		&model.RolePermission{},
		&model.Role{},
		&model.UserRole{},
		&model.UserMaster{})

	return &UserDao{
		log: log,
		NDb: db,
	}
}
