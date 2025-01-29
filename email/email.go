package email

import (
	"fmt"
	"net/smtp"
)

type EmailSender interface {
	SendMail(to, subject, body string) error
}

type emailSender struct {
	from string
	addr string
	auth smtp.Auth
}

func NewMailSender(from, pw, host, port string) EmailSender {
	auth := smtp.PlainAuth("", from, pw, host)
	addr := fmt.Sprintf("%s:%s", host, port)
	return &emailSender{from, addr, auth}
}

func (e *emailSender) SendMail(to, subject, body string) error {
	message := fmt.Sprintf("Subject: %s\r\n\r\n%s\r\n", subject, body)
	err := smtp.SendMail(e.addr, e.auth, e.from, []string{to}, []byte(message))
	if err != nil {
		return err
	}
	return nil
}
