package service_common

import (
	"context"
	"github.com/krilie/lico_alone/module/module-dynamic-share/model"
)

func (a *CommonService) QueryDynamicShare(ctx context.Context, param model.QueryDynamicShareModel) (list *model.QueryDynamicShareResModel, err error) {
	return a.moduleDynamicShare.QueryDynamicShare(ctx, param)
}
