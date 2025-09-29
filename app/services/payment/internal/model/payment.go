package model

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	OrderID   uint      `json:"order_id"`
	Amount    float64   `json:"amount"`
	SessionID string    `json:"session_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type PaymentCreatedEvent struct {
	SessionID string    `json:"session_id"`
	PaymentID uuid.UUID `json:"payment_id"`
	OrderID   string    `json:"order_id"`
	UserID    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}
