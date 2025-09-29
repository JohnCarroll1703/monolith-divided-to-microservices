package service

import (
	"context"
	"monolith-divided-to-microservices/app/sdk/kafka"
	"monolith-divided-to-microservices/app/services/payment/internal/config"
	"monolith-divided-to-microservices/app/services/payment/internal/repository"
	"monolith-divided-to-microservices/app/services/payment/internal/repository/postgres"
	"monolith-divided-to-microservices/app/services/payment/internal/schema"

	"github.com/stripe/stripe-go/v76"
)

type PaymentService struct {
	repo     *postgres.PaymentRepository
	cfg      *config.Config
	producer *kafka.Producer
}

type Payment interface {
	CreatePayment(ctx context.Context, req *schema.SavePaymentRequest) (*stripe.CheckoutSession, error)
	GetPaymentStatus(sessionID string) (*stripe.CheckoutSession, error)
}

func NewServices(repo *repository.Repositories, cfg *config.Config) *Services {
	return &Services{
		PaymentService: NewPaymentService(repo.PaymentRepo, cfg),
	}
}
