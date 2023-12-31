package config

import (
	"errors"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// konfigurasi aplikasi (port, database, jwt, session)
type Config struct {
	Port           string         `env:"PORT" envDefault:"8080"`
	Postgres       PostgresConfig `envPrefix:"POSTGRES_"`
	JWT            JwtConfig      `envPrefix:"JWT_"`
	MidtransConfig MidtransConfig `envPrefix:"MIDTRANS_"`
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
	Database string `env:"DATABASE" envDefault:"depublic-db"`
}

// konfigurasi midtrans
type MidtransConfig struct {
	MidtransServerKey string `env:"SERVER_KEY"`
	MidtransClientKey string `env:"CLIENT_KEY"`
	BaseURL           string `env:"BASE_URL"`
	MidtransEnv       string `env:"ENV"`
}

// NewConfig digunakan untuk membuat config baru
func NewConfig(envPath string) (*Config, error) {
	cfg, err := parseConfig(envPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

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
