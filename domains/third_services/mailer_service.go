package thirdparties

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"path/filepath"
	manifest "pink/settings"
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
	emailPath := manifest.MailerConfigPath("mail_api_token")
	path, err := filepath.Abs(emailPath)
	if err != nil {
		defer panic(err.Error())
		return false
	}
	err = send(from, path, email)
	if err != nil {
		panic(err.Error())
	}
	return true
}

func send(from string, path string, to string) error {
	domain := fmt.Sprintf("%s:%s", address, port)
	auth := smtp.PlainAuth("", from, password, address)
	message, _ := parseHTMLtoString(path, nil)
	err := smtp.SendMail(domain, auth, from, []string{to}, []byte(message))
	return err
}

func parseHTMLtoString(templatePath string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
