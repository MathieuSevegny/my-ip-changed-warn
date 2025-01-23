package env

import "os"

type EnvVariables struct {
	API_ENDPOINT      string
	CACHE_FOLDER_PATH string
	CACHE_FILENAME    string
	EMAIL_TO          string
	EMAIL_FROM        string
	EMAIL_TOKEN       string
	SMTP_HOST         string
	SECONDS_TO_WAIT   string
	DEVICE_NAME       string
	MAX_TRIES         string
}

func ReadEnv() *EnvVariables {
	return &EnvVariables{
		API_ENDPOINT:      os.Getenv("API_ENDPOINT"),
		CACHE_FOLDER_PATH: os.Getenv("CURRENT_FOLDER_PATH"),
		CACHE_FILENAME:    os.Getenv("CACHE_FILENAME"),
		EMAIL_TO:          os.Getenv("EMAIL_TO"),
		EMAIL_FROM:        os.Getenv("EMAIL_FROM"),
		EMAIL_TOKEN:       os.Getenv("EMAIL_TOKEN"),
		SMTP_HOST:         os.Getenv("SMTP_HOST"),
		SECONDS_TO_WAIT:   os.Getenv("SECONDS_TO_WAIT"),
		DEVICE_NAME:       os.Getenv("DEVICE_NAME"),
		MAX_TRIES:         os.Getenv("MAX_TRIES"),
	}
}
