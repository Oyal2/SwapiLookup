package swapi

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/semaphore"
)

type SwapiClient struct {
	client  http.Client
	baseURL string
	sem     *semaphore.Weighted
}

func New(cfg SwapiClientConfig) (*SwapiClient, error) {
	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://swapi.dev/api"
	}

	if cfg.MaxConcurrent == 0 {
		return nil, fmt.Errorf("we cannot have a max concurrent value of 0")
	}

	return &SwapiClient{
		client: http.Client{
			Timeout: time.Second * 10,
		},
		baseURL: cfg.BaseURL,
		sem:     semaphore.NewWeighted(cfg.MaxConcurrent),
	}, nil
}
