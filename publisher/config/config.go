package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("config: %W", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
