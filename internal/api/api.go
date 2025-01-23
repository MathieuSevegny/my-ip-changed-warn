package api

import (
	"io"
	"net/http"

	_ "golang.org/x/crypto/x509roots/fallback"
)

type ApiClient struct {
	Endpoint string
}

const DefaultEndpoint = "https://api.ipify.org/"

func (client ApiClient) GetPublicIp() (string, error) {
	if client.Endpoint == "" {
		client.Endpoint = DefaultEndpoint
	}

	resp, err := http.Get(client.Endpoint)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)

	return string(body), err
}
