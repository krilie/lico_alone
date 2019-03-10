package user_auth

import (
	"github.com/asaskevich/govalidator"
	"github.com/lico603/lico-my-site-user/common/errs"
	"github.com/lico603/lico-my-site-user/common/log"
	"github.com/lico603/lico-my-site-user/common/validator_util"
	"github.com/lico603/lico-my-site-user/model"
)

//用户是否有这个角色,有app权限才能调用这个接口
func UserAuthHasRole(userId, roleName string) (bool, error) {
	//参数检查
	if len(roleName) == 0 ||
		(!govalidator.IsAlpha(roleName)) ||
		!(validator_util.IsIdString(userId)) {
		log.Error("UserAuthHasRole", "参数格式不正确")
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
