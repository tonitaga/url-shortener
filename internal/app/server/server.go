package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tonitaga/url-shortener/internal/app/config"
)

type Server struct {
	config *config.Config
	logger *logrus.Logger
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

	if err := s.startServer(); err != nil {
		s.logger.WithError(err).Error("Failed to start server")
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

func (s *Server) startServer() error {
	const op = "server.Server.startServer"

	binding_address := fmt.Sprintf("%s:%d", s.config.Application.Host, s.config.Application.Port)
	s.logger.Infof("Server is listening on %s", binding_address)

	return fmt.Errorf("%s: %v", op, http.ListenAndServe(binding_address, nil))
}
