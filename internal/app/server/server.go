package server

import (
	"fmt"
	"net/http"

	"github.com/tonitaga/url-shortener/internal/app/config"
)

type Server struct {
	config *config.Config
}

func New(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Run() error {
	binding_address := fmt.Sprintf("%s:%d", s.config.Application.Host, s.config.Application.Port)
	return http.ListenAndServe(binding_address, nil)
}
