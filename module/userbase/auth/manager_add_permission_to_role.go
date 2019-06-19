package auth

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/validator"
	"github.com/krilie/lico_alone/module/userbase/model"
)

//给角色 添加新的permission
func (UserManage) AddPermissionToRole(ctx *context.Context, roleId string, permissionId string) error {
	//检查参数
	if !(validator.IsIdStr(roleId) && validator.IsIdStr(permissionId)) {
		log.Infoln("AddPermissionToRole", "param error:", roleId, permissionId)
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
