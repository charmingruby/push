package config

import (
	"log/slog"

	env "github.com/caarlos0/env/v6"
)

type environment struct {
	MongoURL   string `env:"MONGO_URL,required"`
	ServerPort string `env:"SERVER_PORT,required"`
	ServerHost string `env:"SERVER_HOST,required"`
}

func NewConfig() (*Config, error) {
	slog.Info("Loading environment...")
	environment := environment{}
	if err := env.Parse(&environment); err != nil {
		return nil, err
	}

	slog.Info("Environment loaded successfully!")

	cfg := Config{
		MongoConfig: &mongoConfig{
			URL: environment.MongoURL,
		},
		ServerConfig: &serverConfig{
			Port: environment.ServerPort,
			Host: environment.ServerHost,
		},
	}

	return &cfg, nil
}

type Config struct {
	MongoConfig  *mongoConfig
	ServerConfig *serverConfig
}

type mongoConfig struct {
	URL string
}

type serverConfig struct {
	Port string
	Host string
}
