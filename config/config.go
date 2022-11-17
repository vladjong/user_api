package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Listen struct {
		Port string `env:"PORT" env-default:":3333"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logrus.Print("Read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "Service config"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			logrus.Print(help)
			logrus.Fatal(err)
		}
	})
	return instance
}
