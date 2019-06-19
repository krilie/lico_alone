package auth

import (
	"github.com/asaskevich/govalidator"
	"github.com/krilie/lico_alone/common/comstruct/errs"
	"github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/utils/id_util"
	"github.com/krilie/lico_alone/module/userbase/model"
	"time"
)

//给系统添加新的permission permission项,默认权限检查已经通过
func (UserManage) NewPermission(ctx context.Context, pName string, pDescription string) (p *model.Permission, err error) {
	//检查参数
	if !govalidator.IsAlpha(pName) || len(pDescription) == 0 {
		log.Infoln("", "param error:", pName, pDescription)
		return nil, errs.ErrParam
	}
	//添加一个
	p = new(model.Permission)
	p.ID = id_util.GetUuid()
	p.CreateTime = time.Now()
	p.Description = pDescription
	p.Name = pName
	err = model.Db.Create(p).Error
	if err != nil {
		return nil, err
	} else {
		return p, nil
	}
}
