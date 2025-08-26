package server

import (
	"context"
	"log"
	"monolith-divided-to-microservices/app/services/item/internal/config"
	"net/http"
	"time"
)

type HTTPServer struct {
	srv *http.Server
	cfg *config.Config
}

func NewHTTPServer(cfg *config.Config, handler http.Handler) *HTTPServer {
	srv := &http.Server{
		Addr:    cfg.AppPort,
		Handler: handler,
	}

	return &HTTPServer{srv: srv, cfg: cfg}
}

func (s *HTTPServer) Start() error {
	log.Printf("HTTP server started on %s", s.cfg.AppPort)
	return s.srv.ListenAndServe()
}

func (s *HTTPServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("HTTP server stopped")
	return s.srv.Shutdown(ctx)
}
