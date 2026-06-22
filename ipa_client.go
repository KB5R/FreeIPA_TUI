package main

import (
	"crypto/tls"
	"net/http"

	"github.com/infra-monkey/go-freeipa/freeipa"
)

type IPAConfig struct {
	Host     string
	Username string
	Password string
	Insecure bool
}

type IPAClient struct {
	api *freeipa.Client
}

// * означает что мы работает с одним клиетом а не перезаписываем все
func NewIPAClient(config IPAConfig) (*IPAClient, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.Insecure,
		},
	}

	api, err := freeipa.Connect(
		config.Host,
		transport,
		config.Username,
		config.Password,
	)
	if err != nil {
		return nil, err
	}

	return &IPAClient{
		api: api,
	}, nil
}
