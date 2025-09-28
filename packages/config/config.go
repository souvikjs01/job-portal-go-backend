package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Server   ServerConfig   `koanf:"server" validate:"required"`
	Database DatabaseConfig `koanf:"database" validate:"required"`
	App      AppConfig      `koanf:"app" validate:"required"`
	Auth     AuthConfig     `koanf:"auth" validate:"required"`
}

type ServerConfig struct {
	Port string `koanf:"port" validate:"required"`
}

type DatabaseConfig struct {
	DB_URL string `koanf:"db_url" validate:"required"`
}

type AppConfig struct {
	Env string `koanf:"env" validate:"required"`
}

type AuthConfig struct {
	Jwt_secret  string        `koanf:"jwt_secret" validate:"required"`
	TokenExpiry time.Duration `koanf:"token_expiry" validate:"required"`
}

func LoadConfig() (*Config, error) {
	k := koanf.New(".")

	err := k.Load(env.Provider("JOBAPP_", ".", func(s string) string {
		return strings.ToLower(strings.TrimPrefix(s, "JOBAPP_"))
	}), nil)

	if err != nil {
		return nil, fmt.Errorf("could not load initial env variable")
	}

	mainConfig := &Config{}

	err = k.Unmarshal("", mainConfig)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	mainConfig.Auth.TokenExpiry = 24 * time.Hour

	validate := validator.New()

	err = validate.Struct(mainConfig)
	if err != nil {
		return nil, fmt.Errorf("config validation failed")
	}

	return mainConfig, nil
}
