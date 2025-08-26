package model

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stock       int       `json:"stock"`
	Price       float64   `json:"price"`
	CountryID   int       `json:"country_id"`
	CategoryID  int       `json:"category_id"`
	AddedAt     time.Time `json:"added_at"`
}
