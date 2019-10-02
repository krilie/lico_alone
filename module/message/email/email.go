package email

import "gopkg.in/gomail.v2"

type Email struct {
	Address  string
	Host     string
	Port     int
	UserName string
	Password string
	Dialer   *gomail.Dialer
}

func NewEmail(addr, host string, port int, name, password string) *Email {
	return &Email{
		Address:  addr,
		Host:     host,
		Port:     port,
		UserName: name,
		Password: password,
		Dialer:   gomail.NewDialer(host, port, name, password),
	}
}

// SendServiceUpEmail 发送服务启动消息
func (e *Email) SendEmail(to, subject, msg string) error {
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
