package service

import (
	"context"
	"github.com/krilie/lico_alone/module/module-dynamic-share/model"
)

func (a *DynamicShareModule) QueryDynamicShare(ctx context.Context, param model.QueryDynamicShareModel) (*model.QueryDynamicShareResModel, error) {
	return a.Dao.QueryDynamicShare(ctx, param)
}
