package openai

import (
	"errors"
	"net/http"
)

type Config struct {
	URL   string
	Token string
}

type Client struct {
	baseURL string
	token   string
	http    *http.Client
}

func New(cfg Config) (*Client, error) {
	if cfg.URL == "" {
		return nil, errors.New("empty url")
	}

	if cfg.Token == "" {
		return nil, errors.New("empty token")
	}

	return &Client{
		baseURL: cfg.URL,
		token:   cfg.Token,
		http: &http.Client{
			Transport: http.DefaultTransport,
		},
	}, nil
}
