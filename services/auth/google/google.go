package google

import "tinyUrl/context"

type Config struct{}

type GoogleProvider struct {
	*Config
}

func New(cfg *Config) *GoogleProvider {
	return &GoogleProvider{cfg}
}

func (provider *GoogleProvider) Login(context *context.Context) (string, error) {
	return "", nil
}

func (provider *GoogleProvider) Register(context *context.Context) (string, error) {
	return "", nil
}
