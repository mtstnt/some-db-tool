package config

import (
	"os"
	"path/filepath"
)

var (
	ExecutablePath string
	CurrentPath    string
)

// Loads all configurations.
func LoadGlobals() error {
	var err error

	executablePath, err := os.Executable()
	if err != nil {
		return err
	}
	ExecutablePath = filepath.Join(executablePath, "..")

	CurrentPath, err = os.Getwd()
	if err != nil {
		return err
	}

	return nil
}
