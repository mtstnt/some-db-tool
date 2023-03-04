package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Mog configuration yaml.

type MogConfig struct {
	Database struct {
		Driver   string
		Host     string
		Port     int
		Name     string
		User     string
		Password string
	}
}

func LoadMogConfig(path string) (MogConfig, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return MogConfig{}, fmt.Errorf("error loading mog config on path %s: %w", path, err)
	}

	var conf MogConfig
	if err := yaml.Unmarshal(contents, &conf); err != nil {
		return MogConfig{}, fmt.Errorf("error unmarshalling yaml: %w", err)
	}
	return conf, nil
}
