package postgres

import (
	"context"
	"monolith-divided-to-microservices/app/services/payment/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentRepository struct {
	db *pgxpool.Pool
}

func NewPayment(db *pgxpool.Pool) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) SavePayment(ctx context.Context, p *model.Payment) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}

	query := `INSERT INTO payments (id, user_id, amount, order_id, session_id, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(ctx, query, p.ID, p.UserID, p.Amount, p.OrderID, p.SessionID, p.CreatedAt)
	return err
}

func (r *PaymentRepository) GetSessionByID(ctx context.Context, sID string) (*model.Payment, error) {
	var p model.Payment
	row := r.db.QueryRow(ctx,
		`SELECT id, user_id, order_id, amount, session_id, status, created_at
		FROM payments WHERE session_id = $1`, sID)

	if err := row.Scan(&p.ID, &p.UserID, &p.OrderID, &p.Amount,
		&p.SessionID, &p.Status, &p.CreatedAt); err != nil {
		return nil, nil
	}

	return &p, nil
}

func (r *PaymentRepository) UpdateStatus(ctx context.Context, sessionID, status string) error {
	_, err := r.db.Exec(ctx,
		`UPDATE payments SET status=$1 WHERE session_id=$2`,
		status, sessionID,
	)
	return err
}
