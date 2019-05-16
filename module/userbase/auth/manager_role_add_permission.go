package auth

import (
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/common/validator_util"
	"github.com/krilie/lico_alone/module/userbase/model"
)

//给角色 添加新的permission
func (Manage) ManagerRoleAddPermission(ctx *context_util.Context, roleId string, permissionId string) error {
	//检查参数
	if !(validator_util.IsIdStr(roleId) && validator_util.IsIdStr(permissionId)) {
		log.Infoln("ManagerRoleAddPermission", "param error:", roleId, permissionId)
	}
	//添加关系
	var relation model.RolePermission
	relation.RoleID = roleId
	relation.PermissionID = permissionId
	err := model.Db.Create(&relation).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
