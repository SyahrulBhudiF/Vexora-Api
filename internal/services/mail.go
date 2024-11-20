package services

import (
	"gopkg.in/gomail.v2"
)

type MailService struct {
	mail gomail.Dialer
}

func NewMailService(host string, port int, email string, password string) *MailService {
	mail := gomail.Dialer{
		Host:     host,
		Port:     port,
		Username: email,
		Password: password,
	}

	return &MailService{mail: mail}
}

func (m *MailService) SendMail(to string, subject string, otp string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", "ryu4w@gmail.com")
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", "Your OTP is "+otp)

	if err := m.mail.DialAndSend(mail); err != nil {
		return err
	}

	return nil
}
