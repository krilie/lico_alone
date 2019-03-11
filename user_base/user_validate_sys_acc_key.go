package user_base

import (
	"database/sql"
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/errs"
	"github.com/lico603/lico-my-site-user/model"
	"time"
)

func UserValidateSysAccKey(ctx *context_util.Context, sysAccKey string) (key *model.SysUserAccessToken, err error) {
	row := model.Db.Where("token=?", sysAccKey).Row()
	key = new(model.SysUserAccessToken)
	err = row.Scan(key)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	//检查是否过期
	if key.ExpirationTime.Before(time.Now()) {
		return nil, errs.ErrParam //过期时间 在 当前时间 之前，过期了
	}
	//检查是否可用
	if !key.IsValid {
		return nil, errs.ErrInternal
	}
	return key, nil
}
