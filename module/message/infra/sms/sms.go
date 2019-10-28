package sms

import (
	"context"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/krilie/lico_alone/common/utils/str_util"
)

// 用户登录名称
//AccessKey ID
//AccessKeySecret
//sms@1297016936336588.onaliyun.com
//LTAI4FcUBH2hxNLHaJ6JVZXm
//yMlaDzBBb6ImpwTBidoVjz2B1EA4N6

type AliSms struct {
	Key    string
	Secret string
	client *dysmsapi.Client
}

func NewAliSms(key, secret string) *AliSms {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", key, secret)
	if err != nil {
		panic(err)
	}
	return &AliSms{
		Key:    key,
		Secret: secret,
		client: client,
	}
}

func (a *AliSms) SendRegisterSms(ctx context.Context, code string) error {
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = "18761438228"
	request.SignName = "迅如雨"
	request.TemplateCode = "SMS_173946021"
	request.TemplateParam = str_util.ToJson(map[string]string{"code": code})

	response, err := a.client.SendSms(request)
	if err != nil {
		return err
	}
	if response.Message != "OK" {
		return errors.New(response.Message)
	}
	return nil
}

type IAliSms interface {
	SendRegisterSms(ctx context.Context, code string) error
}
