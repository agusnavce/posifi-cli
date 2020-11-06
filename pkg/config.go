package pkg

import "github.com/aws/aws-sdk-go/aws/credentials"

// Config structure
type Config interface {
	Get() string
	Set(string) error
}

type config struct {
	AWSCredentials credentials.Value
	AWSRegion      string
	Hostname       string
}

// ParseDefaultConfig - Parses default config
func ParseDefaultConfig() (Config, error) {
	return &config{}, nil
}

func (c *config) Get() string      { return "" }
func (c *config) Set(string) error { return nil }
