package server

import (
	"context"
	"net/http"

	"github.com/anton-ag/javacode/internal/config"
)

type Server struct {
	http *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		http: &http.Server{
			Addr:    ":" + cfg.HTTP.Port,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.http.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
