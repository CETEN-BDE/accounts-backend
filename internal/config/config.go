package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DbConfig struct {
		Dsn string `env:"DSN"`
	} `envPrefix:"DB_"`

	Port     string `env:"PORT" envDefault:"8080"`
	Host	 string `env:"HOST" envDefault:"0.0.0.0"`	
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

var config Config

func GetConfig() Config {
	return config
}

func init() {
	godotenv.Load()
	if err := env.Parse(&config); err != nil {
		logrus.Fatal(err)
	}

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logrus.Warnf("Invalid log level '%s', defaulting to 'info'", config.LogLevel)
		level = logrus.InfoLevel
		config.LogLevel = level.String()
	}
	logrus.SetLevel(level)
	logrus.Infof("Log level set to '%s'", level.String())

	logrus.Debugf("Loaded config : %+v", config)
}
