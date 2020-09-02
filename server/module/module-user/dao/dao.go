package dao

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/component/ndb"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/krilie/lico_alone/module/module-user/model"
)

type UserDao struct {
	log *nlog.NLog
	*ndb.NDb
}

func NewUserDao(db *ndb.NDb, log *nlog.NLog) *UserDao {
	log = log.WithField(context_enum.Module.Str(), "module user dao")
	//err := db.GetDb(context2.NewContext()).AutoMigrate(
	//	&model.Permission{},
	//	&model.RolePermission{},
	//	&model.Role{},
	//	&model.UserRole{},
	//	&model.UserMaster{})
	//if err != nil {
	//	panic(err)
	//}
	return &UserDao{
		log: log,
		NDb: db,
	}
}

func (d *UserDao) DeleteAllUserData(ctx context.Context) (err error) {
	err = d.Transaction(ctx, func(ctx context.Context) error {
		err = d.GetDb(ctx).Delete(new(model.UserMaster)).Error
		if err != nil {
			return err
		}
		err = d.GetDb(ctx).Delete(new(model.UserRole)).Error
		if err != nil {
			return err
		}
		err = d.GetDb(ctx).Delete(new(model.Role)).Error
		if err != nil {
			return err
		}
		err = d.GetDb(ctx).Delete(new(model.RolePermission)).Error
		if err != nil {
			return err
		}
		err = d.GetDb(ctx).Delete(new(model.Permission)).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
