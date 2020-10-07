package thirdparties

import (
	"fmt"
	"net/smtp"
	configs "pink/settings/configs"
)

var config = configs.MailerConfig()
var from = config.UserName
var password = config.Password
var domain = config.Domain
var address = config.Address
var port = config.Port

// MailerServiceInf interface
type MailerServiceInf interface {
	MailAPIToken() bool
}

// MailerService struct
type MailerService struct {
}

// MailAPIToken func
func (service MailerService) MailAPIToken(email string) bool {
	fmt.Print(config)
	message := "Hello"
	err := send(from, message, email)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func send(from string, message string, to string) error {
	domain := fmt.Sprintf("%s:%s", address, port)
	err := smtp.SendMail(domain, smtp.PlainAuth("", from, password,
		address), from, []string{to}, []byte(message))
	return err
}
