package server

import (
	"strconv"

	"gopkg.in/gomail.v2"
)

//发送邮件
func Sendmail(mailTo string, subject string, body string) error {
	//邮箱服务器信息
	mailConn := map[string]string{
		"user": "service@ri-co.cn",
		"pass": ConfInfo.Pass,
		"host": "smtp.mxhichina.com",
		"port": "465",
	}
	port, _ := strconv.Atoi(mailConn["port"])
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(mailConn["user"], "Rico"))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err
}
