package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tonitaga/url-shortener/internal/config"
	"github.com/tonitaga/url-shortener/internal/storage"
)

type Server struct {
	config  *config.Config
	logger  *logrus.Logger
	storage *storage.Storage
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
	}
}

func (s *Server) Run() error {
	// Logger configuration
	s.configureLogger()
	s.logger.Info("Logger configuration finished successfully")

	// Database configuration
	s.logger.Info("Starting storage configuration")
	if err := s.configureStorage(); err != nil {
		return err
	}

	// Starting server
	if err := s.startServer(); err != nil {
		return err
	}

	return nil
}

func (s *Server) configureLogger() {
	switch s.config.Application.Environment {
	case config.EnvLocal:
		s.logger.SetLevel(logrus.DebugLevel)
		s.logger.SetFormatter(&logrus.TextFormatter{})
	case config.EnvDev:
		s.logger.SetLevel(logrus.DebugLevel)
		s.logger.SetFormatter(&logrus.JSONFormatter{})
	case config.EnvProd:
		s.logger.SetLevel(logrus.InfoLevel)
		s.logger.SetFormatter(&logrus.JSONFormatter{})
	}

	s.logger.SetOutput(os.Stdout)
}

func (s *Server) configureStorage() error {
	storage := storage.New(&s.config.Database)
	if err := storage.Connect(); err != nil {
		s.logger.WithError(err).Error("Failed to connect to database")
		return err
	}

	s.storage = storage
	s.logger.Info("Database configuration finishes successfully")
	return nil
}

func (s *Server) startServer() error {
	const op = "server.Server.startServer"

	binding_address := fmt.Sprintf("%s:%d", s.config.Application.Host, s.config.Application.Port)
	s.logger.Infof("Server is listening on %s", binding_address)

	return fmt.Errorf("%s: %v", op, http.ListenAndServe(binding_address, nil))
}
