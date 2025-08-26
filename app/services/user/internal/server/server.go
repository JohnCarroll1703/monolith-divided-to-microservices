package server

import (
	"log"
	userpb "monolith-divided-to-microservices/app/sdk/proto/user/v1"
	"monolith-divided-to-microservices/app/services/user/internal/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config, httpHandler http.Handler, userHandler userpb.UserServiceServer) error {
	grpcSrv := NewGRPCServer(cfg, userHandler)

	go func() {
		log.Printf("starting gRPC server on port %s", cfg.GrpcPort)
		if err := grpcSrv.Start(); err != nil {
			log.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

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

	grpcSrv.Stop()
	if err := httpSrv.Stop(); err != nil {
		log.Printf("error while stopping HTTP server: %v", err)
	}

	log.Println("Servers stopped gracefully")
	return nil
}
