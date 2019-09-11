package email

import (
	"context"
	"gopkg.in/gomail.v2"
)

// 阿里云
// 用户 livo@amail.lizo.top
// 密码 asdfa1321321EERWE
// 一天免费两百个

var address = "livo@amail.lizo.top"
var name = "livo@amail.lizo.top"
var password = "asdfa1321321EERWE"
var host = "smtpdm.aliyun.com"

// SendServiceUpEmail 发送服务启动消息
func SendEmail(ctx context.Context, to, subject, msg string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", address)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", msg)
	d := gomail.NewDialer(host, 465, name, password)
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
