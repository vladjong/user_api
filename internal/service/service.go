package service

import (
	"context"
	"os"
	"os/signal"
	"refactoring/config"
	jsondb "refactoring/internal/adapters/db/jsonDB"
	"refactoring/internal/controller/handler"
	"refactoring/pkg/server"
	"syscall"

	"github.com/sirupsen/logrus"
)

type Service struct {
	cfg *config.Config
}

func New(cfg *config.Config) (service Service, err error) {
	return Service{
		cfg: cfg,
	}, nil
}

func (s *Service) Run() error {
	s.startHTTP()
	return nil
}

func (s *Service) startHTTP() {
	server := new(server.Server)
	storage := jsondb.New()
	handler := handler.New(storage)
	go func() {
		if err := server.Run(s.cfg.Listen.Port, handler.NewRouter()); err != nil {
			logrus.Fatalf("error: occured while running HTTP server: %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Info("HTTP server shutdown")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error: occured on server shutdown: %s", err.Error())
	}
}
