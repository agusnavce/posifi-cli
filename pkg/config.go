package pkg

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config structure
type Config interface {
	GetHostname() string
}

type config struct {
	Hostname string
}

// ParseDefaultConfig - Parses default config
func ParseDefaultConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	fmt.Println(viper.Get("tests"))
	viper.Set("tests", "match2")
	err = viper.WriteConfigAs("./config.yml")
	if err != nil {
		return nil, err
	}
	return &config{
		Hostname: viper.GetString("hostname"),
	}, nil
}

func (c *config) GetHostname() string { return c.Hostname }
