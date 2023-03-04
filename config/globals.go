package config

import (
	"os"
	"path/filepath"
)

type GlobalInfo struct {
	ExecutablePath string
	CurrentPath    string
}

// Loads all configurations.
func LoadGlobals() (GlobalInfo, error) {
	var err error

	executablePath, err := os.Executable()
	if err != nil {
		return GlobalInfo{}, err
	}
	executablePath = filepath.Join(executablePath, "..")

	currentPath, err := os.Getwd()
	if err != nil {
		return GlobalInfo{}, err
	}

	return GlobalInfo{
		ExecutablePath: executablePath,
		CurrentPath:    currentPath,
	}, nil
}
