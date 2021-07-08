package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

//config for service
type ServerConfig struct {
	ServiceHost string `env:"SERVICE_HOST"`
	ServicePort string `env:"SERVICE_PORT"`
	DB          Database
}

//Database config
type Database struct {
	Name     string `env:"DB_SCHEMA" default:"db_1"`
	Adapter  string `env:"DB_DRIVER" default:"mysql"`
	Host     string `env:"DB_HOST" default:"localhost"`
	Port     string `env:"DB_PORT" default:"3306"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	SslMode  string `env:"DB_SSL_MODE"`
}

var Config ServerConfig

func init() {
	err := loadConfig()
	if err != nil {
		panic(err)
	}
}

func loadConfig() (err error) {
	err = godotenv.Load()
	if err != nil {
		log.Warn().Msg("Cannot find .env file. OS Environments will be used")
	}
	err = env.Parse(&Config)

	return err
}
