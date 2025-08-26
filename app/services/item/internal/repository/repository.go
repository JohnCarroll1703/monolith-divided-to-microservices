package repository

import (
	"context"
	"monolith-divided-to-microservices/app/services/item/internal/model"
	"monolith-divided-to-microservices/app/services/item/internal/repository/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	ItemRepo *postgres.ItemRepository
}

func NewRepository(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		ItemRepo: postgres.NewItem(db),
	}
}

type Item interface {
	GetItems(ctx context.Context) ([]model.Item, error)
	CreateItem(ctx context.Context, item *model.Item) error
}
