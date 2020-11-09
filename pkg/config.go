package pkg

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/spf13/viper"
)

// Config structure
type Config interface {
	GetHostname() string
	GetAccessKey() string
	GetSecretKey() string
	GetRegion() string
}

type config struct {
	AWSCredentials credentials.Value
	AWSRegion      string
	Hostname       string
}

// ParseDefaultConfig - Parses default config
func ParseDefaultConfig() (Config, error) {
	c := credentials.NewSharedCredentials("", "default")
	creds, err := c.Get()
	if err != nil {
		return nil, err
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		return nil, err // Handle errors reading the config file
	}
	fmt.Println(viper.Get("tests"))
	viper.Set("tests", "match2")
	err = viper.WriteConfigAs("./config.yml")
	if err != nil { // Handle errors reading the config file
		return nil, err
	}
	return &config{
		AWSCredentials: creds,
		AWSRegion:      viper.GetString("aws_region"),
		Hostname:       viper.GetString("hostname"),
	}, nil
}

func (c *config) GetHostname() string  { return c.Hostname }
func (c *config) GetAccessKey() string { return c.AWSCredentials.AccessKeyID }
func (c *config) GetSecretKey() string { return c.AWSCredentials.SecretAccessKey }
func (c *config) GetRegion() string    { return c.AWSRegion }
