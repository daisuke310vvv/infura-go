package infura

import (
	"net/http"
)

type Infura struct {
	id      uint
	apiKey  string
	network Network
	client  *http.Client
}

func New(config *Config) (*Infura, error) {
	infura := &Infura{
		apiKey:  config.ApiKey,
		network: config.Network,
		client:  http.DefaultClient,
	}

	return infura, nil
}
