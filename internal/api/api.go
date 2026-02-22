package api

import (
	"fmt"
	"io"
	"net/http"

	_ "golang.org/x/crypto/x509roots/fallback"
)

type ApiClient struct {
	Endpoint string
}

func (client ApiClient) GetPublicIp() (string, error) {
	resp, err := http.Get(client.Endpoint)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		askedToRetryAfter := resp.Header.Get("Retry-After")
		if askedToRetryAfter != "" {
			return "", fmt.Errorf("too many requests: %d, retry after: %ss", resp.StatusCode, askedToRetryAfter)
		}
		return "", fmt.Errorf("too many requests: %d", resp.StatusCode)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	return string(body), err
}
