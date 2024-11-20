package _interface

type IMailService interface {
	// SendMail sends an email to the specified email address with the specified subject and OTP
	SendMail(to string, subject string, otp string) error
}
