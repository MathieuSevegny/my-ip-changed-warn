package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "golang.org/x/crypto/x509roots/fallback"
)

type ApiClient struct {
	Endpoint string
}

const MAX_RETRY_WAIT_TIME = time.Hour // Maximum wait time of 1 hour to avoid waiting indefinitely in case of persistent issues

func (client ApiClient) GetPublicIp() (string, error) {
	resp, err := http.Get(client.Endpoint)
	if err != nil {
		return "", fmt.Errorf("error making GET request: %w", err)
	}

	// Handle rate limiting (HTTP 429 Too Many Requests)
	if resp.StatusCode == http.StatusTooManyRequests {
		askedToRetryAfter := resp.Header.Get("Retry-After")
		log.Printf("too many requests, asked to retry after: %s seconds, will wait accordingly...", askedToRetryAfter)
		duration, err := time.ParseDuration(askedToRetryAfter + "s")
		if err != nil {
			log.Printf("error parsing Retry-After header: %s, will wait for 1 minute by default", err)
			duration = time.Minute
		}
		if duration > MAX_RETRY_WAIT_TIME {
			log.Printf("Retry-After duration is too long (%s), capping it to %s", duration, MAX_RETRY_WAIT_TIME)
			duration = MAX_RETRY_WAIT_TIME
		}
		time.Sleep(duration)
		resp, err = http.Get(client.Endpoint) // Retry the request after waiting
		if err != nil {
			return "", fmt.Errorf("error making GET request after waiting for rate limit: %w", err)
		}
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return string(body), err
}
