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
	Port string `envconfig:"PORT"` // добавляем тег
}

type ResendConfig struct {
	ApiKey string `envconfig:"APIKEY"` // добавляем тег
}

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

	// Обрабатываем каждую секцию конфигурации отдельно
	if err := envconfig.Process("APP", &config.App); err != nil {
		return Config{}, err
	}

	if err := envconfig.Process("RESEND", &config.Resend); err != nil {
		return Config{}, err
	}

	// Добавим проверку, что порт установлен
	if config.App.Port == "" {
		config.App.Port = "50054" // значение по умолчанию
	}

	return config, nil
}
