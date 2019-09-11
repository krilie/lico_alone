package email

import "gopkg.in/gomail.v2"

// 密码 asdfa1321321EERWE

func SendEmail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "livo@amail.lizo.top")
	m.SetHeader("To", "776334655@qq.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("/home/Alex/lolcat.jpg")
	d := gomail.NewDialer("smtpdm.aliyun.com", 465, "livo", "asdfa1321321EERWE")
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
