package email

import (
	"context"
	"gopkg.in/gomail.v2"
)

type Email struct {
	Dialer *gomail.Dialer

	Address  string
	Host     string
	Port     int
	UserName string
	Password string
}

func NewEmail(addr, host string, port int, name, password string) *Email {
	return &Email{
		Dialer:   gomail.NewDialer(host, port, name, password),
		Address:  addr,
		Host:     host,
		Port:     port,
		UserName: name,
		Password: password,
	}
}

// SendServiceUpEmail 发送服务启动消息
func (e *Email) SendEmail(ctx context.Context, to, subject, msg string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.Address)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", msg)
	if err := e.Dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
