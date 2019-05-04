package user_auth_manager

import (
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/common/validator_util"
	"github.com/krilie/lico_alone/module/user_auth/model"
)

//给用户添加新角色
func ManagerUserAddRole(ctx *context_util.Context, roleId string, userId string) error {
	//检查参数
	if !(validator_util.IsIdStr(roleId) && validator_util.IsIdStr(userId)) {
		log.Infoln("ManagerRoleAddPermission", "param error:", roleId, userId)
	}
	//添加关系
	var relation model.UserRole
	relation.RoleID = roleId
	relation.UserID = userId
	err := model.Db.Create(&relation).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
