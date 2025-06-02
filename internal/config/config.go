package internal

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/vektra/mockery/v3/internal/file"
)

func FindConfig() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("getting current working directory: %w", err)
	}
	currentPath := filepath.ToSlash(cwd)
	for currentPath != "" {
		for _, confName := range []string{".mockery.yaml", ".mockery.yml"} {
			configPath := path.Join(currentPath, confName)
			exists, err := file.Exists(configPath)
			if err != nil {
				return "", fmt.Errorf("checking if %s exists: %w", configPath, err)
			}
			if exists {
				return configPath, nil
			}
		}
		currentPath = path.Dir(strings.TrimRight(currentPath, "/"))
	}
	return "", errors.New("mockery config file not found")
}
