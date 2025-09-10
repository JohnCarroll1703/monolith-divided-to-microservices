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
