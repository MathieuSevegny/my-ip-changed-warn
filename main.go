package main

import (
	"fmt"
	"log"
	"my-ip-changed-warn/internal/api"
	"my-ip-changed-warn/internal/data"
	"my-ip-changed-warn/internal/email"
	"my-ip-changed-warn/internal/env"
	"time"
)

func main() {
	envValues, err := env.ReadEnv()

	if err != nil {
		log.Fatalf("Error reading environment variables: %s", err)
	}

	dataProvider := &data.DataProvider{
		FilePath: envValues.DataFilePath,
	}
	apiClient := &api.ApiClient{
		Endpoint: envValues.ApiEndpoint,
	}
	emailClient := &email.EmailClient{
		From:  envValues.EmailFrom,
		To:    envValues.EmailTo,
		Token: envValues.EmailToken,
		Host:  envValues.SmtpHost,
	}

	ticker := time.NewTicker(envValues.WaitTime)

	log.Printf("Listening for IP changes periodically (every %s). Press Ctrl+C to exit.", envValues.WaitTime.String())
	for range ticker.C {
		oldIp, err := dataProvider.Get()
		if err != nil {
			log.Printf("Error getting public ip from storage: %s", err)
			continue
		}

		currentIp, err := apiClient.GetPublicIp()
		if err != nil {
			log.Printf("Error getting public ip: %s", err)
			continue
		}
		if oldIp == currentIp {
			continue
		}

		log.Printf("Detected IP change! Old IP: %s, current IP: %s", oldIp, currentIp)

		subject := fmt.Sprintf("[WARN] Public IP of %s changed", envValues.DeviceName)
		message := fmt.Sprintf("The new public IPv4 address of %s is now : %s", envValues.DeviceName, currentIp)
		if err = emailClient.SendEmail(subject, message); err != nil {
			log.Printf("Error sending email: %s", err)
		} else {
			log.Printf("Email sent successfully!")
			// Save the new IP to the cache only if the email was sent successfully
			err = dataProvider.Save(currentIp)
			if err != nil {
				log.Printf("error saving public ip to storage: %s", err)
				continue
			}
		}
	}
}
