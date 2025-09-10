package server

import (
	"log"
	"monolith-divided-to-microservices/app/services/payment/internal/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config, httpHandler http.Handler) error {

	httpSrv := NewHTTPServer(cfg, httpHandler)

	go func() {
		log.Printf("starting HTTP server on port: %s", cfg.AppPort)
		if err := httpSrv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start HTTP server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Printf("shutting down servers...")

	if err := httpSrv.Stop(); err != nil {
		log.Printf("error while stopping HTTP server: %v", err)
	}

	log.Println("Servers stopped gracefully")
	return nil
}
