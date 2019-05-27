package auth

import (
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/common/validator"
	"github.com/krilie/lico_alone/module/userbase/model"
)

//给用户添加新角色
func (UserManage) AddRoleToUser(ctx *context.Context, roleId string, userId string) error {
	//检查参数
	if !(validator.IsIdStr(roleId) && validator.IsIdStr(userId)) {
		log.Infoln("AddPermissionToRole", "param error:", roleId, userId)
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
