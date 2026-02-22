package email

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

type EmailClient struct {
	From  string
	To    string
	Token string
	Host  string
}

func (client *EmailClient) SendEmail(subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", client.From)
	m.SetHeader("To", client.To)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(client.Host, 587, client.From, client.Token)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}
