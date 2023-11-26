package httpconfig

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server serverConfig
}

type serverConfig struct {
	Host    string
	Port    string
	Timeout time.Duration `yaml:"timeout,omitempty"`
}

func GetConfig(fileName string) (*Config, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
