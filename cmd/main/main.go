package main

import (
	"refactoring/config"
	"refactoring/internal/service"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("config initializing")
	cfg := config.GetConfig()
	logrus.Info("running service")
	service, err := service.New(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	if err := service.Run(); err != nil {
		logrus.Fatal(err)
	}
}
