package chatgpt

import (
	"errors"
	"net/http"
)

type Config struct {
	URL   string
	Token string
	Model string
}

type Client struct {
	http *http.Client

	baseURL string
	token   string
	model   string
}

func New(cfg Config) (*Client, error) {
	if cfg.URL == "" {
		return nil, errors.New("empty url")
	}

	if cfg.Token == "" {
		return nil, errors.New("empty token")
	}

	return &Client{
		http: &http.Client{
			Transport: http.DefaultTransport,
		},
		baseURL: cfg.URL,
		token:   cfg.Token,
		model:   cfg.Model,
	}, nil
}
