package auth

import (
	"context"
	"github.com/deckarep/golang-set"
	"github.com/krilie/lico_alone/common/model/errs"
	"github.com/krilie/lico_alone/common/utils/validator"
	"github.com/krilie/lico_alone/module/user/dao"
	"github.com/krilie/lico_alone/module/user/model"
)

//获取用户的 permission 列表,连表查询
func (UserAuth) GetPermissions(ctx context.Context, userId string) (set mapset.Set, err error) {
	if !validator.IsIdStr(userId) {
		log.Infoln("GetPermissions", "user id format error:", userId)
		return nil, errs.ErrParam
	}
	return model.GetAllPermissionByUserId(dao.Db, userId)
}
