package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `env:"ENV" env-description:"Environment of application (local | dev | prod)"`

	App struct {
		Name    string `env:"APP_NAME" env-required:"true" env-description:"Application name"`
		Version string `env:"APP_VERSION" env-required:"true" env-description:"Application version"`
		Host    string `env:"APP_HOST" env-required:"true" env-description:"Application host"`
		Port    int    `env:"APP_PORT" env-required:"true" env-description:"Port that application runs on"`
	}

	Database struct {
		Host     string `env:"DB_HOST" env-required:"true" env-description:"Host of database"`
		Port     int    `env:"DB_PORT" env-required:"true" env-description:"Database port"`
		User     string `env:"DB_USER" env-required:"true" env-description:"Database user"`
		Password string `env:"DB_PASS" env-required:"true" env-description:"Database user's password"`
		Name     string `env:"DB_NAME" env-required:"true" env-description:"Database name"`
	}
}

func New() *Config {
	config := new(Config)

	if err := cleanenv.ReadEnv(config); err != nil {
		header := fmt.Sprintf("%s - %s", os.Getenv("APP_NAME"), os.Getenv("APP_VERSION"))

		usage := cleanenv.FUsage(os.Stdout, config, &header)
		usage()

		os.Exit(-1)
	}

	return config
}
