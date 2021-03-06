package config

import (
	"os"
	"os/user"
	"path/filepath"
)

func getConfigDirectory() (string, error) {
	configDir := os.Getenv("XDG_CONFIG_DIR")

	if configDir != "" {
		return configDir, nil
	}

	currentUser, userError := user.Current()

	if userError != nil {
		return "", userError
	}

	return filepath.Join(currentUser.HomeDir, ".config", AppNameLowercase), nil
}
