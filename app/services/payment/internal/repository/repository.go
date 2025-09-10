package repository

import (
	"context"
	"monolith-divided-to-microservices/app/services/payment/internal/model"
	"monolith-divided-to-microservices/app/services/payment/internal/repository/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	PaymentRepo *postgres.PaymentRepository
}

type Payment interface {
	SavePayment(ctx context.Context, p *model.Payment) error
	GetSessionByID(ctx context.Context, sID string) (*model.Payment, error)
	UpdateStatus(ctx context.Context, sessionID, status string) error
}

func NewRepository(db *pgxpool.Pool) *Repositories {
	paymentRepo := postgres.NewPayment(db)
	return &Repositories{
		PaymentRepo: paymentRepo,
	}
}
