package user_base

import (
	"database/sql"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/module/user_auth/model"
	"time"
)

func UserValidateClientAccToken(ctx *context_util.Context, clientAccKey string) (key *model.ClientUserAccessToken, err error) {
	row := model.Db.Where("token=?", clientAccKey).Row()
	key = new(model.ClientUserAccessToken)
	err = row.Scan(key)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	//检查是否过期
	if key.ExpirationTime.Before(time.Now()) {
		return nil, errs.ErrClientAccTokenExp //过期时间 在 当前时间 之前，过期了
	}
	//检查是否可用
	if !key.IsValid {
		return nil, errs.ErrClientAccTokenNotValid
	}
	return key, nil
}
