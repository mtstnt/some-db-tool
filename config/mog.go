package config

import (
	"fmt"
	"io"
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

	Project struct {
		MigrationDir string `yaml:"migration_dir"`
	}
}

func LoadMogConfig(fptr io.Reader) (MogConfig, error) {
	contents, err := io.ReadAll(fptr)
	if err != nil {
		return MogConfig{}, fmt.Errorf("LoadMogConfig: error loading mog config: %w", err)
	}

	var conf MogConfig
	if err := yaml.Unmarshal(contents, &conf); err != nil {
		return MogConfig{}, fmt.Errorf("LoadMogConfig: error unmarshalling yaml: %w", err)
	}
	return conf, nil
}

func ReadConfig(path string) (MogConfig, error) {
	fptr, err := os.Open(path)
	if err != nil {
		return MogConfig{}, fmt.Errorf("ReadConfig: error opening file (path: %s): %w", path, err)
	}

	return LoadMogConfig(fptr)
}
