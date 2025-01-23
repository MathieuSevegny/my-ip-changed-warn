package src

import (
	"log"
	"net/smtp"
)

type EmailClient struct {
	From  string
	To    string
	Token string
}

func (client *EmailClient) SendEmail(subject, body string) error {
	from := client.From
	pass := client.Token
	to := client.To

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	return nil
}
