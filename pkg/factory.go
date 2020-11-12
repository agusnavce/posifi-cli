package pkg

import (
	"github.com/cli/cli/pkg/iostreams"
)

// Factory for the cmd
type Factory struct {
	HTTPClient func() (Client, error)
	IOStreams  *iostreams.IOStreams
	Config     func() (Config, bool, error)
}

// NewFactory returns a new Factory
func NewFactory() *Factory {
	io := iostreams.System()

	var cachedConfig Config
	var configError error
	configFunc := func() (Config, bool, error) {
		if cachedConfig != nil || configError != nil {
			return cachedConfig, false, configError
		}
		cachedConfig, configError = ParseDefaultConfig()
		return cachedConfig, false, nil
	}

	return &Factory{
		HTTPClient: func() (Client, error) {
			cfg, hasConfig, err := configFunc()
			if err != nil || !hasConfig {
				return nil, err
			}
			return NewHTTPClient(cfg), nil
		},
		Config:    configFunc,
		IOStreams: io,
	}
}
