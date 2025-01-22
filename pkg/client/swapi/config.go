package swapi

import (
	"os"
	"strconv"
)

type SwapiClientConfig struct {
	BaseURL       string
	MaxConcurrent int64
}

func DefaultSwapiConfig() SwapiClientConfig {
	cfg := SwapiClientConfig{
		BaseURL:       os.Getenv("SWAPI_BASE_URL"),
		MaxConcurrent: 0,
	}

	maxConcurrent := os.Getenv("SWAPI_MAX_CONCURRENT")
	if num, err := strconv.ParseInt(maxConcurrent, 10, 64); err == nil {
		cfg.MaxConcurrent = num
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://swapi.dev/api"
	}

	if cfg.MaxConcurrent == 0 {
		cfg.MaxConcurrent = 10
	}

	return cfg
}
