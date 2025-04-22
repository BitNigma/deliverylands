package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		PG   string `yaml:"PG" env:"PG"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	log.Println(cfg.Port)

	cfg.PG = os.Getenv("PG")
	return cfg, nil
}
