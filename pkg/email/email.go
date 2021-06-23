package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type SMTPInfo struct {
	Host string
	Port int
	IsSSL bool
	UserName string
	Password string
	Form string
}

type Email struct {
	*SMTPInfo
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{
		SMTPInfo: info,
	}
}

func (e *Email) SendMail(to []string, subject, body string) error {
	//新建一个信息实力
	m := gomail.NewMessage()
	//发件人
	m.SetHeader("Form:", e.Form)
	//收件人
	m.SetHeader("To", to...)
	//主题
	m.SetHeader("Subject", subject)
	//内容
	m.SetBody("text/html", body)

	//创建一个新的SMTP拨号实例
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}

	//打开与stmp服务器的链接并发送电子邮件
	return dialer.DialAndSend(m)
}

