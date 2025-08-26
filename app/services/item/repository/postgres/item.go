package postgres

import (
	"awesomeProject666/app/internal/model"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ItemRepository struct {
	db *pgxpool.Pool
}

func NewItem(db *pgxpool.Pool) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (r *ItemRepository) GetItems(ctx context.Context) ([]model.Item, error) {
	query := `SELECT id, name, description, category_id, country_id, stock, price, added_at FROM items`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err = rows.Scan(&item.ID, &item.Name, &item.Description, &item.CategoryID, &item.CountryID,
			&item.Stock, &item.Price, &item.AddedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
