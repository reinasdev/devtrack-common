package config

import (
	"os"
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv(path string) error {
	if err := godotenv.Load(path); err != nil {
		return fmt.Errorf("arquivo .env não carregado de %s: %w", path, err)
	}

	return nil
}

func Get(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
