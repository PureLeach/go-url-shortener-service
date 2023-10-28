package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `env:"env" env-default:"local"`
	StoragePath string        `env:"storage_path" env-required:"true"`
	Address     string        `env:"address" env-required:"true"`
	Timeout     time.Duration `env:"timeout" env-required:"true"`
	IdleTimeout time.Duration `env:"idle_timeout" env-required:"true"`
	User        string        `env:"user" env-required:"true"`
}

func ConfigLoad() *Config {

	var cfg Config

	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		log.Fatalf("Configuration read failed: %s", err)
	}
	return &cfg
}
