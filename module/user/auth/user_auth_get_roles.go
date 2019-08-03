package auth

import (
	"context"
	"github.com/deckarep/golang-set"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/utils/validator"
	"github.com/krilie/lico_alone/module/user/model"
)

//获取用户的所有角色
func (UserAuth) GetRoles(ctx context.Context, userId string) (roles mapset.Set, err error) {
	if !validator.IsIdStr(userId) {
		return nil, errs.ErrParam
	}
	return model.GetAllRolesByUserId(model.Db, userId)
}
