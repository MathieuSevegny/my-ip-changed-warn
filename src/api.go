package src

import (
	"io"
	"net/http"
)

type ApiClient struct {
	Endpoint string
}

func (client ApiClient) GetPublicIp() (string, error) {
	resp, err := http.Get(client.Endpoint)
	if err == nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)

	return string(body), err
}
