package main

import (
	"monolith-divided-to-microservices/app/services/user/internal/auth"
	"monolith-divided-to-microservices/app/services/user/internal/config"
	"monolith-divided-to-microservices/app/services/user/internal/database"
	grpcdelivery "monolith-divided-to-microservices/app/services/user/internal/delivery/grpc"
	v1 "monolith-divided-to-microservices/app/services/user/internal/delivery/http/v1"
	"monolith-divided-to-microservices/app/services/user/internal/logging"
	"monolith-divided-to-microservices/app/services/user/internal/repository"
	"monolith-divided-to-microservices/app/services/user/internal/server"
	"monolith-divided-to-microservices/app/services/user/internal/service"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg, err := config.GetConfig("../.env")
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

	jwtManager := auth.NewJWTManager(cfg.JWTSecret, 24*time.Hour)
	repos := repository.NewRepository(pgPool)
	services := service.NewServices(repos, cfg)
	h := v1.NewHandler(logg.Logger, services, jwtManager)
	grpcHandler := grpcdelivery.NewUserHandler(services.UserService)

	r := h.Init()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := server.Run(cfg, r, grpcHandler); err != nil {
		logg.Fatalf("server stopped with error: %v", err)
	}
}
