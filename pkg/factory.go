package pkg

import (
	"net/http"

	"github.com/cli/cli/pkg/iostreams"
)

// Factory for the cmd
type Factory struct {
	HTTPClient func() (*http.Client, error)
	DynamoDB   func()
	IOStreams  *iostreams.IOStreams
	Config     func() (*Config, error)
}

// NewFactory returns a new Factory
func NewFactory() *Factory {
	io := iostreams.System()

	var cachedConfig *Config
	var configError error
	configFunc := func() (*Config, error) {
		if cachedConfig != nil || configError != nil {
			return cachedConfig, configError
		}
		cachedConfig, configError = ParseDefaultConfig()
		return &Config{}, nil
	}

	return &Factory{
		HTTPClient: func() (*http.Client, error) {
			return NewHTTPClient(), nil
		},
		DynamoDB:  func() {},
		Config:    configFunc,
		IOStreams: io,
	}
}
