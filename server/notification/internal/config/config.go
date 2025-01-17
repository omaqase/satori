package config

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

}
