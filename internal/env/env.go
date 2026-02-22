package env

import (
	"fmt"
	"log"
	"os"
	"time"
)

type EnvVariables struct {
	ApiEndpoint  string
	DataFilePath string
	EmailTo      string
	EmailFrom    string
	EmailToken   string
	SmtpHost     string
	WaitTime     time.Duration
	DeviceName   string
}

var RequiredEnvVariables = []string{
	"EMAIL_TO",
	"EMAIL_FROM",
	"EMAIL_TOKEN",
	"SMTP_HOST",
}

func ReadEnv() (*EnvVariables, error) {
	waitTime := 5 * time.Minute
	if givenValue, err := time.ParseDuration(os.Getenv("WAIT_TIME")); err != nil {
		log.Printf("error parsing duration: %s, default will be used (%s).", err, waitTime.String())
	} else {
		waitTime = givenValue
	}
	apiEndpoint := "https://api.ipify.org/"
	if givenValue := os.Getenv("API_ENDPOINT"); givenValue != "" {
		apiEndpoint = givenValue
	}
	dataFilePath := "data/current_ip.txt"
	if givenValue := os.Getenv("DATA_FILE_PATH"); givenValue != "" {
		dataFilePath = givenValue
	}
	deviceName := "your server"
	if givenValue := os.Getenv("DEVICE_NAME"); givenValue != "" {
		deviceName = givenValue
	}

	missingEnvVar := false
	for _, envVar := range RequiredEnvVariables {
		if os.Getenv(envVar) == "" {
			log.Printf("required environment variable %s is not set", envVar)
		}
	}

	if missingEnvVar {
		return nil, fmt.Errorf("one or more required environment variables are missing")
	}

	return &EnvVariables{
		ApiEndpoint:  apiEndpoint,
		DataFilePath: dataFilePath,
		EmailTo:      os.Getenv("EMAIL_TO"),
		EmailFrom:    os.Getenv("EMAIL_FROM"),
		EmailToken:   os.Getenv("EMAIL_TOKEN"),
		SmtpHost:     os.Getenv("SMTP_HOST"),
		WaitTime:     waitTime,
		DeviceName:   deviceName,
	}, nil
}
