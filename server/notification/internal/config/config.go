package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"path/filepath"
)

type Config struct {
	App    AppConfig
	Resend ResendConfig
}

type AppConfig struct {
	Port string
}

type ResendConfig struct {
	ApiKey string
}

func NewConfig() (*Config, error) {
	config := &Config{}

	root, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		return nil, err
	}

	if err = envconfig.Process("APP", config.App); err != nil {
		return nil, err
	}

	if err = envconfig.Process("RESEND", config.Resend); err != nil {
		return nil, err
	}

	return config, nil
}
