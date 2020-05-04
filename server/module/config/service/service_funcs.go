package service

import (
	"context"
	"github.com/krilie/lico_alone/common/errs"
	"strconv"
)

func (a *ConfigService) GetValueInt(ctx context.Context, name string) (*int, error) {
	valueStr, err := a.GetValueStr(ctx, name)
	if err != nil {
		return nil, err
	}
	if valueStr == nil {
		return nil, nil
	}
	val, err := strconv.Atoi(*valueStr)
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	return &val, nil
}
func (a *ConfigService) SetValueInt(ctx context.Context, name string, value int) error {
	return a.SetValueStr(ctx, name, strconv.Itoa(value))
}

func (a *ConfigService) GetValueBool(ctx context.Context, name string) (*bool, error) {
	str, err := a.GetValueStr(ctx, name)
	if err != nil {
		return nil, err
	}
	if str == nil {
		return nil, nil
	}
	b, err := strconv.ParseBool(*str)
	if err != nil {
		return nil, errs.NewInternal().WithError(err)
	}
	return &b, err
}
func (a *ConfigService) SetValueBool(ctx context.Context, name string, value bool) error {
	return a.SetValueStr(ctx, name, strconv.FormatBool(value))
}
