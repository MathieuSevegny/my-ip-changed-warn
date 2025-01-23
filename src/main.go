package src

import (
	"log"
	"strconv"
	"time"
)

func main() {
	env := ReadEnv()

	cache := &CacheProvider{
		folder:   env.CACHE_FOLDER_PATH,
		filename: env.CACHE_FILENAME,
	}
	api := &ApiClient{
		Endpoint: env.API_ENDPOINT,
	}
	email := &EmailClient{
		From:  env.EMAIL_FROM,
		To:    env.EMAIL_TO,
		Token: env.EMAIL_TOKEN,
	}

	waitTime, err := time.ParseDuration(env.SECONDS_TO_WAIT + "s")

	if err != nil {
		log.Fatalf("error parsing duration: %s", err)
	}

	numberOfTries := 0
	maxTries, err := strconv.ParseUint(env.MAX_TRIES, 10, 64)

	if err != nil {
		maxTries = 5
		log.Printf("error parsing max tries, default will be used (%d): %s", maxTries, err)
	}

	for {
		oldIp, err := cache.Get()

		if err != nil {
			log.Printf("error getting public ip: %s", err)
			numberOfTries++
			time.Sleep(waitTime)
			continue
		}

		currentIp, err := api.GetPublicIp()

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

		err = email.SendEmail("Public IP of "+env.DEVICE_NAME+" changed", "New public IPv4 address : "+currentIp)

		if err != nil {
			log.Printf("error sending email: %s", err)
			numberOfTries++
		} else {
			// Save the new IP to the cache only if the email was sent successfully
			err = cache.Save(currentIp)
			if err != nil {
				log.Fatalf("error saving cache: %s", err)
			}
		}

		time.Sleep(waitTime)
	}

}
