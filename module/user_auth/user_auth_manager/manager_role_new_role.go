package user_auth_manager

import (
	"github.com/asaskevich/govalidator"
	"github.com/krilie/lico_alone/common/common_struct/errs"
	"github.com/krilie/lico_alone/common/context_util"
	"github.com/krilie/lico_alone/common/log"
	"github.com/krilie/lico_alone/common/uuid_util"
	"github.com/krilie/lico_alone/module/user_auth/model"
	"time"
)

//给系统添加新的角色
func ManagerRoleNewRole(ctx *context_util.Context, roleName string, roleDescription string) (role *model.Role, err error) {
	//检查参数
	if !govalidator.IsAlpha(roleName) || len(roleDescription) == 0 {
		log.Infoln("ManagerRoleNewRole", "param err:", roleName, roleDescription)
		return nil, errs.ErrParam
	}
	//生成新角色并添加
	role = new(model.Role)
	role.ID = uuid_util.GetUuid()
	role.Version = 0
	role.Name = roleName
	role.Description = roleDescription
	role.CreateTime = time.Now()
	err = model.Db.Create(role).Error
	if err != nil {
		return nil, err
	} else {
		return role, nil
	}
}
