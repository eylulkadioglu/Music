package mailer

import (
	"fmt"

	"github.com/eylulkadioglu/Music/appconfig"
	"gopkg.in/gomail.v2"
)

func SendMail(to string, subject string, body string) error {
	config := appconfig.ReadConfig()

	m := gomail.NewMessage()
	m.SetHeader("From", config.MailerFrom)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	//
	d := gomail.NewDialer(
		config.MailerAddress,
		config.MailerPort,
		config.MailerFrom,
		config.MailerPassword,
	)
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("Can't send email, err: %v\n", err)
		return err
	}

	return nil
}
