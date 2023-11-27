package config

import (
	"errors"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// konfigurasi aplikasi (port, database, jwt, session)
type Config struct {
	Port     string         `env:"PORT" envDefault:"8080"`
	Postgres PostgresConfig `envPrefix:"POSTGRES_"`
	JWT      JwtConfig      `envPrefix:"JWT_"`
}

// konfigurasi untuk JWT
type JwtConfig struct {
	SecretKey string `env:"SECRET_KEY"`
}

// konfigurasi database
type PostgresConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Database string `env:"DATABASE" envDefault:"depublic"`
}

// NewConfig digunakan untuk membuat config baru
func NewConfig(envPath string) (*Config, error) {
	cfg, err := parseConfig(envPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// parseConfig parses the configuration file located at envPath and returns a
// Config struct and an error if any. It uses the godotenv package to load the
// environment variables from the file and the env package to parse them into
// the Config struct.
//
// envPath: The path to the environment file.
// Returns: A pointer to the Config struct and an error.

func parseConfig(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, errors.New("failed to load .env file")
	}

	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		return nil, errors.New("failed to parse environment variables Config")
	}
	return cfg, nil
}

//Note
//direktori config digunakan untuk konfigurasi database, jwt, dan port
