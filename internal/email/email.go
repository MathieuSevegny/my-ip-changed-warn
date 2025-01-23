package email

import (
	"crypto/tls"
	"log"

	"gopkg.in/gomail.v2"
)

type EmailClient struct {
	From  string
	To    string
	Token string
	Host  string
}

func (client *EmailClient) SendEmail(subject string, body string) error {
	from := client.From
	pass := client.Token
	to := client.To

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(client.Host, 587, from, pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
