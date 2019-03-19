package manager

import (
	"github.com/lico603/lico-my-site-user/common/context_util"
	"github.com/lico603/lico-my-site-user/common/log"
	"github.com/lico603/lico-my-site-user/common/validator_util"
	"github.com/lico603/lico-my-site-user/model"
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
