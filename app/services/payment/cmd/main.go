package main

import (
	"context"
	"log"
	kf "monolith-divided-to-microservices/app/sdk/kafka"
	"monolith-divided-to-microservices/app/sdk/logging"
	"monolith-divided-to-microservices/app/services/payment/internal/config"
	"monolith-divided-to-microservices/app/services/payment/internal/database"
	v1 "monolith-divided-to-microservices/app/services/payment/internal/delivery/http/v1"
	"monolith-divided-to-microservices/app/services/payment/internal/repository"
	"monolith-divided-to-microservices/app/services/payment/internal/server"
	"monolith-divided-to-microservices/app/services/payment/internal/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.GetConfig(".env")
	if err != nil {
		panic(err)
	}

	logg := logging.InitLogger(cfg.LogLevel, cfg.AppName)
	logg.Info("Logger initialized successfully")

	pgPool, err := database.NewPostgresPool(cfg.Databases.PostgresDSN)
	if err != nil {
		logg.Fatalf("Postgres connection failed: %v", err)
	}

	logg.Info("Postgres connection established")

	producer := kf.NewProducer(
		[]string{"localhost:9092"},
		"payment",
	)

	consumer := kf.NewConsumer(
		[]string{"localhost:9092"},
		"orders", "payment-service")

	go func() {
		err := consumer.ReadMessage(ctx, func(m kafka.Message) {
			logg.Infof("got message at offset %d: key=%s value=%s",
				m.Offset, string(m.Key), string(m.Value))
		})
		if err != nil {
			logg.Infof("consumer stopped: %v", err)
		}
	}()

	err = producer.WriteMessage(ctx, []byte("payment_id"), []byte("Payment service started"))
	if err != nil {
		logg.Infof("Error producing message: %v", err)
	}

	repos := repository.NewRepository(pgPool)
	services := service.NewServices(repos, cfg)
	h := v1.NewHandler(logg.Logger, services)

	r := h.Init()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := server.Run(cfg, r); err != nil {
		logg.Fatalf("server stopped with error: %v", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("Shutting down payment service...")

	cancel() // Cancel the context to stop the consumer

	consumer.Close()
	producer.Close()
	pgPool.Close()
}
