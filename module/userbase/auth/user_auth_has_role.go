package auth

import (
	"github.com/asaskevich/govalidator"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/common/validator"
	"github.com/krilie/lico_alone/module/userbase/model"
)

//用户是否有这个角色,有app权限才能调用这个接口
func (UserAuth) HasRole(ctx *context.Context, userId, roleName string) (bool, error) {
	//参数检查
	if len(roleName) == 0 ||
		(!govalidator.IsAlpha(roleName)) ||
		!(validator.IsIdStr(userId)) {
		log.Error("HasRole", "参数格式不正确")
		return false, errs.ErrParam
	}
	//直接取到role id
	roleId, err := model.GetRoleIdByName(model.Db, roleName)
	if err != nil {
		return false, err
	}
	//检查是否存在对应关系
	exist, err := model.IsUserRoleRelationExist(model.Db, userId, roleId)
	if err != nil {
		return false, err
	}
	return exist, nil
}
