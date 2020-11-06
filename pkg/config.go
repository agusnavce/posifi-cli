package pkg

import "github.com/aws/aws-sdk-go/aws/credentials"

// Config structure
type Config struct {
	AWSCredentials credentials.Value
	AWSRegion      string
	Hostname       string
}

// ParseDefaultConfig - Parses default config
func ParseDefaultConfig() (*Config, error) {
	return &Config{}, nil
}
