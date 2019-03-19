package manager

import (
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/log"
	"github.com/lico603/lico-my-site-user/common/validator_util"
	"github.com/lico603/lico-my-site-user/model"
)

//给角色 添加新的permission
func ManagerRoleAddPermission(ctx *context_util.Context, roleId string, permissionId string) error {
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
