package main

import (
	"monolith-divided-to-microservices/app/services/payment/internal/config"
	"monolith-divided-to-microservices/app/services/payment/internal/database"
	v1 "monolith-divided-to-microservices/app/services/payment/internal/delivery/http/v1"
	"monolith-divided-to-microservices/app/services/payment/internal/logging"
	"monolith-divided-to-microservices/app/services/payment/internal/repository"
	"monolith-divided-to-microservices/app/services/payment/internal/server"
	"monolith-divided-to-microservices/app/services/payment/internal/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg, err := config.GetConfig(".env")
	if err != nil {
		panic(err)
	}

	logg := logging.InitLogger(cfg.LogLevel)
	logg.Info("Logger initialized successfully")

	pgPool, err := database.NewPostgresPool(cfg.Databases.PostgresDSN)
	if err != nil {
		logg.Fatalf("Postgres connection failed: %v", err)
	}
	defer pgPool.Close()
	logg.Info("Postgres connection established")

	repos := repository.NewRepository(pgPool)
	services := service.NewServices(repos, cfg)
	h := v1.NewHandler(logg.Logger, services)

	r := h.Init()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := server.Run(cfg, r); err != nil {
		logg.Fatalf("server stopped with error: %v", err)
	}
}
