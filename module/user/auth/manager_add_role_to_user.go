package auth

import (
	"context"
	"github.com/krilie/lico_alone/common/utils/validator"
	"github.com/krilie/lico_alone/module/user/dao"
	"github.com/krilie/lico_alone/module/user/model"
)

//给用户添加新角色
func (UserManage) AddRoleToUser(ctx context.Context, roleId string, userId string) error {
	//检查参数
	if !(validator.IsIdStr(roleId) && validator.IsIdStr(userId)) {
		log.Infoln("AddPermissionToRole", "param error:", roleId, userId)
	}
	//添加关系
	var relation model.UserRole
	relation.RoleID = roleId
	relation.UserID = userId
	err := dao.Db.Create(&relation).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
