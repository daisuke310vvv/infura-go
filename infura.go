package infura

import (
	"net/http"
)

type Infura struct {
	id        uint
	projectID string
	network   Network
	client    *http.Client
}

func New(config *Config) (*Infura, error) {
	infura := &Infura{
		projectID: config.ProjectID,
		network:   config.Network,
		client:    http.DefaultClient,
	}

	return infura, nil
}

func (i *Infura) CurrentNetwork() Network {
	return i.network
}
