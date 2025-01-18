package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
	"path/filepath"
)

type (
	Config struct {
		App                 AppConfiguration
		NotificationService NotificationServiceConfiguration
	}

	AppConfiguration struct {
		Port string
	}

	NotificationServiceConfiguration struct {
		Host string
		Port int
	}
)

func NewConfig() (Config, error) {
	var config Config

	root, err := os.Getwd()
	if err != nil {
		return Config{}, err
	}

	err = godotenv.Load(filepath.Join(root, ".env"))
	if err != nil {
		return Config{}, err
	}

	if err = envconfig.Process("NOTIFICATION_SERVICE", &config.NotificationService); err != nil {
		return Config{}, err
	}

	if err = envconfig.Process("APP", &config.App); err != nil {
		return Config{}, err
	}

	return config, nil
}
