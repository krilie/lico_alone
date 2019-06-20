package user

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/module/userbase/model"
	"time"
)

func (User) ValidateClientAccToken(ctx context.Context, clientAccKey string) (key *model.ClientUserAccessToken, err error) {
	key = new(model.ClientUserAccessToken)
	err = model.Db.Find(key, "token=?", clientAccKey).Error
	if err != nil && err == gorm.ErrRecordNotFound {
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
