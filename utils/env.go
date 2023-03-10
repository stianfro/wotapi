// Package utils stores utility functions
package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

func loadEnv(fileName string) error {
	file, err := os.Open(filepath.Clean(fileName))
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to open file")
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := parts[0]
		value := parts[1]
		err := os.Setenv(key, value)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Failed to set environment variable")
		}
	}

	if err := file.Close(); err != nil {
		log.Error().
			Msg("Failed to close file")
	}

	return scanner.Err()
}

// SetEnv loads environment variables from a file if ENVFILE is set
func SetEnv() {
	envFile := os.Getenv("ENVFILE")
	if envFile != "" {
		if err := loadEnv(envFile); err != nil {
			log.Error().
				Err(err).
				Msg("Failed to load environment variables from file")
		} else {
			log.Info().
				Msg("Loaded environment variables from file")
		}
	}
}
