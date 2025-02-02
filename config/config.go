package config

import (
	"github.com/caarlos0/env/v10"
)

type Config struct {
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port int    `env:"PORT" envDefault:"80"`
}

func New() (cfg *Config, err error) {
	cfg = &Config{}
	err = env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
