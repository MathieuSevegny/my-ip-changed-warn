package main

import (
	"log"
	"my-ip-changed-warn/internal/api"
	"my-ip-changed-warn/internal/cache"
	"my-ip-changed-warn/internal/email"
	"my-ip-changed-warn/internal/env"
	"strconv"
	"time"
)

func main() {
	envValues := env.ReadEnv()

	cacheProvider := &cache.CacheProvider{
		Folder:   envValues.CACHE_FOLDER_PATH,
		Filename: envValues.CACHE_FILENAME,
	}
	apiClient := &api.ApiClient{
		Endpoint: envValues.API_ENDPOINT,
	}
	emailClient := &email.EmailClient{
		From:  envValues.EMAIL_FROM,
		To:    envValues.EMAIL_TO,
		Token: envValues.EMAIL_TOKEN,
		Host:  envValues.SMTP_HOST,
	}

	var waitTime time.Duration = 5 * time.Second

	givenValue, err := time.ParseDuration(envValues.SECONDS_TO_WAIT + "s")

	if err != nil {
		log.Printf("error parsing duration: %s", err)
	} else {
		waitTime = givenValue
	}

	numberOfTries := 0
	var maxTries uint64 = 5

	number, err := strconv.ParseUint(envValues.MAX_TRIES, 10, 64)

	if err != nil {
		maxTries = number
		log.Printf("error parsing max tries, default will be used (%d): %s", maxTries, err)
	}

	log.Print("Listening for IP changes...")
	for {
		oldIp, err := cacheProvider.Get()

		if err != nil {
			log.Printf("error getting public ip: %s", err)
			numberOfTries++
			time.Sleep(waitTime)
			continue
		}

		currentIp, err := apiClient.GetPublicIp()

		if err != nil {
			log.Printf("error getting public ip: %s", err)
			numberOfTries++
			time.Sleep(waitTime)
			continue
		}

		if oldIp == currentIp {
			time.Sleep(waitTime)
			continue
		}
		log.Printf("old ip: %s, current ip: %s", oldIp, currentIp)

		err = emailClient.SendEmail("Public IP of "+envValues.DEVICE_NAME+" changed", "New public IPv4 address : "+currentIp)

		if err != nil {
			log.Printf("error sending email: %s", err)
			numberOfTries++
		} else {
			log.Printf("email sent successfully!")
			// Save the new IP to the cache only if the email was sent successfully
			err = cacheProvider.Save(currentIp)
			if err != nil {
				log.Fatalf("error saving cache: %s", err)
			}
		}

		time.Sleep(waitTime)
	}

}
